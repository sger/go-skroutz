package skroutz

import (
	"encoding/json"
	"strconv"
)

// GetProductOutput request output.
type GetProductOutput struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	SkuID           int64    `json:"sku_id"`
	ShopID          int64    `json:"shop_id"`
	CategoryID      int64    `json:"category_id"`
	Availability    string   `json:"availability"`
	ClickURL        string   `json:"click_url"`
	ShopUID         string   `json:"shop_uid"`
	Expenses        string   `json:"expenses"`
	WebURI          string   `json:"web_uri"`
	Price           float64  `json:"price"`
	Sizes           []string `json:"sizes"`
	ImmediatePickup bool     `json:"immediate_pickup"`
}

// GetProductsCollectionOutput request output.
type GetProductsCollectionOutput struct {
	Products []GetProductOutput `json:"products"`
	Meta     GetMetaOutput      `json:"meta"`
}

// GetSingleProductOutput request output.
type GetSingleProductOutput struct {
	Product GetProductOutput `json:"product"`
}

// Products client.
type Products struct {
	*Client
}

// NewProducts client.
func NewProducts(config *Config) *Products {
	return &Products{
		Client: &Client{
			Config: config,
		},
	}
}

// GetSingleProduct retrieve a single product.
// https://developer.skroutz.gr/api/v3/product/#retrieve-a-single-product
func (c *Products) GetSingleProduct(productID int) (out *GetSingleProductOutput, err error) {
	body, err := c.call("GET", "/products/"+strconv.Itoa(productID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// SearchProducts search for products.
// https://developer.skroutz.gr/api/v3/product/#search-for-products
func (c *Products) SearchProducts(shopID int, shopUID string) (out *GetProductsCollectionOutput, err error) {
	body, err := c.call("GET", "/shops/"+strconv.Itoa(shopID)+"/products/search?shop_uid="+shopUID, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
