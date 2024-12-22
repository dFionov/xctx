package xctx

import (
	"context"
	"testing"
)

// BenchmarkWithField measures the performance of adding a single key/value.
func BenchmarkWithField(b *testing.B) {
	baseCtx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Each iteration: create a new context with one additional field
		_ = WithField(baseCtx, "key", i)
	}
}

// BenchmarkWithFields measures adding multiple fields at once.
func BenchmarkWithFields(b *testing.B) {
	baseCtx := context.Background()
	newFields := Fields{
		"trace_id":   "abc123",
		"request_id": "req-001",
		"user_id":    "user-foo",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = WithFields(baseCtx, newFields)
	}
}

// BenchmarkGetFields measures the speed of retrieving the entire map.
func BenchmarkGetFields(b *testing.B) {
	// Prepare a context with some fields
	ctx := context.Background()
	ctx = WithFields(ctx, Fields{
		"trace_id":   "abc123",
		"request_id": "req-001",
	})

	b.ReportAllocs()
	b.ResetTimer()

	var f Fields
	for i := 0; i < b.N; i++ {
		f = GetFields(ctx)
	}

	// Just to use f so it won't get optimized away
	_ = f
}

// BenchmarkGetField measures the speed of retrieving a single key from the map.
func BenchmarkGetField(b *testing.B) {
	// Prepare a context with some fields
	ctx := context.Background()
	ctx = WithFields(ctx, Fields{
		"trace_id":   "abc123",
		"request_id": "req-001",
	})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = GetField(ctx, "trace_id")
	}
}
