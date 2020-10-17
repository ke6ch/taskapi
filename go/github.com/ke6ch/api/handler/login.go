package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/ke6ch/api/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// func newRediStore() *redisStore.RediStore {
// 	store, err := redisStore.NewRediStore(10, "tcp", "localhost:6379", "", []byte("secret-key"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return store
// }

// Login ログインページ
func Login(c echo.Context) error {
	// loginチェック
	cookie, err := c.Cookie("logged_in")
	if err != nil {
		// Cookieが見つからない場合
		if err == http.ErrNoCookie {
			cookie := new(http.Cookie)
			cookie.Name = "logged_in"
			cookie.Value = "no"
			cookie.Path = "/"
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)
			return c.JSON(http.StatusOK, nil)
		}
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// loginしている場合、ページ遷移する
	if cookie.Value == "yes" {
		// 有効期限を更新
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		sess, err := session.Get("user-session", c)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusServiceUnavailable, nil)
		}
		//ログインしているか
		if b, _ := sess.Values["authenticated"]; b != true {
			return c.String(http.StatusUnauthorized, "401")
		}
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// Session ログイン処理
func Session(c echo.Context) error {
	// redisStore
	// store := newRediStore()
	// defer store.Close()

	// Get a session.
	sess, err := session.Get("user-session", c)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	// 認証処理
	auth := new(model.Auth)
	if err := c.Bind(auth); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	db := open()
	defer db.Close()

	var count int32
	if err := db.QueryRow("SELECT count(email) FROM users where email = '" + auth.Email + "' and `password` = '" + auth.Password + "'").Scan(&count); err != nil {
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	// ユーザが存在しない場合
	if count == 0 {
		m := model.Message{Message: "ユーザが存在しません。"}
		return c.JSON(http.StatusUnauthorized, m)
	}

	// logged_in cookie設定
	cookie := new(http.Cookie)
	cookie.Name = "logged_in"
	cookie.Value = "yes"
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	// ユーザーを認証済みに設定する。
	sess.Values["authenticated"] = true

	// Save.
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		fmt.Println("Error saving session: %v", err)
		return c.JSON(http.StatusServiceUnavailable, err)
	}

	return c.JSON(http.StatusOK, nil)
}
