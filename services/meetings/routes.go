package meetings

import (
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
}

func (h *Handler) getAllMeetings(w http.ResponseWriter, r *http.Request) {
    meetings, err := h.store.GetAllMeetings()
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, meetings)
}

