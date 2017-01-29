package skroutz

// Favorite request output.
type Favorite struct {
	ID                   int    `json:"id"`
	HaveIt               bool   `json:"have_it"`
	UserID               int64  `json:"user_id"`
	UserNotes            string `json:"user_notes"`
	SkuID                int64  `json:"sku_id"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	GetAbsoluteThreshold string `json:"get_absolute_threshold"`
}

// SingleFavorite request output.
type SingleFavorite struct {
	Favorite Favorite `json:"favorite"`
}

// FavoritesCollection request output.
type FavoritesCollection struct {
	Favorite []Favorite `json:"favorites"`
	Meta     Meta       `json:"meta"`
}
