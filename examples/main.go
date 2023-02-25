package main

import (
	"fmt"

	"github.com/maetthu/go-idsort"
)

func main() {
	unsorted := []int{3, 2, 1}
	sorted := idsort.IDsort(unsorted)

	fmt.Printf("\\o/ all sorted (you just don't understand this specific ordering): %v\n", sorted)
}
