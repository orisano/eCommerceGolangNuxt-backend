package main

// go:generate sqlboiler --wipe postgres
import (
	"bongo/model"
	"bongo/myadmin"
	"bongo/myauth"
	"bongo/mynonuser"
	"bongo/myseller"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	model.InitDatabase()
	//db.Init()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Open handle to database like normal
	//bun.InitDatabase()
	//conn.Connect()
	// If you don't want to pass in db to all generated methods
	// you can use boil.SetDB to set it globally, and then use
	// the G variant methods like so (--add-global-variants to enable)
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "header:X-Csrf-Token",
	//	CookieName:     "csrf_bongo",
	//	CookieSameSite: "Strict",
	//}))

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Dhaka",
	}))
	//app.Use("/ws", func(c *fiber.Ctx) error {
	//	// IsWebSocketUpgrade returns true if the client
	//	// requested upgrade to the WebSocket protocol.
	//	if websocket.IsWebSocketUpgrade(c) {
	//		c.Locals("allowed", true)
	//		return c.Next()
	//	}
	//	return fiber.ErrUpgradeRequired
	//})
	myadmin.AdminSocketRoutes(app)
	myadmin.AdminRoutes(app)
	myauth.AuthRoutes(app)
	myseller.SellerRoutes(app)
	mynonuser.NonAuthRoutes(app)
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "header:X-Csrf-Token",
	//	CookieName:     "csrf_bongo",
	//	CookieSameSite: "Strict",
	//	Expiration:     24 * time.Hour,
	//	KeyGenerator:   utils.UUID,
	//}))

	app.Static("/static", "./public")
	os.Setenv("HOSTNAME", "http://localhost:8000")
	os.Setenv("WEBSOCKET_HOST", "ws://localhost:8000")

	log.Fatal(app.Listen(":8000"))
	//go run main.go
	//err := app.ListenTLS(":8000", "/etc/letsencrypt/live/www.alifnuryana.software/fullchain.pem", "/etc/letsencrypt/live/www.alifnuryana.software/privkey.pem")
	//if err != nil {
	//	fmt.Println("Cannot start server.")
	//}

}
func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
