package skroutz

// SKUSSpecificationsCollections request output.
type SKUSSpecificationsCollections struct {
	Specification []Specification `json:"specifications,omitempty"`
	Group         []Group         `json:"groups,omitempty"`
}
