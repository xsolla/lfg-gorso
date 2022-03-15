package gorso

import (
	"net/http"
	"time"
)

//
type Client struct {
	//
	ID string
	//
	Secret string
	//
	Redirect string
	//
	Timeout time.Duration
}

// addAuthHeader appends basic authorization header encoded in base64
func (c *Client) addAuthHeader(req *http.Request) {
	req.SetBasicAuth(c.ID, c.Secret)
}

// getTimeout returns default timeout if client timeout is not specified
func (c *Client) getTimeout() time.Duration {
	if c.Timeout == 0 {
		return DEFAULT_TIMEOUT
	}

	return c.Timeout
}
