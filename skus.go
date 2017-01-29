package skroutz

import (
	"encoding/json"
	"strconv"
)

// GetAverageOutput request output.
type GetAverageOutput struct {
	Date     string  `json:"date"`
	Price    float64 `json:"price"`
	ShopName string  `json:"shop_name"`
}

// GetAverageCollectionOutput request output.
type GetAverageCollectionOutput struct {
	Average []GetAverageOutput `json:"average"`
	Lowest  []GetAverageOutput `json:"lowest"`
}

// GetPriceHistoryCollectionOutput request output.
type GetPriceHistoryCollectionOutput struct {
	AverageCollection GetAverageCollectionOutput `json:"history"`
}

// GetFlagQueryInput request input.
type GetFlagQueryInput struct {
	Flag GetFlagOutput `url:"flag"`
}

// GetFlagOutput request output.
type GetFlagOutput struct {
	Reason string `url:"reason,omitempty"`
}

// GetVoteQueryInput request input.
type GetVoteQueryInput struct {
	Embed string        `url:"embed,omitempty"`
	Vote  GetVoteOutput `url:"vote"`
}

// GetVoteOutput request output.
type GetVoteOutput struct {
	Helpful bool `url:"helpful,omitempty"`
}

// GetSKUOutput request output.
type GetSKUOutput struct {
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

// GetSingleSKUOutput request output.
type GetSingleSKUOutput struct {
	SKU GetSKUOutput `json:"sku"`
}

// GetSKUSCollectionOutput request output.
type GetSKUSCollectionOutput struct {
	SKU           []GetSKUOutput `json:"skus"`
	GetMetaOutput GetMetaOutput  `json:"meta"`
}

// SKUS client.
type SKUS struct {
	*Client
}

// NewSKUS client.
func NewSKUS(config *Config) *SKUS {
	return &SKUS{
		Client: &Client{
			Config: config,
		},
	}
}

// GetCategorySKUS list SKUs of specific category.
func (c *SKUS) GetCategorySKUS(categoryID int, sq *GetSearchQueryInput) (out *GetSKUSCollectionOutput, err error) {
	u := "/categories/" + strconv.Itoa(categoryID) + "/skus"
	u, err = addURLOptions(u, sq)
	if err != nil {
		return nil, err
	}
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSingleSKU retrieve a single SKU.
func (c *SKUS) GetSingleSKU(skuID int) (out *GetSingleSKUOutput, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID), nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSimilarSKUS retrieve similar SKUs.
func (c *SKUS) GetSimilarSKUS(skuID int) (out *GetSKUSCollectionOutput, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID)+"/similar", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSKUSProducts retrieve an SKU's products.
func (c *SKUS) GetSKUSProducts(skuID int) (out *GetProductsCollectionOutput, err error) {
	body, err := c.call("GET", "/skus/"+strconv.Itoa(skuID)+"/products", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSKUSReviews retrieve an SKU's reviews.
func (c *SKUS) GetSKUSReviews(skuID int, sq *GetSearchQueryInput) (out *GetReviewsCollectionOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/reviews"
	u, err = addURLOptions(u, sq)
	if err != nil {
		return nil, err
	}
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// VoteSKUReview vote a SKU's review.
func (c *SKUS) VoteSKUReview(skuID int, reviewID int, v *GetVoteQueryInput) (out *GetSKUReviewVoteOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/reviews/" + strconv.Itoa(reviewID) + "/votes"
	u, err = addURLOptions(u, v)
	if err != nil {
		return nil, err
	}
	body, err := c.call("POST", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// FlagSKUReview flag a SKU's review.
func (c *SKUS) FlagSKUReview(skuID int, reviewID int, f *GetFlagQueryInput) (out *GetSKUReviewFlagOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/reviews/" + strconv.Itoa(reviewID) + "/flags"
	u, err = addURLOptions(u, f)
	if err != nil {
		return nil, err
	}
	body, err := c.call("POST", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSKUSSpecifications retrieve an SKU's specifications.
func (c *SKUS) GetSKUSSpecifications(skuID int, sq *GetSearchQueryInput) (out *GetSKUSSpecificationsCollectionsOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/specifications/"
	u, err = addURLOptions(u, sq)
	if err != nil {
		return nil, err
	}
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSKUSPriceHistory retrieve a SKU's price history.
func (c *SKUS) GetSKUSPriceHistory(skuID int) (out *GetPriceHistoryCollectionOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/price_history/"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSKUSFavorite retrieve an SKU's favorite.
func (c *SKUS) GetSKUSFavorite(skuID int) (out *GetFavoritesCollectionOutput, err error) {
	u := "/skus/" + strconv.Itoa(skuID) + "/favorite/"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
