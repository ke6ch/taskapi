package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ke6ch/api/model"
	"github.com/labstack/echo"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "user:pass@tcp(db:3306)/clear")
	if err != nil {
		panic(err.Error())
	}
}

// GetTasks GET /tasks
func GetTasks(c echo.Context) error {
	rows, err := db.Query("SELECT * FROM tasks order by status desc, `order`")
	// defer db.Close()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	tasks := []*model.Task{}
	var Status []uint8

	for rows.Next() {
		task := new(model.Task)
		err := rows.Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusServiceUnavailable, nil)
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
	task := new(model.Task)
	tasks := []*model.Task{}
	var Status []uint8

	if err := db.QueryRow("SELECT * FROM tasks where id = "+c.Param("id")).Scan(&task.ID, &task.Name, &Status, &task.Order, &task.Timestamp); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}
	defer db.Close()

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
	// データ取得
	task := new(model.Task)
	if err = c.Bind(task); err != nil {
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
	if err := db.QueryRow("SELECT max(id) FROM tasks").Scan(&id); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, id)
}
