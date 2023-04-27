//go:build amd64

package echosonic

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	Name  string
	Int64 int64
}

func NewContext(e *echo.Echo) echo.Context {
	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{\"name\":\"test\"}"))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(request, recorder)
}

func TestSonicJSONSerializer_Serialize(t *testing.T) {
	e := echo.New()

	type args struct {
		c      echo.Context
		i      interface{}
		indent string
	}
	tests := []struct {
		name    string
		s       SonicJSONSerializer
		args    args
		wantErr bool
	}{
		{
			name: "test case 0",
			s:    SonicJSONSerializer{},
			args: args{
				c: NewContext(e),
				i: TestStruct{
					Name:  "test_case_0",
					Int64: int64(1),
				},
				indent: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SonicJSONSerializer{}
			if err := s.Serialize(tt.args.c, tt.args.i, tt.args.indent); (err != nil) != tt.wantErr {
				t.Errorf("SonicJSONSerializer.Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
