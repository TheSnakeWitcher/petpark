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
    DbUri := fmt.Sprintf(
        "%s://%s:%s@%s:%d/%s?sslmode=disable",
        Config.DbProtocol,
        Config.DbUser,
        Config.DbPassword,
        Config.DbHost,
        Config.DbPort,
        Config.DbName,
    )

    db ,InitDbErr = sql.Open(Config.DbDriver,DbUri)
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
