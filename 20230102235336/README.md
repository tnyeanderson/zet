# Prompt user for input in go

The generic `fmt` tools can't capture an entire line of response in one
variable. To do that, try this:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func ask(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s ", prompt)
	scanner.Scan()
	return scanner.Text()
}
```

    #go #prompt #cli
