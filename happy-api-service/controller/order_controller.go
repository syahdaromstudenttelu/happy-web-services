package controller

import "github.com/gofiber/fiber/v2"

type OrderController interface {
	Create(*fiber.Ctx) error
	UpdateByStatusPayment(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	FindByUserId(*fiber.Ctx) error
	FindByUserAndOrderId(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}
