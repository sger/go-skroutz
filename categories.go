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

// CategoriesCollection request output.
type CategoriesCollection struct {
	Categories []Category `json:"categories"`
	Meta       Meta       `json:"meta"`
}

// CategoryItem request output.
type CategoryItem struct {
	Categories Category `json:"category"`
}

// Specification request output.
type Specification struct {
	ID     float64  `json:"id"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
	Order  float64  `json:"order"`
	Unit   string   `json:"unit"`
}

// SpecificationCollection request output.
type SpecificationCollection struct {
	Specification []Specification `json:"specifications"`
}

// SpecificationGroupCollection request output.
type SpecificationGroupCollection struct {
	Specification []Specification `json:"groups"`
}

// Favorite request output.
type Favorite struct {
	ID                   int    `json:"id"`
	HaveIt               bool   `json:"have_it"`
	UserID               int64  `json:"user_id"`
	UserNotes            string `json:"user_notes"`
	SkuID                int64  `json:"sku_id"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	GetAbsoluteThreshold string `json:"get_absolute_threshold"`
}

// SingleFavorite request output.
type SingleFavorite struct {
	Favorite Favorite `json:"favorite"`
}

// FavoritesCollection request output.
type FavoritesCollection struct {
	Favorite []Favorite `json:"favorites"`
	Meta     Meta       `json:"meta"`
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
func (c *Categories) GetCategories() (out *CategoriesCollection, err error) {
	body, err := c.call("GET", "/categories", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSingleCategory retrieve a single category.
func (c *Categories) GetSingleCategory(categoryID int) (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetParentCategory retrieve the parent of a category.
func (c *Categories) GetParentCategory(categoryID int) (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/parent", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetRootCategory retrieve the root category
func (c *Categories) GetRootCategory() (out *CategoryItem, err error) {
	body, err := c.call("GET", "/categories/root", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetChildrenCategories list the children categories of a category.
func (c *Categories) GetChildrenCategories(categoryID int) (out *CategoriesCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/children", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategorySpecifications list a category's specifications.
func (c *Categories) GetCategorySpecifications(categoryID int) (out *SpecificationCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategorySpecificationsByGroup list a category's specifications.
func (c *Categories) GetCategorySpecificationsByGroup(categoryID int) (out *SpecificationGroupCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/specifications?include=group", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategoryManufacturers list a category's manufacturers.
func (c *Categories) GetCategoryManufacturers(categoryID int) (out *ManufacturersCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/manufacturers", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCategoryFavorites list a category's favorites.Requires user token with the 'favorites' permission.
func (c *Categories) GetCategoryFavorites(categoryID int) (out *ManufacturersCollection, err error) {
	body, err := c.call("GET", "/categories/"+strconv.Itoa(categoryID)+"/favorites", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
