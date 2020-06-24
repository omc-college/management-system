package api

import (
	"github.com/gorilla/mux"
	"github.com/omc-college/management-system/pkg/ri/service"
	"net/http"
)

func NewRoomsRouter(service *service.RoomsService) *mux.Router {

	roomsHandler := NewRoomsHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/rooms", roomsHandler.GetAllRooms).Methods(http.MethodGet)
	router.HandleFunc("/rooms", roomsHandler.InsertRoom).Methods(http.MethodPost)
	router.HandleFunc("/rooms/{id}", roomsHandler.GetRoom).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{id}", roomsHandler.UpdateRoom).Methods(http.MethodPut)
	router.HandleFunc("/rooms/{id}", roomsHandler.DeleteRoom).Methods(http.MethodDelete)
	return router
}
