package skroutz

import (
	"encoding/json"
	"strconv"
)

// Manufacturers client.
type Manufacturers struct {
	*Client
}

// NewManufacturers client.
func NewManufacturers(config *Config) *Manufacturers {
	return &Manufacturers{
		Client: &Client{
			Config: config,
		},
	}
}

// Manufacturer request output.
type Manufacturer struct {
	ID       float64 `json:"id"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
}

// SingleManufacturer request output.
type SingleManufacturer struct {
	Manufacturer Manufacturer `json:"manufacturer"`
}

// ManufacturersCollection request output.
type ManufacturersCollection struct {
	Specification []Specification `json:"manufacturers"`
	Meta          Meta            `json:"meta"`
}

// GetManufacturers list manufacturers.
func (c *Manufacturers) GetManufacturers() (out *ManufacturersCollection, err error) {
	u := "/manufacturers"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSingleManufacturer retrieve a single manufacturer.
func (c *Manufacturers) GetSingleManufacturer(manufacturerID int) (out *SingleManufacturer, err error) {
	u := "/manufacturers/" + strconv.Itoa(manufacturerID)
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetManufacturersCategories retrieve a manufacturer's categories.
func (c *Manufacturers) GetManufacturersCategories(manufacturerID int) (out *CategoriesCollection, err error) {
	u := "/manufacturers/" + strconv.Itoa(manufacturerID) + "/categories"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetManufacturersSKUS retrieve a manufacturer's SKUs.
func (c *Manufacturers) GetManufacturersSKUS(manufacturerID int) (out *SKUSCollection, err error) {
	u := "/manufacturers/" + strconv.Itoa(manufacturerID) + "/skus"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
