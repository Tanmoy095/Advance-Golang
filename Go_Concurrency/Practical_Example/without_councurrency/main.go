package main

import (
	"fmt"
	"log"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	comments, err := getComments(id)
	if err != nil {
		log.Fatal(err)

	}
	likes, err := getLikes(id)
	if err != nil {
		log.Fatal(err)

	}

	friends, err := getFriends(id)
	if err != nil {
		log.Fatal(err)
	}

	return &UserProfile{
		ID:       id,
		Comments: comments,
		Likes:    likes,
		Friends:  friends,
	}, nil
}
func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(userProfile)
	fmt.Println("fetching user profile took", time.Since(start))

}

func getComments(id int) ([]string, error) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Messi is the G.O.A.T",
		"Cr7 is a hard worker",
		"Nicola Tesla is the G.O.A.T",
	}
	return comments, nil
}
func getLikes(id int) (int, error) {
	time.Sleep(time.Millisecond * 200)

	likes := 32
	return likes, nil
}
func getFriends(id int) ([]int, error) {
	time.Sleep(time.Millisecond * 100)
	friends := []int{0, 1, 3, 4, 5}

	return friends, nil
}

// if we run go run main.go fetching user profile took 503.241667ms

// but with concurrency in async way it takes less time 200ms
