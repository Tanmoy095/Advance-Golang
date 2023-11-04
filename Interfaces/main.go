package main

import "fmt"

//declaration of the interface

type NumberStorer interface {
	GetAll() ([]int, error)
	put(number int) error
}

// implementation of interface..implement these two function getall and pu
type MongoDBNumberStore struct {
	//some values
}

// connect to posgres
type PosgresNumberStore struct {
	//posgres values (db connection)
}

// attach function to structure.....which called golang function reciver
func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}
func (m MongoDBNumberStore) put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}
func (p PosgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 4, 6, 7, 8, 9, 9, 3}, nil
}
func (p PosgresNumberStore) put(number int) error {
	fmt.Println("store the number into the Postgres storage...")
	return nil
}

// strp 2.. we are telling api server that we have a interface numberstorer
type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: PosgresNumberStore{},
	}
	//logic........

	if err := apiServer.numberStore.put(1); err != nil {
		panic(err)
	}
	//get all returns two values slice of intiger and err and we assign it as numbers and error
	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)

	}

	//no err

	fmt.Println(numbers)
}
