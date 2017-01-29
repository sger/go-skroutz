package skroutz

// GetSpecificationOutput request output.
type GetSpecificationOutput struct {
	ID     float64  `json:"id"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
	Order  float64  `json:"order"`
	Unit   string   `json:"unit"`
}

// GetSpecificationCollectionOutput request output.
type GetSpecificationCollectionOutput struct {
	Specifications []GetSpecificationOutput `json:"specifications"`
}

// GetSpecificationGroupCollectionOutput request output.
type GetSpecificationGroupCollectionOutput struct {
	Groups []GetSpecificationOutput `json:"groups"`
}

// GetSKUSSpecificationsCollectionsOutput request output.
type GetSKUSSpecificationsCollectionsOutput struct {
	Specifications []GetSpecificationOutput `json:"specifications,omitempty"`
	Groups         []GetGroupsOutput        `json:"groups,omitempty"`
}
