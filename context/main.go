package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	context := context.WithValue(context.Background(), "username", "lololol")
	res, err := fetchUserID(context)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Time took: %v, value got: %v", time.Since(now), res)
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	fmt.Println(ctx.Value("username"))

	type user struct {
		userID string
		err    error
	}

	msgch := make(chan user, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		msgch <- user{
			userID: res,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case msg := <-msgch:
		return msg.userID, msg.err
	}

}

func thirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 101)
	return "User ID: 1", nil
}
