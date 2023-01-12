package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients/dto"
)

const PAGE_SIZE int = 250

type HsClient struct {
	token string
}

type Creds struct {
	clientId     string
	clientSecret string
}

func NewHsClient() (*HsClient, error) {
	token, err := getToken()
	if err != nil {
		return &HsClient{}, err
	}

	hsClient := HsClient{
		token: token,
	}
	return &hsClient, nil
}

func (hc *HsClient) GetAllCards() ([]domain.Card, error) {
	var cardsList []domain.Card
	page := 1
	for {
		cards, err := hc.GetCardsWithPagination(page, PAGE_SIZE)
		if err != nil {
			return []domain.Card{}, err
		}
		if len(cards) == 0 {
			break
		}

		for _, card := range cards {
			cardsList = append(cardsList, card)
		}
		page += 1
	}

	return cardsList, nil
}

func (hc *HsClient) GetCardsWithPagination(page int, pageSize int) ([]domain.Card, error) {
	apiUrl := "https://eu.api.blizzard.com/hearthstone/cards?locale=en_US"

	queryParams := url.Values{}
	queryParams.Set("page", strconv.Itoa(page))
	queryParams.Set("pageSize", strconv.Itoa(pageSize))

	queryString := queryParams.Encode()

	url := fmt.Sprintf("%s?%s", apiUrl, queryString)
	//fmt.Println(url)

	response, err := hc.executeCreateGetRequest(url)
	if err != nil {
		return []domain.Card{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Card{}, err
	}

	var cardsDto dto.CardsDto
	err = json.Unmarshal(body, &cardsDto)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to decode response from /cards. Body was %+v", body)
	}

	cards := dto.MapToCards(cardsDto)

	return cards, nil
}

func (hc *HsClient) GetSets() ([]domain.Set, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/sets?locale=en_US"

	response, err := hc.executeCreateGetRequest(url)
	if err != nil {
		return []domain.Set{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Set{}, err
	}

	var setsDto dto.SetsDto
	err = json.Unmarshal(body, &setsDto)
	if err != nil {
		return []domain.Set{}, err
	}

	sets := dto.MapToSets(setsDto)

	return sets, nil
}

func (hc *HsClient) executeCreateGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to create new GET-request for /cards")
	}
	req.Header.Set("Authorization", "Bearer "+hc.token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to GET url: %s", url)
	}

	return resp, nil
}

func getToken() (string, error) {
	urlAdress := "https://oauth.battle.net/token"
	creds, err := getClientCredentials()
	if err != nil {
		return "", err
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", urlAdress, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(creds.clientId, creds.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenDto dto.Token
	err = json.Unmarshal(body, &tokenDto)
	if err != nil {
		return "", err
	}

	log.Println("Received new token")
	return tokenDto.AccessToken, nil
}

func getClientCredentials() (Creds, error) {
	clientId, present := os.LookupEnv("client_id")
	if !present {
		return Creds{}, fmt.Errorf("client_id is not present in .env")
	}
	clientSecret, present := os.LookupEnv("client_secret")
	if !present {
		return Creds{}, fmt.Errorf("client_secret is not present in .env")
	}
	return Creds{
		clientId:     clientId,
		clientSecret: clientSecret}, nil
}
