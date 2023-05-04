package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Auth(*fiber.Ctx) error
}
