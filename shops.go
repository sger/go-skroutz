package skroutz

import (
	"encoding/json"
	"strconv"
)

// Shops client.
type Shops struct {
	*Client
}

// NewShops client.
func NewShops(config *Config) *Shops {
	return &Shops{
		Client: &Client{
			Config: config,
		},
	}
}

// GetShopOutput request output.
type GetShopOutput struct {
	ID                 int64                  `json:"id"`
	Name               string                 `json:"name"`
	Link               string                 `json:"link"`
	Phone              string                 `json:"phone"`
	ImageURL           string                 `json:"image_url"`
	ThumbshotURL       string                 `json:"thumbshot_url"`
	ReviewsCount       int64                  `json:"reviews_count"`
	LatestReviewsCount int64                  `json:"latest_reviews_count"`
	ReviewScore        float64                `json:"review_score"`
	PaymentMethods     map[string]interface{} `json:"payment_methods"`
	Shipping           map[string]interface{} `json:"shipping"`
	WebURI             string                 `json:"web_uri"`
	ExtraInfo          map[string]interface{} `json:"extra_info"`
	TopPositiveReasons []string               `json:"top_positive_reasons"`
}

// GetShopsCollectionOutput request output.
type GetShopsCollectionOutput struct {
	Shops []GetShopOutput `json:"shops"`
	Meta  GetMetaOutput   `json:"meta"`
}

// GetSingleShopOutput request output.
type GetSingleShopOutput struct {
	Shop GetShopOutput `json:"shop"`
}

// GetShopReviewOutput request output.
type GetShopReviewOutput struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Review    string   `json:"review"`
	Rating    float64  `json:"rating"`
	ShopReply string   `json:"shop_reply"`
	CreatedAt string   `json:"created_at"`
	Negative  bool     `json:"negative"`
	Reasons   []string `json:"reasons"`
}

// GetShopReviewsCollectionOutput requets output.
type GetShopReviewsCollectionOutput struct {
	ShopReviews []GetShopReviewOutput `json:"reviews"`
	Meta        GetMetaOutput         `json:"meta"`
}

// GetLocationOutput request output.
type GetLocationOutput struct {
	ID          int64    `json:"id"`
	Headquarter bool     `json:"headquarter"`
	Phones      []string `json:"phones"`
	PickupPoint bool     `json:"pickup_point"`
	Store       bool     `json:"store"`
	FullAddress string   `json:"full_address"`
	Structured  string   `json:"structured"`
	Lat         string   `json:"lat"`
	Lng         string   `json:"lng"`
	Info        string   `json:"info"`
	Address     Address  `json:"address"`
}

// GetSingleShopLocationOutput request output.
type GetSingleShopLocationOutput struct {
	Location GetLocationOutput `json:"location"`
}

// Address request output.
type Address struct {
	ID           int64  `json:"id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	PostCode     string `json:"postcode"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
}

// GetLocationsCollectionOutput request output.
type GetLocationsCollectionOutput struct {
	Locations []GetLocationOutput `json:"locations"`
	Meta      GetMetaOutput       `json:"meta"`
}

// GetSingleShop retrieve a single shop.
func (c *Shops) GetSingleShop(shopID int) (out *GetSingleShopOutput, err error) {
	u := "/shops/" + strconv.Itoa(shopID)
	body, err := c.call("GET", u, nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetShopReviews retrieve a shop's reviews.
func (c *Shops) GetShopReviews(shopID int, sq *GetSearchQueryInput) (out *GetShopReviewsCollectionOutput, err error) {
	u := "/shops/" + strconv.Itoa(shopID) + "/reviews"
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

// GetShopLocations list shop locations.
func (c *Shops) GetShopLocations(shopID int, sq *GetSearchQueryInput) (out *GetLocationsCollectionOutput, err error) {
	u := "/shops/" + strconv.Itoa(shopID) + "/locations"
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

// GetSingleShopLocation retrieve a single shop location.
func (c *Shops) GetSingleShopLocation(shopID int, locationID int, sq *GetSearchQueryInput) (out *GetSingleShopLocationOutput, err error) {
	u := "/shops/" + strconv.Itoa(shopID) + "/locations/" + strconv.Itoa(locationID)
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

// SearchShop search for shops
func (c *Shops) SearchShop(sq *GetSearchQueryInput) (out *GetShopsCollectionOutput, err error) {
	u := "/shops/search"
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
