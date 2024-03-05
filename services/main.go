package services

import (
	"database/sql"
	"time"
)

// time for db to process any transaction
const dbTimeout = time.Second * 3

var db *sql.DB

// create a new Pool DB connection
func New(dbPool *sql.DB) Models {
    db = dbPool
    return Models{}
}

type Config struct {
    Port string
}

type Models struct {
    
}
