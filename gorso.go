package gorso

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lf-group/gorso/rsoerror"
)

//
type CodeResponse struct {
	//
	Scope string `json:"scope"`
	//
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	SubSID       string `json:"sub_sid"`
	AccessToken  string `json:"access_token"`
}

// GetToken
func (c *Client) GetToken(code string) (*CodeResponse, error) {
	client := http.Client{Timeout: c.getTimeout()}

	req, err := http.NewRequest(http.MethodPost, "https://auth.riotgames.com/authorize", nil)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	c.addAuthHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Form.Add("grant_type", "authorization_code")
	req.Form.Add("code", code)
	req.Form.Add("redirect_uri", c.ClientRedirect)

	res, err := client.Do(req)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	if res.StatusCode != http.StatusOK {
		errRequestFailed := fmt.Errorf("BattlenetGetTokenByCode: request failed: resp status %d, msg: %s", res.StatusCode, string(body))
		return nil, errRequestFailed
	}

	data := CodeResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	return &data, nil
}

// RefreshToken
func (c *Client) RefreshToken(refreshToken string) (*CodeResponse, error) {
	client := http.Client{Timeout: c.getTimeout()}

	req, err := http.NewRequest(http.MethodPost, "https://auth.riotgames.com/authorize", nil)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	c.addAuthHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Form.Add("grant_type", "refresh_token")
	req.Form.Add("refresh_token", refreshToken)
	req.Form.Add("redirect_uri", c.Redirect)

	res, err := client.Do(req)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	if res.StatusCode != http.StatusOK {
		errRequestFailed := fmt.Errorf("BattlenetGetTokenByCode: request failed: resp status %d, msg: %s", res.StatusCode, string(body))
		return nil, errRequestFailed
	}

	data := CodeResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, rsoerror.New(rsoerror.System, err)
	}

	return &data, nil
}

var client = Client{
	ID:       "CLIENT_ID",
	Secret:   "CLIENT_SECRET",
	Redirect: "REDIRECT_URL",
}

func ExampleAuthUser() {
	code := "CLIENT_CODE" // code is obtained on a client side

	data, err := client.GetToken(code)
	if err != nil {
		if errors.Is(err, rsoerror.ErrSystem) {
			panic(err)
		}

		return
	}

	fmt.Println(data.AccessToken)
}
