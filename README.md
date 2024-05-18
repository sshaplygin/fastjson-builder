# fastjson-builder

fastjson it is low-level library for working with fastjson objects. We have task for parse part of biggest json and insert some changes in parent data.

This library resolve once case - Transform golang objects to fastjson values without a lot of monkey jobs.

## Disclaimer

It is non procutions ready solution, because solution is WIP

## Problem

I am programmer and i am lazy :)

## Usage

```go
```

## Benchmarks

```
goos: darwin
goarch: arm64
pkg: github.com/sshaplygin/fastjson-builder/cmd
BenchmarkStd-10                       5713700        201.1 ns/op       96 B/op        1 allocs/op
BenchmarkManualFastJSONBuild-10       2095821        680.0 ns/op     3283 B/op        7 allocs/op
BenchmarkReflectFastJSONBuild-10      1000000       1105 ns/op     3646 B/op       19 allocs/op
```
