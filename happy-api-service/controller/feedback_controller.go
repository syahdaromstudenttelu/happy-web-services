package controller

import "github.com/gofiber/fiber/v2"

type FeedbackController interface {
	Create(*fiber.Ctx) error
	FindByProductId(*fiber.Ctx) error
}
