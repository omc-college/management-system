package service

import (
	"context"
	"encoding/json"

	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

type RolesService struct {
	RolesRepository  *postgres.RolesRepository
	PubSubRepository *pubsub.Client
}

func NewRolesService(rolesRepository *postgres.RolesRepository, pubsubRepository *pubsub.Client) *RolesService {
	return &RolesService{
		RolesRepository:  rolesRepository,
		PubSubRepository: pubsubRepository,
	}
}

func (service *RolesService) GetAllRoles(ctx context.Context) ([]rbac.Role, error) {
	roles, err := service.RolesRepository.GetAllRoles(ctx)
	if err != nil {
		return []rbac.Role{}, err
	}

	return roles, nil
}

func (service *RolesService) GetRole(ctx context.Context, id int) (rbac.Role, error) {
	role, err := service.RolesRepository.GetRole(ctx, id)
	if err != nil {
		return rbac.Role{}, err
	}

	return role, nil
}

func (service *RolesService) CreateRole(ctx context.Context, role rbac.Role) error {
	err := service.RolesRepository.CreateRole(ctx, &role)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(role, rbac.RoleOperationCreate, rbac.RoleType)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, rbac.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) UpdateRole(ctx context.Context, role rbac.Role, id int) error {
	err := service.RolesRepository.UpdateRole(ctx, role, id)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(role, rbac.RoleOperationUpdate, rbac.RoleType)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, rbac.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) DeleteRole(ctx context.Context, id int) error {
	err := service.RolesRepository.DeleteRole(ctx, id)
	if err != nil {
		return err
	}

	msg, err := pubsub.NewEnvelope(id, rbac.RoleOperationDelete, rbac.RoleType)

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = service.PubSubRepository.Publish(bytes, rbac.RolesTopicName)
	if err != nil {
		return err
	}

	return nil
}

func (service *RolesService) GetRoleTmpl(ctx context.Context) (rbac.RoleTmpl, error) {
	roleTmpl, err := service.RolesRepository.GetRoleTmpl(ctx)
	if err != nil {
		return rbac.RoleTmpl{}, err
	}

	return roleTmpl, nil
}
