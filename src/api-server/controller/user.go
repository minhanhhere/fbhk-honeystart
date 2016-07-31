package controller

import (
    "github.com/labstack/echo"
    "net/http"
    "fmt"
    "time"
    "github.com/dgrijalva/jwt-go"
    "gitlab.com/hs-api-go/database/model"
    "gitlab.com/hs-api-go/secure/password"
    "gitlab.com/hs-api-go/manager"
    "gitlab.com/hs-api-go/config"
    "gitlab.com/hs-api-go/secure/auth"
)

var User struct {
    Login         echo.HandlerFunc
    Register      echo.HandlerFunc
    LoginFacebook echo.HandlerFunc
    TestAuth      echo.HandlerFunc
}

func init() {

    fmt.Println("init controller.user")

    User.Login = func(c echo.Context) error {
        req := &struct {
            Email    string `json: "email"`
            Password string `json: "password"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }
        user := manager.UserManager.GetByEmail(req.Email)
        if password.ValidatePassword(req.Password, user.Password) {
            // Create token
            token := jwt.New(jwt.SigningMethodHS256)

            // Set claims
            claims := token.Claims.(jwt.MapClaims)
            claims["id"] = user.Id.Hex()
            claims["email"] = user.Email
            claims["roles"] = user.Roles
            claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

            // Generate encoded token and send it as response.
            t, err := token.SignedString(config.SecretByte())
            if err != nil {
                return err
            }
            return c.JSON(http.StatusOK, map[string]interface{}{
                "id": user.Id,
                "email": user.Email,
                "avatar": user.Avatar,
                "firstName": user.FirstName,
                "lastName": user.LastName,
                "roles": user.Roles,
                "address": user.Address,
                "phone": user.Phone,
                "token": t,
            })
        }
        return echo.ErrUnauthorized
    }

    User.Register = func(c echo.Context) error {
        req := &struct {
            FirstName string `json: "firstName"`
            LastName  string `json: "lastName"`
            Email     string `json: "email"`
            Password  string `json: "password"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }
        user := model.NewUser()
        user.FirstName = req.FirstName
        user.LastName = req.LastName
        user.Email = req.Email
        user.Password, _ = password.CreateHash(req.Password)
        manager.UserManager.Save(user)
        return c.JSON(http.StatusOK, user)
    }

    User.LoginFacebook = func(c echo.Context) error {
        req := &struct {
            FirstName  string `json: "firstName"`
            LastName   string `json: "lastName"`
            Email      string `json: "email"`
            FacebookId string `json: "facebookId"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }
        user := manager.UserManager.GetByFacebookId(req.FacebookId)
        if user == nil {
            user = model.NewUser()
            user.FacebookId = req.FacebookId
        }
        user.FirstName = req.FirstName
        user.LastName = req.LastName
        user.Email = req.Email
        manager.UserManager.Save(user)

        token := jwt.New(jwt.SigningMethodHS256)

        // Set claims
        claims := token.Claims.(jwt.MapClaims)
        claims["id"] = user.Id.Hex()
        claims["email"] = user.Email
        claims["roles"] = user.Roles
        claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

        // Generate encoded token and send it as response.
        t, err := token.SignedString(config.SecretByte())
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]interface{}{
            "id": user.Id,
            "email": user.Email,
            "avatar": user.Avatar,
            "firstName": user.FirstName,
            "lastName": user.LastName,
            "roles": user.Roles,
            "address": user.Address,
            "phone": user.Phone,
            "token": t,
        })
    }

    User.TestAuth = func(c echo.Context) error {
        user := auth.CurrentUser(c)
        return c.JSON(http.StatusOK, user)
    }
}

func RegisterUserRoute(api *echo.Group) {
    api.POST("/users/login", User.Login)
    api.POST("/users/loginFacebook", User.LoginFacebook)
    api.POST("/users", User.Register)
    api.GET("/users/testAuth", User.TestAuth, auth.Secured(), auth.RolesAccepted(auth.ROLE_USER))
}