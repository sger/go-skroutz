package skroutz

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "1"
	defaultBaseURL = "http://api.skroutz.gr"
	authBaseURL    = "https://www.skroutz.gr"
	mediaType      = "application/vnd.skroutz+json; version=3.1"
)

// Client struct.
type Client struct {
	*Config
	User          *User
	Search        *Search
	Categories    *Categories
	SKUS          *SKUS
	Products      *Products
	FlagContent   *FlagContent
	FilterGroups  *FilterGroups
	Shops         *Shops
	Manufacturers *Manufacturers
}

// New client.
func New(config *Config) *Client {
	client := &Client{Config: config}
	client.Categories = &Categories{client}
	client.User = &User{client}
	client.Search = &Search{client}
	client.SKUS = &SKUS{client}
	client.Products = &Products{client}
	client.FlagContent = &FlagContent{client}
	client.FilterGroups = &FilterGroups{client}
	client.Shops = &Shops{client}
	client.Manufacturers = &Manufacturers{client}
	return client
}

// addURLOptions adds the parameters for the url query.
func addURLOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// call api request GET, POST, DELETE, PATCH, PUT.
func (c *Client) call(method string, path string, in interface{}) (io.ReadCloser, error) {
	url := defaultBaseURL + path

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Accept", mediaType)
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
