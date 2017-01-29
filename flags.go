package skroutz

import "encoding/json"

// FlagsContent client.
type FlagsContent struct {
	*Client
}

// NewFlagsContent client.
func NewFlagsContent(config *Config) *FlagsContent {
	return &FlagsContent{
		Client: &Client{
			Config: config,
		},
	}
}

// GetFlagsOptionOutput request output.
type GetFlagsOptionOutput struct {
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// GetFlagsOptionCollectionOutput request output.
type GetFlagsOptionCollectionOutput struct {
	Flags []GetFlagsOptionOutput `json:"flags"`
}

// GetAllFlags retrieve all flags.
// https://developer.skroutz.gr/api/v3/flag/#retrieve-all-flags
func (c *FlagsContent) GetAllFlags() (out *GetFlagsOptionCollectionOutput, err error) {
	u := "/flags"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
