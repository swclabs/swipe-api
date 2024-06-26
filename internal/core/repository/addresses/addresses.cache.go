package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	address IAddressRepository
}

var _ IAddressRepository = (*cache)(nil)

func useCache(repo IAddressRepository) IAddressRepository {
	return &cache{address: repo}
}

// Insert implements IAddressRepository.
func (c *cache) Insert(ctx context.Context, data domain.Addresses) error {
	return c.address.Insert(ctx, data)
}
