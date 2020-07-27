package postgres

import (
	"context"
	"database/sql"
	"errors"
	"sort"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/omc-college/management-system/pkg/rbac"
)

type RolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *RolesRepository {
	return &RolesRepository{
		db: db,
	}
}

func (repository *RolesRepository) GetAllRoles(ctx context.Context) ([]rbac.Role, error) {
	query := `SELECT roles.id, roles.name, features.id, features.name, features.description, endpoints.id, endpoints.name, endpoints.path, endpoints.method
			  FROM roles LEFT JOIN roles_to_features
			  ON roles.id = roles_to_features.role_id
			  LEFT JOIN features
			  ON roles_to_features.feature_id = features.id
			  LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	rows, err := repository.db.QueryxContext(ctx, query)
	if err != nil {
		return []rbac.Role{}, QueryError{queryErrorMessage, err}
	}

	var tmpRoles = make(map[int]role)

	for rows.Next() {
		var tmpRole role
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		var isRoleExisting bool
		var isFeatureExisting bool

		err := rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return []rbac.Role{}, ScanError{scanErrorMessage, err}
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
		return []rbac.Role{}, ScanError{scanErrorMessage, err}
	}

	return toRoles(tmpRoles), nil
}

func (repository *RolesRepository) GetRole(ctx context.Context, id int) (rbac.Role, error) {
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

	rows, err := repository.db.QueryxContext(ctx, query, id)
	if err != nil {
		return rbac.Role{}, QueryError{queryErrorMessage, err}
	}

	var tmpRole role
	tmpRole.Entries = make(map[sql.NullInt64]featureEntry)

	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		var isFeatureExisting bool

		err = rows.Scan(&tmpRole.ID, &tmpRole.Name, &tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return rbac.Role{}, ScanError{scanErrorMessage, err}
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
		return rbac.Role{}, ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		return rbac.Role{}, ScanError{scanErrorMessage, err}
	}

	return toRole(tmpRole), nil
}

func (repository *RolesRepository) CreateRole(ctx context.Context, role *rbac.Role) error {
	tx, err := repository.db.Beginx()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	var roleId int

	query := `INSERT INTO roles(name) VALUES($1) RETURNING(id)`

	err = tx.GetContext(ctx, &roleId, query, role.Name)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == errCodeUniqueViolation {
				return ErrNotUniqueRole
			}
		}

		return QueryError{queryErrorMessage, err}
	}

	query = `INSERT INTO roles_to_features(role_id, feature_id) VALUES($1, $2)`

	for _, feature := range role.Entries {
		_, err = tx.ExecContext(ctx, query, roleId, feature.ID)
		if err != nil {
			if pgErr, ok := err.(pgx.PgError); ok {
				if pgErr.Code == errCodeForeignKeyViolation {
					return ErrNonexistentFeature
				}
			}

			return QueryError{queryErrorMessage, err}
		}
	}

	err = tx.Commit()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	role.ID = roleId

	return nil
}

func (repository *RolesRepository) UpdateRole(ctx context.Context, role rbac.Role, id int) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := `SELECT FROM roles WHERE id = $1`

	err = repository.db.QueryRowxContext(ctx, query, id).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoRows
		}

		return QueryError{queryErrorMessage, err}
	}

	query = `UPDATE roles SET name = $1 WHERE id = $2`

	_, err = repository.db.ExecContext(ctx, query, role.Name, id)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == errCodeUniqueViolation {
				return ErrNotUniqueRole
			}
		}

		return QueryError{queryErrorMessage, err}
	}

	query = `DELETE FROM roles_to_features WHERE role_id = $1`

	_, err = repository.db.ExecContext(ctx, query, id)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	query = `INSERT INTO roles_to_features(role_id, feature_id) VALUES ($1, $2)`

	for _, feature := range role.Entries {
		_, err = repository.db.ExecContext(ctx, query, id, feature.ID)
		if err != nil {
			if pgErr, ok := err.(pgx.PgError); ok {
				if pgErr.Code == errCodeForeignKeyViolation {
					return ErrNonexistentFeature
				}
			}

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

func (repository *RolesRepository) GetRoleTmpl(ctx context.Context) (rbac.RoleTmpl, error) {
	query := `SELECT features.id, features.name, features.description, endpoints.id, endpoints.name, endpoints.path, endpoints.method
			  FROM features LEFT JOIN features_to_endpoints
			  ON features.id = features_to_endpoints.feature_id
			  LEFT JOIN endpoints
			  ON features_to_endpoints.endpoint_id = endpoints.id`

	rows, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		return rbac.RoleTmpl{}, QueryError{queryErrorMessage, err}
	}

	var tmpFeatures = make(map[sql.NullInt64]featureEntry)

	// Get all features and connect them to endpoints
	for rows.Next() {
		var tmpFeature featureEntry
		var tmpEndpoint endpoint

		err := rows.Scan(&tmpFeature.ID, &tmpFeature.Name, &tmpFeature.Description, &tmpEndpoint.ID, &tmpEndpoint.Name, &tmpEndpoint.Path, &tmpEndpoint.Method)
		if err != nil {
			return rbac.RoleTmpl{}, ScanError{scanErrorMessage, err}
		}

		var isFeatureExisting bool

		_, isFeatureExisting = tmpFeatures[tmpFeature.ID]
		if !isFeatureExisting {
			tmpFeature.Endpoints = make(map[sql.NullInt64]endpoint)
			tmpFeatures[tmpFeature.ID] = tmpFeature
		}

		tmpFeatures[tmpFeature.ID].Endpoints[tmpEndpoint.ID] = tmpEndpoint
	}

	err = rows.Err()
	if err != nil {
		return rbac.RoleTmpl{}, ScanError{scanErrorMessage, err}
	}

	var tmpRoleTmpl roleTmpl
	tmpRoleTmpl.Entries = tmpFeatures

	return toRoleTmpl(tmpRoleTmpl), nil
}

func (repository *RolesRepository) CreateRoleTmpl(ctx context.Context, roleTmpl rbac.RoleTmpl) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}
	defer tx.Rollback()

	query := "SELECT name FROM features"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	var existingFeatures = make(map[string]bool)

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

	query = "SELECT name FROM endpoints"
	rows, err = tx.QueryContext(ctx, query)
	if err != nil {
		return QueryError{queryErrorMessage, err}
	}

	var existingEndpoints = make(map[string]bool)

	for rows.Next() {
		var endpointName string

		err := rows.Scan(&endpointName)
		if err != nil {
			return ScanError{scanErrorMessage, err}
		}

		existingEndpoints[endpointName] = true
	}

	err = rows.Err()
	if err != nil {
		return ScanError{scanErrorMessage, err}
	}

	for _, feature := range roleTmpl.Entries {
		var currentFeatureId int
		_, isFeatureExisting := existingFeatures[feature.Name]

		if isFeatureExisting {
			query := "UPDATE features SET description = $1 WHERE name = $2 RETURNING (id)"
			err := tx.QueryRowContext(ctx, query, feature.Description, feature.Name).Scan(&currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}

			query = "DELETE FROM features_to_endpoints WHERE feature_id = $1"
			_, err = tx.ExecContext(ctx, query, currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}
		}
		if !isFeatureExisting {
			query := "INSERT INTO features(name, description) VALUES ($1, $2) RETURNING (id)"
			err := tx.QueryRowContext(ctx, query, feature.Name, feature.Description).Scan(&currentFeatureId)
			if err != nil {
				return QueryError{queryErrorMessage, err}
			}
		}

		for _, endpoint := range feature.Endpoints {
			var currentEndpointId int
			_, isEndpointExisting := existingEndpoints[endpoint.Name]

			if isEndpointExisting {
				query = `UPDATE endpoints SET path = $1, method = $2 WHERE name = $3 RETURNING id`
				err := tx.QueryRowContext(ctx, query, endpoint.Path, endpoint.Method, endpoint.Name).Scan(&currentEndpointId)
				if err != nil {
					return QueryError{queryErrorMessage, err}
				}
			}

			if !isEndpointExisting {
				query := "INSERT INTO endpoints(name, path, method) VALUES ($1, $2, $3) RETURNING (id)"
				err := tx.QueryRowContext(ctx, query, endpoint.Name, endpoint.Path, endpoint.Method).Scan(&currentEndpointId)
				if err != nil {
					return QueryError{queryErrorMessage, err}
				}

				existingEndpoints[endpoint.Name] = true
			}

			query := "INSERT INTO features_to_endpoints(feature_id, endpoint_id) VALUES ($1, $2)"
			_, err := tx.ExecContext(ctx, query, currentFeatureId, currentEndpointId)
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

func toRoles(tmpRoles map[int]role) []rbac.Role {
	var genericRoles []rbac.Role

	for _, tmpRole := range tmpRoles {
		genericRoles = append(genericRoles, toRole(tmpRole))
	}

	sort.SliceStable(genericRoles, func(i, j int) bool {
		return genericRoles[i].ID < genericRoles[j].ID
	})

	return genericRoles
}

func toRole(tmpRole role) rbac.Role {
	var genericFeatures = []rbac.FeatureEntry{}

	for _, tmpFeature := range tmpRole.Entries {
		var genericEndpoints = []rbac.Endpoint{}

		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := rbac.Endpoint{ID: int(tmpEndpoint.ID.Int64), Name: tmpEndpoint.Name.String, Path: tmpEndpoint.Path.String, Method: tmpEndpoint.Method.String}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}

		sort.SliceStable(genericEndpoints, func(i, j int) bool {
			return genericEndpoints[i].ID < genericEndpoints[j].ID
		})

		var genericFeature = rbac.FeatureEntry{ID: int(tmpFeature.ID.Int64), Name: tmpFeature.Name.String, Description: tmpFeature.Description.String, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}

	sort.SliceStable(genericFeatures, func(i, j int) bool {
		return genericFeatures[i].ID < genericFeatures[j].ID
	})

	var genericRole = rbac.Role{ID: tmpRole.ID, Name: tmpRole.Name, Entries: genericFeatures}

	return genericRole
}

func toRoleTmpl(tmpRoleTmpl roleTmpl) rbac.RoleTmpl {
	var genericFeatures = []rbac.FeatureEntry{}

	for _, tmpFeature := range tmpRoleTmpl.Entries {
		var genericEndpoints = []rbac.Endpoint{}

		for _, tmpEndpoint := range tmpFeature.Endpoints {
			genericEndpoint := rbac.Endpoint{ID: int(tmpEndpoint.ID.Int64), Name: tmpEndpoint.Name.String, Path: tmpEndpoint.Path.String, Method: tmpEndpoint.Method.String}
			genericEndpoints = append(genericEndpoints, genericEndpoint)
		}

		sort.SliceStable(genericEndpoints, func(i, j int) bool {
			return genericEndpoints[i].ID < genericEndpoints[j].ID
		})

		genericFeature := rbac.FeatureEntry{ID: int(tmpFeature.ID.Int64), Name: tmpFeature.Name.String, Description: tmpFeature.Description.String, Endpoints: genericEndpoints}
		genericFeatures = append(genericFeatures, genericFeature)
	}

	sort.SliceStable(genericFeatures, func(i, j int) bool {
		return genericFeatures[i].ID < genericFeatures[j].ID
	})

	var genericRoleTmpl = rbac.RoleTmpl{Entries: genericFeatures}

	return genericRoleTmpl
}
