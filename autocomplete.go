package skroutz

type Autocomplete struct {
	K string `json:"k"`
	I int64  `json:"i"`
	D struct {
		N  string `json:"n"`
		ID int64  `json:"id"`
	} `json:"d"`
}

type AutocompleteCollection struct {
	Autocomplete []Autocomplete `json:"autocomplete"`
}
