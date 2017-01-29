package skroutz

// GetReviewOutput request output.
type GetReviewOutput struct {
	ID                int64  `json:"id"`
	UserID            int64  `json:"user_id"`
	Review            string `json:"review"`
	Rating            int64  `json:"rating"`
	CreatedAt         string `json:"created_at"`
	Demoted           bool   `json:"demoted"`
	VotesCount        int64  `json:"votes_count"`
	HelpfulVotesCount int64  `json:"helpful_votes_count"`
	Voted             bool   `json:"voted"`
	Flagged           bool   `json:"flagged"`
	Helpful           bool   `json:"helpful"`
}

// GetReviewsCollectionOutput request output.
type GetReviewsCollectionOutput struct {
	Reviews       []GetReviewOutput `json:"reviews"`
	GetMetaOutput GetMetaOutput     `json:"meta"`
}

// GetSKUReviewVoteOutput request output.
type GetSKUReviewVoteOutput struct {
	ID          int64           `json:"id"`
	SKUReviewID int64           `json:"sku_review_id"`
	UserID      int64           `json:"user_id"`
	Helpful     bool            `json:"helpful"`
	CreatedAt   string          `json:"created_at"`
	Review      GetReviewOutput `json:"sku_review"`
}

// SKUReviewFlag request output.
type GetSKUReviewFlagOutput struct {
	ID            int64  `json:"id"`
	FlaggableID   int64  `json:"flaggable_id"`
	FlaggableType string `json:"flaggable_type"`
	UserID        int64  `json:"user_id"`
	Reason        string `json:"reason"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
