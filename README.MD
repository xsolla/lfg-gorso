[![Go Report Card](https://goreportcard.com/badge/github.com/lf-group/gorso)](https://goreportcard.com/report/github.com/lf-group/gorso)
[![codecov](https://codecov.io/gh/lf-group/gorso/branch/master/graph/badge.svg)](https://codecov.io/gh/KnutZuidema/golio)

# GORSO

GORSO is a Riot OAuth API wrapper written in pure Go.

### Prerequisites

For the moment, OAuth API does not provide any official documentation. All endpoints, parameters and info provided in GORSO are based on the [Henrik-3 gist](https://gist.github.com/Henrik-3/d6b631fb7c61821bc16b17cd347a3811).

This package does not attempt to implement Riot API in its entirety, covering only the newest public RSO feature. To access remaining methods, we recommend checking out packages, such as [golio](https://github.com/KnutZuidema/golio).

### Usage

```go
import (
	"github.com/lf-group/gorso"
)

var rso = gorso.NewClient(gorso.Client{
	ID:       "CLIENT_ID",
	Secret:   "CLIENT_SECRET",
	Redirect: "REDIRECT_URL",
	Shard: 		gorso.ShardEU,
})

func ExampleAuthUser() {
	code := "CLIENT_CODE" // code is obtained on a client side

	data, err := rso.GetToken(code)
	if err != nil {
		if errors.Is(err, gorso.ErrSystem) {
			panic(err)
		}

		return
	}

	fmt.Println(data.AccessToken)
}
```
