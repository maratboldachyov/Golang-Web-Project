package model

import "time"

type Transaction struct {
	Id           uint64    `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	SenderCard   string    `gorm:"type:text;unique;NOT NULL" json:"sender_card"`
	ReceiverCard string    `gorm:"type:text;unique;NOT NULL" json:"receiver_card"`
	Amount       uint64    `gorm:"type:bigint;NOT NULL; check:amount > 100" json:"amount"`
	ProvidedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"provided_at"`
}
