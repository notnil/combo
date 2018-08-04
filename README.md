# combos

[![Build Status](https://travis-ci.org/notnil/combos.svg?branch=master)](https://travis-ci.org/notnil/combos)
[![Go Doc](https://godoc.org/github.com/notnil/combos?status.svg)](https://godoc.org/github.com/notnil/combos)
[![Go Report Card](https://goreportcard.com/badge/github.com/notnil/combos)](https://goreportcard.com/report/github.com/notnil/combos)
[![codecov](https://codecov.io/gh/notnil/combos/branch/master/graph/badge.svg)](https://codecov.io/gh/notnil/combos)

combos is a simple combinations package

## Usage

```go
package main

import (
	"fmt"

	"github.com/notnil/combos"
)

func main() {
	fmt.Println(combos.New(7, 5))
	// [[0 1 2 3 4] [0 1 2 3 5] [0 1 2 3 6] [0 1 2 4 5] [0 1 2 4 6] [0 1 25 6] [0 1 3 4 5] [0 1 3 4 6] [0 1 3 5 6] [0 1 4 5 6] [0 2 3 4 5] [0 2 3 4 6] [0 2 3 5 6] [0 2 4 5 6] [0 3 4 5 6] [1 2 3 4 5] [1 2 3 4 6] [1 2 3 5 6] [1 2 4 5 6] [1 3 4 5 6] [23 4 5 6]]
}
```

## Benchmark

```bash
~ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/notnil/combos
Benchmark10C2-8     	  500000	      2628 ns/op	    4080 B/op	      62 allocs/op
Benchmark100C2-8    	   10000	    219910 ns/op	  556416 B/op	    4980 allocs/op
Benchmark1000C2-8   	      20	  89627977 ns/op	78746426 B/op	  499554 allocs/op
PASS
ok  	github.com/notnil/combos	5.582s
```