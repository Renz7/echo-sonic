package echosonic

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"net/http"
)

// JsoniterSerializer is JosnSerializer using github.com/json-iterator/go
type JsoniterSerializer struct {
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (j JsoniterSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (j JsoniterSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Deserialize error: %v", err)).SetInternal(err)

	}
	return err
}
