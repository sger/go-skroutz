package skroutz

// GetMetaOutput struct contains all the metadata for the api response.
type GetMetaOutput struct {
	Q             string                 `json:"q,omitempty"`
	Alternatives  []GetAlternativeOutput `json:"alternatives,omitempty"`
	StrongMatches struct {
		SKU               GetSKUOutput          `json:"sku,omitempty"`
		Manufacturer      GetManufacturerOutput `json:"manufacturer,omitempty"`
		GetCategoryOutput GetCategoryOutput     `json:"category,omitempty"`
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
	ShowInTiles        bool                          `json:"show_in_tiles,omitempty"`
	SKURatingBreakdown []GetSKURatingBreakdownOutput `json:"sku_rating_breakdown,omitempty"`
	Pagination         GetPaginationOutput           `json:"pagination"`
}

// GetAlternativeOutput struct for meta.
type GetAlternativeOutput struct {
	Term      string          `json:"term"`
	Count     int             `json:"count"`
	Important bool            `json:"important"`
	Drops     []GetDropOutput `json:"drop"`
}

// GetDropOutput request output.
type GetDropOutput struct {
	Token   string `json:"token"`
	Dropped bool   `json:"dropped"`
}

// TotalResults request output.
type GetPaginationOutput struct {
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Page         int `json:"page"`
	Per          int `json:"per"`
}

// GetSKURatingBreakdownOutput request output.
type GetSKURatingBreakdownOutput struct {
	Star       int64 `json:"star"`
	Percentage int   `json:"percentage"`
	Count      int   `json:"count"`
}
