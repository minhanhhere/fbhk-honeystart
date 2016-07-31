package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/fasthttp"
    "gitlab.com/hs-api-go/controller"
    "github.com/labstack/echo/middleware"
    "gitlab.com/hs-api-go/database"
    "fmt"
    "gitlab.com/hs-api-go/manager"
    _ "gitlab.com/hs-api-go/config"
)

const (
    APP_NAME = "hs-api-go"
    VERSION = "1.0.0"
)

func main() {

    fmt.Println("Start server", APP_NAME, VERSION)

    conn := database.GetConnection()
    defer fmt.Println("server terminating...")
    defer conn.Session.Close()
    defer fmt.Println("server terminated!")

    fmt.Printf("Total user: %d\n", manager.UserManager.CountAll())

    server := echo.New()
    server.SetDebug(true)
    server.Pre(middleware.RemoveTrailingSlash())
    server.Use(middleware.CORS())

    api := server.Group("/api/v1")
    controller.RegisterIndexRoute(api)
    controller.RegisterUserRoute(api)
    controller.RegisterProjectRoute(api)
    controller.RegisterDisqusRoute(api)
    controller.RegisterTransactionRoute(api)

    server.Run(fasthttp.New(":5000"))
}
