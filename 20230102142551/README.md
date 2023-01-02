# Random number of given length

Some languages only provide `rand()` functions that accept a maximum value, not
a minimum. If you need to generate a random number that is a specific number of
digits, the following algorithm will be helpful.

For example, if you need a 4-digit random number, the acceptable range is from
1000 to 9999. To ensure the floor is met, the minimum value (1000) will end up
being added to whatever random number is generated. This means that our random
number can have a value between 0 and 8999 (`max - min => 9999 - 1000 =>
8999`).

These values can all be determined in the following manner (psuedo-code):
```
length = 4
min = 10^(length - 1)
max = 10^length - 1
value = min + rand(max)
```

Below is an example implementation in Go (type conversions are annoying):
```go
package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
)

func getRand(length int) *big.Int {
	l := float64(length)
	min := int64(math.Pow(10, l-1))
	max := int64(math.Pow(10, l) - 1)
	r, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Fatal(err)
	}
	return big.NewInt(0).Add(r, big.NewInt(min))
}

func main() {
	// Get a 10 character random string to use as a transaction ID (nonce)
	fmt.Println(getRand(10).String())
}
```

    #go #rand #random #algo
