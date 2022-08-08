package handler

import (
	srvc "astro/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *srvc.Service
}

func NewHandler(services *srvc.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/picday", h.getPicOfTheDay).Methods(http.MethodGet)
	router.HandleFunc("/stored", h.getPicturesFromStorage).Methods(http.MethodGet)

	return router
}
