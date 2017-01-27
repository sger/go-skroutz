package skroutz

// Meta struct contains all the metadata for api response
type Meta struct {
	Q             string        `json:"q,omitempty"`
	Alternatives  []Alternative `json:"alternatives,omitempty"`
	StrongMatches struct {
		SKU          SKU          `json:"sku,omitempty"`
		Manufacturer Manufacturer `json:"manufacturer,omitempty"`
		Category     Category     `json:"category,omitempty"`
	} `json:"strong_matches,omitempty"`
	OrderedBy        string `json:"ordered_by,omitempty"`
	AvailableFilters struct {
		Filters       map[string]int64 `json:"filters,omitempty"`
		Manufacturers map[string]int64 `json:"manufacturers,omitempty"`
	} `json:"available_filters,omitempty"`
	AppliedFilters struct {
		Filters       []int64 `json:"filters,omitempty"`
		Manufacturers []int64 `json:"manufacturers,omitempty"`
	} `json:"applied_filters,omitempty"`
	ShowInTiles bool       `json:"show_in_tiles,omitempty"`
	Pagination  Pagination `json:"pagination"`
}

type Alternative struct {
	Term      string `json:"term"`
	Count     int    `json:"count"`
	Important bool   `json:"important"`
	Drops     []Drop `json:"drop"`
}

type Drop struct {
	Token   string `json:"token"`
	Dropped bool   `json:"dropped"`
}

type Pagination struct {
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Page         int `json:"page"`
	Per          int `json:"per"`
}
