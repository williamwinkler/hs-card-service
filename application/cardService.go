package application

import "github.com/williamwinkler/hs-card-service/application/interfaces"

type CardService struct {
	HsClient *interfaces.HsClient
	CardRepo *interfaces.CardRepository
}

func NewCardService(hsclient *interfaces.HsClient, cardRepo interfaces.CardRepository) *CardService {
	return &CardService{
		HsClient: hsclient,
		CardRepo: &cardRepo,
	}
}
