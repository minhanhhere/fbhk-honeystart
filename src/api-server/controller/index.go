package controller

import (
    "github.com/labstack/echo"
    "net/http"
    "fmt"
)

var Index struct {
    Get echo.HandlerFunc
}

func init() {
    fmt.Println("init controller.index")
    Index.Get = func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Welcome to API v1!",
        })
    }
}

func RegisterIndexRoute(api *echo.Group) {
    api.GET("", Index.Get)
}