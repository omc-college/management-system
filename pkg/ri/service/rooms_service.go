package service

import (
	"github.com/omc-college/management-system/pkg/ri/models"
	postgresql2 "github.com/omc-college/management-system/pkg/ri/repository/postgresql"
)

type RoomsService struct {
	RoomsRepository *postgresql2.RoomsRepository
}

func NewRoomsService(roomsRepository *postgresql2.RoomsRepository) *RoomsService {
	return &RoomsService{
		RoomsRepository: roomsRepository,
	}
}

func (service *RoomsService) GetAllRooms() (rooms []models.Room, err error) {
	rooms, err = service.RoomsRepository.GetAllRooms()
	if err != nil {
		return []models.Room{}, err
	}

	return rooms, nil
}

func (service *RoomsService) GetRoom(id int) (room *models.Room, err error) {
	room, err = service.RoomsRepository.GetRoom(id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (service *RoomsService) InsertRoom(room models.Room) (err error) {
	err = service.RoomsRepository.InsertRoom(room)
	if err != nil {
		return err
	}

	return nil
}

func (service *RoomsService) UpdateRoom(room models.Room, id int) (err error) {
	err = service.RoomsRepository.UpdateRoom(room, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *RoomsService) DeleteRoom(RoomID int) (err error) {
	err = service.RoomsRepository.DeleteRoom(RoomID)
	if err != nil {
		return err
	}

	return nil
}

func (service *RoomsService) GetRoomByName(name string) (room *models.Room, err error) {
	room, err = service.RoomsRepository.GetRoomByName(name)
	if err != nil {
		return nil, err
	}

	return room, nil
}

