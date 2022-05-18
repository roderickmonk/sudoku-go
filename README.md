# Demo Sudoku Software (Go)

## Purpose

The purpose of this repo is to demonstrate typical Go implementation of a Sudoku server and testing of such.  A test client is also available that puts the server through its paces.

## Theory of Operation
The signing in process assigns a token to the session and this token is constant throughout the session and is returned to client software as a cookie.  Thereafter all subsequent API calls require the use of this token.  Also, at the point of the sign-in, a new game board is created and also returned to the user. This board is  recorded to a Redis database instance using the assigned token as key.

## Software Versions
    Go: 1.18.1
    Redis: Any version

## Running the test software
After cloning, the test software can be run as follows:

    $ cd sudoku-go
    $ go get -u ./...   # Install all required packages
    $ source project  
    $ goserver          # An alias for go run cmd/server/server.go      

## API

### POST /signin
/signin requires a username and a password. Two users are known to the server, `user1` and `user2`, both having the same password. Besides creating a game for a newly signed in user, /signin returns a Json Web Token (JWT) and thereafter the JWT, which is managed as a cookie, is used to locate the board that the user is currently playing.

### GET /game/refresh
There is an underlying assumption that client software is managing a parallel game board, which means that the server is not routinely returning the board to the client every time that the user places a number.  Nevertheless, /game/refresh allows the client software to retrieve the server's board should a resynch be necessary.  

### POST /game/place
/game/place allows client software to place a number on the board.  However, illegal placements are not permitted and will be rejected as a bad request.  Checks that are applied are the following:

    1. The placement cell is available.
    2. The placement value does not already exist in the placement row.
    3. The placement value does not already exist in the placement column.
    4. The placement value does not already exist in the placement box.

### POST /game/set
/game/set is strictly for testing purposes.  It allows tests to be composed using canned boards which /game/set then sends to the server.  The server overwrites any existing board that may already exist with the incoming board.  To use this endpoint, the user must be signed-in.

## Candidate Security Improvements
*   Use https
*   User 2FA
*   Timeout the tokens
*   Whitelist the AWS ec2's source IP addresses for port 8090

## Candidate Performance Improvements
*   Assuming each board requires 400 bytes, this would require 4Gb for 10m users, which is doable with a sufficiently large ec2.
*   Box conflict testing improvements: memorize the boxes so that they do not need to be redetermined on each placement.  Possibly this could be done on an LRU basis, that is, retaining boxes in a Map() and only re-determining the boxes should the board need to be retrieved from Redis.
*   Use a load balancer that would farm off connections to child servers.

# Testing

    $ go test ./test








