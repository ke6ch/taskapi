package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ke6ch/api/model"
	"github.com/labstack/echo/v4"
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
	defer db.Close()

	// データ取得
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	// insert
	result, err := db.Exec("INSERT INTO tasks(id, name, status, `order`, timestamp) VALUES( ?, ?, ?, ?, ? )", task.ID, task.Name, task.Status, task.Order, time.Now())
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, nil)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// select
	tasks := []*model.Task{}
	var Status []uint8

	if err := db.QueryRow("SELECT * FROM tasks where id = "+strconv.FormatInt(lastInsertID, 10)).Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	if Status[0] == 0 {
		task.Status = false
	} else {
		task.Status = true
	}
	tasks = append(tasks, task)

	return c.JSON(http.StatusOK, tasks)

}

// UpdateTask PATCH /tasks/:id
func UpdateTask(c echo.Context) error {
	db := open()
	defer db.Close()

	// データ取得
	id := c.Param("id")
	var status uint8
	if c.QueryParam("status") == "false" {
		status = 0
	} else {
		status = 1
	}

	// update
	result, err := db.Exec("UPDATE tasks SET status = ? where id = ?", status, id)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	fmt.Println(rowsAffected)

	// select
	task := new(model.Task)
	tasks := []*model.Task{}
	var Status []uint8

	if err := db.QueryRow("SELECT * FROM tasks where id = "+id).Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	if Status[0] == 0 {
		task.Status = false
	} else {
		task.Status = true
	}
	tasks = append(tasks, task)

	return c.JSON(http.StatusOK, tasks)
}

// DeleteTask DELETE /tasks/:id
func DeleteTask(c echo.Context) error {
	db := open()
	defer db.Close()

	id := c.Param("id")

	// delete
	result, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	fmt.Println(rowsAffected)

	m := model.Message{}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}
	m.Message = "ID : " + id + " is deleted"
	return c.JSON(http.StatusOK, m)
}

// DeleteTasks DELETE /tasks
func DeleteTasks(c echo.Context) error {
	db := open()
	defer db.Close()

	// delete
	result, err := db.Exec("DELETE FROM tasks WHERE status = 0")
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	fmt.Println(rowsAffected)

	m := model.Message{}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}
	m.Message = "Tasks is deleted"
	return c.JSON(http.StatusOK, m)
}

// GetMaxID GET /id
func GetMaxID(c echo.Context) error {
	db := open()
	defer db.Close()

	var id int32
	if err := db.QueryRow("SELECT max(id) FROM tasks").Scan(&id); err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}
	return c.JSON(http.StatusOK, id)
}
