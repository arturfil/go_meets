package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
    store types.UserStore // interface type
}

func NewHandler(store types.UserStore) *Handler {
    return &Handler {
        store: store,
    }
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
    router.Get("/v1/healthcheck", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Msg  string
			Code int
		}{
			Msg:  "Making sure this works",
			Code: 200,
		}

		jsonStr, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonStr)
}
