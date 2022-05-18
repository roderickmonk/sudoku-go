package test

import (
	// "bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/stretchr/testify/assert"
)

func signIn(t *testing.T) (*http.Cookie, *sudoku.Board, error) {

	var (
		JWT_Cookie *http.Cookie
	)

	fmt.Println("Test: Test_SignIn")

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

	fmt.Println("Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)

	board := sudoku.Board{}
	json.Unmarshal(body, &board)

	return JWT_Cookie, &board, nil
}

func place(JWT_Cookie *http.Cookie, placement sudoku.Placement) error {

	fmt.Println("place")

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
	fmt.Printf("place StatusCode: %d\n", resp.StatusCode)

	if resp.StatusCode == 200 {
		return nil
	} else {
		return fmt.Errorf("StatusCode: %v", resp.StatusCode)
	}
}

func refresh(JWT_Cookie *http.Cookie) (*sudoku.Board, error) {

	fmt.Println("refresh")

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

	fmt.Println("refresh status code:", resp.StatusCode)

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

func Test_1(t *testing.T) {

	var (
		err        error
		JWT_Cookie *http.Cookie
		board      *sudoku.Board
	)

	if JWT_Cookie, board, err = signIn(t); err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	fmt.Println(JWT_Cookie)
	sudoku.PrintBoard(board)

	if err = place(JWT_Cookie, sudoku.Placement{I: 0, J: 1, Value: 4}); err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	if board, err = refresh(JWT_Cookie); err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	sudoku.PrintBoard(board)
}
