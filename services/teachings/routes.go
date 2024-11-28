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

	router.Route("/v1/teachings", func(router chi.Router) {
		router.Get("/{id}", h.getAllTeachings)
		router.Post("/create", h.createTeaching)
		router.Delete("/delete/{teachingId}", h.deleteTeaching)
	})

	router.Route("/v1/schedules", func(router chi.Router) {
		router.Get("/{id}", h.getUsersSchedule)
		router.Post("/schedule", h.createSchedule)
		router.Delete("/{id}", h.deleteSchedule)
	})
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

func (h *Handler) createTeaching(w http.ResponseWriter, r *http.Request) {
	var teaching types.TeachingSubmission

	err := json.NewDecoder(r.Body).Decode(&teaching)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateTeaching(teaching)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Successfully created a schedule for a day")
}

func (h *Handler) deleteTeaching(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "teachingId")

	err := h.store.DeleteTeaching(id)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Successfully deleted the teaching")
}

func (h *Handler) getUsersSchedule(w http.ResponseWriter, r *http.Request) {
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

	helpers.WriteJSON(w, http.StatusOK, "Successfully created a schedule for a day")
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
