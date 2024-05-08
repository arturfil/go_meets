package subjects

import (
	"net/http"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
    store types.SubjectStore
}

func NewHandler(store types.SubjectStore) *Handler {
    return &Handler{
        store: store,
    }
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
    router.Route("/v1/subjects", func(router chi.Router) {
        router.Get("/", h.getAllSubjects)
    })
}

func (h *Handler) getAllSubjects(w http.ResponseWriter, r *http.Request) {
    subjects, err := h.store.GetAllSubjects()
    if err != nil {
        helpers.WriteERROR(w, http.StatusInternalServerError, err)
        return 
    }

    helpers.WriteJSON(w, http.StatusOK, subjects)
}
