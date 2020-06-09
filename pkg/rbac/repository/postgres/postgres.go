package postgres

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/omc-college/management-system/pkg/rbac/models"
)

type RolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *RolesRepository {
	return &RolesRepository{
		db: db,
	}
}

func (repository *RolesRepository) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	query := `SELECT roles.id, roles.name, features.id, features.name, features.description, endpoints.id, endpoints.name, endpoints.path, endpoints.method
			  FROM roles LEFT JOIN roles_to_features
			  ON roles.id = roles_to_features.role_id
			  LEFT JOIN features
			  ON roles_to_features.feature_id = features.id
			  LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	var isRoleExisting bool
	var isFeatureExisting bool

	tmpRoles := make(map[int]role)

	rows, err := repository.db.QueryxContext(ctx, query)
	if err != nil {
		return []models.Role{}, QueryError{queryErrorMessage, err}
	}

	for rows.Next() {
		var tmpRole role
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isRoleExisting = false
		isFeatureExisting = false

		err := rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return []models.Role{}, ScanError{scanErrorMessage, err}
		}

		_, isRoleExisting = tmpRoles[tmpRole.ID]
		if !isRoleExisting {
			tmpRole.Entries = make(map[sql.NullInt64]featureEntry)
			tmpRoles[tmpRole.ID] = tmpRole
		}

		_, isFeatureExisting = tmpRoles[tmpRole.ID].Entries[tmpFeature.ID]
		if !isFeatureExisting && tmpFeature.ID.Valid {
			tmpFeature.Endpoints = make(map[sql.NullInt64]endpoint)
			tmpRoles[tmpRole.ID].Entries[tmpFeature.ID] = tmpFeature
		}

		if tmpFeature.ID.Valid {
			tmpRoles[tmpRole.ID].Entries[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
		}
	}

	err = rows.Err()
	if err != nil {
		return []models.Role{}, ScanError{scanErrorMessage, err}
	}

	return toRoles(tmpRoles), nil
}

func (repository *RolesRepository) GetRole(ctx context.Context, id int) (models.Role, error) {
	query := `SELECT roles.id, roles.name, features.id, features.name, features.description, endpoints.id, endpoints.name, endpoints.path, endpoints.method
			  FROM roles LEFT JOIN roles_to_features
			  ON roles.id = roles_to_features.role_id
			  LEFT JOIN features
			  ON roles_to_features.feature_id = features.id
			  LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id
			  WHERE roles.id = $1`

	var isFeatureExisting bool
	var tmpRole role
	tmpRole.Entries = make(map[sql.NullInt64]featureEntry)

	rows, err := repository.db.QueryxContext(ctx, query, id)
	if err != nil {
		return models.Role{}, QueryError{queryErrorMessage, err}
	}

	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isFeatureExisting = false

		err = rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return models.Role{}, ScanError{scanErrorMessage, err}
		}

		_, isFeatureExisting = tmpRole.Entries[tmpFeature.ID]
		if !isFeatureExisting && tmpFeature.ID.Valid {
			tmpFeature.Endpoints = make(map[sql.NullInt64]endpoint)
			tmpRole.Entries[tmpFeature.ID] = tmpFeature
		}

		if tmpFeature.ID.Valid {
			tmpRole.Entries[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
		}
	}

	// rows.Scan after db.Query doesn't return sql.ErrNoRows
	if tmpRole.ID == 0 {
		return models.Role{}, ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		return models.Role{}, ScanError{scanErrorMessage, err}
	}

	return toRole(tmpRole), nil
}

func (repository *RolesRepository) CreateRole(ctx context.Context, role models.Role) error {
	var roleId int

	tx, err := repository.db.Beginx()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := `INSERT INTO roles(name) VALUES($1) RETURNING(id)`

	err = tx.GetContext(ctx, &roleId, query, role.Name)
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
		_, err = tx.ExecContext(ctx, query, roleId, feature.ID)
		if err != nil {
			return QueryError{queryErrorMessage, err}
		}
	}

	err = tx.Commit()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func (repository *RolesRepository) UpdateRole(ctx context.Context, role models.Role, id int) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := `SELECT FROM roles WHERE id = $1`

	err = repository.db.QueryRowxContext(ctx, query, id).Scan()
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

	_, err = repository.db.ExecContext(ctx, query, role.Name, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	// Delete all connections with the role
	query = `DELETE FROM roles_to_features WHERE role_id = $1`

	_, err = repository.db.ExecContext(ctx, query, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	// Establish new connection between the role and it's features
	query = `INSERT INTO roles_to_features(role_id, feature_id) VALUES ($1, $2)`

	for _, feature := range role.Entries {
		_, err = repository.db.ExecContext(ctx, query, id, feature.ID)
		if err != nil {
			return QueryError{queryErrorMessage, err}
		}
	}

	err = tx.Commit()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func (repository *RolesRepository) DeleteRole(ctx context.Context, id int) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := `SELECT FROM roles WHERE id = $1`

	err = repository.db.QueryRowContext(ctx, query, id).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNoRows
		} else {
			err = QueryError{queryErrorMessage, err}
		}
		return err
	}

	query = `DELETE FROM roles WHERE id = $1`

	_, err = repository.db.ExecContext(ctx, query, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	err = tx.Commit()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func (repository *RolesRepository) GetRoleTmpl(ctx context.Context) (models.RoleTmpl, error) {
	var isFeatureExisting bool
	var tmpRoleTmpl roleTmpl
	var tmpFeatures = make(map[sql.NullInt64]featureEntry)

	query := `SELECT features.id, features.name, features.description, endpoints.id, endpoints.name, endpoints.path, endpoints.method
			  FROM features LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	rows, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		return models.RoleTmpl{}, QueryError{queryErrorMessage, err}
	}

	// Get all features and connect them to endpoints
	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		isFeatureExisting = false

		err := rows.Scan(&tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Method, &tmpEndpoint.Path)
		if err != nil {
			return models.RoleTmpl{}, ScanError{scanErrorMessage, err}
		}

		_, isFeatureExisting = tmpFeatures[tmpFeature.ID]
		if !isFeatureExisting {
			tmpFeature.Endpoints = make(map[sql.NullInt64]endpoint)
			tmpFeatures[tmpFeature.ID] = tmpFeature
		}

		tmpFeatures[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
	}

	err = rows.Err()
	if err != nil {
		return models.RoleTmpl{}, ScanError{scanErrorMessage, err}
	}

	tmpRoleTmpl.Entries = tmpFeatures

	return toRoleTmpl(tmpRoleTmpl), nil
}

func (repository *RolesRepository) CreateRoleTmpl(ctx context.Context, roleTmpl models.RoleTmpl) error {
	var addedEndpoints []models.Endpoint
	var isEndpointExisting bool
	var currentFeatureId int
	var currentEndpointId int
	var isFeatureUpdated bool
	var existingFeatures map[string]bool = make(map[string]bool)

	tx, err := repository.db.Begin()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := "SELECT name FROM features"
	rows, err := repository.db.Query(query)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	// make a set of existing features
	for rows.Next() {
		var featureName string

		err := rows.Scan(&featureName)
		if err != nil {
			return ScanError{scanErrorMessage, err}
		}

		existingFeatures[featureName] = true
	}

	err = rows.Err()
	if err != nil {
		return ScanError{scanErrorMessage, err}
	}

	for _, feature := range roleTmpl.Entries {
		isFeatureUpdated = false

		_, isFeatureUpdated = existingFeatures[feature.Name]
		if isFeatureUpdated {
			query := "UPDATE features SET description = $1 WHERE name = $2 RETURNING (id)"
			err := repository.db.QueryRow(query, feature.Description, feature.Name).Scan(&currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}

			query = "DELETE FROM features_to_endpoints WHERE feature_id = $1"
			_, err = repository.db.Exec(query, currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}
		} else {
			query := "INSERT INTO features(name, description) VALUES ($1, $2) RETURNING (id)"
			err := repository.db.QueryRow(query, feature.Name, feature.Description).Scan(&currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}
		}

		for _, endpoint := range feature.Endpoints {
			// delete endpoint of existing feature
			if isFeatureUpdated {
				query := "DELETE FROM endpoints WHERE name = $1"
				_, err := repository.db.Exec(query, endpoint.Name)
				if err != nil {
					return QueryError{queryErrorMessage, err}
				}

				for existingEndpointIndex := range addedEndpoints {
					addedEndpoints = append(addedEndpoints[:existingEndpointIndex], addedEndpoints[existingEndpointIndex+1:]...)
				}
			}

			// check whether endpoint exists
			isEndpointExisting = false
			for _, existingEndpoint := range addedEndpoints {
				if existingEndpoint.Name == endpoint.Name {
					isEndpointExisting = true
					break
				}
			}

			if !isEndpointExisting {
				query := "INSERT INTO endpoints(name, path, method) VALUES ($1, $2, $3) RETURNING (id)"
				err := repository.db.QueryRow(query, endpoint.Name, endpoint.Path, endpoint.Method).Scan(&currentEndpointId)
				if err != nil {
					return QueryError{queryErrorMessage, err}
				}

				addedEndpoints = append(addedEndpoints, endpoint)
			} else {
				query := "SELECT id FROM endpoints WHERE name = $1"
				err := repository.db.QueryRow(query, endpoint.Name).Scan(&currentEndpointId)
				if err != nil {
					return QueryError{queryErrorMessage, err}
				}
			}

			query := "INSERT INTO features_to_endpoints(feature_id, endpoint_id) VALUES ($1, $2)"
			_, err := repository.db.Exec(query, currentFeatureId, currentEndpointId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	return nil
}

func toRoles(tmpRoles map[int]role) (genericRoles []models.Role) {
	for _, tmpRole := range tmpRoles {
		genericRoles = append(genericRoles, toRole(tmpRole))
	}

	return genericRoles
}

func toRole(tmpRole role) (genericRole models.Role) {
	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint

	genericFeatures = []models.FeatureEntry{}
	for _, tmpFeature := range tmpRole.Entries {
		genericEndpoints = []models.Endpoint{}
		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := models.Endpoint{ID: int(tmpEndpoint.ID.Int64), Name: tmpEndpoint.Name.String, Path: tmpEndpoint.Path.String, Method: tmpEndpoint.Method.String}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := models.FeatureEntry{ID: int(tmpFeature.ID.Int64), Name: tmpFeature.Name.String, Description: tmpFeature.Description.String, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRole = models.Role{ID: tmpRole.ID, Name: tmpRole.Name, Entries: genericFeatures}

	return genericRole
}

func toRoleTmpl(tmpRoleTmpl roleTmpl) (genericRoleTmpl models.RoleTmpl) {
	var genericFeatures []models.FeatureEntry
	var genericEndpoints []models.Endpoint

	genericFeatures = []models.FeatureEntry{}
	for _, tmpFeature := range tmpRoleTmpl.Entries {
		genericEndpoints = []models.Endpoint{}
		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := models.Endpoint{Name: tmpEndpoint.Name.String, Path: tmpEndpoint.Path.String, Method: tmpEndpoint.Method.String}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}
		genericFeature := models.FeatureEntry{Name: tmpFeature.Name.String, Description: tmpFeature.Description.String, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}
	genericRoleTmpl = models.RoleTmpl{Entries: genericFeatures}

	return genericRoleTmpl
}
