package skroutz

import "encoding/json"

// Category request output.
type Category struct {
	ID                 float64 `json:"id"`
	Name               string  `json:"name"`
	ChildrenCount      float64 `json:"children_count"`
	ImageURL           string  `json:"image_url"`
	ParentID           float64 `json:"parent_id"`
	Fashion            bool    `json:"fashion"`
	LayoutMode         string  `json:"layout_mode"`
	WebURI             string  `json:"web_uri"`
	Code               string  `json:"code"`
	Path               string  `json:"path"`
	ShowSpecifications bool    `json:"show_specifications"`
	ManufacturerTitle  string  `json:"manufacturer_title"`
}

// CategoryItems request output.
type CategoryItems struct {
	Categories []Category `json:"categories"`
}

// Categories client for categories struct
type Categories struct {
	*Client
}

// NewCategories client.
func NewCategories(config *Config) *Categories {
	return &Categories{
		Client: &Client{
			Config: config,
		},
	}
}

// GetCategories lists all categories
func (c *Categories) GetCategories() (out *CategoryItems, err error) {
	body, err := c.call("GET", "/categories", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
