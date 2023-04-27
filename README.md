# echo-sonic

[Echo](https://echo.labstack.com/) framework `JSONSerializer` [sonic](https://github.com/bytedance/sonic) implement.

## Requirement

amd64 supported only,cause depending on sonic.

- Go 1.15~1.20
- Linux/MacOS/Windows
- Amd64 ARCH

## Installing

```shell
go get github.com/renz7/echo-sonic
```

## Use in echo project

```go
package main

import (
    "github.com/labstack/echo/v4
)


e := echo.New()
e.JSONSerializer = &SonicJSONSerializer{}

```
