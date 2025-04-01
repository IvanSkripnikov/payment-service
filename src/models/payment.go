package models

const ServiceDatabase = "NotificationService"
const TypeDeposit = "deposit"
const TypePayment = "payment"

type Deposit struct {
	Amount float32 `json:"amount"`
}

type PaymentParams struct {
	UserID    int     `gorm:"index;type:int" json:"userId"`
	Amount    float32 `gorm:"type:float" json:"amount"`
	RequestID string  `gorm:"index;type:string" json:"requestId"`
}

type Payment struct {
	ID        int     `gorm:"index;type:int" json:"id"`
	UserID    int     `gorm:"index;type:int" json:"userId"`
	Type      string  `gorm:"type:string" json:"type"`
	Amount    float32 `gorm:"type:float" json:"amount"`
	Created   string  `gorm:"index;type:string" json:"created"`
	Status    int8    `gorm:"index;type:tinyint;default:0" json:"status"`
	RequestID string  `gorm:"index;type:string" json:"requestId"`
}

func (s Payment) TableName() string { return "payments" }

type UniquePayment struct {
	ID        int    `gorm:"index;type:int" json:"id"`
	RequestID string `gorm:"index;type:string" json:"requestId"`
	Response  string `gorm:"type:string" json:"response"`
}

func (s UniquePayment) TableName() string { return "unique_payments" }
