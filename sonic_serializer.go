//go:build amd64

package echosonic

import (
	"net/http"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/labstack/echo/v4"
)

type SonicJSONSerializer struct {
}

func (s SonicJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := encoder.NewStreamEncoder(c.Response())
	enc.SetIndent("", indent)
	return enc.Encode(i)
}

func (s SonicJSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := decoder.NewStreamDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*decoder.MismatchTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, ute.Description()).SetInternal(err)
	} else if se, ok := err.(*decoder.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, se.Description()).SetInternal(err)
	}
	return err
}
