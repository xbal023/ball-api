package controlers 

import (
	"github.com/gofiber/fiber/v2"
	)

func TestControler(s *fiber.Ctx) error {
	return s.JSON(fiber.Map{
		"Code": fiber.StatusOK,
		"Data": "ini aku",
	});
};