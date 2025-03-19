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
	// notifications
	newRoute(http.MethodPut, "/v1/payment/pay", controllers.PaymentPayV1),
	newRoute(http.MethodPut, "/v1/payment/rollback", controllers.RollbackPaymentV1),
	newRoute(http.MethodPut, "/v1/payment/deposit", controllers.DepositV1),
}
