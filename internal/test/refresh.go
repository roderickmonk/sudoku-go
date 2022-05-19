package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
)

func Refresh(JWT_Cookie *http.Cookie) (*sudoku.Board, error) {

	// fmt.Println("test.Refresh")

	req, err := http.NewRequest(
		"GET",
		"http://localhost:8090/game/refresh",
		nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(JWT_Cookie)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("refresh status code:", resp.StatusCode)

	if resp.StatusCode == 200 {

		var body []byte
		body, _ = ioutil.ReadAll(resp.Body)

		board := sudoku.Board{}
		json.Unmarshal(body, &board)

		return &board, nil

	} else {
		return nil, fmt.Errorf("StatusCode: %v", resp.StatusCode)
	}
}
