package skroutz

// GetAutocompleteOutput struct.
type GetAutocompleteOutput struct {
	K string `json:"k"`
	I int64  `json:"i"`
	D struct {
		N  string `json:"n"`
		ID int64  `json:"id"`
	} `json:"d"`
}

// GetAutocompleteCollectionOutput struct.
type GetAutocompleteCollectionOutput struct {
	Autocomplete []GetAutocompleteOutput `json:"autocomplete"`
}
