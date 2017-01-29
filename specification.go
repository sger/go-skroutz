package skroutz

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

// SKUSSpecificationsCollections request output.
type SKUSSpecificationsCollections struct {
	Specification []Specification `json:"specifications,omitempty"`
	Group         []Group         `json:"groups,omitempty"`
}
