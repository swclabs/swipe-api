// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"context"
	"log"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"

	"gorm.io/gorm"
)

type Products struct {
	data *domain.Products
	conn *gorm.DB
}

func NewProducts() domain.IProductRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Products{
		data: &domain.Products{},
		conn: _conn,
	}
}

func (product *Products) Insert(ctx context.Context, prd *domain.Products) error {
	return db.SafeWriteQuery(
		ctx,
		product.conn,
		queries.InsertIntoProducts,
		prd.Image, prd.Name, prd.Description, prd.Available, prd.SupplierID, prd.CategoryID,
	)
}

func (product *Products) GetLitmit(ctx context.Context, limit int) ([]domain.Products, error) {
	var products []domain.Products
	if err := product.conn.Table(domain.ProductsTable).Find(&products).Limit(limit).Error; err != nil {
		return nil, err
	}
	return products, nil
}
