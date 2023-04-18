// TODO: implement super tokens auth integrations
// TODO: implement api gateway
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheSnakeWitcher/petpark/internal/adoptions"
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


    svc := adoptions.NewService(db)
    srv := adoptions.NewServer(svc)


    errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
        errs <- http.ListenAndServe(Config.ServerHost + ":" + Config.ServerPort,srv)
	}()


    fmt.Println("listening...")
    fmt.Println(<-errs)

}
