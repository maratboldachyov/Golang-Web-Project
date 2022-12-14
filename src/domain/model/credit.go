package model

type Credit struct {
	Id        uint64  `gorm:"primary_key;autoincrement;NOT NULL;" json:"id"`
	UserID    uint    `gorm:"index:unique_review;" json:"user_id"`
	ProductID uint    `gorm:"index:unique_review;" json:"product_id"`
	Product   Product `gorm:"foreign_key:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
	User      User    `gorm:"foreign_key:SubcategoryUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
}
