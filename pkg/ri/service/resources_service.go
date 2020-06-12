package service

import (
	"github.com/omc-college/management-system/pkg/ri/models"
	postgresql2 "github.com/omc-college/management-system/pkg/ri/repository/postgresql"
)

type ResourcesService struct {
	ResourcesRepository *postgresql2.ResourcesRepository
}

func NewResourcesService(resourcesRepository *postgresql2.ResourcesRepository) *ResourcesService {
	return &ResourcesService{
		ResourcesRepository: resourcesRepository,
	}
}

func (service *ResourcesService) GetAllResources() (resources []models.Resources, err error) {
	resources, err = service.ResourcesRepository.GetAllResources()
	if err != nil {
		return []models.Resources{}, err
	}

	return resources, nil
}

func (service *ResourcesService) GetResource(id int) (resource *models.Resources, err error) {
	resource, err = service.ResourcesRepository.GetResource(id)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (service *ResourcesService) InsertResource(resource models.Resources) (err error) {
	err = service.ResourcesRepository.InsertResource(resource)
	if err != nil {
		return err
	}

	return nil
}

func (service *ResourcesService) UpdateResource(resource models.Resources, id int) (err error) {
	err = service.ResourcesRepository.UpdateResource(resource, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *ResourcesService) DeleteResource(ResourceID int) (err error) {
	err = service.ResourcesRepository.DeleteResource(ResourceID)
	if err != nil {
		return err
	}

	return nil
}

func (service *ResourcesService) GetResourceByName(name string) (resource *models.Resources, err error) {
	resource, err = service.ResourcesRepository.GetResourceByName(name)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

