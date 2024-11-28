package requests

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.RequestStore
}

func NewHandler(store types.RequestStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Get("/v1/requests", h.getAllRequests)
	router.Get("/v1/requests/{id}", h.getRequestById)
	router.Post("/v1/requests/create", h.createRequest)
	router.Put("/v1/request/update", h.updateRequest)
}

func (h *Handler) getAllRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := h.store.GetAllRequests()
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, requests)
}

func (h *Handler) getRequestById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	request, err := h.store.GetRequestById(id)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		log.Println("Error", err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, request)
}

func (h *Handler) createRequest(w http.ResponseWriter, r *http.Request) {
	request := types.Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateRequest(request)
	if err != nil {
		log.Println("Error\t", log.Lshortfile)
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusNoContent, "Successfully created a request")
}

func (h *Handler) updateRequest(w http.ResponseWriter, r *http.Request) {
	var request types.Request

	token, claims, err := helpers.GetTokenClaims(r)

	if !token.Valid {
		helpers.WriteERROR(w, http.StatusUnauthorized, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.UpdateRequest(claims.Sub, request)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusNoContent, "Successfully updated your request!")

}
