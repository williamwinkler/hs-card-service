package application

import "github.com/williamwinkler/hs-card-service/internal/application/interfaces"

type KeywordService struct {
	keywordRepo interfaces.KeywordRepository
	hsClient    interfaces.HsClient
}

func NewKeywordService(keywordRepo interfaces.KeywordRepository, hsClient interfaces.HsClient) *KeywordService {
	return &KeywordService{
		keywordRepo: keywordRepo,
		hsClient:    hsClient,
	}
}

func (c *KeywordService) Update() error {
	keywords, err := c.hsClient.GetKeywords()
	if err != nil {
		return err
	}

	err = c.keywordRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.keywordRepo.InsertMany(keywords)
}
