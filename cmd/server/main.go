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
	"github.com/arturfil/meetings_app_server/services/user"
	"github.com/go-chi/chi/v5"
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
    router.Mount("/api/", router)

    // user entity
    userStore := user.NewStore(app.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(router)

    // meetings entity
    meetingsStore := meetings.NewStore(app.db)
    meetingsHandler := meetings.NewHandler(meetingsStore)
    meetingsHandler.RegisterRoutes(router)

    srv := &http.Server{
        Addr: fmt.Sprintf("%s", app.addr),
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





