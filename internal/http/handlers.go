package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheSnakeWitcher/petplanet/pets"
	"github.com/labstack/echo/v4"
)


func Home(ctx echo.Context) error {
    return ctx.String(http.StatusOK,"welcome")
}

func ListPets(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        pets , err := svc.ListPets(ctx.Request().Context())
        if err != nil {
            ctx.Logger().Error(err)
            return ctx.JSON(http.StatusNoContent,fmt.Sprintf("\"error\":\"%s\"",err))
        }
        return ctx.JSON(http.StatusOK,pets)
    }
}

func GetPet(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        id ,err := strconv.Atoi(ctx.Param("id"))
        if err != nil {
            ctx.JSON(http.StatusNotModified,fmt.Sprintf("\"error\": \"%s\"",err))
        }

        var pet pets.Pet
        pet , err = svc.GetPet(ctx.Request().Context(),int32(id))
        return ctx.JSON(http.StatusOK,pet)
    }
}

func AddPet(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        name := ctx.QueryParam("name")
        loc := ctx.QueryParam("location")
        pet , err := svc.AddPet(ctx.Request().Context(),name,loc)
        if err != nil {
            ctx.String(http.StatusNotImplemented,fmt.Sprintf(" \"error\" : \"%s\" ",err))
        }
        return ctx.JSON(http.StatusOK,pet)
    }
}

func DelPet(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        id ,err := strconv.Atoi(ctx.Param("id"))
        if err != nil {
            ctx.JSON(http.StatusNotModified,fmt.Sprintf("\"error\": \"%s\"",err))
        }
        svc.DelPet(ctx.Request().Context(),int32(id))
        return ctx.JSON(http.StatusOK,fmt.Sprintf("\"id\" : \"%d\"",id))
    }
}
