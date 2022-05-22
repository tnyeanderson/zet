# Embedded struct methods calling containing struct methods in golang

See `golang-embedded-structs-question.go` for an example.

In go, since there is no such thing as inheritance, methods of an embedded struct cannot
be made to call methods of the containing struct.

Key things to remember:
>When working with structs and embedding, everything is STATICALLY LINKED. All references
>are resolved at compile time.

>It is important to note that all inherited methods are called on the hidden-field-struct.
>It means that a base method cannot see or know about derived methods or fields. Everything
>is non-virtual.

Related:
* https://github.com/luciotato/golang-notes/blob/master/OOP.md
* https://medium.com/@gianbiondi/interfaces-in-go-59c3dc9c2d98#.7l72q6qwg
* https://gist.github.com/tnyeanderson/bbc50995e5d447fa6c3819f72f46618d

    #go #tips #polymorphism #oo #inheritance
