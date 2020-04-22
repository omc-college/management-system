package postgres

import (
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/omc-college/management-system/pkg/rbac/models"
)

type RolesRepository struct {
	DB *sql.DB
}

func NewRolesRepository(dsn string) (*RolesRepository, error) {
	db, err := sql.Open("pgx", dsn)

	return &RolesRepository{
		DB: db,
	}, err
}

func GetAllRoles(repository *RolesRepository) ([]models.Role, error) {
	query := `SELECT roles.id, roles.name, features.id, features.name, endpoints.id, endpoints.path, endpoints.method
			  FROM roles LEFT JOIN roles_to_features
			  ON roles.id = roles_to_features.role_id
			  LEFT JOIN features
			  ON roles_to_features.feature_id = features.id
			  LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	var genericRoles []models.Role
	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint
	var isRoleExisting bool
	var isFeatureExisting bool

	tmpRoles := make(map[int]role)

	rows, err := repository.DB.Query(query)
	if err != nil {
		return []models.Role{}, QueryError{queryErrorMessage, err}
	}

	for rows.Next() {
		var tmpRole role
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isRoleExisting = false
		isFeatureExisting = false

		err := rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpEndpoint.ID, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return []models.Role{}, ScanError{scanErrorMessage, err}
		}

		_, isRoleExisting = tmpRoles[tmpRole.ID]
		if !isRoleExisting {
			tmpRole.Entries = make(map[int]featureEntry)
			tmpRoles[tmpRole.ID] = tmpRole
		}

		_, isFeatureExisting = tmpRoles[tmpRole.ID].Entries[tmpFeature.ID]
		if !isFeatureExisting {
			tmpFeature.Endpoints = make(map[int]endpoint)
			tmpRoles[tmpRole.ID].Entries[tmpFeature.ID] = tmpFeature
		}

		tmpRoles[tmpRole.ID].Entries[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
	}

	err = rows.Err()
	if err != nil {
		return []models.Role{}, ScanError{scanErrorMessage, err}
	}

	for _, tmpRole := range tmpRoles {
		genericFeatures = []models.FeatureEntry{}
		for _, tmpFeature := range tmpRole.Entries {
			genericEndpoints = []models.Endpoint{}
			for _, tmpEndpoint := range tmpFeature.Endpoints {
				genericEndpoint := models.Endpoint{ID: tmpEndpoint.ID, Path: tmpEndpoint.Path, Method: tmpEndpoint.Method}
				genericEndpoints = append(genericEndpoints, genericEndpoint)
			}
			genericFeature := models.FeatureEntry{ID: tmpFeature.ID, Name: tmpFeature.Name, Endpoints: genericEndpoints}
			genericFeatures = append(genericFeatures, genericFeature)
		}
		genericRole := models.Role{ID: tmpRole.ID, Name: tmpRole.Name, Entries: genericFeatures}
		genericRoles = append(genericRoles, genericRole)
	}

	return genericRoles, nil
}

func GetRole(repository *RolesRepository, id int) (models.Role, error) {
	query := `SELECT roles.id, roles.name, features.id, features.name, endpoints.id, endpoints.path, endpoints.method
			  FROM roles LEFT JOIN roles_to_features
			  ON roles.id = roles_to_features.role_id
			  LEFT JOIN features
			  ON roles_to_features.feature_id = features.id
			  LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id
			  WHERE roles.id = $1`

	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint
	var isFeatureExisting bool
	var tmpRole role
	tmpRole.Entries = make(map[int]featureEntry)

	rows, err := repository.DB.Query(query, id)
	if err != nil {
		return models.Role{}, QueryError{queryErrorMessage, err}
	}

	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isFeatureExisting = false

		err = rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpEndpoint.ID, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return models.Role{}, ScanError{scanErrorMessage, err}
		}

		_, isFeatureExisting = tmpRole.Entries[tmpFeature.ID]
		if !isFeatureExisting {
			tmpFeature.Endpoints = make(map[int]endpoint)
			tmpRole.Entries[tmpFeature.ID] = tmpFeature
		}

		tmpRole.Entries[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
	}

	// rows.Scan after db.Query doesn't return sql.ErrNoRows
	if tmpRole.ID == 0 {
		return models.Role{}, ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		return models.Role{}, ScanError{scanErrorMessage, err}
	}

	genericFeatures = []models.FeatureEntry{}
	for _, tmpFeature := range tmpRole.Entries {
		genericEndpoints = []models.Endpoint{}
		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := models.Endpoint{ID: tmpEndpoint.ID, Path: tmpEndpoint.Path, Method: tmpEndpoint.Method}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := models.FeatureEntry{ID: tmpFeature.ID, Name: tmpFeature.Name, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRole := models.Role{ID: tmpRole.ID, Name: tmpRole.Name, Entries: genericFeatures}

	return genericRole, nil
}

func GetRoleTemplate(repository *RolesRepository) (models.Role, error) {
	query := `SELECT features.id, features.name, endpoints.id, endpoints.path, endpoints.method
			  FROM features LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	rows, err := repository.DB.Query(query)
	if err != nil {
		return models.Role{}, QueryError{queryErrorMessage, err}
	}

	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint
	var isFeatureExisting bool
	var tmpRoleTemplate role
	var tmpFeatures = make(map[int]featureEntry)

	// Get all features and connect them to endpoints
	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isFeatureExisting = false

		err := rows.Scan(&tmpFeature.ID, &tmpFeature.Name, &tmpEndpoint.ID, &tmpEndpoint.Method, &tmpEndpoint.Path)
		if err != nil {
			return models.Role{}, ScanError{scanErrorMessage, err}
		}

		_, isFeatureExisting = tmpFeatures[tmpFeature.ID]
		if !isFeatureExisting {
			tmpFeature.Endpoints = make(map[int]endpoint)
			tmpFeatures[tmpFeature.ID] = tmpFeature
		}

		tmpFeatures[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
	}

	err = rows.Err()
	if err != nil {
		return models.Role{}, ScanError{scanErrorMessage, err}
	}

	tmpRoleTemplate.Entries = tmpFeatures

	genericFeatures = []models.FeatureEntry{}
	for _, tmpFeature := range tmpRoleTemplate.Entries {
		genericEndpoints = []models.Endpoint{}
		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := models.Endpoint{ID: tmpEndpoint.ID, Path: tmpEndpoint.Path, Method: tmpEndpoint.Method}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := models.FeatureEntry{ID: tmpFeature.ID, Name: tmpFeature.Name, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRoleTemplate := models.Role{ID: tmpRoleTemplate.ID, Name: tmpRoleTemplate.Name, Entries: genericFeatures}

	return genericRoleTemplate, nil
}

func CreateRole(repository *RolesRepository, role models.Role) error {
	query := `INSERT INTO roles(name) VALUES($1) RETURNING(id)`

	var roleId int

	err := repository.DB.QueryRow(query, role.Name).Scan(&roleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	// Establish connection between the role and it's features
	query = `INSERT INTO roles_to_features(role_id, feature_id) VALUES($1, $2)`

	for _, feature := range role.Entries {
		_, err = repository.DB.Exec(query, roleId, feature.ID)
		if err != nil {
			return QueryError{queryErrorMessage, err}
		}
	}

	return nil
}

func UpdateRole(repository *RolesRepository, role models.Role, id int) error {
	query := `SELECT FROM roles WHERE id = $1`

	err := repository.DB.QueryRow(query, id).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	// Update role's id and name
	query = `UPDATE roles SET name = $1 WHERE id = $2`

	_, err = repository.DB.Exec(query, role.Name, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	// Delete all connections with the role
	query = `DELETE FROM roles_to_features WHERE role_id = $1`

	_, err = repository.DB.Exec(query, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	// Establish new connection between the role and it's features
	query = `INSERT INTO roles_to_features(role_id, feature_id) VALUES ($1, $2)`

	for _, feature := range role.Entries {
		_, err = repository.DB.Exec(query, id, feature.ID)
		if err != nil {
			return QueryError{queryErrorMessage, err}
		}
	}

	return nil
}

func DeleteRole(repository *RolesRepository, id int) error {
	query := `SELECT FROM roles WHERE id = $1`

	err := repository.DB.QueryRow(query, id).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `DELETE FROM roles WHERE id = $1`

	_, err = repository.DB.Exec(query, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}
