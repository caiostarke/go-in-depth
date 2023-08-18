package main

import (
	"fmt"
	"sync"
)

type UserProfile struct {
	ID       int
	Likes    int
	Comments []string
	Friends  []int
}

type UserData struct {
	data any
	err  error
}

func GetLikes(id int, data chan UserData, wg *sync.WaitGroup) {
	data <- UserData{
		data: 239,
		err:  nil,
	}
	wg.Done()
}

func GetComments(id int, data chan UserData, wg *sync.WaitGroup) {
	data <- UserData{
		data: []string{"c1", "c2", "c3", "c4", "c5"},
		err:  nil,
	}

	wg.Done()
}

func GetFriends(id int, data chan UserData, wg *sync.WaitGroup) {
	data <- UserData{
		data: []int{1, 2, 3, 4},
		err:  nil,
	}

	wg.Done()
}

func handleUserProfile(id int) *UserProfile {
	var (
		wg    *sync.WaitGroup = &sync.WaitGroup{}
		msgch chan UserData   = make(chan UserData, 3)
	)

	go GetLikes(1, msgch, wg)
	go GetComments(1, msgch, wg)
	go GetFriends(1, msgch, wg)
	wg.Add(3)
	wg.Wait()
	close(msgch)

	up := &UserProfile{}
	up.ID = id

	for resp := range msgch {
		switch val := resp.data.(type) {
		case int:
			up.Likes = val
		case []string:
			up.Comments = val
		case []int:
			up.Friends = val
		}
	}

	return up
}

func main() {
	up := handleUserProfile(1)

	fmt.Println(up)
}
