package controllers

import (
	"net/http"

	"payment-service/helpers"
)

func PaymentPayV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.PayPayment(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/payment/pay")
	}
}

func RollbackPaymentV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.RollbackPayment(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/payment/rollback")
	}
}
