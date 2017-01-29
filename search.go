package skroutz

import "encoding/json"

// Search client.
type Search struct {
	*Client
}

// NewSearch client.
func NewSearch(config *Config) *Search {
	return &Search{
		Client: &Client{
			Config: config,
		},
	}
}

// Search query with less than 2 characters
// https://developer.skroutz.gr/api/v3/search/#search
func (c *Search) Search(query string) (out *GetCategoriesCollectionOutput, err error) {
	body, err := c.call("GET", "/search?q="+query, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// Autocomplete query with less than 2 characters
// https://developer.skroutz.gr/api/v3/search/#autocomplete
func (c *Search) Autocomplete(query string) (out *GetAutocompleteCollectionOutput, err error) {
	body, err := c.call("GET", "/autocomplete?q="+query, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
