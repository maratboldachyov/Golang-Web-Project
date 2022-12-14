package model

type User struct {
	Id           uint64 `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	Login        string `gorm:"type:varchar;unique;NOT NULL" json:"login"`
	Password     string `gorm:"NOT NULL;" json:"password"`
	CardNumber   string `gorm:"type:text;unique;NOT NULL" json:"card_number"`
	Photo        string `gorm:"type:text;default:NULL" json:"photo"`
	CardBalance  uint64 `gorm:"type:bigint;default:0" json:"card_balance"`
	BonusBalance uint64 `gorm:"type:bigint;default:0" json:"bonus_balance"`
	Credits      uint16 `gorm:"type: smallint;default:0" json:"credits"`
}
