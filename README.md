# go-idsort

Intelligent Design Sort implementation for Go

Also see http://www.dangermouse.net/esoteric/intelligentdesignsort.html for a detailed analysis of this algorithm.

## Usage

```shell
$ go get github.com/maetthu/go-idsort
```

```go
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
```

```shell
$ go run main.go
\o/ all sorted (you just don't understand this specific ordering): [3 2 1]
```