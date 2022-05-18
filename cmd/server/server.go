package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/roderickmonk/sudoku-go/internal/sudoku"
)

var users = make(map[string]string)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret []byte

func signin(w http.ResponseWriter, req *http.Request) {

	type Signin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var signin Signin

	switch req.Method {

	default:
		w.WriteHeader(http.StatusNotImplemented)

	case "POST":

		err := json.NewDecoder(req.Body).Decode(&signin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if password, ok := users[signin.Username]; ok {

			if signin.Password == password {

				fmt.Println("Signed In!")

				// Create a new token object, specifying signing method and the claims
				// you would like it to contain.
				// https://pkg.go.dev/github.com/golang-jwt/jwt#example-New-Hmac
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"foo": "bar",
					"nbf": time.Now().Unix(),
				})

				// Sign and get the complete encoded token as a string using the secret
				tokenString, err := token.SignedString(hmacSampleSecret)

				fmt.Println(tokenString, err)

				board := sudoku.MakeGame()

				sudoku.PrintBoard(&board)

				cookie := http.Cookie{
					Name:     "JWT",
					Value:    tokenString,
					Path:     "/",
					Expires:  time.Now().Add(100 * time.Minute),
					Secure:   false,
					HttpOnly: true,
					MaxAge:   90000}
				http.SetCookie(w, &cookie)

				fmt.Println("cookie: ", cookie)

				json_board, _ := json.Marshal(board)
				// Save to redis
				_redis.Set(context.TODO(), tokenString, json_board, 0)

				w.Write(json_board)

			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

		} else {
			fmt.Println("Bad Request")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// fmt.Fprintf(w, "Signin: %+v\n", user)
		// fmt.Fprintf(w, "Signin: %v\n", user)
		// fmt.Fprintf(w, "Signin.Username: %v\n", signin.Username)

	}
}

func place(w http.ResponseWriter, req *http.Request) {

	var (
		ctx   = context.Background()
		JWT   string
		ok    bool
		board *sudoku.Board
		err   error
	)

	switch req.Method {

	default:
		w.WriteHeader(http.StatusNotImplemented)

	case "POST":

		fmt.Println("place")

		// Use the token to retrieve the board from redis
		if JWT, ok = sudoku.GetRequestJWT(req); !ok {
			http.Error(w, "JWT missing", http.StatusBadRequest)
			return
		}

		if _, board, err = sudoku.GetBoard(_redis, JWT); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		sudoku.PrintBoard(board)

		// Get Placement
		var placement sudoku.Placement
		err = json.NewDecoder(req.Body).Decode(&placement)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		fmt.Println("place!")
		fmt.Println("placement", placement)

		if board[placement.I][placement.J] != 0 {
			w.WriteHeader(http.StatusForbidden)
		} else {
			// Apply the placement and then save the board back to redis
			board[placement.I][placement.J] = placement.Value
			data, _ := json.Marshal(board)
			_redis.Set(ctx, JWT, data, 0)
		}
	}
}

func refresh(w http.ResponseWriter, req *http.Request) {

	var (
		JWT   string
		ok    bool
	)

	switch req.Method {

	default:
		w.WriteHeader(http.StatusNotImplemented)

	case "GET":
		fmt.Println("refresh")

		// Use the token to retrieve the board from redis
		if JWT, ok = sudoku.GetRequestJWT(req); !ok {
			http.Error(w, "JWT missing", http.StatusBadRequest)
			return
		}

		if json_board, _, err := sudoku.GetBoard(_redis, JWT); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			// Return the current board
			w.Write([]byte(json_board))
		}
	}
}

// setBoard is required for testing purposes
func setBoard(w http.ResponseWriter, req *http.Request) {

	var (
		ctx = context.Background()
		JWT string
		ok  bool
	)

	switch req.Method {

	default:
		w.WriteHeader(http.StatusNotImplemented)

	case "GET":
		fmt.Println("refresh")

		// Use the token to ensure the user is known
		if JWT, ok = sudoku.GetRequestJWT(req); !ok {
			http.Error(w, "JWT missing", http.StatusBadRequest)
			return
		}

		_, err := _redis.Get(ctx, JWT).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if json_board, err := ioutil.ReadAll(req.Body); err == nil {
			// Save to redis
			_redis.Set(context.TODO(), JWT, json_board, 0)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var _redis *redis.Client

func main() {

	_redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Load up some canned user data
	users["user1"] = "%jL1Jt0Irq$Y"
	users["user2"] = "%jL1Jt0Irq$Y"

	http.HandleFunc("/signin", signin)
	http.HandleFunc("/game/place", place)
	http.HandleFunc("/game/refresh", refresh)
	http.HandleFunc("/game/set", setBoard)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}