# Testable examples

This directory contains a number of testable examples that are used to validate the `rwxrob/ssh` package in a comprehensive way (that cannot be done through other Go testing).

## RunOnAny: all up

1. Spin up three containers running SSH on different ports with different host names.
1. Spin up a fourth container that just runs the `runonany` binary pointing to each of the other three servers.
1. Run `hostname` and `cat` (with three lines of standard input) on each and capture to log.

## RunOnAny: one down temporarily

### RunOnAny: one down permanently

### RunOnAny: two down

### RunOnAny: all down
