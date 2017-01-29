package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_GetSingleProduct(t *testing.T) {
	c := client()
	_, err := c.Products.GetSingleProduct(12176638)
	assert.NoError(t, err)
}
