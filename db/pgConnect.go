package db

import (
    "fmt"
    "log"
    "database/sql"

    _ "github.com/lib/pq"
    "github.com/ecuyle/reservationsapi/config"
)

func PGConnect() *sql.DB {
    connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", config.USER, config.DBNAME, config.PASSWORD, config.HOST)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

