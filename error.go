package gorso

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	StatusCode  int    `json:"-"`
	Type        string `json:"error"`
	Description string `json:"error_description"`
}

func NewError(code int, body []byte) *Error {
	data := Error{}

	if err := json.Unmarshal(body, &data); err != nil {
		return &Error{
			StatusCode:  -1,
			Type:        "UNKNOWN",
			Description: string(body),
		}
	}

	data.StatusCode = code

	return &data
}

func (e *Error) Error() string {
	return fmt.Sprintf("[goRSO Error]: (%d %s) %s\n", e.StatusCode, e.Type, e.Description)
}

func NewErrorCustom(errType string, desc string) *Error {
	return &Error{
		StatusCode:  -1,
		Type:        errType,
		Description: desc,
	}
}

func NewErrorSystem(errType string, err error) *Error {
	return &Error{
		StatusCode:  -1,
		Type:        errType,
		Description: err.Error(),
	}
}
