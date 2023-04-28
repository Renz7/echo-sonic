package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
)

func NewTestContext(e *echo.Echo) echo.Context {
	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(request, recorder)
}

func NewTestContextWithBody(e *echo.Echo, body string) echo.Context {
	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(request, recorder)
}
