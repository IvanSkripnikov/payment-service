package models

import (
	"os"
)

type Config struct {
	BillingServiceUrl string
}

func LoadConfig() (*Config, error) {
	return &Config{
		BillingServiceUrl: os.Getenv("BILLING_SERViCE_URL"),
	}, nil
}

func GetRequiredVariables() []string {
	return []string{
		// Url сервиса платежей
		"BILLING_SERViCE_URL",
	}
}
