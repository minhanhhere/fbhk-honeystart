package controller

import (
    "github.com/labstack/echo"
    "net/http"
    "fmt"
    "gitlab.com/hs-api-go/manager"
    "gitlab.com/hs-api-go/database/model"
)

var Project struct {
    GetHome    echo.HandlerFunc
    FindById   echo.HandlerFunc
    Save       echo.HandlerFunc
    GetUpdates echo.HandlerFunc
    GetBackers echo.HandlerFunc
}

func init() {
    fmt.Println("init controller.project")
    Project.GetHome = func(c echo.Context) error {
        projects := manager.ProjectManager.GetAll()
        return c.JSON(http.StatusOK, projects)
    }
    Project.FindById = func(c echo.Context) error {
        id := c.Param("id")
        project := manager.ProjectManager.FindById(id)
        return c.JSON(http.StatusOK, project)
    }
    Project.Save = func(c echo.Context) error {
        req := &struct {
            Name    string `json: "name"`
            Summary string `json: "summary"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }
        project := model.NewProject()
        manager.ProjectManager.SaveProject(project)
        return c.JSON(http.StatusOK, project)
    }
    Project.GetUpdates = func(c echo.Context) error {
        id := c.Param("id")
        updates := manager.ProjectManager.GetUpdates(id)
        return c.JSON(http.StatusOK, updates)
    }
    Project.GetBackers = func(c echo.Context) error {
        id := c.Param("id")
        backers := manager.ProjectManager.GetBackers(id)
        return c.JSON(http.StatusOK, backers)
    }
}

func RegisterProjectRoute(api *echo.Group) {
    api.GET("/projects/home", Project.GetHome)
    api.GET("/projects/:id", Project.FindById)
    api.GET("/projects/:id/updates", Project.GetUpdates)
    api.GET("/projects/:id/backers", Project.GetBackers)
    api.POST("/projects", Project.Save)
}