package skroutz

import (
	"testing"

	"github.com/segmentio/go-env"
	"github.com/stretchr/testify/assert"
)

func client() *Client {
	accessToken := env.MustGet("SKROUTZ_ACCESS_TOKEN")
	return New(NewConfig(accessToken))
}

func TestClient_error_json(t *testing.T) {
	c := client()

	_, err := c.Search.Search("s")
	assert.Error(t, err)

	e := err.(*Error)

	assert.Contains(t, e.Error(), "Η αναζήτησή σου πρέπει να περιέχει τουλάχιστον 2 χαρακτήρες")
	assert.Equal(t, "Unprocessable Entity", e.Status)
	assert.Equal(t, 422, e.StatusCode)
}
