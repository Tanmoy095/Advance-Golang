package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello Mr Messi", UpperCase))
	fmt.Println(transformString("hello Mr Messi", Prefixer("OMG G.O.A.T of football! ")))

}
