# Fix Google Coral USB Vendor ID

It is very stupid expected behavior that when a Google Coral USB device is first
plugged in, it will initially be listed in `lsusb` as:

```
1a6e:089a Global Unichip Corp.
```

Only once the firmware is initialized to the device, usually by the
pycoral/libcoral library, will the ID be corrected to:

```
18d1:9302 Google Inc.
```

After much searching, here is a `Dockerfile` for an image which initializes the
Coral when it is run with the Coral USB passed into the container.

```Dockerfile
FROM ubuntu

WORKDIR /app

RUN apt-get update && apt-get install -y wget dfu-util

RUN wget https://github.com/google-coral/libedgetpu/raw/master/driver/usb/apex_latest_single_ep.bin

CMD ["dfu-util", "-D", "apex_latest_single_ep.bin", "-d", "1a6e:089a", "-R"]
```

To run it:

```bash
docker build -t coral-init
docker run --rm --device /dev/bus/usb coral-init
```

Once it completes, your Coral will have the correct USB Vendor ID.

## Resources

- <https://github.com/google-coral/edgetpu/issues/536#issuecomment-1033056048>
- <https://github.com/google-coral/webcoral/blob/615e9fe27ea63d45a97316d231273a74bbb3ea3e/Makefile#L66-L69>

    #coral #ai #ml #gpu #google
