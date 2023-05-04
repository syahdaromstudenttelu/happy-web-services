package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Create(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	FindByUserName(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
