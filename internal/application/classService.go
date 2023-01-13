package application

import "github.com/williamwinkler/hs-card-service/internal/application/interfaces"

type ClassService struct {
	classRepo interfaces.ClassRepository
	hsClient  interfaces.HsClient
}

func NewClassService(ClassRepo interfaces.ClassRepository, hsClient interfaces.HsClient) *ClassService {
	return &ClassService{
		classRepo: ClassRepo,
		hsClient:  hsClient,
	}
}

func (c *ClassService) Update() error {
	classes, err := c.hsClient.GetClasses()
	if err != nil {
		return err
	}

	err = c.classRepo.DeleteAll()
	if err != nil {
		return err
	}

	return c.classRepo.InsertMany(classes)
}
