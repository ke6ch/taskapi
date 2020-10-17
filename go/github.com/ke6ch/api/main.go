package main

import (
	"net/http"

	"github.com/ke6ch/api/handler"
	"github.com/ke6ch/api/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	redisStore "gopkg.in/boj/redistore.v1"
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://app:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	store, err := redisStore.NewRediStore(10, "tcp", "localhost:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	e.Use(session.Middleware(store))
	e.Use(serverHeader)

	e.GET("/", func(c echo.Context) error {
		m := model.Greeting{Message: "Hello World"}
		return c.JSON(http.StatusOK, m)
	})

	e.GET("/login", handler.Login)
	e.POST("/session", handler.Session)

	e.GET("/tasks", handler.GetTasks)
	e.GET("/tasks/:id", handler.GetTask)
	e.POST("/tasks", handler.CreateTask)
	e.PATCH("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)
	e.DELETE("/tasks", handler.DeleteTasks)
	e.GET("/id", handler.GetMaxID)

	e.Logger.Fatal(e.Start(":1323"))
}
