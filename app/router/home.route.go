package router 

import (
	"github.com/gofiber/fiber/v2"
	)

func HomeRoute(r *fiber.App) {
	r.Get("/home", func(s *fiber.Ctx) error {
		return s.Render("home", nil)
	});
};