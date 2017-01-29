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

// GetUserFavoriteListsOutput request output.
type GetUserFavoriteListsOutput struct {
	FavoriteList  []FavoriteList `json:"favorite_lists"`
	GetMetaOutput GetMetaOutput  `json:"meta"`
}

// FavoriteList request output.
type FavoriteList struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
}

// GetFavoriteListOutput request output.
type GetFavoriteListOutput struct {
	FavoriteList FavoriteList `json:"favorite_list"`
}

// FavoriteListQuery request input.
type FavoriteListQuery struct {
	FL `url:"favorite_list"`
}

// FL request output.
type FL struct {
	Name string `url:"name,omitempty"`
}

// FavoriteQuery request input.
type FavoriteQuery struct {
	F `url:"favorite"`
}

// F request output.
type F struct {
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
func (c *User) GetUserFavoriteLists() (out *GetUserFavoriteListsOutput, err error) {
	body, err := c.call("GET", "/favorite_lists", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetUserFavorites ...
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
func (c *User) CreateFavoriteList(fl *FavoriteListQuery) (out *GetFavoriteListOutput, err error) {
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
func (c *User) CreateFavorite(f *FavoriteQuery) (out *GetSingleFavoriteOutput, err error) {
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
