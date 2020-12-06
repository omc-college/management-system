package postgresql

import (
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ri/models"
)

type ResourcesRepository struct {
	db sqlx.Ext
}

func NewResourcesRepository(db sqlx.Ext) *ResourcesRepository {
	return &ResourcesRepository{
		db: db,
	}
}

//InsertResource inserts Resource's data into DB
func (rr *ResourcesRepository) InsertResource(r models.Resources) error {
	var err error

	result, err := rr.db.Queryx("INSERT INTO resources (resource_name, resource_description, modified_at) VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING id", r.ResourceName, r.ResourceDescription)
	if err != nil {
		return QueryError{QueryErrorMessage, err}
	}

	for result.Next() {
		err = result.Scan(&r.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rr *ResourcesRepository) GetAllResources() ([]models.Resources, error) {

	var Resources []models.Resources

	result, err := rr.db.Queryx("SELECT * FROM resources ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var r models.Resources

		err := result.StructScan(&r)
		if err != nil {
			return Resources, err
		}

		rs, err := ToResource(r)
		if err != nil {
			return Resources, err
		}

		Resources = append(Resources, *rs)
	}

	return Resources, nil
}

func (rr *ResourcesRepository) GetResource(ID int) (*models.Resources, error) {

	r := models.Resources{}

	result, err := rr.db.Queryx("SELECT * FROM resources WHERE id= $1", ID)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err = result.StructScan(&r)
		if err != nil {
			return nil, err
		}
	}

	if r.ID == 0 {
		return nil, ErrNoRows
	}

	return ToResource(r)
}

func (rr *ResourcesRepository) UpdateResource(r models.Resources, id int) error {

	Resource := models.Resources{}

	result, err := rr.db.Queryx("SELECT * FROM resources WHERE id= $1", id)
	if err != nil {
		return err
	}

	for result.Next() {
		err = result.StructScan(&Resource)
		if err != nil {
			return err
		}
	}

	_, err = rr.db.Exec("UPDATE resources SET resource_name = $1, resource_description = $2, modified_at= CURRENT_TIMESTAMP WHERE id = $3", r.ResourceName, r.ResourceDescription, id)
	if err != nil {
		return QueryError{QueryErrorMessage, err}
	}

	return nil
}

func (rr *ResourcesRepository) DeleteResource(ResourceID int) error {

	_, err := rr.db.Exec("DELETE FROM resources WHERE id= $1", ResourceID)
	if err != nil {
		return err
	}

	return nil
}

func ToResource(r models.Resources) (*models.Resources, error) {

	var genericResource *models.Resources

	genericResource = &models.Resources{
		ID:                  r.ID,
		ResourceName:        r.ResourceName,
		ResourceDescription: r.ResourceDescription,
		ModifiedAt:			 r.ModifiedAt,
	}

	return genericResource, nil
}
