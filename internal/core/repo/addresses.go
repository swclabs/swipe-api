// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"context"
	"errors"
	"log"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"

	"gorm.io/gorm"
)

type Addresses struct {
	data *domain.Addresses
	conn *gorm.DB
}

func NewAddresses() domain.IAddressRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Addresses{
		data: &domain.Addresses{},
		conn: _conn,
	}
}

func (addr *Addresses) Insert(ctx context.Context, data *domain.Addresses) error {
	if data == nil {
		return errors.New("input data invalid (nil)")
	}
	if (data.SupplierId != "" && data.UserId != "") ||
		(data.SupplierId == "" && data.UserId == "") {
		return errors.New("unknown data from supplier or user")
	}
	return db.SafeWriteQuery(
		ctx,
		addr.conn,
		queries.InsertIntoAddresses,
		data.UserId, data.SupplierId, data.Street, data.Ward, data.District, data.City,
	)
}