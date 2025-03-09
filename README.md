## Payment Service Overview 

The repository of payment service

## Endpoints

Method | Path                             | Description                                   |                                                                         
---    |----------------------------------|------------------------------------------------
GET    | `/health`                        | Health page                                   |
GET    | `/metrics`                       | Страница с метриками                          |
PUT    | `/v1/payment/pay`                | Оплата заказа                                 |
PUT    | `/v1/payment/rollback`           | Возврат средств                               |