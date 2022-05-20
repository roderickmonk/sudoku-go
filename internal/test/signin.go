package test

import (

	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"testing"
	"errors"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
)

func SignIn(t *testing.T) (*http.Cookie, *sudoku.Board, error) {

	var (
		JWT_Cookie *http.Cookie
	)

	requestBody, _ := json.Marshal(map[string]string{
		"username": "user1",
		"password": "%jL1Jt0Irq$Y",
	})

	resp, err := http.Post(
		"http://localhost:8090/signin",
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if JWT_Cookie, err = sudoku.GetJWTCookie(resp); err != nil {
		return nil, nil, errors.New("JWT Cookie Not Received")
	}

	body, err := ioutil.ReadAll(resp.Body)

	board := sudoku.Board{}
	json.Unmarshal(body, &board)

	return JWT_Cookie, &board, nil
}