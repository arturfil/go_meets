package teachings

import (
	"net/http"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.TeachingsStore
}

func NewHandler(store types.TeachingsStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
    router.Get("/v1/teachings/{id}", h.getAllTeachings)
    router.Get("/v1/teachings/schedule/{id}", h.getAvailableTimes)
}

func (h *Handler) getAllTeachings(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    teachings, err := h.store.GetAllTeachings(id)
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, teachings)
}

func (h *Handler) getAvailableTimes(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    times, err := h.store.GetSchedule(id)
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, times)
}

