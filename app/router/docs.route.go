package router 

import (
	"github.com/gofiber/fiber/v2"
	)

func DocsRoute(r *fiber.App) {
	r.Get("/docs", func(s *fiber.Ctx) error {
		return s.JSON(fiber.Map{
			"Message": "Ini Docs API",
		})
	});
};