package lib

import (
	"github.com/labstack/echo"
)

// ServerHeader 共通ヘッダーを設定する
// Content-Type: application/json を設定する
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(c)
	}
}
