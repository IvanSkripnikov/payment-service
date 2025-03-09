package models

type PaymentParams struct {
	UserID int     `gorm:"index;type:int" json:"userId"`
	Amount float32 `gorm:"type:float" json:"amount"`
}
