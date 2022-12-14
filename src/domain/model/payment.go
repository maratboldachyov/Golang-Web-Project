package model

import "time"

type Payment struct {
	Id         uint64    `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	Card       string    `gorm:"type:text;unique;NOT NULL" json:"card"`
	Amount     uint64    `gorm:"type:bigint;NOT NULL;" json:"amount"`
	ProvidedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"provided_at"`
}
