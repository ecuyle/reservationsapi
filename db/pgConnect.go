package db

import (
    "log"
    "database/sql"

    _ "github.com/lib/pq"
)

func PGConnect() *sql.DB {
    connStr := "user= dbname= password= host=ec2-52-14-155-234.us-east-2.compute.amazonaws.com sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

