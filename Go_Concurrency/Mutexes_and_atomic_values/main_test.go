package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := &State{}
	//wg=wait group
	wg := sync.WaitGroup{}
	// for i := 0; i < 10; i++ {
	// 	state.count = i
	// }

	// instead of updating state we use go routine
	for i := 0; i < 10; i++ {
		wg.Add(1) //doing work
		//for each loop we are gonna add 1 to the wait group
		//go routine will be scheduled somewhere else async
		go func(i int) {
			state.setState(i + 1)
			wg.Done() //done work
			// and each function done its job
		}(i)
		//we can get any output if we test this...
	}
	wg.Wait()
	// wait group is gonna wait until 10 times increment or decrement
	// its gonna wait until its back to 0
	fmt.Printf("%+v\n", state)
	//go test ./... -v -count=1 if we run this command count gonna be once 9 once 5 once 1..
	//because we dont know the order of the execution

}
