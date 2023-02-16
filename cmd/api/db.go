package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

var (
  db *sql.DB
  InitDbErr error
)

func initDB() {
    DbOpts := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        Config.DbHost,
        Config.DbPort,
        Config.DbUser,
        Config.DbPassword,
        Config.DbName,
    )
    fmt.Println(DbOpts)


    db ,InitDbErr = sql.Open(Config.DbDriver,Config.DbUri)
    if InitDbErr != nil {
        InitErrCheck(InitDbErr)
    }
    Logger.Trace().Msg("db connection seted")

    InitDbErr = db.Ping()
    if InitDbErr != nil {
        InitErrCheck(InitDbErr)
    }
    Logger.Trace().Msg("db connection verified")
}
