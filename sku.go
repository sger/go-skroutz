package skroutz

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SearchQuery struct {
	Q               string   `url:"q,omitempty"`
	ManufacturerIDS []string `url:"manufacturer_ids,brackets"`
	FilterIDS       []string `url:"filter_ids,brackets"`
	IncludeMeta     []string `url:"include_meta,brackets"`
}

type SKU struct {
	ID                   int64   `json:"id"`
	EAN                  string  `json:"ean"`
	PN                   string  `json:"pn"`
	Name                 string  `json:"name"`
	DisplayName          string  `json:"display_name"`
	CategoryID           int64   `json:"category_id"`
	FirstProductShopInfo string  `json:"first_product_shop_info"`
	ClickURL             string  `json:"click_url"`
	PriceMax             float64 `json:"price_max"`
	PriceMin             float64 `json:"price_min"`
	ReviewScore          float64 `json:"review_score"`
	ShopCount            int     `json:"shop_count"`
	PlainSpecSummary     string  `json:"plain_spec_summary"`
	ManufacturerID       int64   `json:"manufacturer_id"`
	Future               bool    `json:"future"`
	ReviewsCount         int     `json:"reviews_count"`
	Virtual              bool    `json:"virtual"`
	WebURI               string  `json:"web_uri"`
	Comparable           bool    `json:"comparable"`
	Images               struct {
		Main         string   `json:"main"`
		Alternatives []string `json:"alternatives"`
	} `json:"images"`
}

type SingleSKU struct {
	SKU SKU `json:"sku"`
}

type SKUSCollection struct {
	SKU  []SKU `json:"skus"`
	Meta Meta  `json:"meta"`
}

// Categories client for categories struct
type SKUS struct {
	*Client
}

// NewCategories client.
func NewSKUS(config *Config) *SKUS {
	return &SKUS{
		Client: &Client{
			Config: config,
		},
	}
}

func (c *SKUS) GetCategorySKUS(categoryID int, sq *SearchQuery) (out *SKUSCollection, err error) {
	u := "/categories/" + strconv.Itoa(categoryID) + "/skus"
	u, err = addURLOptions(u, sq)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *SKUS) GetSingleSKU(skuID int) (out *SingleSKU, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *SKUS) GetSimilarSKUS(skuID int) (out *SKUSCollection, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID)+"/similar", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

func (c *SKUS) GetSKUSProducts(skuID int) (out *ProductsCollection, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID)+"/products", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
