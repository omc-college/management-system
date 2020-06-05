package service

import (
<<<<<<< HEAD
	"encoding/json"

	"github.com/omc-college/management-system/pkg/pubsub"
=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

type RolesService struct {
	RolesRepository *postgres.RolesRepository
<<<<<<< HEAD
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
=======
}

func (service *RolesService) GetAllRoles() (roles []models.Role, err error) {
	roles, err = postgres.GetAllRoles(service.RolesRepository)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}

func (service *RolesService) GetRole(id int) (role models.Role, err error) {
<<<<<<< HEAD
	role, err = service.RolesRepository.GetRole(id)
=======
	role, err = postgres.GetRole(service.RolesRepository, id)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (service *RolesService) CreateRole(role models.Role) (err error) {
<<<<<<< HEAD
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
=======
	err = postgres.CreateRole(service.RolesRepository, role)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) UpdateRole(role models.Role, id int) (err error) {
<<<<<<< HEAD
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
=======
	err = postgres.UpdateRole(service.RolesRepository, role, id)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) DeleteRole(id int) (err error) {
<<<<<<< HEAD
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
=======
	err = postgres.DeleteRole(service.RolesRepository, id)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) GetRoleTmpl() (roleTmpl models.RoleTmpl, err error) {
<<<<<<< HEAD
	roleTmpl, err = service.RolesRepository.GetRoleTmpl()
=======
	roleTmpl, err = postgres.GetRoleTmpl(service.RolesRepository)
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	if err != nil {
		return models.RoleTmpl{}, err
	}

	return roleTmpl, nil
}
