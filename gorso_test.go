package gorso

import (
	"encoding/json"
	"fmt"
	"testing"
)

const CLIENT_ID = "lfgroup"
const REDIRECT = "https://beta.lf.group/auth/riot"
const SHARD = ShardEU
const SECRET = ""

func TestGetToken(t *testing.T) {
	code := ""

	client := NewClient(&Params{
		ID:       CLIENT_ID,
		Secret:   SECRET,
		Redirect: REDIRECT,
		Shard:    SHARD,
	})

	response, err := client.GetToken(code)
	if err != nil {
		t.Error(err)
		return
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Println(string(str))
}

func TestRefreshToken(t *testing.T) {
	code := ""

	client := Client{
		ID:       CLIENT_ID,
		Secret:   SECRET,
		Redirect: REDIRECT,
	}

	response, err := client.RefreshToken(code)
	if err != nil {
		t.Error(err)
		return
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Println(string(str))
}

func TestGetUserInfo(t *testing.T) {
	acessToken := ""

	client := Client{}

	response, err := client.GetUserInfo(acessToken)
	if err != nil {
		t.Error(err)
		return
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Println(string(str))
}

func TestGetAccount(t *testing.T) {
	acessToken := ""

	client := Client{}

	response, err := client.GetAccount(acessToken)
	if err != nil {
		t.Error(err)
		return
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Println(string(str))
}
