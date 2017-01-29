package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags_GetAllFlags(t *testing.T) {
	c := client()
	_, err := c.FlagsContent.GetAllFlags()
	assert.NoError(t, err)
}
