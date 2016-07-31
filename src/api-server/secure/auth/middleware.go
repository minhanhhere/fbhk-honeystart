package auth

import (
    "github.com/labstack/echo"
    "gitlab.com/hs-api-go/config"
    "github.com/labstack/echo/middleware"
)

func Secured() echo.MiddlewareFunc {
    return middleware.JWT(config.SecretByte())
}

func RolesAccepted(neededRoles ...string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            if hasRoles(c, neededRoles...) {
                return next(c)
            }
            return echo.ErrUnauthorized
        }
    }
}