package handler

import (
	"fmt"
	"net/http"

	"github.com/ke6ch/api/lib"
	"github.com/ke6ch/api/model"
	"github.com/labstack/echo"
)

func init() {
}

// Login ログインページ
func Login(c echo.Context) error {
	// CookieのsessionIDを取得する
	cookie, err := lib.ReadCookie(c)
	if err != nil {
		return c.JSON(http.StatusOK, nil)
	}

	// Sessionが存在するかチェックする
	if err := lib.GetSession(cookie.Value); err != nil {

		// Sessionが存在しない場合は、リダイレクトしない
		return c.JSON(http.StatusOK, nil)
	}
	// Sessionが存在する場合は、リダイレクトする
	return c.JSON(http.StatusMovedPermanently, nil)
}

// Session ログイン処理
func Session(c echo.Context) error {
	// POSTデータを取得
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		fmt.Println("No POST Data")
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	// uuidを作成する
	uuid := lib.CreateUUID()

	// Session開始
	if err := lib.SetSession(uuid, u.Email); err != nil {
		fmt.Println("Error Set Session")
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	// Cookieにsessionを追加する
	lib.WriteCookie(c, uuid)

	m := new(model.Payload)
	m.Message = "Success"

	return c.JSON(http.StatusOK, m)
}
