package skroutz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// AuthorizationResponse represents an individual Skroutz authorization.
// Skroutz API docs: https://developer.skroutz.gr/authorization/
type AuthorizationResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Expires     float64 `json:"expires_in"`
}

// Authorization requires clientID and clientSecret request one from http://skroutz.it/API_access.
func Authorization(clientID string, clientSecret string) (*AuthorizationResponse, error) {
	resource := "/oauth2/token/"
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "public")

	u, _ := url.ParseRequestURI(authBaseURL)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response AuthorizationResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
