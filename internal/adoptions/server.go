package adoptions

import (
    "github.com/gorilla/mux"
    "net/http"
)

type Server struct {
    *mux.Router
}

func NewServer(svc PetService) (http.Handler) {
    router := mux.NewRouter()
    srv := &Server { router }
    srv.InitHandlers(svc)
    return srv
}
