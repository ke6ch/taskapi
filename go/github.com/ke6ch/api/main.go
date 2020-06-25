package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ke6ch/api/handler"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	Order     int    `json:"order"`
	Timestamp string `json:"timestamp"`
}
type Tasks []Task

var (
	task = Task{}
	err  error
)

var db *sql.DB

func init() {
	db, err = sql.Open("mysql", "user:pass@tcp(db:3306)/clear")
	if err != nil {
		panic(err.Error())
	}
}

// middleware
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(c)
	}
}

func main() {
	// echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(ServerHeader)

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

	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}
