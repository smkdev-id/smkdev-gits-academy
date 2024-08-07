package models

type Product struct{
	Id int64 `gorm:"PrimaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"name"`
	ProductDescription string `gorm:"type:text" json:"description"`
	ProductPrice string `gorm:"type:varchar(300)" json:"price"`
	ProductCategory string `gorm:"type:varchar(300)" json:"category"`
}