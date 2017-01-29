package skroutz

import (
	"encoding/json"
	"strconv"
)

// GetCategoryOutput request output.
type GetCategoryOutput struct {
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

// GetCategoriesCollectionOutput request output.
type GetCategoriesCollectionOutput struct {
	Categories []GetCategoryOutput `json:"categories"`
	Meta       GetMetaOutput       `json:"meta"`
}

// GetSingleCategoryOutput request output.
type GetSingleCategoryOutput struct {
	Category GetCategoryOutput `json:"category"`
}

// Categories client.
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

// GetCategories lists all categories.
func (c *Categories) GetCategories() (out *GetCategoriesCollectionOutput, err error) {
	body, err := c.call("GET", "/categories", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSingleCategory retrieve a single category.
func (c *Categories) GetSingleCategory(categoryID int) (out *GetSingleCategoryOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetParentCategory retrieve the parent of a category.
func (c *Categories) GetParentCategory(categoryID int) (out *GetSingleCategoryOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/parent", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetRootCategory retrieve the root category
func (c *Categories) GetRootCategory() (out *GetSingleCategoryOutput, err error) {
	body, err := c.call("GET", "/categories/root", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetChildrenCategories list the children categories of a category.
func (c *Categories) GetChildrenCategories(categoryID int) (out *GetCategoriesCollectionOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/children", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategorySpecifications list a category's specifications.
func (c *Categories) GetCategorySpecifications(categoryID int) (out *GetSpecificationCollectionOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategorySpecificationsByGroup list a category's specifications.
func (c *Categories) GetCategorySpecificationsByGroup(categoryID int) (out *GetSpecificationGroupCollectionOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications?include=group", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategoryManufacturers list a category's manufacturers.
func (c *Categories) GetCategoryManufacturers(categoryID int) (out *GetManufacturersCollectionOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/manufacturers", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategoryFavorites list a category's favorites.Requires user token with the 'favorites' permission.
func (c *Categories) GetCategoryFavorites(categoryID int) (out *GetManufacturersCollectionOutput, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/favorites", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
