package main

import (
	"log"
	"nidus-server/internal/routes"
	"nidus-server/pkg/infrastructure"
	"nidus-server/pkg/repository"
	"nidus-server/pkg/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	// Environments variables
	// FIXME: Improve dev/build workflow with environment variables
	// With Docker environment variables are loaded into the container.
	// With Go we load the local .env file
	_, exist := os.LookupEnv("DATABASE_URI")
	if !exist {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Fiber Application
	app := fiber.New()

	// Cors
	app.Use(cors.New())

	// Logging
	app.Use("/", logger.New(logger.Config{
		Format:     "${cyan}[${time}] [${ip}]:${port} ${white}${pid} ${red}${status} - ${blue}${method} ${white}${path}\n",
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "Europe/Brussels",
	}))

	// Database
	db, err := infrastructure.MongoDBConnection()
	if err != nil {
		log.Fatal("Database connection Error $s", err)
	}

	// Collections
	userCollection := db.Collection("users")
	deviceCollection := db.Collection("devices")
	zoneCollection := db.Collection("zones")
	capabilityCollection := db.Collection("capabilities")
	temperaturesCollection := db.Collection("readings.temperature")
	humidityCollection := db.Collection("readings.humidity")
	pressureCollection := db.Collection("readings.pressure")
	// settingsCollection := db.Collection("settings")

	// Repositories
	userRepo := repository.NewUserRepo(userCollection)
	deviceRepo := repository.NewDeviceRepo(deviceCollection)
	zoneRepo := repository.NewZoneRepo(zoneCollection)
	capabilityRepo := repository.NewCapabilityRepo(capabilityCollection)
	measurementRepo := repository.NewMeasurementRepo(temperaturesCollection,
		humidityCollection,
		pressureCollection)
	// settingsRepository := repository.NewSettingsRepo(settingsCollection)

	// Services
	userService := service.NewUserService(userRepo)
	deviceService := service.NewDeviceService(deviceRepo)
	zoneService := service.NewZoneService(zoneRepo)
	capabilityService := service.NewCapabilityService(capabilityRepo)
	measurementService := service.NewMeasurementService(measurementRepo)
	networkService := service.NewNetworkService()
	// settingsService := service.NewSettingsService()

	// Routes
	api := app.Group("/")
	routes.UserRouter(api, userService)
	routes.DeviceRouter(api, deviceService)
	routes.ZoneRouter(api, zoneService)
	routes.CapabilityRouter(api, capabilityService)
	routes.MeasurementRouter(api, measurementService)
	routes.NetworkRouter(api, networkService)
	// routes.SettingsRouter(api, settingsService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nidus API v0.0.1")
	})

	// Start the server
	log.Fatal(app.Listen(":3333"))
}
