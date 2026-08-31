package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients/dto"
	"github.com/williamwinkler/hs-card-service/internal/observability"
)

const (
	PAGE_SIZE          = 250
	httpRequestTimeout = 15 * time.Second
)

type HsClient struct {
	token token

	client  *http.Client
	tokenMu sync.Mutex
}

type token struct {
	accessToken string
	expireDate  time.Time
}

func NewHsClient() (*HsClient, error) {
	client := newHTTPClient()
	initialToken, err := fetchToken(context.Background(), client)
	if err != nil {
		return &HsClient{}, err
	}

	return &HsClient{token: initialToken, client: client}, nil
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   httpRequestTimeout,
		Transport: http.DefaultTransport,
	}
}

func (hc *HsClient) GetAllCards(ctx context.Context) ([]domain.Card, error) {
	cardMap := make(map[string]domain.Card)

	for page := 1; ; page++ {
		cards, err := hc.GetCardsWithPagination(ctx, page, PAGE_SIZE)
		if err != nil {
			return nil, err
		}
		if len(cards) == 0 {
			break
		}
		for _, card := range cards {
			cardMap[card.Name] = card
		}
	}

	cardsList := make([]domain.Card, 0, len(cardMap))
	for _, card := range cardMap {
		cardsList = append(cardsList, card)
	}
	return cardsList, nil
}

func (hc *HsClient) GetCardsWithPagination(ctx context.Context, page int, pageSize int) ([]domain.Card, error) {
	queryParams := url.Values{}
	queryParams.Set("locale", "en_US")
	queryParams.Set("page", strconv.Itoa(page))
	queryParams.Set("pageSize", strconv.Itoa(pageSize))

	response, err := hc.executeGetRequest(ctx, "https://eu.api.blizzard.com/hearthstone/cards?"+queryParams.Encode())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := wait(ctx); err != nil {
		return nil, err
	}

	var cardsDTO dto.CardsDto
	if err := json.NewDecoder(response.Body).Decode(&cardsDTO); err != nil {
		return nil, fmt.Errorf("decode cards response: %w", err)
	}
	return dto.MapToCards(cardsDTO), nil
}

func (hc *HsClient) GetSets(ctx context.Context) ([]domain.Set, error) {
	response, err := hc.executeGetRequest(ctx, "https://us.api.blizzard.com/hearthstone/metadata/sets?locale=en_US")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := wait(ctx); err != nil {
		return nil, err
	}

	var setsDTO dto.SetsDto
	if err := json.NewDecoder(response.Body).Decode(&setsDTO); err != nil {
		return nil, fmt.Errorf("decode sets response: %w", err)
	}
	return dto.MapToSets(setsDTO), nil
}

func (hc *HsClient) GetClasses(ctx context.Context) ([]domain.Class, error) {
	response, err := hc.executeGetRequest(ctx, "https://us.api.blizzard.com/hearthstone/metadata/classes?locale=en_US")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := wait(ctx); err != nil {
		return nil, err
	}

	var classesDTO dto.ClassesDto
	if err := json.NewDecoder(response.Body).Decode(&classesDTO); err != nil {
		return nil, fmt.Errorf("decode classes response: %w", err)
	}
	return dto.MapToClasses(classesDTO), nil
}

func (hc *HsClient) GetRarities(ctx context.Context) ([]domain.Rarity, error) {
	response, err := hc.executeGetRequest(ctx, "https://us.api.blizzard.com/hearthstone/metadata/rarities?locale=en_US")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := wait(ctx); err != nil {
		return nil, err
	}

	var raritiesDTO dto.RaritiesDto
	if err := json.NewDecoder(response.Body).Decode(&raritiesDTO); err != nil {
		return nil, fmt.Errorf("decode rarities response: %w", err)
	}
	return dto.MapToRariteis(raritiesDTO), nil
}

func (hc *HsClient) GetTypes(ctx context.Context) ([]domain.Type, error) {
	response, err := hc.executeGetRequest(ctx, "https://us.api.blizzard.com/hearthstone/metadata/types?locale=en_US")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := wait(ctx); err != nil {
		return nil, err
	}

	var typesDTO dto.TypesDto
	if err := json.NewDecoder(response.Body).Decode(&typesDTO); err != nil {
		return nil, fmt.Errorf("decode types response: %w", err)
	}
	return dto.MapToTypes(typesDTO), nil
}

func (hc *HsClient) GetKeywords(ctx context.Context) ([]domain.Keyword, error) {
	response, err := hc.executeGetRequest(ctx, "https://us.api.blizzard.com/hearthstone/metadata/keywords?locale=en_US")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := wait(ctx); err != nil {
		return nil, err
	}

	var keywordsDTO dto.KeywordsDto
	if err := json.NewDecoder(response.Body).Decode(&keywordsDTO); err != nil {
		return nil, fmt.Errorf("decode keywords response: %w", err)
	}
	return dto.MapToKeywords(keywordsDTO), nil
}

func (hc *HsClient) executeGetRequest(ctx context.Context, endpoint string) (*http.Response, error) {
	ctx, span := observability.Start(ctx, "blizzard.api.request", attribute.String("http.request.method", http.MethodGet))
	defer span.End()

	accessToken, err := hc.getToken(ctx)
	if err != nil {
		observability.Fail(span, "upstream_auth_failed")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		observability.Fail(span, "upstream_request_invalid")
		return nil, fmt.Errorf("create Blizzard API request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	response, err := hc.client.Do(req)
	if err != nil {
		observability.Fail(span, "upstream_request_failed")
		return nil, fmt.Errorf("request Blizzard API: %w", err)
	}
	span.SetAttributes(attribute.Int("http.response.status_code", response.StatusCode))
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		response.Body.Close()
		observability.Fail(span, "upstream_non_success")
		return nil, fmt.Errorf("Blizzard API returned HTTP %d", response.StatusCode)
	}
	return response, nil
}

func (hc *HsClient) getToken(ctx context.Context) (string, error) {
	hc.tokenMu.Lock()
	defer hc.tokenMu.Unlock()

	if hc.token.expireDate.After(time.Now().Add(30 * time.Second)) {
		return hc.token.accessToken, nil
	}
	newToken, err := fetchToken(ctx, hc.client)
	if err != nil {
		return "", err
	}
	hc.token = newToken
	return hc.token.accessToken, nil
}

func fetchToken(ctx context.Context, client *http.Client) (token, error) {
	ctx, span := observability.Start(ctx, "blizzard.oauth.token", attribute.String("http.request.method", http.MethodPost))
	defer span.End()

	clientID, clientSecret, err := getClientCredentials()
	if err != nil {
		observability.Fail(span, "upstream_auth_configuration_failed")
		return token{}, err
	}

	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://oauth.battle.net/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return token{}, fmt.Errorf("create Blizzard token request: %w", err)
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	response, err := client.Do(req)
	if err != nil {
		observability.Fail(span, "upstream_request_failed")
		return token{}, fmt.Errorf("request Blizzard access token: %w", err)
	}
	defer response.Body.Close()
	span.SetAttributes(attribute.Int("http.response.status_code", response.StatusCode))
	if response.StatusCode != http.StatusOK {
		observability.Fail(span, "upstream_non_success")
		return token{}, fmt.Errorf("Blizzard token endpoint returned HTTP %d", response.StatusCode)
	}

	var tokenDTO dto.Token
	if err := json.NewDecoder(response.Body).Decode(&tokenDTO); err != nil {
		observability.Fail(span, "upstream_decode_failed")
		return token{}, fmt.Errorf("decode Blizzard access token response: %w", err)
	}
	return token{
		accessToken: tokenDTO.AccessToken,
		expireDate:  time.Now().Add(time.Duration(tokenDTO.ExpiresIn) * time.Second),
	}, nil
}

func wait(ctx context.Context) error {
	timer := time.NewTimer(200 * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// getClientCredentials returns client ID, client secret, and an error.
func getClientCredentials() (string, string, error) {
	clientID, present := os.LookupEnv("CLIENT_ID")
	if !present {
		return "", "", fmt.Errorf("CLIENT_ID is not present in the environment")
	}
	clientSecret, present := os.LookupEnv("CLIENT_SECRET")
	if !present {
		return "", "", fmt.Errorf("CLIENT_SECRET is not present in the environment")
	}
	return clientID, clientSecret, nil
}
