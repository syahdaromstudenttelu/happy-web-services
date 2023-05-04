package controller

import "github.com/gofiber/fiber/v2"

type FeedbackController interface {
	Create(*fiber.Ctx) error
	FindByIdProduct(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
