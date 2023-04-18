package adoptions

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")

func (srv Server) InitHandlers(svc PetService) {
    endpoints := MakeServerEndpoints(svc)
    options := []httptransport.ServerOption{
        httptransport.ServerErrorEncoder(encodeError),
    }

    srv.Methods(http.MethodGet).Path("/pets").Handler(httptransport.NewServer(
        endpoints.ListPetsEndpoint,
        decodeListPetsReq,
        encodeResponse,
        options...,
    ))

    srv.Methods(http.MethodGet).Path("/pets/{id}").Handler(httptransport.NewServer(
        endpoints.GetPetEndpoint,
        decodeGetPetReq,
        encodeResponse,
        options...,
    ))

    srv.Methods(http.MethodPost).Path("/pets").Handler(httptransport.NewServer(
        endpoints.AddPetEndpoint,
        decodeAddPetReq,
        encodeResponse,
        options...,
    ))

    srv.Methods(http.MethodDelete).Path("/pets/{id}").Handler(httptransport.NewServer(
        endpoints.DelPetEndpoint,
        decodeDelPetReq,
        encodeResponse,
        options...,
    ))

}


////////////////////////////////////////////////////////////////////////////////
// decode functions
////////////////////////////////////////////////////////////////////////////////


func decodeListPetsReq(ctx context.Context, r *http.Request) (any, error) {
	return r.Body , nil
}

func decodeGetPetReq(ctx context.Context, r *http.Request) (any, error) {
    vars := mux.Vars(r)
    id,ok := vars["id"]
    if !ok {
        return GetPetReq{},ErrBadRouting
    }
    return GetPetReq{Id:id} , nil
}

func decodeAddPetReq(ctx context.Context, r *http.Request) (any, error) {
    var addPetParams AddPetParams
    println("decoding json body")
	if err := json.NewDecoder(r.Body).Decode(&addPetParams); err != nil {
		return nil, err
	}
    println("after decode json body")
    return AddPetReq{Arg:addPetParams} , nil
}

func decodeDelPetReq(ctx context.Context, r *http.Request) (any, error) {
    vars := mux.Vars(r)
    id,ok := vars["id"]
    if !ok {
        return DelPetReq{},ErrBadRouting
    }
    return DelPetReq{Id:id} , nil
}


////////////////////////////////////////////////////////////////////////////////
// encode functions
////////////////////////////////////////////////////////////////////////////////


func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}


////////////////////////////////////////////////////////////////////////////////
// utily functions
////////////////////////////////////////////////////////////////////////////////


func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrBadRouting:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
