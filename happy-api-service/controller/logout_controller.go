package controller

import "github.com/gofiber/fiber/v2"

type LogoutController interface {
	Logout(*fiber.Ctx) error
}
