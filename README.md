## Payment Service Overview 

The repository of payment service

## Endpoints

Method | Path                                        | Description                                   |                                                                         
---    |---------------------------------------------|------------------------------------------------
GET    | `/health`                                   | Health page                                   |
GET    | `/metrics`                                  | Страница с метриками                          |
PUT    | `/v1/payment/pay`                           | Оплата заказа                                 |
PUT    | `/v1/payment/rollback`                      | Возврат средств                               |
PUT    | `/v1/payment/deposit`                       | Внесение депозита                             |
GET    | `/v1/payment/get-deposits-by-user/{userId}` | Получение всех депозитов пользователя         |