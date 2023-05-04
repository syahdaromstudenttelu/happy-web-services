package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	Update(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
