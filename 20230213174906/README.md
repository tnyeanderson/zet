# Performance difference between well-defined structs and interfaces for json.Unmarshal

I recently heard the claim that unmarshaling JSON data to a raw `interface` has
some performance benefit over unmarshaling into a well-defined `struct` type.
The benchmarks don't bear this out.

I started with a [wonderful JSON
dataset](https://github.com/dr5hn/countries-states-cities-database/blob/master/states%2Bcities.json)
to simulate the data in question. This dataset consists of a top-level array of
objects representing geographic states, with each state having a `cities` array
containing objects representing cities. There are 4989 states and 150710 nested
cities in the data set, so it is a pretty good analog for a large API response.

For demonstration purposes, I also created a "short" version of this dataset
with only one state object along with its contained cities.

Here are the benchmark results (`go test -bench=.`):

```
BenchmarkUnmarshalStruct-12             1000000000           0.3961 ns/op
BenchmarkUnmarshalInterface-12          1000000000           0.3914 ns/op
BenchmarkUnmarshalStructShort-12        1000000000           0.0000415 ns/op
BenchmarkUnmarshalInterfaceShort-12     1000000000           0.0000640 ns/op
```

As we can see, even a huge dataset can be processed into a well-defined struct
in less than half of a nanosecond, and there is no practical difference between
using an interface and a well-defined data structure.

See the [source code](./src) for how the test is set up.

Go is fast!

    #golang #go #struct #json #benchmark
