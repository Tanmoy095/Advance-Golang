package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}
type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {

	var (
		respchn = make(chan Response, 3)
		wg      = &sync.WaitGroup{}
	)
	// we are doing 3 request inside their own go routine
	go getComments(id, respchn, wg)
	go getLikes(id, respchn, wg)
	go getFriends(id, respchn, wg)
	// adding  3  to the waitgroup
	wg.Add(3)
	wg.Wait() //block until the  wg counter == 0 we unlock
	close(respchn)

	//keep rangeing. But when to stop??

	userProfile := &UserProfile{}
	for resp := range respchn {
		if resp.err != nil {
			return nil, resp.err

		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []string:
			userProfile.Comments = msg
		case []int:
			userProfile.Friends = msg
		}

	}

	return userProfile, nil

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

func getComments(id int, respchn chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	comments := []string{
		"Messi is the G.O.A.T",
		"Cr7 is a hard worker",
		"Nicola Tesla is the G.O.A.T",
	}
	respchn <- Response{
		data: comments,
		err:  nil,
	}
	//work is done
	wg.Done()
}
func getLikes(id int, respchn chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	likes := 32

	respchn <- Response{
		data: likes,
		err:  nil,
	}
	// work is done
	wg.Done()
}
func getFriends(id int, respchn chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	friends := []int{0, 1, 3, 4, 5}
	respchn <- Response{
		data: friends,
		err:  nil,
	}
	// work is done
	wg.Done()

}
