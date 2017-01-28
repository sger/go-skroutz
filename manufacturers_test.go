package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturers_GetManufacturers(t *testing.T) {
	c := client()
	_, err := c.Manufacturers.GetManufacturers()
	assert.NoError(t, err)
}
