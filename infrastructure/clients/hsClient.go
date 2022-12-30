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

	"github.com/williamwinkler/hs-card-service/domain"
	"github.com/williamwinkler/hs-card-service/infrastructure/clients/dto"
)

type HsClient struct {
	Token string
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
		Token: token,
	}
	return &hsClient, nil
}

func (hc *HsClient) GetCardsWithPagination(page int, pageSize int) ([]domain.Card, error) {
	apiUrl := "https://eu.api.blizzard.com/hearthstone/cards"

	queryParams := url.Values{}
	queryParams.Set("locale", "en_US")
	queryParams.Set("page", string(page))
	queryParams.Set("pageSize", string(pageSize))

	queryString := queryParams.Encode()

	url := fmt.Sprintf("%s?%s", apiUrl, queryString)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to create new GET-request for /cards")
	}
	req.Header.Set("Authorization", "Bearer "+hc.Token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to do GET-request for /cards")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []domain.Card{}, err
	}

	fmt.Println(body)

	var cardsReponse dto.GetCardsResponse
	err = json.Unmarshal(body, &cardsReponse)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to decode response from /cards. Body was %+v", body)
	}

	//fmt.Println(cardsReponse)
	return []domain.Card{}, nil
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
