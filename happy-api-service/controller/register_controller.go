package controller

import "github.com/gofiber/fiber/v2"

type RegisterController interface {
	Register(*fiber.Ctx) error
}
