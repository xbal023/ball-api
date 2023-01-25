package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bolaxd/dumn/app/controlers"	
	)
	
func ApiRoute(r *fiber.App) {
	r.Get("/test", controlers.TestControler)
}