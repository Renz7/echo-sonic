package echosonic

import (
	"github.com/labstack/echo/v4"
	"github.com/renz7/echo-sonic/util"
	"testing"
)

var e = echo.New()

func TestJsoniterSerializer_Deserialize(t *testing.T) {
	type args struct {
		c echo.Context
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				c: util.NewTestContextWithBody(e, "{\"str\":\"test short str\"}"),
				i: &struct {
					Str string
				}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsoniterSerializer{}
			if err := j.Deserialize(tt.args.c, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsoniterSerializer_Serialize(t *testing.T) {
	type args struct {
		c      echo.Context
		i      interface{}
		indent string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				c: util.NewTestContext(e),
				i: &struct {
					Str string
					I32 int32
				}{
					"short test str",
					123,
				},
				indent: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsoniterSerializer{}
			if err := j.Serialize(tt.args.c, tt.args.i, tt.args.indent); (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.c.Response().Size < 0 {
				t.Errorf("Serialize() error = 0 size body")
			}
		})
	}
}
