package helpers

import (
	"encoding/json"
	"net/http"

	"payment-service/models"
)

func PayPayment(w http.ResponseWriter, r *http.Request) {
	category := "/v1/payment/pay"

	// получаем входные параметры
	var paymentParams models.PaymentParams
	err := json.NewDecoder(r.Body).Decode(&paymentParams)
	if checkError(w, err, category) {
		return
	}

	newAccount := models.Account{UserID: paymentParams.UserID, Balance: paymentParams.Amount}
	response, err := CreateQueryWithScalarResponse(http.MethodPut, Config.BillingServiceUrl+"/v1/account/buy", newAccount)
	if checkError(w, err, category) {
		return
	}

	if response != "success" {
		response = "failure"
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func RollbackPayment(w http.ResponseWriter, r *http.Request) {
	category := "/v1/payment/rollback"
	// получаем входные параметры
	var paymentParams models.PaymentParams
	err := json.NewDecoder(r.Body).Decode(&paymentParams)
	if checkError(w, err, category) {
		return
	}

	newAccount := models.Account{UserID: paymentParams.UserID, Balance: paymentParams.Amount}
	response, err := CreateQueryWithScalarResponse(http.MethodPut, Config.BillingServiceUrl+"/v1/account/deposit", newAccount)
	if checkError(w, err, category) {
		return
	}

	if response != "success" {
		response = "failure"
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}
