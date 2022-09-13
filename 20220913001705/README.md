# Search for a host in known_hosts

The `.ssh/known_hosts` file contains only hashes of hostname associations (on
many systems). To find which line contains the match in `known_hosts`, use:
```bash
ssh-keygen -F hostname
```

From the manpage:

> Search for the specified hostname (with optional port number) in a known_hosts
file, listing any occurrences found.  This option is useful to find hashed host
names or addresses and may also be used in conjunction with the -H option to
print found keys in a hashed format. 

    #ssh #tips #bash
