package products

import (
	"context"
	Products "mh-api/app/internal/domain/products"
	"mh-api/app/internal/driver/mysql"
)

type productRepository struct{}

func NewMonsterRepository() *productRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, p Products.Product) error {
	product := mysql.Product{
		MonsterId:   p.GetID(),
		Name:        p.GetName(),
		PublishYear: p.GetYear(),
		TotalSales:  p.GetSales(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&product).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Remove(ctx context.Context, productId string) error {
	product := mysql.Product{
		ProductId: productId,
	}
	err := mysql.CtxFromDB(ctx).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
