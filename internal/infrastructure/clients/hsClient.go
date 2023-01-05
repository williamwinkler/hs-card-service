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
	log.Println("Getting all cards...")

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

	log.Printf("Finished getting all cards. In total: %d", len(cardsList))
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
	//fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to create new GET-request for /cards")
	}
	req.Header.Set("Authorization", "Bearer "+hc.token)
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

	var cardsDto dto.CardsDto
	err = json.Unmarshal(body, &cardsDto)
	if err != nil {
		return []domain.Card{}, fmt.Errorf("failed to decode response from /cards. Body was %+v", body)
	}

	cards := mapToCards(cardsDto)

	return cards, nil
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

func mapToCards(cardResp dto.CardsDto) []domain.Card {
	var cards []domain.Card
	for _, p := range cardResp.Cards {
		if p.CopyOfCardID != 0 {
			continue // if CopyOfCard is not 0, it's an unwanted outdated version
		}

		var c domain.Card
		c.ID = p.ID
		c.Collectible = p.Collectible
		c.Slug = p.Slug
		c.ClassID = p.ClassID
		c.MultiClassIds = p.MultiClassIds
		c.CardTypeID = p.CardTypeID
		c.CardSetID = p.CardSetID
		c.RarityID = p.RarityID
		c.ArtistName = p.ArtistName
		c.Health = p.Health
		c.Attack = p.Attack
		c.ManaCost = p.ManaCost
		c.Name = p.Name
		c.Text = p.Text
		c.Image = p.Image
		c.ImageGold = p.ImageGold
		c.FlavorText = p.FlavorText
		c.CropImage = p.CropImage
		c.ParentID = p.ParentID
		c.KeywordIds = p.KeywordIds
		c.Duels = p.Duels

		cards = append(cards, c)
	}
	return cards
}
