package Products

import "gorm.io/gorm"

type Products []*Product

type Product struct {
	gorm.Model
	ProductId   string `gorm:"column:product_id;primaryKey;type:varchar(255);not null"`
	MonsterId   string `gorm:"column:monster_id;type:varchar(10);not null"`
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	PublishYear string `gorm:"column:publish_year;type:varchar(20)"`
	TotalSales  string `gorm:"column:total_sales;type:varchar(255)"`
}
