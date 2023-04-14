package http

import (
	"fmt"
	"net/http"

	"github.com/TheSnakeWitcher/petpark/pets"
	"github.com/labstack/echo/v4"
)


func Root(ctx echo.Context) error {
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
        id := ctx.Param("id")
        pet , err := svc.GetPet(ctx.Request().Context(),id)
        if err != nil {
            ctx.Logger().Error(err)
            ctx.JSON(http.StatusNotModified,fmt.Sprintf("\"error\": \"%s\"",err))
        }
        return ctx.JSON(http.StatusOK,pet)
    }
}

func AddPet(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        var petParams pets.AddPetParams
        ctx.Bind(&petParams)
        fmt.Println("pets:\n",petParams)

        pet , err := svc.AddPet(ctx.Request().Context(),petParams)
        if err != nil {
            ctx.Logger().Error(err)
            ctx.String(http.StatusNotImplemented,fmt.Sprintf(" \"error\" : \"%s\" ",err))
        }

        return ctx.JSON(http.StatusOK,pet)
    }
}

func DelPet(svc pets.Service) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        id := ctx.Param("id")
        err := svc.DelPet(ctx.Request().Context(),id)
        if err != nil {
            ctx.JSON(http.StatusNotModified,fmt.Sprintf("\"error\": \"%s\"",err))
        }
        return ctx.JSON(http.StatusOK,fmt.Sprintf("\"id\" : \"%s\"",id))
    }
}
