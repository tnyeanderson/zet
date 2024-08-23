# Generate a 32 byte base64 encoded string

Use this python one-liner:

```bash
python3 -c 'import secrets;import base64;print(base64.b64encode(secrets.token_bytes(32)).decode("ASCII"))'
```

This is helpful for generating ESPHome device encryption keys.

    #iot #random #secret
