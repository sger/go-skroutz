package skroutz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup_GetFilterGroups(t *testing.T) {
	c := client()
	_, err := c.FilterGroups.GetFilterGroups(40)
	assert.NoError(t, err)
}
