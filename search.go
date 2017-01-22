package skroutz

import "encoding/json"

// Search client for search struct
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

func (c *Search) Search(query string) (out *CategoriesCollection, err error) {
	body, err := c.call("GET", "/search?q="+query, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Search) Autocomplete(query string) (out *AutocompleteCollection, err error) {
	body, err := c.call("GET", "/autocomplete?q="+query, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
