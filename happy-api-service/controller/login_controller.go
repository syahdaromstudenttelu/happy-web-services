package controller

import "github.com/gofiber/fiber/v2"

type LoginController interface {
	Login(*fiber.Ctx) error
}
