# POSIX specification for command line tools

CLI tools need to define the flags, options, arguments, and conventions
accepted by the program. In other words, the interface needs to be defined.
This can be done many ways and there are many libraries which introduce their
own conventions.

POSIX has defined standards for how its utilities must (or *shall*) be written
in Section 12 of POSIX.1-2017, entitled "Utility Conventions".

POSIX.1-2017 is another name for IEEE Std 1003.1-2017. Browse the standard at
the link below:

<https://pubs.opengroup.org/onlinepubs/9699919799/nframe.html>

Also see the GNU extension to the standard, which adds "long options" (e.g.
`--option value`):

<https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html>

    #posix #unix #standards #cli
