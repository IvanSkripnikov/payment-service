package httphandler

import (
	"net/http"
	"regexp"

	"payment-service/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// payment
	newRoute(http.MethodPut, "/v1/payment/pay", controllers.PaymentPayV1),
	newRoute(http.MethodPut, "/v1/payment/rollback", controllers.RollbackPaymentV1),
	newRoute(http.MethodPut, "/v1/payment/deposit", controllers.DepositV1),
	newRoute(http.MethodGet, "/v1/payment/get-deposits-by-user/([0-9]+)", controllers.GetDepositsByUserV1),
}
