package service

import (
	"encoding/json"

	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

type RolesService struct {
	RolesRepository *postgres.RolesRepository
	PubSubRepository *pubsub.GroupClient
}

func NewRolesService(rolesRepository *postgres.RolesRepository, pubsubRepository *pubsub.GroupClient) *RolesService {
	return &RolesService{
		RolesRepository: rolesRepository,
		PubSubRepository: pubsubRepository,
	}
}

func (service *RolesService) GetAllRoles() (roles []models.Role, err error) {
	roles, err = service.RolesRepository.GetAllRoles()
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}

func (service *RolesService) GetRole(id int) (role models.Role, err error) {
	role, err = service.RolesRepository.GetRole(id)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (service *RolesService) CreateRole(role models.Role) (err error) {
	err = service.RolesRepository.CreateRole(role)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(role, models.RoleOperationCreate, models.RoleType)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(msg)
	if  err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, models.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) UpdateRole(role models.Role, id int) (err error) {
	err = service.RolesRepository.UpdateRole(role, id)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(role, models.RoleOperationUpdate, models.RoleType)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(msg)
	if  err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, models.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) DeleteRole(id int) (err error) {
	err = service.RolesRepository.DeleteRole(id)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(id, models.RoleOperationDelete, models.RoleType)

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, models.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) GetRoleTmpl() (roleTmpl models.RoleTmpl, err error) {
	roleTmpl, err = service.RolesRepository.GetRoleTmpl()
	if err != nil {
		return models.RoleTmpl{}, err
	}

	return roleTmpl, nil
}
