package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "Ad Tanmoy")
	userId, err := fetchUserId(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v:%+v\n", time.Since(start), userId)
}

func fetchUserId(cntxt context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(cntxt, time.Millisecond*100)
	defer cancel()

	//context.timeout means how long we want this context to be valid
	//max 100 milisesond if the function takes 100 millisecond we are gonna cancel this

	// we can pass any values in our own context and can receive the value from another function
	val := cntxt.Value("username")
	fmt.Println("the value = ", val)
	type result struct {
		userId string
		err    error
	}
	resultchan := make(chan result, 1)

	go func() {
		res, err := thirdPartyHttpCall()
		resultchan <- result{
			userId: res,
			err:    err,
		}
	}()
	// we spin a go routine and it will work in  non blocking way
	select {
	//Done()
	//-> means the context  timeout is exceeded because it takes too long than 100 milliseccond
	//or the context has been manually cancelled...->cancel()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultchan:
		return res.userId, res.err

	}

}

func thirdPartyHttpCall() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "User id 1", nil
}
