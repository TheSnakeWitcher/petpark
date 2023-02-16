package main

import (
	"github.com/TheSnakeWitcher/petplanet/internal/http"
	"github.com/TheSnakeWitcher/petplanet/pets"
)

var (
    InitErr error
)

func InitErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}


func init() {
    InitErr = InitConfig()
    InitErrCheck(InitErr)

    InitErr = InitLogger()
    InitErrCheck(InitErr)

    initDB()
}


func main() {
	Logger.Trace().Msg("execution starts")
	defer db.Close()
	defer LogFile.Close()

    initDB()
    svc := pets.NewService(db)
    srv := http.NewServer(*svc)
    srv.Logger.Fatal(srv.Start(":" + Config.ServerPort))

	Logger.Trace().Msg("execution ends")
}
