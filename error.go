package skroutz

// Error response struct.
type Error struct {
	Status      string
	StatusCode  int
	Description []Description `json:"errors"`
}

// Error contains the basic description for the error string.
func (e *Error) Error() string {
	return e.Description[0].Messages[0]
}

// Description error.
type Description struct {
	Code     string   `json:"code"`
	Messages []string `json:"messages"`
}

// ErrorDescription contains error json from the skroutz api.
type ErrorDescription struct {
	Description []Description `json:"errors"`
}
