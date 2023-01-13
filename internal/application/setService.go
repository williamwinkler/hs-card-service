package application

import "github.com/williamwinkler/hs-card-service/internal/application/interfaces"

type SetService struct {
	setRepo  interfaces.SetRepository
	hsClient interfaces.HsClient
}

func NewSetService(setRepo interfaces.SetRepository, hsClient interfaces.HsClient) *SetService {
	return &SetService{
		setRepo:  setRepo,
		hsClient: hsClient,
	}
}

func (c *SetService) Update() error {
	sets, err := c.hsClient.GetSets()
	if err != nil {
		return err
	}

	err = c.setRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.setRepo.InsertMany(sets)
}
