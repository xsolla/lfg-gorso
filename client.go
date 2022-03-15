package gorso

import (
	"net/http"
	"time"
)

// Client is a client for Riot API service
type Client struct {
	// ClientID: Unique name for your product without punctation or spaces (provided by Riot)
	ID string
	// Client secret: API Key (provided by Riot)
	Secret string
	// Redirect URL for Riot callback you have to set up at your own server.
	// This route needs to be able to process a code query parameter
	// that is added to the URI on when Riot Sign On redirects the player back to our URI.
	// We must also be sure we have this URI added as one of the redirect_uris during client registration
	Redirect string
	// Timeout is maximum request waiting time
	// If not provided, default is 10sec
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
