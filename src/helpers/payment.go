package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"payment-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
	"github.com/IvanSkripnikov/go-logger"
	"gorm.io/gorm"
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
		response = models.Failure
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
		response = models.Failure
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

	var newDeposit models.Payment
	var uniquePayment models.UniquePayment
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Where("request_id = ?", paymentParams.RequestID).First(&newDeposit).Error

	// такого депозита раньше не было, создаём новый, иначе возвращаем его результат
	response := models.Success
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Производим начисление средств через сервис платежей
		response := models.Success
		newDepositObject := models.Account{UserID: paymentParams.UserID, Balance: paymentParams.Amount}
		newDepositResponse, err := CreateQueryWithScalarResponse(http.MethodPut, Config.BillingServiceUrl+"/v1/account/deposit", newDepositObject)
		if checkError(w, err, category) || newDepositResponse != models.Success {
			response = models.Failure
		}

		newDeposit := models.Payment{UserID: paymentParams.UserID, Type: models.TypeDeposit, Amount: paymentParams.Amount, Created: int(GetCurrentTimestamp()), RequestID: paymentParams.RequestID}

		if response != models.Success {
			response = models.Failure
		} else {
			newDeposit.Status = 1
		}

		// записываем сообщение в БД
		err = db.Create(&newDeposit).Error
		if err != nil {
			logger.Errorf("Cant create deposit %v", err)
			response = models.Failure
		}

		// создаём запись с уникальным ID от пользователя
		uniquePayment.RequestID = paymentParams.RequestID
		uniquePayment.Response = response
		err = db.Create(&uniquePayment).Error
		if err != nil {
			logger.Errorf("Cant create unique payment record %v", err)
			response = models.Failure
		}
	} else {
		err = db.Where("request_id = ?", paymentParams.RequestID).First(&uniquePayment).Error
		if checkError(w, err, category) {
			return
		}
		response = uniquePayment.Response
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetDepositsByUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/payment/get-deposits-by-user"
	var deposits []models.Payment

	userID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Where("user_id = ? AND type = ?", userID, models.TypeDeposit).Find(&deposits).Error
	if checkError(w, err, category) && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	data := ResponseData{
		"response": deposits,
	}
	SendResponse(w, data, category, http.StatusOK)
}
