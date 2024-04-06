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

12. If the existing value of a map is a pointer to a struct and the same key
    exists in the YAML, a new struct will be initialized and the pointer will
    be changed. This is probably the least intuitive item in this list,
    especially when contrasted with #1 and #13.

13. If the existing value of a map is a struct, that struct will be updated
    in-place with any values set in the YAML.

14. Slices will always be replaced with the new slice if present in the YAML. 

All of the above points are demonstrated by the `yamlupdate` package
[embedded](./yamlupdate) in this zet.

Each test corresponds to a number in the above list and is a self-contained
demonstration of the expected behavior. A test with an added layer of nesting
is available for each as well. To run the tests:

```bash
go test -v .
```

    #go #yaml #tip
