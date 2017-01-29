package skroutz

import (
	"encoding/json"
	"strconv"
)

// GetGroupsOutput request output.
type GetGroupsOutput struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Order float64 `json:"order"`
}

// GetGroupsCollectionOutput request output.
type GetGroupsCollectionOutput struct {
	Groups []GetGroupsOutput `json:"groups"`
}

// GetFilterGroupOutput request output.
type GetFilterGroupOutput struct {
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

// GetFilterGroupCollectionOutput request output.
type GetFilterGroupCollectionOutput struct {
	FilterGroups  []GetFilterGroupOutput `json:"filter_groups"`
	GetMetaOutput GetMetaOutput          `json:"meta"`
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
func (c *FilterGroups) GetFilterGroups(categoryID int) (out *GetFilterGroupCollectionOutput, err error) {
	u := "/categories/" + strconv.Itoa(categoryID) + "/filter_groups"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
