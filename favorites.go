package skroutz

// GetFavoriteOutput request output.
type GetFavoriteOutput struct {
	ID                   int    `json:"id"`
	HaveIt               bool   `json:"have_it"`
	UserID               int64  `json:"user_id"`
	UserNotes            string `json:"user_notes"`
	SkuID                int64  `json:"sku_id"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	GetAbsoluteThreshold string `json:"get_absolute_threshold"`
}

// GetSingleFavoriteOutput request output.
type GetSingleFavoriteOutput struct {
	Favorite GetFavoriteOutput `json:"favorite"`
}

// GetFavoritesCollectionOutput request output.
type GetFavoritesCollectionOutput struct {
	Favorites []GetFavoriteOutput `json:"favorites"`
	Meta      GetMetaOutput       `json:"meta"`
}
