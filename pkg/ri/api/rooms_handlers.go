package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omc-college/management-system/pkg/ri/models"
	"github.com/omc-college/management-system/pkg/ri/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RoomsHandler struct {
	RoomsService *service.RoomsService
}

func NewRoomsHandler(service *service.RoomsService) *RoomsHandler {
	return &RoomsHandler{
		RoomsService: service,
	}
}

func (handler *RoomsHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	rooms, err := handler.RoomsService.GetAllRooms()
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RoomsHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	room, err := handler.RoomsService.GetRoom(id)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RoomsHandler) InsertRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room models.Room
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &room)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RoomsService.InsertRoom(room)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RoomsHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room models.Room
	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &room)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RoomsService.UpdateRoom(room, id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RoomsHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RoomsService.DeleteRoom(id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
