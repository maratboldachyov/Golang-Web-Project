package model

type Review struct {
	UserID    uint    `gorm:"type:bigint" json:"user_id"`
	ProductID uint    `gorm:"type:bigint" json:"product_id"`
	Product   Product `gorm:"foreign_key:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
	User      User    `gorm:"foreign_key:SubcategoryUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
	Content   string  `gorm:"type:text" json:"content"`
	Rating    uint    `gorm:"type:smallint" json:"rating"`
}
