package sudoku

import (
	"errors"
	"fmt"
	"net/http"
)

func GetJWTCookie(resp *http.Response) (*http.Cookie, error) {

	for _, c := range resp.Cookies() {
		if c.Name == "JWT" {
			return c, nil
		}
	}

	return nil, errors.New("JWT Not Received")
}

func GetRequestJWT(req *http.Request) (string, bool) {

	fmt.Println("GetRequestJWT")

	for _, c := range req.Cookies() {
		if c.Name == "JWT" {
			return c.Value, true
		}
	}

	return "", false
}
