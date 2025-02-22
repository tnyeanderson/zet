# Copy a Hashicorp Vault secret to new path

Say you accidentally created a secret in Hashicorp Vault at the incorrect path:

```
kv/my/wrong/path
```

You really wanted it at:

```
kv/my/correct/path
```

You can use the following command to copy the secret contents to a new secret.
Importantly, this will NOT copy older versions of the secret, or additional
metadata associated with the original secret path!

```
vault kv get -format=json kv/my/wrong/path | jq .data.data | vault kv put kv/my/correct/path -
```

This essentially just grabs the key/value pairs from the secret in JSON format,
and creates a new secret in the desired location based on that JSON.

    #hashicorp #vault #tips #secrets
