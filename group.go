package skroutz

import (
	"encoding/json"
	"strconv"
)

// Group request output.
type Group struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Order float64 `json:"order"`
}

// GroupCollection request output.
type GroupCollection struct {
	Group []Group `json:"groups"`
}

// FilterGroup request output.
type FilterGroup struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	CategoryID int64  `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Hint       string `json:"hint"`
	Combined   bool   `json:"combined"`
	FilterType int    `json:"filter_type"`
}

// FilterGroupCollection request output.
type FilterGroupCollection struct {
	FilterGroup []FilterGroup `json:"filter_groups"`
	Meta        Meta          `json:"meta"`
}

// FilterGroups client.
type FilterGroups struct {
	*Client
}

// NewFilterGroups client.
func NewFilterGroups(config *Config) *FilterGroups {
	return &FilterGroups{
		Client: &Client{
			Config: config,
		},
	}
}

// GetFilterGroups list FilterGroups.
func (c *FilterGroups) GetFilterGroups(categoryID int) (out *FilterGroupCollection, err error) {
	u := "/categories/" + strconv.Itoa(categoryID) + "/filter_groups"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
