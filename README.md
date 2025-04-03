# benchmark-gin-and-fiber

My simple benchmark between Gin and Fiber.

Each test file use 1 endpoint with GET method that returning static object contain just 1 field.

## Test Command

```bash
go test -bench='.*' gin_test.go > gin.txt
```

```bash
go test -bench='.*' fiber_test.go > fiber.txt
```

## Device

```log
goos: linux
goarch: amd64
pkg: benchmark-gin-and-fiber
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
```

## Benchmark Gin Result

```log
BenchmarkGin-8   	 1943730	       616.8 ns/op	     525 B/op	       7 allocs/op
BenchmarkGin-8   	 1822292	       596.6 ns/op	     492 B/op	       7 allocs/op
BenchmarkGin-8   	 1967044	       613.6 ns/op	     524 B/op	       7 allocs/op
BenchmarkGin-8   	 1845427	       606.0 ns/op	     492 B/op	       7 allocs/op
BenchmarkGin-8   	 1862421	       602.4 ns/op	     492 B/op	       7 allocs/op
BenchmarkGin-8   	 1864510	       661.3 ns/op	     528 B/op	       7 allocs/op
BenchmarkGin-8   	 1782824	       621.0 ns/op	     493 B/op	       7 allocs/op
BenchmarkGin-8   	 1717359	       607.0 ns/op	     495 B/op	       7 allocs/op
BenchmarkGin-8   	 1863752	       612.9 ns/op	     492 B/op	       7 allocs/op
BenchmarkGin-8   	 1761206	       651.7 ns/op	     494 B/op	       7 allocs/op
PASS
ok  	command-line-arguments	18.079s
```

## Bechmark Fiber Result

```log
BenchmarkFiber-8   	  367953	      3899 ns/op	    3888 B/op	      35 allocs/op
BenchmarkFiber-8   	  391921	      3949 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  350497	      3696 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  384408	      3275 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  415675	      3198 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  379567	      3183 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  384247	      3317 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  380810	      3234 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  328096	      3488 ns/op	    3894 B/op	      36 allocs/op
BenchmarkFiber-8   	  358549	      3397 ns/op	    3894 B/op	      36 allocs/op
PASS
ok  	command-line-arguments	13.286s
```
