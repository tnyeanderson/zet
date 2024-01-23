# Print environment variables which match a pattern in Windows

Say you want to print the `http_proxy`/`https_proxy`/`no_proxy` variables in
linux (along with their uppercase equivalents). This is, of course, simple:

```
env | grep -i proxy
```

Let's say you want to do this on Windows (in Powershell). This is, of course,
needlessly and frustratingly complicated:

```
Get-ChildItem -Path Env: | Select-String proxy
```

Which can luckily be shortened to:

```
gci Env: | sls proxy
```

    #windows #why #powershell #env
