package test

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
)

func Place(JWT_Cookie *http.Cookie, placement sudoku.Placement) error {

	requestBody, _ := json.Marshal(placement)

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8090/game/place",
		bytes.NewBuffer(requestBody))
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

	if resp.StatusCode == 200 {
		return nil
	} else {
		return fmt.Errorf("StatusCode: %v", resp.StatusCode)
	}
}