package main

import (
	"payment-service/helpers"
	"payment-service/httphandler"
	"payment-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatalf("Config error: %v", err)
	}

	// сделать конфиг глобальным для хэлпера
	helpers.InitConfig(config)

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
