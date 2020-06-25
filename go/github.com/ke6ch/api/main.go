package main

import (
	"net/http"

	"github.com/ke6ch/api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(c)
	}
}

func main() {
	// echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(serverHeader)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTask)
	e.POST("/tasks", handler.CreateTask)
	e.PATCH("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)
	e.DELETE("/tasks", handler.DeleteTasks)
	e.GET("/id", handler.GetMaxId)

	e.Logger.Fatal(e.Start(":1323"))
}
