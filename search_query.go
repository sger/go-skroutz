package skroutz

// GetSearchQueryInput request input.
type GetSearchQueryInput struct {
	Q               string   `url:"q,omitempty"`
	ManufacturerIDS []string `url:"manufacturer_ids,brackets"`
	FilterIDS       []string `url:"filter_ids,brackets"`
	IncludeMeta     []string `url:"include_meta,brackets"`
	Include         string   `url:"include,omitempty"`
	Embed           string   `url:"embed,omitempty"`
	HaveIt          bool     `url:"have_it,omitempty"`
}
