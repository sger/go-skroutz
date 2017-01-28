package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategories_GetCategories(t *testing.T) {
	c := client()
	_, err := c.Categories.GetCategories()
	assert.NoError(t, err)
}
