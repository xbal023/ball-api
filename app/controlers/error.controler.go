package controlers 

import (
	"github.com/gofiber/fiber/v2"
	)

func ErrorControler(s *fiber.Ctx) error {
	return s.Render("error", nil)
};