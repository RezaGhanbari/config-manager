package main

import (
	"config-manager/controller"
	"config-manager/domain"
	"config-manager/muxinator"
	"config-manager/service"
	"log"
	"time"
)

func main() {
	config := domain.Config{}

	configService := service.ConfigService{
		Config:   &config,
		Location: "config.yaml",
	}

	go configService.Watch(time.Second * 30)
	c := controller.Controller{
		Config: &config,
	}

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", c.ReadConfig)

	log.Fatal(router.ListenAndServe(":8080"))
}
