package main

import (
	"log"
	"nidus-server/internal/config"
	"nidus-server/internal/routes"
	"nidus-server/pkg/infrastructure"
	"nidus-server/pkg/repository"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
Nidus API
https://docs.gofiber.io/
*/
func main() {

	// Environments variables
	log.Println("[Config] Checking configuration ...")
	err := config.CheckConfig()
	if err != nil {
		log.Fatal("[CheckConfig] ", err)
	}

	// Database
	log.Println("[Database] Initializing ...")
	db, err := infrastructure.MongoDBConnection()
	if err != nil {
		log.Fatal("[MongoDB] ", err)
	}

	// Fiber Application
	// https://docs.gofiber.io/api/fiber#new
	// https://docs.gofiber.io/api/fiber#config
	log.Println("[App] Initializing ...")
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Nidus API v1.0.1",
	})

	// Cors
	app.Use(cors.New())

	// Logging
	app.Use("/", logger.New(logger.Config{
		Format:     "${cyan}[${time}] [${ip}]:${port} ${white}${pid} ${red}${status} - ${blue}${method} ${white}${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Europe/Brussels",
	}))

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

	// TODO: create new route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nidus API")
	})

	// Start the server
	err = app.Listen(":3333")
	if err != nil {
		log.Fatal("[App] ", err)
	}

	// MQTT Client

}
