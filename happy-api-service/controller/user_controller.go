package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindByUserName(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
