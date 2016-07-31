package auth

import (
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
)

type AuthorizedUser struct {
    Id string
    Email string
    Roles []interface{}
}

func CurrentUser(c echo.Context) *AuthorizedUser {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    return &AuthorizedUser{
        Id: claims["id"].(string),
        Email: claims["email"].(string),
        Roles: claims["roles"].([]interface{}),
    }
}

func hasRoles(c echo.Context, neededRoles ...string) bool {

    _, currentRoles := getRoles(c)

    for _, role := range neededRoles {
        if containsInSlice(currentRoles, role) {
            return true
        }
    }
    return false
}

func getRoles(c echo.Context) (*jwt.Token, []interface{}) {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    roles := claims["roles"].([]interface{})
    return user, roles
}
