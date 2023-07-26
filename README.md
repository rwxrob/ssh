# Pure Go secure shell (SSH) functions

[![GoDoc](https://godoc.org/github.com/rwxrob/ssh?status.svg)](https://godoc.org/github.com/rwxrob/ssh)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Performance benchmarks

Caching the TCP/IP dial-up connection (`*ssh.Client`) is 67.72 times faster than recreating a new connection every time:

```
Uncached:   4059  88907623 ns/op  109533 B/op  570 allocs/op
Cached:   277924   1312706 ns/op   70661 B/op  145 allocs/op
```

```bc
scale=2; 88907623 / 1312706
```

```sh
go test -bench . -benchtime=5m -benchmem
goos: linux
goarch: amd64
pkg: github.com/rwxrob/ssh
cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
BenchmarkRun-8              4059          88907623 ns/op          109533 B/op        570 allocs/op
PASS
```

```sh
go test -bench . -benchtime=5m -benchmem
goos: linux
goarch: amd64
pkg: github.com/rwxrob/ssh
cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
BenchmarkRun-8            277924           1312706 ns/op           70661 B/op        145 allocs/op
PASS
```
