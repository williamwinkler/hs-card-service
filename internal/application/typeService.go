package application

import "github.com/williamwinkler/hs-card-service/internal/application/interfaces"

type TypeService struct {
	typeRepo  interfaces.TypeRepository
	hsClient interfaces.HsClient
}

func NewTypeService(typeRepo interfaces.TypeRepository, hsClient interfaces.HsClient) *TypeService {
	return &TypeService{
		typeRepo:  typeRepo,
		hsClient: hsClient,
	}
}

func (c *TypeService) Update() error {
	types, err := c.hsClient.GetTypes()
	if err != nil {
		return err
	}

	err = c.typeRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.typeRepo.InsertMany(types)
}
