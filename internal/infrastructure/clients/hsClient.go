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
	"time"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/clients/dto"
)

const PAGE_SIZE int = 250

type HsClient struct {
	token token
}

type token struct {
	accessToken string
	expireDate  time.Time
}

func NewHsClient() (*HsClient, error) {
	token, err := fetchToken()
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
		// log.Printf("Getting cards with page %d limit %d", page, PAGE_SIZE)
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
	apiUrl := "https://eu.api.blizzard.com/hearthstone/cards"

	queryParams := url.Values{}
	queryParams.Set("locale", "en_US")
	queryParams.Set("page", strconv.Itoa(page))
	queryParams.Set("pageSize", strconv.Itoa(pageSize))

	queryString := queryParams.Encode()

	url := fmt.Sprintf("%s?%s", apiUrl, queryString)

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	response, err := hc.executeGetRequest(url)
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
		return []domain.Card{}, fmt.Errorf("failed to decode response from /cards: %v", err)
	}

	cards := dto.MapToCards(cardsDto)

	return cards, nil
}

func (hc *HsClient) GetSets() ([]domain.Set, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/sets?locale=en_US"

	response, err := hc.executeGetRequest(url)
	if err != nil {
		return []domain.Set{}, err
	}
	defer response.Body.Close()

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

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

func (hc *HsClient) GetClasses() ([]domain.Class, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/classes?locale=en_US"

	response, err := hc.executeGetRequest(url)
	if err != nil {
		return []domain.Class{}, err
	}
	defer response.Body.Close()

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Class{}, err
	}

	var classesDto dto.ClassesDto
	err = json.Unmarshal(body, &classesDto)
	if err != nil {
		return []domain.Class{}, err
	}

	classes := dto.MapToClasses(classesDto)

	return classes, nil
}

func (hc *HsClient) GetRarities() ([]domain.Rarity, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/rarities?locale=en_US"

	response, err := hc.executeGetRequest(url)
	if err != nil {
		return []domain.Rarity{}, err
	}
	defer response.Body.Close()

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Rarity{}, err
	}

	var raritiesDto dto.RaritiesDto
	err = json.Unmarshal(body, &raritiesDto)
	if err != nil {
		return []domain.Rarity{}, err
	}

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	rarities := dto.MapToRariteis(raritiesDto)

	return rarities, nil
}

func (hc *HsClient) GetTypes() ([]domain.Type, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/types?locale=en_US"

	response, err := hc.executeGetRequest(url)
	if err != nil {
		return []domain.Type{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Type{}, err
	}

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	var typesDto dto.TypesDto
	err = json.Unmarshal(body, &typesDto)
	if err != nil {
		return []domain.Type{}, err
	}

	types := dto.MapToTypes(typesDto)

	return types, nil
}

func (hc *HsClient) GetKeywords() ([]domain.Keyword, error) {
	url := "https://us.api.blizzard.com/hearthstone/metadata/keywords?locale=en_US"

	response, err := hc.executeGetRequest(url)
	if err != nil {
		return []domain.Keyword{}, err
	}
	defer response.Body.Close()

	log.Println(url)
	time.Sleep(200 * time.Millisecond)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []domain.Keyword{}, err
	}

	var keywordsDto dto.KeywordsDto
	err = json.Unmarshal(body, &keywordsDto)
	if err != nil {
		return []domain.Keyword{}, err
	}

	keywords := dto.MapToKeywords(keywordsDto)

	return keywords, nil
}

func (hc *HsClient) executeGetRequest(url string) (*http.Response, error) {
	token, err := hc.getToken()
	if err != nil {
		return &http.Response{}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to create new GET-request for /cards")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to GET url: %s", url)
	}

	return resp, nil
}

func (hc *HsClient) getToken() (string, error) {
	//hc.token.expireDate = hc.token.expireDate.AddDate(0, 0, -1)
	if hc.token.expireDate.Before(time.Now()) {
		newToken, err := fetchToken()
		if err != nil {
			return "", err
		}
		hc.token = newToken
	}

	return hc.token.accessToken, nil
}

func fetchToken() (token, error) {
	urlAdress := "https://oauth.battle.net/token"
	clientId, clientSecret, err := getClientCredentials()
	if err != nil {
		return token{}, err
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", urlAdress, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return token{}, err
	}
	req.SetBasicAuth(clientId, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return token{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token{}, err
	}

	var tokenDto dto.Token
	err = json.Unmarshal(body, &tokenDto)
	if err != nil {
		return token{}, err
	}

	expireDate := time.Now().Add(time.Duration(tokenDto.ExpiresIn) * time.Second)
	token := token{
		accessToken: tokenDto.AccessToken,
		expireDate:  expireDate,
	}

	log.Println("Received new token")
	return token, nil
}

// returns clientId, clientSecret, error
func getClientCredentials() (string, string, error) {
	clientId, present := os.LookupEnv("client_id")
	if !present {
		return "", "", fmt.Errorf("client_id is not present in .env")
	}
	clientSecret, present := os.LookupEnv("client_secret")
	if !present {
		return "", "", fmt.Errorf("client_secret is not present in .env")
	}
	return clientId, clientSecret, nil
}
