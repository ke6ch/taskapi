package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
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

// GET /tasks
func getTasks(c echo.Context) error {
	rows, err := db.Query("SELECT * FROM tasks order by status desc, `order`")
	if err != nil {
		panic(err.Error())
	}

	var tasks []Task
	var Status []uint8

	for rows.Next() {
		err := rows.Scan(&task.Id, &task.Name, &Status, &task.Order, &task.Timestamp)
		if err != nil {
			panic(err.Error())
		}

		if Status[0] == 0 {
			task.Status = false
		} else {
			task.Status = true
		}
		tasks = append(tasks, task)
	}
	return c.JSON(http.StatusOK, tasks)
}

// GET /tasks/:id
func getTask(c echo.Context) error {
	var tasks []Task
	var Status []uint8

	if err := db.QueryRow("SELECT * FROM tasks where id = "+c.Param("id")).Scan(&task.Id, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
		panic(err.Error())
	}

	if Status[0] == 0 {
		task.Status = false
	} else {
		task.Status = true
	}
	tasks = append(tasks, task)

	return c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func createTask(c echo.Context) error {
	// データ取得
	task := new(Task)
	if err = c.Bind(task); err != nil {
		panic(err.Error())
	}

	// insert
	stmt, err := db.Prepare("INSERT INTO tasks(id, name, status, `order`, timestamp) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(task.Id, task.Name, task.Status, task.Order, time.Now())

	// TODO: 構造体の使いまわしできたらコメントアウト外す
	// 登録したデータを返す
	// var tasks []Task
	// var Status []uint8
	// if err := db.QueryRow("SELECT * FROM tasks where id = (SELECT max(id) FROM tasks)").Scan(&task.Id, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
	// 	panic(err.Error())
	// }

	// if Status[0] == 0 {
	// 	task.Status = false
	// } else {
	// 	task.Status = true
	// }
	// tasks = append(tasks, task)

	// // return c.JSON(http.StatusCreated, tasks)
	return c.JSON(http.StatusCreated, nil)
}

// PATCH /tasks/:id
func updateTask(c echo.Context) error {
	// TODO: idとtimestampはPOSTにいれたくない
	// データ取得
	id := c.Param("id")
	var status uint8
	if c.QueryParam("status") == "false" {
		status = 0
	} else {
		status = 1
	}

	// update
	stmt, err := db.Prepare("UPDATE tasks SET status = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(status, id)

	// TODO: 構造体の使いまわしできたらコメントアウト外す
	// 登録したデータを返す
	// var tasks []Task
	// var Status []uint8
	// if err := db.QueryRow("SELECT * FROM tasks where id = (SELECT max(id) FROM tasks)").Scan(&task.Id, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
	// 	panic(err.Error())
	// }

	// if Status[0] == 0 {
	// 	task.Status = false
	// } else {
	// 	task.Status = true
	// }
	// tasks = append(tasks, task)

	// return c.JSON(http.StatusCreated, tasks)
	return c.JSON(http.StatusCreated, nil)
}

// DELETE /tasks/:id
func deleteTask(c echo.Context) error {
	// delete
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}

// DELETE /tasks
func deleteTasks(c echo.Context) error {
	// delete
	stmt, err := db.Prepare("DELETE FROM tasks WHERE status = 0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stmt)
	fmt.Println("データ削除")
	stmt.Exec()
	return c.JSON(http.StatusOK, nil)
}

// GET /id
func getMaxId(c echo.Context) error {
	var id int32
	if err := db.QueryRow("SELECT max(id) FROM tasks").Scan(&id); err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func main() {
	// echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(ServerHeader)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)
	e.POST("/tasks", createTask)
	e.PATCH("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.DELETE("/tasks", deleteTasks)
	e.GET("/id", getMaxId)

	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}
