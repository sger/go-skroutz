package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShops_GetSingleShop(t *testing.T) {
	c := client()
	_, err := c.Shops.GetSingleShop(452)
	assert.NoError(t, err)
}
