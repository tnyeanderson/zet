# Klipsch ProMedia 2.1 control pod cable repair

The Klipsch ProMedia 2.1 control pod has cables that are not easily replaceable
by the user because they are wired directly into the PCB using wire-to-board
(WTB) connectors. This includes a standard 3.5mm aux cable and an RS232 cable
that they use to communicate with the subwoofer's control panel.

My cables got damaged from where they enter the control pod to about 1ft down
the cable. I decided to shorten the cable by trimming the damaged section,
crimping new pins to the wires, and re-using the existing connectors. In the
process, I had to spend significant time finding the part numbers for these
connectors and pins, because Klipsch support refused to provide them even
though the product is discontinued.

My controller pod part number is 1071692, and the part number for the board itself reads:

```text
54-KO2-53-F1-51 2020.07.08, Klipsch PM2.1 VR PCB REV1.1
```

Here are the part numbers for the connector housings and pins:

| Manufacturer    | Series  | Part number    | Description                                                  |
|-----------------|---------|----------------|--------------------------------------------------------------|
| TE Connectivity | AMP/HPI | 1735801-1      | Pins for female (cable end) WTB connector for RS232 cable    |
| TE Connectivity | AMP/HPI | 440129-6       | Female (cable end) WTB connector housing for RS232 cable     |
| TE Connectivity | AMP/HPI | 440054-6       | Male (board end) WTB connector housing for RS232 cable       |
| JST             | PA      | SPHD-001T-P0.5 | Pins for female (wire end) WTB connector for 3.5mm aux cable |
| JST             | PA      | PAP-03V-S      | Female (wire end) WTB connector housing for 3.5mm aux cable  |
| JST             | PA      | B03B-PASK-1    | Male (board end) WTB connector housing for 3.5mm aux cable   |

**In most cases, you only need to purchase the pins for the damaged cable, and
you can reuse the OEM connector housings.**

> IMPORTANT: My 3.5mm female connector housing was disfigured so it was
> slightly difficult to tell, but I believe the SPHD-001T-P0.5 pins are a direct
> fit to the OEM connector. If replacing the connector housings like I did, the
> PAP-03V-S female connector housing *does not* appear to fit in the OEM male
> housing on the board. The B03B-PASK-1 male connector housing will need to be
> modified by removing the manufacturer-included pins (push up from the bottom)
> and cutting off the small plastic board alignment tab on the bottom of the
> connector. Then remove the OEM connector housing by sliding it off the pins,
> and slide the replacement onto the pins.

Finally, if you don't have a crimp tool, **you will need one**. These cannot be
crimped with pliers. The iCrimp IWS-3220M will work for both pin types
mentioned above.

## Repair steps

1.  Remove the control pod from the speaker.
2.  Remove the four (#1 phillips) screws in the corners of top of the control
    pod. The screws are located underneath the foam.
3.  Remove the two screws holding the board in place.
4.  Remove the board by carefully moving it away from the 3.5mm ports on the
    side of the control pod until they are dislodged, then pull up and out of
    the housing.
5.  Unplug the existing cable(s) from the board.
6.  Note the existing order of wires in the housing. Use a needle to lift the
    white retaining tabs which are holding each of the pins in place within the
    female housing, and remove all the existing wires/pins from the housing. Be
    careful not to damage the housing.
7.  Cut off the damaged part of the cable.
8.  Use a razor blade along the length of the cable to cut off an appropriate
    amount of insulation.
9.  Unravel the stranded shielding and twist to form a wire. Remove any
    additional shielding (only on RS232 cable).
10. Put heatshrink around the main cable and around the shielding wire.
11. Strip about 1-1.5mm off the end of each remaining wire.
12. Crimp a pin onto each wire (including the shielding) using a crimp tool
    (practice this on your damaged cable portion first).
13. Push the pins into the housing, ensuring they lock in place.
14. Heatshrink to fit.
15. Reverse the disassembly steps to reassemble.

See the attached `./images` directory for more information and diagrams.

    #audio #repair #wire
