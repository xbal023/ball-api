package router 

import (
	"github.com/gofiber/fiber/v2"
	)

func DonateRoute(r *fiber.App) {
	r.Get("/donate", func(s *fiber.Ctx) error {
		// s.Download()
		return s.JSON(fiber.Map{
			"Message": "Ini Donate api",
		})
	});
};