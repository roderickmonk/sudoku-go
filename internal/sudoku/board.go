package sudoku

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func GetBoard(_redis *redis.Client, JWT string) (string, *Board, error) {

	var ctx = context.Background()

	if json_board, err := _redis.Get(ctx, JWT).Result(); err != nil {
		return "", nil, err

	} else {

		board := Board{}
		json.Unmarshal([]byte(json_board), &board)

		return json_board, &board, nil
	}
}
