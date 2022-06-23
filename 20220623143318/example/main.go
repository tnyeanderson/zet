package main

import (
	"childmodule"
	"commonmodule"
	"fmt"
)

func main() {
	fmt.Println("-- IN MAIN MODULE --")
	fmt.Println(commonmodule.ValueTest)
	childmodule.ChangeCommonModuleValue()
	fmt.Println("-- IN MAIN MODULE --")
	fmt.Println(commonmodule.ValueTest)
}
