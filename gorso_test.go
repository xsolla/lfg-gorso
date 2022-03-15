package gorso

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestAccessTokenGet(t *testing.T) {
	code := "INVALID_CODE"

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	client := Client{
		ID:       clientID,
		Secret:   clientSecret,
		Redirect: "",
	}

	response, err := client.GetToken(code)
	if err != nil {
		t.Error(err)
		return
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Println(str)
}
