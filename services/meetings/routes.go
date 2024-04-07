package meetings

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
    store types.MeetingStore // this is an interface
}

func NewHandler(store types.MeetingStore) *Handler {
    return &Handler{
        store: store,
    }
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
    router.Get("/v1/meetings", h.getAllMeetings)
    router.Post("/v1/meetings/create", h.createMeeting)
}

func (h *Handler) getAllMeetings(w http.ResponseWriter, r *http.Request) {
    meetings, err := h.store.GetAllMeetings()
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, meetings)
}

func (h *Handler) createMeeting(w http.ResponseWriter, r *http.Request) {
    meeting := types.Meeting{}
    err := json.NewDecoder(r.Body).Decode(&meeting)
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    err = h.store.CreateMeeting(meeting)  
    if err != nil {
        log.Println("Error\t", log.Lshortfile)
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusNoContent, "Successfuly created a meeting")
}

