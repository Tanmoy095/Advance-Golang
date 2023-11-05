package main

import "fmt"

// // i want to map username with their id
//
//	type CustomMap struct {
//		data map[string]int
//	}
//
// we are gonna use generics because key should be comparable
type CustomMap[k comparable, v any] struct {
	data map[k]v
}

// // method of CustomMAp
//
//	func (m *CustomMap) Insert(k string, v int) error {
//		m.data[k] = v
//		return nil
//	}
func (m *CustomMap[k, v]) Insert(K k, V v) error {
	m.data[K] = V
	return nil
}

// //initialize there data

//	func NewCustomMap() *CustomMap {
//		return &CustomMap{
//			data: make(map[string]int),
//		}
//	}
func NewCustomMap[k comparable, v any]() *CustomMap[k, v] {
	return &CustomMap[k, v]{
		data: make(map[k]v),
	}
}
func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("Messi", 10)
	m1.Insert("Cristiano", 7)
	m2 := NewCustomMap[int, float64]()

	m2.Insert(7, 2.4444)
	fmt.Println(m1)
	fmt.Println(m2)

}
