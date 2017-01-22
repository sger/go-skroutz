package skroutz

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
