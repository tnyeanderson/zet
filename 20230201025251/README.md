# G602 mouse programming

The Logitech G602 no longer has any support for MacOS 11+. Logi Options, G HUB,
or Logitech Gaming Software 9.02.22 **do not work**.

There is no linux solution for pushing the XML-formatted `.dat` file to the
persistent storage on the mouse. Piper/libratbag do not work with the G602.

The only way is to create a Windows VM (7 preferred, less bloat in the logi
app) and Logitech Gaming Software (e.g. version 9.04.49).

Pass the mouse through Virtualbox using the graphical interface:

- Enable USB Controller
  - USB 1.1 (OHCI) Controller
- USB Device Filters
  - Plug in device, click add filter, select the device

> NOTE: If using two of the same mouse, one for the host machine and one for
the VM, you'll need to add a `Port` to the filter. Find the port with
`vboxmanage list usbhost`

Now open the application, don't allow it internet access (why would you?) and
import the `.dat` file into the device. Done! So stupid! The protocol should be
well-documented by the manufacturer and the tools should be open source!

    #virtualbox #vbox #mouse #g602 #macros
