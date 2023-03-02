# Difference between execvp in C and unix.Exec in Go

When a C program calls `execvp()` on an executable file which has no content,
the calling application crashes. However, when an empty executable file is
called with `unix.Exec()` in golang, *the calling application continues as if
nothing happened!*

See the contained `src` directory for a demonstration, or see the results below.

`cexec/main.c`:
```c
#include <stdio.h>
#include <unistd.h>

int main() {
  printf("There should be nothing printed after this due to a crash!\n");

  // noop is an empty executable file
  char* a[] = {"./noop", NULL};
  execvp(a[0], a);

  printf("This line should not be printed after noop!\n");
}
```

Output:
```
$ (cd cexec && ./test)
There should be nothing printed after this due to a crash!
```

`goexec/main.go`:
```go
$ cat goexec/main.go
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
```

Output:
```
$ (cd goexec && ./test)
There should be nothing printed after this due to a crash!
This line should not be printed after noop!
```

Not sure if this is expected behavior, a regression, or a bug, but I definitely
didn't expect it!

    #c #golang #go #unexpected #bug
