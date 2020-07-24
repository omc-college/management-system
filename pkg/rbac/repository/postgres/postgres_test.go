package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"

	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

var dbURL = fmt.Sprintf(
	"%s://postgres:%s@%s:%s/%s?sslmode=%s",
	"postgres", "superuser", "localhost", "5432", "roles", "disable",
)

func TestGetAllRoles_AllExistingRolesShouldBeReturned(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := [][]rbac.Role{
		{
			{
				ID:   2,
				Name: "admin",
				Entries: []rbac.FeatureEntry{
					{
						ID:          1,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
			{
				ID:   3,
				Name: "user",
				Entries: []rbac.FeatureEntry{
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			retrievedRoles, err := repository.GetAllRoles(context.Background())
			if err != nil {
				t.Fatalf("cannot get roles: %s", err.Error())
			}

			if !reflect.DeepEqual(testCase, retrievedRoles) {
				t.Fatalf("expected: %v, got: %v", testCase, retrievedRoles)
			}
		})
	}
}
func TestGetAllRoles_EmptyRolesSliceShouldBeReturnedIfNoRolesExist(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	repository := postgres.NewRolesRepository(db)

	retrievedRoles, err := repository.GetAllRoles(context.Background())
	if err != nil {
		t.Fatalf("cannot get roles: %s", err.Error())
	}

	if len(retrievedRoles) != 0 {
		t.Fatalf("expected: %v, got: %v", []rbac.Role{}, retrievedRoles)
	}
}

func TestGetRole_ExistingRoleShouldBeReturnedByID(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := [][]rbac.Role{
		{
			{
				ID:   2,
				Name: "admin",
				Entries: []rbac.FeatureEntry{
					{
						ID:          1,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
			{
				ID:   3,
				Name: "user",
				Entries: []rbac.FeatureEntry{
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			for _, testRole := range testCase {
				retrievedRole, err := repository.GetRole(context.Background(), testRole.ID)
				if err != nil {
					t.Fatalf("cannot get role: %s", err.Error())
				}

				if !reflect.DeepEqual(testRole, retrievedRole) {
					t.Fatalf("expected: %v, got: %v", testCase, testRole)
				}
			}
		})
	}
}
func TestGetRole_ErrorShouldBeReturnedWhenGettingNonexistentRole(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		GettingRoleID   int
		PredefinedRoles []rbac.Role
	}{
		{
			GettingRoleID: 404,
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			_, err := repository.GetRole(context.Background(), testCase.GettingRoleID)
			if err != nil {
				if !errors.Is(err, postgres.ErrNoRows) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNoRows.Error(), err.Error())
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNoRows.Error(), err)
			}
		})
	}
}

func TestDeleteRole_ExistingRoleShouldBeDeletedByID(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		DeletingRoleID  int
		PredefinedRoles []rbac.Role
	}{
		{
			DeletingRoleID: 2,
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.DeleteRole(context.Background(), testCase.DeletingRoleID)
			if err != nil {
				t.Fatalf("cannot delete role: %s", err.Error())
			}

			checkIsRoleDeleted(t, db, testCase.DeletingRoleID)
		})
	}
}
func TestDeleteRole_ErrorShouldBeReturnedWhenDeletingNonexistentRole(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		DeletingRoleID  int
		PredefinedRoles []rbac.Role
	}{
		{
			DeletingRoleID: 404,
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			err := repository.DeleteRole(context.Background(), testCase.DeletingRoleID)
			if err != nil {
				if !errors.Is(err, postgres.ErrNoRows) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNoRows.Error(), err.Error())
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNoRows.Error(), err)
			}
		})
	}
}

func TestCreateRole_NonexistentRoleShouldBeCreated(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.Role{
		{
			ID:   1,
			Name: "superuser",
			Entries: []rbac.FeatureEntry{
				{
					ID:          1,
					Name:        "rolesManagement",
					Description: "description of rolesManagement",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          2,
					Name:        "rolesReading",
					Description: "description of rolesReading",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.CreateRole(context.Background(), &testCase)
			if err != nil {
				t.Fatalf("cannot create role: %s", err.Error())
			}

			checkIsValidRoleCreated(t, db, testCase)
		})
	}
}
func TestCreateRole_ErrorShouldBeReturnedWhenCreatingNotUniqueRole(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.Role{
		{
			ID:   2,
			Name: "admin",
			Entries: []rbac.FeatureEntry{
				{
					ID:          1,
					Name:        "rolesManagement",
					Description: "description of rolesManagement",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          2,
					Name:        "rolesReading",
					Description: "description of rolesReading",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.CreateRole(context.Background(), &testCase)
			if err != nil {
				if !errors.Is(err, postgres.ErrNotUniqueRole) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNotUniqueRole, err)
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNotUniqueRole, err)
			}
		})
	}
}
func TestCreateRole_ErrorShouldBeReturnedWhenCreatingRoleWithNonexistentFeature(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.Role{
		{
			ID:   1,
			Name: "ceo",
			Entries: []rbac.FeatureEntry{
				{
					ID:          90,
					Name:        "rolesManagement",
					Description: "description of rolesManagement",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          91,
					Name:        "rolesReading",
					Description: "description of rolesReading",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.CreateRole(context.Background(), &testCase)
			if err != nil {
				if !errors.Is(err, postgres.ErrNonexistentFeature) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNonexistentFeature, err)
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNonexistentFeature, err)
			}
		})
	}
}

func TestUpdateRole_ExistingRoleShouldBeUpdatedByID(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		NewRoleID       int
		NewRole         rbac.Role
		PredefinedRoles []rbac.Role
	}{
		{
			NewRoleID: 2,
			NewRole: rbac.Role{
				ID:   2,
				Name: "superuser",
				Entries: []rbac.FeatureEntry{
					{
						ID:          1,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.UpdateRole(context.Background(), testCase.NewRole, testCase.NewRoleID)
			if err != nil {
				t.Fatalf("cannot update role: %s", err.Error())
			}

			checkIsValidRoleCreated(t, db, testCase.NewRole)
		})
	}
}
func TestUpdateRole_ErrorShouldBeReturnedWhenUpdatingNonexistentRole(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		NewRoleID       int
		NewRole         rbac.Role
		PredefinedRoles []rbac.Role
	}{
		{
			NewRoleID: 404,
			NewRole: rbac.Role{
				ID:   404,
				Name: "superuser",
				Entries: []rbac.FeatureEntry{
					{
						ID:          1,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			err := repository.UpdateRole(context.Background(), testCase.NewRole, testCase.NewRoleID)
			if err != nil {
				if !errors.Is(err, postgres.ErrNoRows) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNoRows.Error(), err.Error())
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNoRows.Error(), err)
			}
		})
	}
}
func TestUpdateRole_ErrorShouldBeReturnedWhenUpdatingToNotUniqueRole(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		NewRoleID       int
		NewRole         rbac.Role
		PredefinedRoles []rbac.Role
	}{
		{
			NewRoleID: 3,
			NewRole: rbac.Role{
				ID:   3,
				Name: "admin",
				Entries: []rbac.FeatureEntry{
					{
						ID:          1,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          2,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
			PredefinedRoles: []rbac.Role{
				{
					ID:   2,
					Name: "admin",
					Entries: []rbac.FeatureEntry{
						{
							ID:          1,
							Name:        "rolesManagement",
							Description: "description of rolesManagement",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
								{
									ID:     2,
									Name:   "CreateRole",
									Path:   "/roles",
									Method: "POST",
								},
							},
						},
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
				{
					ID:   3,
					Name: "user",
					Entries: []rbac.FeatureEntry{
						{
							ID:          2,
							Name:        "rolesReading",
							Description: "description of rolesReading",
							Endpoints: []rbac.Endpoint{
								{
									ID:     1,
									Name:   "GetAllRoles",
									Path:   "/roles",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.UpdateRole(context.Background(), testCase.NewRole, testCase.NewRoleID)
			if err != nil {
				if !errors.Is(err, postgres.ErrNotUniqueRole) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNotUniqueRole.Error(), err.Error())
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNotUniqueRole.Error(), err)
			}
		})
	}
}
func TestUpdateRole_ErrorShouldBeReturnedWhenUpdatingToRoleWithNonexistentFeature(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []struct {
		NewRoleID       int
		NewRole         rbac.Role
	}{
		{
			NewRoleID: 3,
			NewRole: rbac.Role{
				ID:   3,
				Name: "ceo",
				Entries: []rbac.FeatureEntry{
					{
						ID:          90,
						Name:        "rolesManagement",
						Description: "description of rolesManagement",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
							{
								ID:     2,
								Name:   "CreateRole",
								Path:   "/roles",
								Method: "POST",
							},
						},
					},
					{
						ID:          91,
						Name:        "rolesReading",
						Description: "description of rolesReading",
						Endpoints: []rbac.Endpoint{
							{
								ID:     1,
								Name:   "GetAllRoles",
								Path:   "/roles",
								Method: "GET",
							},
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.UpdateRole(context.Background(), testCase.NewRole, testCase.NewRoleID)
			if err != nil {
				if !errors.Is(err, postgres.ErrNonexistentFeature) {
					t.Fatalf("expected: %s, got: %s", postgres.ErrNonexistentFeature, err)
				}
			}

			if err == nil {
				t.Fatalf("expected: %s, got: %v", postgres.ErrNonexistentFeature, err)
			}
		})
	}
}

func TestGetRoleTmpl_RoleTmplWithAllExistingFeaturesAndEndpointsShouldBeReturned(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.RoleTmpl{
		{
			Entries: []rbac.FeatureEntry{
				{
					ID:          1,
					Name:        "rolesManagement",
					Description: "description of rolesManagement",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          2,
					Name:        "rolesReading",
					Description: "description of rolesReading",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)

			retrievedRoleTmpl, err := repository.GetRoleTmpl(context.Background())
			if err != nil {
				t.Fatalf("cannot get role template: %s", err.Error())
			}

			if !reflect.DeepEqual(testCase, retrievedRoleTmpl) {
				t.Fatalf("expected: %v, got: %v", testCase, retrievedRoleTmpl)
			}
		})
	}
}

func TestCreateRoleTmpl_NonexistentFeaturesAndEndpointsShouldBeCreated(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.RoleTmpl{
		{
			Entries: []rbac.FeatureEntry{
				{
					ID:          1,
					Name:        "rolesManagement",
					Description: "description of rolesManagement",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          2,
					Name:        "rolesReading",
					Description: "description of rolesReading",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			repository := postgres.NewRolesRepository(db)
			err := repository.CreateRoleTmpl(context.Background(), testCase)
			if err != nil {
				t.Fatalf("cannot create role template: %s", err.Error())
			}

			checkIsValidRoleTmplCreated(t, db, testCase)
		})
	}
}
func TestCreateRoleTmpl_ExistingFeaturesAndEndpointsShouldBeUpdated(t *testing.T) {
	db := getDB(t, dbURL)

	resetSchema(t, dbURL)

	testCases := []rbac.RoleTmpl{
		{
			Entries: []rbac.FeatureEntry{
				{
					ID:          1,
					Name:        "rolesManagement",
					Description: "description of rolesManagement updated",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
						{
							ID:     2,
							Name:   "CreateRole",
							Path:   "/roles",
							Method: "POST",
						},
					},
				},
				{
					ID:          2,
					Name:        "rolesReading",
					Description: "description of rolesReading updated",
					Endpoints: []rbac.Endpoint{
						{
							ID:     1,
							Name:   "GetAllRoles",
							Path:   "/roles",
							Method: "GET",
						},
					},
				},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCaseIndex), func(t *testing.T) {
			insertTestData(t, testCaseIndex, dbURL)

			repository := postgres.NewRolesRepository(db)
			err := repository.CreateRoleTmpl(context.Background(), testCase)
			if err != nil {
				t.Fatalf("cannot create role template: %s", err.Error())
			}

			checkIsValidRoleTmplCreated(t, db, testCase)
		})
	}
}

func getDB(t *testing.T, dbURL string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		t.Fatalf("cannot open and connect to DB: %s", err.Error())
	}

	return db
}

func resetSchema(t *testing.T, dbURL string) error {
	migration, err := migrate.New(
		"file://./migrations",
		dbURL,
	)
	if err != nil {
		t.Fatalf("cannot read schema migrations and connect to DB: %s", err.Error())
	}

	err = migration.Down()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			t.Fatalf("cannot down schema migrations: %s", err.Error())
		}
	}

	err = migration.Up()
	if err != nil {
		t.Fatalf("cannot up schema migrations. error: %s", err.Error())
	}

	return nil
}

func insertTestData(t *testing.T, index int, dbURL string) {
	t.Helper()

	sourceURL := fmt.Sprintf("file://./test_fixtures/migrations/%v", index)

	migration, err := migrate.New(
		sourceURL,
		dbURL,
	)
	if err != nil {
		t.Fatalf("cannot read test data inserting migrations and connect to DB: %s", err.Error())
	}

	lastVersion, _, err := migration.Version()
	if err != nil {
		t.Fatalf("cannot get test data inserting migration version: %s", err.Error())
	}

	migration.Force(1)
	if err != nil {
		t.Fatalf("cannot force test data inserting migration version to 0: %s", err.Error())
	}
	defer migration.Force(int(lastVersion))

	err = migration.Down()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			t.Fatalf("cannot down test data inserting migrations: %s", err.Error())
		}
	}

	err = migration.Up()
	if err != nil {
		t.Fatalf("cannot up test data inserting migrations: %s", err.Error())
	}
}

func checkIsRoleDeleted(t *testing.T, db *sqlx.DB, id int) {
	t.Helper()

	query := "SELECT id FROM roles WHERE id = $1"
	err := db.QueryRowxContext(context.Background(), query, id).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			t.Fatalf("role is not deleted. error: %s", err.Error())
		}
	}
}

func checkIsValidRoleCreated(t *testing.T, db *sqlx.DB, role rbac.Role) {
	t.Helper()

	query := `SELECT roles.id, roles.name, roles_to_features.feature_id
					  FROM roles LEFT JOIN roles_to_features
					  ON roles.id = roles_to_features.role_id
					  WHERE roles.id = $1`

	rows, err := db.QueryxContext(context.Background(), query, role.ID)
	if err != nil {
		t.Fatalf("cannot query DB: %s", err.Error())
	}

	var createdRole rbac.Role

	for rows.Next() {
		var feature rbac.FeatureEntry

		err := rows.Scan(&createdRole.ID, &createdRole.Name, &feature.ID)
		if err != nil {
			t.Fatalf("cannot scan query result: %s", err.Error())
		}

		createdRole.Entries = append(createdRole.Entries, feature)
	}

	if createdRole.ID != role.ID {
		t.Fatalf("role's id is wrong. expected: %v, got: %v",
			role.ID, createdRole.ID)
	}

	if createdRole.Name != role.Name {
		t.Fatalf("role's name is wrong. expected: %v, got: %v",
			role.ID, createdRole.ID)
	}

	for featureIndex, feature := range createdRole.Entries {
		if feature.ID != role.Entries[featureIndex].ID {
			t.Fatalf("role's features are wrong. expected: %v, got: %v",
				role.Entries, createdRole.Entries)
		}
	}
}

func checkIsValidRoleTmplCreated(t *testing.T, db *sqlx.DB, roleTmpl rbac.RoleTmpl) {
	t.Helper()

	query := `SELECT features.id, features.name, features.description, 
					  endpoints.id, endpoints.name, endpoints.path, endpoints.method
					  FROM features LEFT JOIN features_to_endpoints
					  ON features.id = features_to_endpoints.feature_id
					  LEFT JOIN endpoints
					  ON features_to_endpoints.endpoint_id = endpoints.id`

	rows, err := db.QueryxContext(context.Background(), query)
	if err != nil {
		t.Fatalf("cannot query DB: %s", err.Error())
	}

	var retrievedRoleTmpl rbac.RoleTmpl

	for rows.Next() {
		var feature rbac.FeatureEntry
		var endpoint rbac.Endpoint

		err := rows.Scan(&feature.ID, &feature.Name, &feature.Description,
			&endpoint.ID, &endpoint.Name, &endpoint.Path, &endpoint.Method)
		if err != nil {
			t.Fatalf("cannot scan query result: %s", err.Error())
		}

		var isFeatureExisting bool

		for existingFeatureIndex, existingFeature := range retrievedRoleTmpl.Entries {
			if feature.ID == existingFeature.ID {
				isFeatureExisting = true

				retrievedRoleTmpl.Entries[existingFeatureIndex].Endpoints = append(
					retrievedRoleTmpl.Entries[existingFeatureIndex].Endpoints,
					endpoint,
				)
			}
		}

		if !isFeatureExisting {
			feature.Endpoints = append(feature.Endpoints, endpoint)
			retrievedRoleTmpl.Entries = append(retrievedRoleTmpl.Entries, feature)
		}
	}

	if !reflect.DeepEqual(roleTmpl, retrievedRoleTmpl) {
		t.Fatalf("expected: %v, got: %v", roleTmpl, retrievedRoleTmpl)
	}
}
