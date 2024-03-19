# Package scoped variables in go

When you create an exported variable at the package level in golang, its value
can be changed by any other package that imports it. Since only one instance of
the package is loaded for all imports in the project, any manipulations of that
variable will be visible to other packages which import it.

A test project was created to demonstrate this, and is contained in this zet
directory. It shows that separate modules importing the same package will all
read/write the same single instance of the exported package variable.

To verify this:

```bash
$ cd ./pkgvarstest/outside
$ go test -v
=== RUN   ExamplePackageVarsScope
--- PASS: ExamplePackageVarsScope (0.00s)
PASS
ok    outside 0.002s
```

    #golang #scope #import
