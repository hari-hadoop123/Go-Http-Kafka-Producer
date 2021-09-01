package main

import (
	"go-projects/http-kafka-producer/app"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type userValidator struct {
	validator *validator.Validate
}

func (cv *userValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &userValidator{validator: validator.New()}

	e.POST("/test", app.PostTestMessage)

	e.Logger.Fatal(e.Start(":" + "1323"))
}
