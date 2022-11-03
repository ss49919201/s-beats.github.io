package main

import (
	"fmt"

	"github.com/ahmetb/go-linq/v3"
)

func main() {
	s := []string{"a", "b", "c"}

	linq.From(s).
		Select(func(v interface{}) interface{} {
			return v.(string) + "!"
		}).ToSlice(&s)

	fmt.Println(s) // [a! b! c!]
}
