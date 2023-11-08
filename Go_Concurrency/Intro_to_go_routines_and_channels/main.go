package main

import (
	"fmt"
	"time"
)

func FetchRecource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("the result is : %d", n)
}
func main() {
	// result := FetchRecource(1)
	// fmt.Println(result)

	// we are going to get an output or can function call and see what output is

	//but problem is we are no gonna get result with go routine
	//then we need chnnel
	go func() {
		result := FetchRecource(1)
		fmt.Println(result)
	}()
	//go FetchRecource() // this is an Async Call
	// Here we are not gonna get the result directly ..because its asynchronous
	//we need channel
	// Channel in more like tunnels .. we can put some in channel and recive otherside
	// More like way of communicating
}

// ALL THESE OPERATION TAKES 5 SECONDS ..SO WE ARE GOING TO WAIT VERY LONG TIME....
// WE can turned fetchRecource into go Routine
// we need to prepend go
// By this we are going to tell golang that fetchRecource right now is going to get scheduled
//somewhere else as a go routione
