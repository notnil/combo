# combos

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
	// [[0 1 2 3 4] [0 1 2 3 5] [0 1 2 3 6] [0 1 2 4 5] [0 1 2 4 6] [0 1 2 5 6] [0 1 3 4 5] [0 1 3 4 6] [0 1 3 5 6] [0 1 4 5 6] [0 2 3 4 5] [0 2 3 4 6] [0 2 3 5 6] [0 2 4 5 6] [0 3 4 5 6] [1 2 3 4 5] [1 2 3 4 6] [1 2 3 5 6] [1 2 4 5 6] [1 3 4 5 6] [2 3 4 5 6]]
}
```
