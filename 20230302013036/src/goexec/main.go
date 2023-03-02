package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Println("There should be nothing printed after this due to a crash!")

	// noop is an empty executable file
	args := []string{"./noop"}
	unix.Exec(args[0], args, nil)

	fmt.Println("This line should not be printed after noop!")
}
