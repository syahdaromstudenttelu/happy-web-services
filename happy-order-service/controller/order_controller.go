package controller

import "github.com/gofiber/fiber/v2"

type OrderController interface {
	Create(*fiber.Ctx) error
	UpdateStatusPayment(*fiber.Ctx) error
	UpdateFeedbackDone(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	FindByIdOrder(*fiber.Ctx) error
	FindByIdUserAndIdOrder(*fiber.Ctx) error
	FindAllByIdUser(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
