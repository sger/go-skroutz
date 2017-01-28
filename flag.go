package skroutz

import "encoding/json"

// FlagContent client.
type FlagContent struct {
	*Client
}

// NewFlagContent client.
func NewFlagContent(config *Config) *FlagContent {
	return &FlagContent{
		Client: &Client{
			Config: config,
		},
	}
}

// FlagOption request output.
type FlagOption struct {
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// FlagOptionCollection request output.
type FlagOptionCollection struct {
	FlagOption []FlagOption `json:"flags"`
}

// GetAllFlags retrieve all flags.
func (c *FlagContent) GetAllFlags() (out *FlagOptionCollection, err error) {
	u := "/flags"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
