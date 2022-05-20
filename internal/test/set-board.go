package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
)

func SetBoard(JWT_Cookie *http.Cookie, board *sudoku.Board) error {

	requestBody, _ := json.Marshal(board)

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8090/game/setboard",
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
