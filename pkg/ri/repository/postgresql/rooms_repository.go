package postgresql

import (
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
	"github.com/omc-college/management-system/pkg/ri/models"
)

type RoomsRepository struct {
	db sqlx.Ext
}

func NewRoomsRepository(db sqlx.Ext) *RoomsRepository {
	return &RoomsRepository{
		db: db,
	}
}

//InsertRoom inserts Room's data into DB
func (rr *RoomsRepository) InsertRoom(r models.Room) error {
	var err error

	result, err := rr.db.Queryx("INSERT INTO rooms (room, modified_at) VALUES ($1, CURRENT_TIMESTAMP) RETURNING id", r.Room)
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

func (rr *RoomsRepository) GetAllRooms() ([]models.Room, error) {

	var Rooms []models.Room

	result, err := rr.db.Queryx("SELECT * FROM rooms ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var r models.Room

		err := result.StructScan(&r)
		if err != nil {
			return Rooms, err
		}

		rs, err := ToRoom(r)
		if err != nil {
			return Rooms, err
		}

		Rooms = append(Rooms, *rs)
	}

	return Rooms, nil
}

func (rr *RoomsRepository) GetRoom(ID int) (*models.Room, error) {

	r := models.Room{}

	result, err := rr.db.Queryx("SELECT * FROM rooms WHERE id= $1", ID)
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
		return nil, postgresql.ErrNoRows
	}

	return ToRoom(r)
}

func (rr *RoomsRepository) UpdateRoom(r models.Room, id int) error {

	Room := models.Room{}

	result, err := rr.db.Queryx("SELECT * FROM rooms WHERE id= $1", r.ID)
	if err != nil {
		return err
	}

	for result.Next() {
		err = result.StructScan(&Room)
		if err != nil {
			return err
		}
	}

	_, err = rr.db.Exec("UPDATE rooms SET room = $1, modified_at= CURRENT_TIMESTAMP WHERE id = $2", r.Room, r.ID)
	if err != nil {
		return QueryError{QueryErrorMessage, err}
	}

	return nil
}

func (rr *RoomsRepository) DeleteRoom(RoomID int) error {

	_, err := rr.db.Exec("DELETE FROM rooms WHERE id= $1", RoomID)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RoomsRepository) GetRoomByName(name string) (*models.Room, error) {
	r := models.Room{}

	result, err := rr.db.Queryx("SELECT * FROM rooms WHERE room= $1", name)
	if err != nil {
		return nil, QueryError{QueryErrorMessage, err}
	}

	for result.Next() {
		err = result.StructScan(&r)
		if err != nil {
			return nil, err
		}
	}

	return ToRoom(r)
}

func ToRoom(r models.Room) (*models.Room, error) {

	var genericRoom *models.Room

	genericRoom = &models.Room{
		ID:              r.ID,
		Room:        	 r.Room,
		ModifiedAt:		 r.ModifiedAt,
	}

	return genericRoom, nil
}
