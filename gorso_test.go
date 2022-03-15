package gorso

import "testing"

func TestAccessTokenGet(t *testing.T) {
	val := AccessTokenGet()
	if val != "" {
		t.Error()
	}
}
