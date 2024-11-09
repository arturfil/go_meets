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
		router.Get("/{userId}", h.getUserSubjects)
		router.Get("/bycategory/{categoryId}", h.getAllSubjectsByCategory)
		router.Get("/subject/{subjectId}", h.getSubjectById)
	})

	router.Route("/v1/categories", func(router chi.Router) {
		router.Get("/", h.getSubjectCategories)
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

func (h *Handler) searchSubject(w http.ResponseWriter, r *http.Request) {
	queryWord := r.URL.Query().Get("queryWord")

	subjects, err := h.store.SearchSubject(queryWord)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

    helpers.WriteJSON(w, http.StatusOK, subjects)
}

func (h *Handler) getUserSubjects(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	userSubjects, err := h.store.GetUserSubjects(userId)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, userSubjects)
}

func (h *Handler) getAllSubjectsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "categoryId")

	subjects, err := h.store.GetAllSubjectsByCategory(categoryId)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, subjects)
}

func (h *Handler) getSubjectById(w http.ResponseWriter, r *http.Request) {
	subjectId := chi.URLParam(r, "subjectId")

	subject, err := h.store.GetSubjectById(subjectId)
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, subject)
}

func (h *Handler) getSubjectCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.store.GetSubjectCategories()
	if err != nil {
		helpers.WriteERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, categories)
}
