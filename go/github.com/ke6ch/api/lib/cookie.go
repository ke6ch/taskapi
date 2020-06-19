package lib

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var name = "_login_session"

// ReadCookie read cookie
func ReadCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		fmt.Println("No Cookie")
		return nil, err
	}
	return cookie, nil
}

// WriteCookie write cookie
func WriteCookie(c echo.Context, sessionID string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = sessionID
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}
