package model

type Subcategory struct {
	Id         uint64   `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	Name       string   `gorm:"type:text" json:"name"`
	CategoryID uint     `gorm:"default:NULL" json:"category_id"`
	Category   Category `gorm:"foreign_key:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
}
