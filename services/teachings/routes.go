package teachings

import (
	"encoding/json"
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
    router.Get("/v1/schedules/{id}", h.getAvailableTimes)
	router.Post("/v1/create/schedule", h.createSchedule)
	router.Delete("/v1/schedule/{id}", h.deleteSchedule)
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

	times, err := h.store.GetSchedules(id)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, times)
}

func (h *Handler) createSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule types.Schedule

	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateSchedule(schedule)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusNoContent, "Successfully created a schedule for a day")
}

func (h *Handler) deleteSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.store.DeleteSchedule(id)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusNoContent, "Successfully deleted the schedule")
}
