package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"payment-service/models"
	"time"

	logger "github.com/IvanSkripnikov/go-logger"
)

var Config *models.Config

func InitConfig(cfg *models.Config) {
	Config = cfg
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatResponse(w http.ResponseWriter, httpStatus int, category string) {
	w.WriteHeader(httpStatus)

	data := ResponseData{
		"error": "Unsuccessfull request",
	}
	SendResponse(w, data, category, httpStatus)
}

func CreateQueryWithScalarResponse(method, url string, data interface{}) (string, error) {
	var err error
	var response string

	jsonData, err := json.Marshal(data)
	if err != nil {
		return response, err
	}
	logger.Infof("json data: %v", string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return response, err
	}

	resp, err := client.Do(req)
	logger.Infof("response for make deposit: %v", resp.Body)
	if err != nil {
		return response, err
	}

	var result map[string]string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &result)

	logger.Infof("Data from create payment %v", result)

	// Преобразуем JSON-строку в map
	if err != nil {
		return response, err
	}

	response, ok := result["response"]
	if !ok {
		return "", errors.New("failed to create payment")
	}

	return response, nil
}
