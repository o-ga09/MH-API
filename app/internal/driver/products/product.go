package products

import (
	"context"
	Products "mh-api/app/internal/domain/products"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type productRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *productRepository {
	return &productRepository{
		conn: conn,
	}
}

func (r *productRepository) Get(ctx context.Context, monsterId string) (*Products.Products, error) {
	monster := []mysql.Product{}
	err := r.conn.Find(&monster).Error
	if err != nil {
		return nil, err
	}

	res := Products.Products{}
	for _, r := range monster {
		res = append(res, *Products.NewProduct(r.ProductId, r.Name, r.PublishYear, r.TotalSales))
	}

	return &res, nil
}

func (r *productRepository) Save(ctx context.Context, p Products.Product) error {
	product := mysql.Product{
		MonsterId:   p.GetID(),
		Name:        p.GetName(),
		PublishYear: p.GetYear(),
		TotalSales:  p.GetSales(),
	}
	err := r.conn.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Remove(ctx context.Context, productId string) error {
	product := mysql.Product{
		ProductId: productId,
	}
	err := r.conn.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
