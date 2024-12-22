# xctx

xctx is a Go package for managing context fields in a structured and efficient manner.

## Installation

To install xctx, use `go get`:

```sh
go get github.com/dFionov/xctx
```

## Usage

### Adding Fields to Context

You can add fields to a context using `WithField` or `WithFields`:

```go
ctx := context.Background()
ctx = xctx.WithField(ctx, "key", "value")

fields := xctx.Fields{
    "trace_id": "abc123",
    "user_id": "user-foo",
}
ctx = xctx.WithFields(ctx, fields)
```

### Retrieving Fields from Context

To retrieve fields from a context, use `GetFields` or `GetField`:

```go
fields := xctx.GetFields(ctx)
traceID := xctx.GetField(ctx, "trace_id")
```

## Benchmarks

The package includes benchmarks to measure performance:

```sh
go test -bench .
```

```
goos: windows
goarch: amd64
pkg: github.com/dFionov/xctx
cpu: AMD Ryzen 7 6800H with Radeon Graphics
BenchmarkWithField-16            6226188               169.0 ns/op           392 B/op          3 allocs/op
BenchmarkWithFields-16           4907156               235.0 ns/op           384 B/op          3 allocs/op
BenchmarkGetFields-16           249881409                4.354 ns/op           0 B/op          0 allocs/op
BenchmarkGetField-16            142807279                8.275 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/dFionov/xctx 6.465s
```