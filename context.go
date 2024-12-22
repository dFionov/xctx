package xctx

import (
	"context"
)

// fieldsKey is a private type to avoid collisions in the context.
type fieldsKey struct{}

// Fields is a dynamic map for storing arbitrary key/value data in context.
type Fields map[string]interface{}

// GetFields retrieves the Fields map from ctx. Returns nil if none is set.
func GetFields(ctx context.Context) Fields {
	if f, ok := ctx.Value(fieldsKey{}).(Fields); ok {
		return f
	}
	return nil
}

// GetField returns the value stored in ctx's Fields under the given key.
// It returns nil if the key doesn't exist or if no Fields are set.
func GetField(ctx context.Context, key string) interface{} {
	f := GetFields(ctx)
	if f == nil {
		return nil
	}
	return f[key]
}

// WithFields merges the provided newFields into a *copy* of the existing Fields
// and returns a brand new context with the updated map.
func WithFields(ctx context.Context, newFields Fields) context.Context {
	oldFields := GetFields(ctx)
	merged := cloneFields(oldFields)
	for k, v := range newFields {
		merged[k] = v
	}
	return context.WithValue(ctx, fieldsKey{}, merged)
}

// WithField is a convenience function to merge a single key-value pair into ctx.
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	oldFields := GetFields(ctx)
	merged := cloneFields(oldFields)
	merged[key] = value
	return context.WithValue(ctx, fieldsKey{}, merged)
}

// cloneFields creates a shallow copy of the provided map.
func cloneFields(orig Fields) Fields {
	if orig == nil {
		return make(Fields)
	}
	copy := make(Fields, len(orig))
	for k, v := range orig {
		copy[k] = v
	}
	return copy
}
