# Print only if line matches

With the following input (`test.txt`):
```
Sacramento   California   USA
Detroit      Michigan     USA
Vancouver    BC           Canada
```

## Using `awk`:

By default, awk prints the current line on true conditions:
```bash
# Print only lines in USA
awk '$3 == "USA"' test.txt
Sacramento   California   USA
Detroit      Michigan     USA
```

Or perform a regex match:
```bash
awk '$3 ~ /[A-Z]{3}/' test.txt
Sacramento   California   USA
Detroit      Michigan     USA
```

To only print certain columns of matching lines:
```bash
# Print only city names in USA
awk '$3 == "USA" { print $1 }' test.txt
Sacramento
Detroit
```

## Using `sed`:

* `-n` means do not print by default
* `/p` command prints the matching line

Print entire matching line:
```bash
# Print only lines containing "USA"
sed -n -e '/USA/p' test.txt
Sacramento   California   USA
Detroit      Michigan     USA
```

Replace some of the text before printing only matching lines:
```bash
# Replace 'i' with 'x' on lines which contain USA
sed -n -e '/USA/s/i/x/gp' test.txt
Sacramento   Calxfornxa   USA
Detroxt      Mxchxgan     USA
```

# Print all lines, but only replace on matching lines
```bash
sed -e '/USA/s/i/x/g' test.txt
Sacramento   Calxfornxa   USA
Detroxt      Mxchxgan     USA
Vancouver    BC           Canada
```

>Notice the regex (`/USA/`) just precedes the `s/` portion of a normal `sed` replacement.

    #bash #sed #awk #tips
