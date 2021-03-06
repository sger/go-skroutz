package skroutz

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// GetUserOutput request output.
type GetUserOutput struct {
	User struct {
		ID       float64 `json:"id"`
		Username string  `json:"username"`
		Sex      string  `json:"sex"`
		Avatar   string  `json:"avatar"`
		Type     string  `json:"type"`
	} `json:"user"`
}

// GetUserFavoriteListsCollectionOutput request output.
type GetUserFavoriteListsCollectionOutput struct {
	FavoriteList []GetFavoriteListOutput `json:"favorite_lists"`
	Meta         GetMetaOutput           `json:"meta"`
}

// GetFavoriteListOutput request output.
type GetFavoriteListOutput struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
}

// GetSingleFavoriteListOutput request output.
type GetSingleFavoriteListOutput struct {
	FavoriteList GetFavoriteListOutput `json:"favorite_list"`
}

// GetFavoriteListQueryInput request input.
type GetFavoriteListQueryInput struct {
	GetFLOutput `url:"favorite_list"`
}

// GetFLOutput request output.
type GetFLOutput struct {
	Name string `url:"name,omitempty"`
}

// GetFavoriteQueryInput request input.
type GetFavoriteQueryInput struct {
	GetFOutput `url:"favorite"`
}

// GetFOutput request output.
type GetFOutput struct {
	SkuID int64 `url:"sku_id,omitempty"`
}

// User client.
type User struct {
	*Client
}

// NewUser client.
func NewUser(config *Config) *User {
	return &User{
		Client: &Client{
			Config: config,
		},
	}
}

// GetUser retrieve the profile of the authenticated user.
// https://developer.skroutz.gr/api/v3/user/#retrieve-the-profile-of-the-authenticated-user
func (c *User) GetUser() (out *GetUserOutput, err error) {
	body, err := c.call("GET", "/user", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetUserFavoriteLists list favorite lists.
// https://developer.skroutz.gr/api/v3/favorites/#list-favorite-lists
func (c *User) GetUserFavoriteLists() (out *GetUserFavoriteListsCollectionOutput, err error) {
	body, err := c.call("GET", "/favorite_lists", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetUserFavorites ...
// https://developer.skroutz.gr/api/v3/favorites/#list-favorites
func (c *User) GetUserFavorites() (out *GetFavoritesCollectionOutput, err error) {
	body, err := c.call("GET", "/favorites", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetUserSingleFavorite ...
// https://developer.skroutz.gr/api/v3/favorites/#retrieve-a-single-favorite
func (c *User) GetUserSingleFavorite(favoriteListID int) (out *GetSingleFavoriteOutput, err error) {
	body, err := c.call("GET", "/favorites/"+strconv.Itoa(favoriteListID), nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// CreateFavoriteList ...
// https://developer.skroutz.gr/api/v3/favorites/#create-a-favoritelist
func (c *User) CreateFavoriteList(fl *GetFavoriteListQueryInput) (out *GetFavoriteListQueryInput, err error) {
	u := "/favorite_lists/"
	u, err = addURLOptions(u, fl)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	body, err := c.call("POST", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// CreateFavorite ...
// https://developer.skroutz.gr/api/v3/favorites/#create-a-new-favorite
func (c *User) CreateFavorite(f *GetFavoriteQueryInput) (out *GetSingleFavoriteOutput, err error) {
	u := "/favorites/"
	u, err = addURLOptions(u, f)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	body, err := c.call("POST", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetListFavoritesBelongingToList ...
// https://developer.skroutz.gr/api/v3/favorites/#list-favorites-belonging-to-list
func (c *User) GetListFavoritesBelongingToList(favoriteListID int) (out *GetFavoritesCollectionOutput, err error) {
	u := "/favorite_lists/" + strconv.Itoa(favoriteListID) + "/favorites"
	body, err := c.call("GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// DeleteFavoriteList ...
// https://developer.skroutz.gr/api/v3/favorites/#destroy-a-favoritelist
func (c *User) DeleteFavoriteList(favoriteListID int) (io.ReadCloser, error) {
	u := "/favorite_lists/" + strconv.Itoa(favoriteListID)
	body, err := c.call("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	return body, nil
}

// DeleteFavorite ...
// https://developer.skroutz.gr/api/v3/favorites/#destroy-a-favorite
func (c *User) DeleteFavorite(favoriteID int) (io.ReadCloser, error) {
	u := "/favorites/" + strconv.Itoa(favoriteID)
	body, err := c.call("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	return body, nil
}

// UpdateFavorite ...
// https://developer.skroutz.gr/api/v3/favorites/#update-a-favorite
func (c *User) UpdateFavorite(favoriteID int, sq *GetSearchQueryInput) (out *GetSingleFavoriteOutput, err error) {
	u := "/favorites/" + strconv.Itoa(favoriteID)
	u, err = addURLOptions(u, sq)
	if err != nil {
		return nil, err
	}
	body, err := c.call("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
