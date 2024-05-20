package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/meetings_app_server/db"
	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/services/meetings"
	"github.com/arturfil/meetings_app_server/services/requests"
	"github.com/arturfil/meetings_app_server/services/subjects"
	"github.com/arturfil/meetings_app_server/services/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type AppServer struct {
    addr string
    db *sql.DB
}

// NewAppServer - constructor for AppServer
func NewAppServer(addr string, db *sql.DB) *AppServer {
    return &AppServer{
        addr: addr,
        db: db,
    }
}

// Serve - will run the application and 
// no error as long as it is running corectly
func (app *AppServer) Serve() error {
    helpers.MessageLogs.InfoLog.Println("API listenting on port", app.addr)

    router := chi.NewRouter()

    router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))


    router.Mount("/api/", router)

    // user entity
    userStore := user.NewStore(app.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(router)

    // meetings entity
    meetingsStore := meetings.NewStore(app.db)
    meetingsHandler := meetings.NewHandler(meetingsStore)
    meetingsHandler.RegisterRoutes(router)

    // subjects entity 
    subjectsStore := subjects.NewStore(app.db)
    subjectsHandler := subjects.NewHandler(subjectsStore)
    subjectsHandler.RegisterRoutes(router)

    // requests entity
    requestStore := requests.NewStore(app.db)
    requestHandler := requests.NewHandler(requestStore)
    requestHandler.RegisterRoutes(router)

    srv := &http.Server{ Addr: fmt.Sprintf("%s", app.addr),
        Handler: router,
    }

    return srv.ListenAndServe()
}

// Main func
func main() {

    dsn := os.Getenv("DSN")

    db, err := db.NewDatabase(dsn)
    if err != nil {
         log.Fatal("Cannot connect to database", err)
    }

    server := NewAppServer(":8080", db.Client)
    if err := server.Serve(); err != nil {
       log.Fatal(err) 
    }
}





