// GORSO is a Riot OAuth API wrapper written in pure Go. Provides idiomatic access to RSO API endpoints
// Available at https://github.com/lf-group/gorso
//
// Example:
//   var rso = gorso.Client{
//     ID:       "CLIENT_ID",
//   	 Secret:   "CLIENT_SECRET",
//   	 Redirect: "REDIRECT_URL",
//   }
//
//   func ExampleAuthUser() {
//   	 code := "CLIENT_CODE" // code is obtained on a client side
//
//   	 data, err := rso.GetToken(code)
//   	 if err != nil {
//   	   if errors.Is(err, gorso.ErrSystem) {
//   		   panic(err)
//    		}
//
//   	    return
//    	}
//
//   	 fmt.Println(data.AccessToken)
//   }
package gorso

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CodeResponse contains tokens to access user private data
type CodeResponse struct {
	// A predefined data scope
	Scope Scope `json:"scope"`
	// Life span of the access token in ms
	ExpiresIn int `json:"expires_in"` // TODO: time.Duration
	// Method of authorization token provides
	TokenType TokenType `json:"token_type"`
	// Issued for the purpose of obtaining new access tokens when an older one expires
	// To reissue an access token, use client.RefreshToken method
	RefreshToken string `json:"refresh_token"`
	// Decryptable JWT Token. Provides information to authenticate a player’s identity
	IDToken string `json:"id_token"`
	// The identifier of an existing session (SID) for the subject (player)
	SubSID string `json:"sub_sid"`
	// Undecryptable JWT Token
	// Used for scoped authentication of a client and player to a resource
	AccessToken string `json:"access_token"`
}

// GetToken returns access&refresh tokens based on a provided code
func (c *Client) GetToken(code string) (*CodeResponse, error) {
	client := http.Client{Timeout: c.getTimeout()}

	req, err := http.NewRequest(http.MethodPost, "https://auth.riotgames.com/authorize", nil)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	c.addAuthHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Form.Add("grant_type", "authorization_code")
	req.Form.Add("code", code)
	req.Form.Add("redirect_uri", c.Redirect)

	res, err := client.Do(req)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	if res.StatusCode != http.StatusOK {
		// TODO: handle errors
		return nil, errorCreate(ErrUnhandled, errors.New("status code not 200"))
	}

	data := CodeResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	return &data, nil
}

// RefreshResponse contains a token info to access user private data
type RefreshResponse struct {
	// A predefined data scope
	Scope Scope `json:"scope"`
	// Life span of the access token in ms
	ExpiresIn int `json:"expires_in"` // TODO: time.Duration
	// Method of authorization token provides
	TokenType TokenType `json:"token_type"`
	// Issued for the purpose of obtaining new access tokens when an older one expires
	// To reissue an access token, use client.RefreshToken method
	RefreshToken string `json:"refresh_token"`
	// Decryptable JWT Token. Provides information to authenticate a player’s identity
	IDToken string `json:"id_token"`
	// The identifier of an existing session (SID) for the subject (player)
	SubSID string `json:"sub_sid"`
	// Undecryptable JWT Token
	// Used for scoped authentication of a client and player to a resource
	AccessToken string `json:"access_token"`
}

// GetToken returns a new refresh token based on a provided refresh token
func (c *Client) RefreshToken(refreshToken string) (*CodeResponse, error) {
	client := http.Client{Timeout: c.getTimeout()}

	req, err := http.NewRequest(http.MethodPost, "https://auth.riotgames.com/authorize", nil)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	c.addAuthHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Form.Add("grant_type", "refresh_token")
	req.Form.Add("refresh_token", refreshToken)
	req.Form.Add("redirect_uri", c.Redirect)

	res, err := client.Do(req)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	if res.StatusCode != http.StatusOK {
		// TODO: handle errors
		return nil, errorCreate(ErrUnhandled, errors.New("status code not 200"))
	}

	data := CodeResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errorCreate(ErrSystem, err)
	}

	return &data, nil
}
