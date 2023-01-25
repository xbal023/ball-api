package app;

import (
	"fmt"
	"log"
	"time"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/goccy/go-json"
	
	// FROM LOCAL
	"github.com/bolaxd/dumn/app/router"
	"github.com/bolaxd/dumn/app/controlers"
	config "github.com/bolaxd/dumn/config"
	)
	
func Run(wg *sync.WaitGroup) {
	defer wg.Done()
	/* Dotenv variable */
	AppName := os.Getenv("APP_NAME")
	Header := os.Getenv("HEADER")
	PORT := os.Getenv("PORT")
	PORT2 := os.Getenv("PORT2")
	tmpl := html.New("./views", ".html")
	tmpl.Reload(true)
	// configCors := cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	// 	AllowCredentials: true,
	// }
	app := fiber.New(fiber.Config{
		Views:						tmpl,
		ViewsLayout:			"Main/main",
		Prefork:					false,
		CaseSensitive:		true,
		StrictRouting:		true,
		ServerHeader:			Header,
		AppName:					AppName,
		JSONEncoder:			json.Marshal,
		JSONDecoder:			json.Unmarshal,
	});
	/* DEBUG LOGS */
	app.Use(logger.New());
	/* CORS */
	app.Use(cors.New());
	app.Use(recover.New());
	app.Use(limiter.New(limiter.Config{
		Expiration:				config.ExpiredLimit * time.Second,
		Max:							config.MaxLimitRefresh,
	}));
	app.Use(cache.New(cache.Config{
		Next: 						func(c *fiber.Ctx) bool { return c.Query("refresh") == "true" },
		Expiration:				config.ExpiredCache * time.Second,
		CacheControl:			true,
	}));
	app.Static("/js", "./views/JS")
	// router
	app.Get("/", func(c *fiber.Ctx) error {
		return c.RedirectBack("/home")
	});
	app.Get("/monitor", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	router.HomeRoute(app);
	router.DocsRoute(app);
	router.ApiRoute(app);
	// if not found
	app.Use("*", controlers.ErrorControler);
	
	err := app.Listen(":" +PORT);
	if err != nil {
		fmt.Println("Mengganti port cadangan...")
		err2 := app.Listen(":" +PORT2)
		if err2 != nil {
			log.Print("Tolong Rubah Port anda, karena port sudah di blokir")
		}
	}
}