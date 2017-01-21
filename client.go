package skroutz

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Client contains Config, User, Categories
type Client struct {
	*Config
	User       *User
	Categories *Categories
}

// New client.
func New(config *Config) *Client {
	client := &Client{Config: config}
	client.Categories = &Categories{client}
	client.User = &User{client}
	return client
}

// call api request GET, POST.
func (c *Client) call(method string, path string, in interface{}) (io.ReadCloser, error) {
	url := "http://api.skroutz.gr" + path

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Accept", "application/vnd.skroutz+json; version=3.1")
	req.Header.Set("Content-Type", "application/json")

	r, _, err := c.do(req)
	return r, err
}

// do initialize the HTTPClient request.
func (c *Client) do(req *http.Request) (io.ReadCloser, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode < 400 {
		return res.Body, res.ContentLength, err
	}

	decoder := json.NewDecoder(res.Body)

	var errorDesc ErrorDescription
	err = decoder.Decode(&errorDesc)

	if err != nil {
		return nil, 0, err
	}

	e := &Error{
		Status:      http.StatusText(res.StatusCode),
		StatusCode:  res.StatusCode,
		Description: errorDesc.Description,
	}

	defer res.Body.Close()

	return nil, 0, e
}
