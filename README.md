# benchmark-gin-and-fiber

My simple benchmark between Gin and Fiber.

Each test file use 1 endpoint with GET method that returning static object contain just 1 field.

## Device

```log
goos: linux
goarch: amd64
pkg: benchmark-gin-and-fiber
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
```

## Benchmark Gin Result

```log
BenchmarkGin
BenchmarkGin-8           1806726               622.0 ns/op           493 B/op          7 allocs/op
PASS
ok      benchmark-gin-and-fiber 1.810s
```

## Bechmark Fiber Result

```log
BenchmarkFiber
BenchmarkFiber-8          327609              3551 ns/op            3891 B/op         35 allocs/op
PASS
ok      benchmark-gin-and-fiber 1.211s
```
