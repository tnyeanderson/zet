# Using local versions of go modules

## The right way

Since writing this zet, I realized how dumb I was. The correct answer to this
is to use go workspaces. For example:

```
go work init
go work use .
go work use ./gin
go work use ../otherpackage
```

Done. Much Better. But see the original content of the zet below if you dare...

## The wrong way

> WARNING: This is stupid, don't do it. Use workspaces.

Sometimes a remote module needs to be loaded locally from a directory. For
instance, when testing potential upstream changes to a dependency of your
project.

> NOTE: You can also replace a remote module with another remote module

## Example

Let's say you want to submt a patch to the `gin` library. You create the patch
and decide to try it out with your current project.

The module source code (or stubs, or whatever) can be placed, cloned, etc into
the `./gin` directory.

Then add this line in `go.mod`:
```
replace github.com/gin-gonic/gin v1.8.1 => ./gin
```

The version number is **required**.

## Related

* https://go.dev/ref/mod#go-mod-file-replace

    #go #tips #dev
