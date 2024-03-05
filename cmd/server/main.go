package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/meetings_app_server/db"
	"github.com/arturfil/meetings_app_server/handlers"
	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/services"
)

type application struct {
    config services.Config
    models services.Models
}

// serve - will run the application and 
// no error as long as it is running corectly
func (app *application) serve() error {
    port := os.Getenv("PORT")
    helpers.MessageLogs.InfoLog.Println("API listenting on port", port)

    srv := &http.Server{
        Addr: fmt.Sprintf(":%s", port),
        Handler: handlers.Routes(),
    }

    return srv.ListenAndServe()

}

// Main func
func main() {
    // create server here
    var cfg services.Config
    port := os.Getenv("PORT")
    cfg.Port = port

    dsn := os.Getenv("DSN")
    dbConn, err := db.ConnectPostgres(dsn)
    if err != nil {
        log.Fatal("Cannot connect to database", err)
    }

    defer dbConn.DB.Close()

    var app = &application{
        config: cfg,
        models: services.New(dbConn.DB),
    }

    err = app.serve()
    if err != nil {
        log.Fatal(err)
    }
}



