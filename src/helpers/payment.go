package helpers

import (
	"encoding/json"
	"net/http"

	"payment-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
	"github.com/IvanSkripnikov/go-logger"
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

	newPayment := models.Payment{UserID: paymentParams.UserID, Type: models.TypePayment, Amount: paymentParams.Amount, Created: int(GetCurrentTimestamp())}

	if response != models.Success {
		response = "failure"
	} else {
		newPayment.Status = 1
	}

	// записываем сообщение в БД
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Create(&newPayment).Error
	if err != nil {
		logger.Errorf("Cant create payment %v", err)
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

	newDeposit := models.Payment{UserID: paymentParams.UserID, Type: models.TypeDeposit, Amount: paymentParams.Amount, Created: int(GetCurrentTimestamp())}

	if response != models.Success {
		response = "failure"
	} else {
		newDeposit.Status = 1
	}

	// записываем сообщение в БД
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Create(&newDeposit).Error
	if err != nil {
		logger.Errorf("Cant create deposit %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	category := "/v1/payment/deposit"
	// получаем входные параметры
	var paymentParams models.PaymentParams
	err := json.NewDecoder(r.Body).Decode(&paymentParams)
	if checkError(w, err, category) {
		return
	}

	// Производим начисление средств через сервис платежей
	response := models.Success
	newDepositObject := models.Account{UserID: paymentParams.UserID, Balance: paymentParams.Amount}
	newDepositResponse, err := CreateQueryWithScalarResponse(http.MethodPut, Config.BillingServiceUrl+"/v1/account/deposit", newDepositObject)
	if checkError(w, err, category) || newDepositResponse != models.Success {
		response = models.Failure
	}

	newDeposit := models.Payment{UserID: paymentParams.UserID, Type: models.TypeDeposit, Amount: paymentParams.Amount, Created: int(GetCurrentTimestamp())}

	if response != models.Success {
		response = "failure"
	} else {
		newDeposit.Status = 1
	}

	// записываем сообщение в БД
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Create(&newDeposit).Error
	if err != nil {
		logger.Errorf("Cant create deposit %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}
