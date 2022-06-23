# Package and module loading in go

In go, when one module is imported by many modules within a single program,
only one instance of the imported module is actually created. This instance is
shared between any other modules which import it.

For example, for a `commonmodule`:
```go
package commonmodule

var ValueTest string

func init() {
  ValueTest = "initial value"
}
```

When `commonmodule` is first loaded, `ValueTest` has a value of `initial
value`. If this is changed later by some `childmodule` within the same app:
```go
package childmodule

import (
  "commonmodule"
)

func ChangeCommonModuleValue() {
  commonmodule.ValueTest = "new value"
}
```

The `new value` will be reflected by *any other module which reads
`commonmodule.ValueTest` anywhere else in the project*.

See the enclosed `example` directory:
```bash
$ cd 20220623143318/example
$ go run main.go
-- IN MAIN MODULE --
initial value
-- IN CHILD MODULE --
initial value
-- VALUE CHANGED --
new value
-- IN MAIN MODULE --
new value
```

    #go #module #package #import #tips
