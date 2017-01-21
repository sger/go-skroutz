package skroutz

import "encoding/json"

type GetUserOutput struct {
	User struct {
		ID float64 `json:"id"`
	} `json:"user"`
}

type User struct {
	*Client
}

// NewUser client.
func NewUser(config *Config) *User {
	return &User{
		Client: &Client{
			Config: config,
		},
	}
}

// GetUser contains basic information about the authenticated user.
func (c *User) GetUser() (out *GetUserOutput, err error) {
	body, err := c.call("GET", "/user", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
