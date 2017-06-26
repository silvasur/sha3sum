sha3sum
=======

Simple program to expose `golang.org/x/crypto/sha3` on the command line (similar to  the `sha1sum` utility)

Usage
-----

Call the program with a list of filenames, these will be hashed. The filename '-' represents stdin. When no argument is passed, hash stdin. Optionally you can specify the hash size (e.g. 256, 512) with the optional `-algo` flag. See `sha3sum -help`.
