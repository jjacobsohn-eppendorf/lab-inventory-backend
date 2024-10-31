package main

import (
	"applicant-backend/internal/controller/devicesController"
	"applicant-backend/internal/helper"
	"applicant-backend/internal/internal"
	"applicant-backend/internal/services/devices_service"
	"log"
)

func main() {
	Config, err := helper.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := internal.InitializeRouter()
	devicesService := devices_service.NewDeviceService()

	_ = devicesController.NewDeviceController(devicesService, r)

	err = r.Run("0.0.0.0:" + Config.Port)
	if err != nil {
		log.Fatalf("Encountered Gin error: %d", err)
		return
	}
}
