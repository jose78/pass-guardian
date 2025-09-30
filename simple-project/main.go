package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var logger echo.Logger

func main() {
	server := echo.New()
	server.GET("/env/:envVar", handlerEventVar)
	logger = server.Logger
	logger.Debug(server.Start(":1323"))
}

func handlerEventVar(ctx echo.Context) error {
	nameEnvVar := ctx.Param("envVar")
	valueEnvVar := os.Getenv(nameEnvVar)


    data := map[string]string{nameEnvVar: valueEnvVar}
	statusCodeReturn := http.StatusOK
	if valueEnvVar == "" {
		statusCodeReturn = http.StatusBadRequest
		data = map[string]string{"error": fmt.Sprintf("Not found env var %s", nameEnvVar)}
	}
	return ctx.JSON(statusCodeReturn, data)
}