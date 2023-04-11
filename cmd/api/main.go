// TODO: implement super tokens auth integrations
package main

import (
	"time"
	"github.com/TheSnakeWitcher/PetPark/internal/http"
	"github.com/TheSnakeWitcher/PetPark/pets"
	mw "github.com/labstack/echo/v4/middleware"
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
	defer LogFile.Close()
	defer db.Close()

    svc := pets.NewService(db)
    srv := http.NewServer(*svc)
    srv.Use(mw.TimeoutWithConfig(mw.TimeoutConfig{
        Timeout: Config.BaseTimeout + time.Second,
    }))
    srv.Logger.Fatal(srv.Start(Config.ServerHost + ":" +  Config.ServerPort))
	Logger.Trace().Msg("execution ends")
}
