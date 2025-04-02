# benchmark-gin-and-fiber
My simple benchmark between Gin and Fiber.

Each test file use 1 endpoint with GET method that returning static object contain just 1 field.

## Device

```
goos: linux
goarch: amd64
pkg: benchmark-gin-and-fiber
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
```

## Benchmark Gin Result

```
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.166µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |        1.16µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |        1.35µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.145µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.147µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.186µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.159µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.143µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.165µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.313µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |        1.15µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.176µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.155µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |        1.15µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.134µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       2.025µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.186µs |       192.0.2.1 | GET      "/ping"
[GIN] 2025/04/02 - 21:31:09 | 200 |       1.169µs |       192.0.2.1 | GET      "/ping"
BenchmarkGin-8             30892             51899 ns/op             773 B/op         18 allocs/op
PASS
ok      benchmark-gin-and-fiber 2.040s
```

## Bechmark Fiber Result

```
BenchmarkFiber
BenchmarkFiber-8          327609              3551 ns/op            3891 B/op         35 allocs/op
PASS
ok      benchmark-gin-and-fiber 1.211s
```
