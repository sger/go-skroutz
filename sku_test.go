package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSKUS_GetCategorySKUS(t *testing.T) {
	c := client()
	_, err := c.SKUS.GetCategorySKUS(40, nil)
	assert.NoError(t, err)
}
