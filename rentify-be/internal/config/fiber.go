package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewFiberApp() *fiber.App {
	fiberConfig := fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "rentify",
		ErrorHandler:  errorHandler(),
	}

	app := fiber.New(fiberConfig)
	app.Use(requestid.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowCredentials: true,
	// }))

	return app
}

func errorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
