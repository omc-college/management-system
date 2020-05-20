package service

import (
	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

type RolesService struct {
	RolesRepository *postgres.RolesRepository
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

	return nil
}

func (service *RolesService) UpdateRole(role models.Role, id int) (err error) {
	err = service.RolesRepository.UpdateRole(role, id)
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

	return nil
}

func (service *RolesService) GetRoleTmpl() (roleTmpl models.RoleTmpl, err error) {
	roleTmpl, err = service.RolesRepository.GetRoleTmpl()
	if err != nil {
		return models.RoleTmpl{}, err
	}

	return roleTmpl, nil
}
