package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ke6ch/api/model"
	"github.com/labstack/echo"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func open() *sql.DB {
	c := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Net:    "tcp",
	}

	dsn := c.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetTasks GET /tasks
func GetTasks(c echo.Context) error {
	db := open()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks order by status desc, `order`")
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	tasks := []*model.Task{}
	var Status []uint8

	for rows.Next() {
		task := new(model.Task)
		err := rows.Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusServiceUnavailable, err)
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

// GetTask /tasks/:id
func GetTask(c echo.Context) error {
	db := open()
	defer db.Close()

	task := new(model.Task)
	tasks := []*model.Task{}
	var Status []uint8

	if err := db.QueryRow("SELECT * FROM tasks where id = "+c.Param("id")).Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	if Status[0] == 0 {
		task.Status = false
	} else {
		task.Status = true
	}
	tasks = append(tasks, task)

	return c.JSON(http.StatusOK, tasks)
}

// CreateTask POST /tasks
func CreateTask(c echo.Context) error {
	db := open()

	// データ取得
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	// insert
	stmt, err := db.Prepare("INSERT INTO tasks(id, name, status, `order`, timestamp) VALUES( ?, ?, ?, ?, ? )")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	stmt.Exec(task.ID, task.Name, task.Status, task.Order, time.Now())
	return c.JSON(http.StatusCreated, nil)
}

// UpdateTask PATCH /tasks/:id
func UpdateTask(c echo.Context) error {
	db := open()

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
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
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

// DeleteTask DELETE /tasks/:id
func DeleteTask(c echo.Context) error {
	db := open()

	// delete
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	stmt.Exec(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}

// DeleteTasks DELETE /tasks
func DeleteTasks(c echo.Context) error {
	db := open()

	// delete
	stmt, err := db.Prepare("DELETE FROM tasks WHERE status = 0")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	stmt.Exec()
	return c.JSON(http.StatusOK, nil)
}

// GetMaxID GET /id
func GetMaxID(c echo.Context) error {
	var id int32
	// FIXME: Get data with mysql
	id = 5
	return c.JSON(http.StatusOK, id)
}
