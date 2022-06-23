package childmodule

import (
	"commonmodule"
	"fmt"
)

func ChangeCommonModuleValue() {
	fmt.Println("-- IN CHILD MODULE --")
	fmt.Println(commonmodule.ValueTest)
	commonmodule.ValueTest = "new value"
	fmt.Println("-- VALUE CHANGED --")
	fmt.Println(commonmodule.ValueTest)
}
