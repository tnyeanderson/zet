# Re-unmarshaling YAML into the same golang struct

If you unmarshal YAML into a struct using [yaml.v3](gopkg.in/yaml.v3), then
unmarshal different YAML into the same struct, what happens? Are nested structs
reinitialized? Will a map contain only the keys from the latest unmarshaling?
Turns out the answer is no for both questions.

>If an internal pointer within a struct is not initialized, the yaml package
will initialize it if necessary for unmarshalling the provided data.

The package docs indicate that nested structs/maps will only be initialized
once (and only if "necessary"). What are the implications of this?

The following statements are true at any level of nesting:

1.  If a pointer within a struct was previously initialized (e.g. by a previous
    unmarshaling), its address will never change from subsequent unmarshaling,
    unless the value is explicitly set to `null` in the YAML.

2.  An initialized pointer within a struct will be set to `nil` if it is
    explicitly set to `null` in the YAML.

3.  Fields of an initialized struct may be updated to have new values, but
    a new struct/map will never be initilized and replace an existing one.

4.  Struct fields that are not present in the YAML will be unchanged when
    unmarshaling.

5.  Non-nullable struct fields that are explicitly set to `null` in the YAML
    will be unchanged when unmarshaling.

6.  Map keys may be updated and added, but never removed.

7.  Map keys which have pointer values that are explicitly set to `null` in the
    YAML will result in a key with a value of `nil` in the map.

8.  Non-nullable keys which are explicitly set to `null` in the YAML will have
    be unchanged when unmarshaling.

9.  Existing map keys that are missing from the YAML will be unchanged when
    unmarshaling.

10. Map fields in structs will not be initialized if missing from the YAML.

11. Map fields in structs will be initialized even if set to an empty object in
    the YAML.

All of the above points are demonstrated by the `yamlupdate` package
[embedded](./yamlupdate) in this zet.

Each test corresponds to a number in the above list and is a self-contained
demonstration of the expected behavior. A test with an added layer of nesting
is available for each as well. To run the tests:

```bash
$ go test -v .
=== RUN   Test1
--- PASS: Test1 (0.00s)
=== RUN   Test1Nested
--- PASS: Test1Nested (0.00s)
=== RUN   Test2
--- PASS: Test2 (0.00s)
=== RUN   Test2Nested
--- PASS: Test2Nested (0.00s)
=== RUN   Test3
--- PASS: Test3 (0.00s)
=== RUN   Test3Nested
--- PASS: Test3Nested (0.00s)
=== RUN   Test4
--- PASS: Test4 (0.00s)
=== RUN   Test4Nested
--- PASS: Test4Nested (0.00s)
=== RUN   Test5
--- PASS: Test5 (0.00s)
=== RUN   Test5Nested
--- PASS: Test5Nested (0.00s)
=== RUN   Test6
--- PASS: Test6 (0.00s)
=== RUN   Test6Nested
--- PASS: Test6Nested (0.00s)
=== RUN   Test7
--- PASS: Test7 (0.00s)
=== RUN   Test7Nested
--- PASS: Test7Nested (0.00s)
=== RUN   Test8
--- PASS: Test8 (0.00s)
=== RUN   Test8Nested
--- PASS: Test8Nested (0.00s)
=== RUN   Test9
--- PASS: Test9 (0.00s)
=== RUN   Test9Nested
--- PASS: Test9Nested (0.00s)
=== RUN   Test10
--- PASS: Test10 (0.00s)
=== RUN   Test10Nested
--- PASS: Test10Nested (0.00s)
=== RUN   Test11
--- PASS: Test11 (0.00s)
=== RUN   Test11Nested
--- PASS: Test11Nested (0.00s)
PASS
ok  	yamlupdate	0.003s
```

    #go #yaml #tip
