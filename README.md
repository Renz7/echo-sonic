# echo-sonic

[Echo](https://echo.labstack.com/) framework `JSONSerializer` [sonic](https://github.com/bytedance/sonic) and [jsoniter](https://github.com/json-iterator/go) implement.

## Requirement

amd64 supported only if sonic.

- Go 1.15~1.20
- Linux/MacOS/Windows
- Amd64 ARCH(sonic)
- Amd64 Arm Darwin(jsoniter)

## Installing

```shell
go get github.com/renz7/echo-sonic
```

## Use in echo project

```go
package echosonic

import (
	"github.com/labstack/echo/v4"
	"github.com/renz7/echo-sonic"
)

func main() {
	e := echo.New()
	e.JSONSerializer = &JsoniterSerializer{}
	//e.JSONSerializer = &SonicJSONSerializer{}

}

```
