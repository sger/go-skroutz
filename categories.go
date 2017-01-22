package skroutz

import (
	"encoding/json"
	"strconv"
)

// Category request output.
type Category struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	ChildrenCount      int    `json:"children_count"`
	ImageURL           string `json:"image_url"`
	ParentID           int    `json:"parent_id"`
	Fashion            bool   `json:"fashion"`
	LayoutMode         string `json:"layout_mode"`
	WebURI             string `json:"web_uri"`
	Code               string `json:"code"`
	Path               string `json:"path"`
	ShowSpecifications bool   `json:"show_specifications"`
	ManufacturerTitle  string `json:"manufacturer_title"`
}

// CategoryItems request output.
type CategoriesCollection struct {
	Categories []Category `json:"categories"`
	Meta       Meta       `json:"meta"`
}

type CategoryItem struct {
	Categories Category `json:"category"`
}

type Specification struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Order float64 `json:"order"`
	Unit  string  `json:"unit"`
}

type SpecificationCollection struct {
	Specification []Specification `json:"specifications"`
}

type SpecificationGroupCollection struct {
	Specification []Specification `json:"groups"`
}

type Manufacturer struct {
	ID       float64 `json:"id"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
}

type ManufacturersCollection struct {
	Specification []Specification `json:"manufacturers"`
	Meta          Meta            `json:"meta"`
}

type Favorite struct {
	ID int `json:"id"`
}

type FavoritesCollection struct {
	Favorite []Favorite `json:"favorites"`
	Meta     Meta       `json:"meta"`
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
func (c *Categories) GetCategories() (out *CategoriesCollection, err error) {
	body, err := c.call("GET", "/categories", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetSingleCategory(categoryID int) (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetParentCategory(categoryID int) (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/parent", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetRootCategory() (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/root", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetChildrenCategories(categoryID int) (out *CategoriesCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/children", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetCategorySpecifications(categoryID int) (out *SpecificationCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetCategorySpecificationsByGroup(categoryID int) (out *SpecificationGroupCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications?include=group", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetCategoryManufacturers(categoryID int) (out *ManufacturersCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/manufacturers", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *Categories) GetCategoryFavorites(categoryID int) (out *ManufacturersCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/favorites", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
