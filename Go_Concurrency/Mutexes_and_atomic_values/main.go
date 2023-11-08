package main

import (
	"sync"
)

// state set or update or delete....
type State struct {
	mu    sync.Mutex
	count int
	//count int32//for atomic values
}

func (s *State) setState(i int) {
	s.mu.Lock()         //u have the yellow flag if u dont unlock he will never leave yeallow flag
	defer s.mu.Unlock() //executes just before the function return
	s.count = i
	//s.mu.Unlock()
}

// we can also use atomic values for race condition..
// func (s *State) setState(i int) {
// 	atomic.AddInt32(&count, i)

// }

func main() {

}
