package lithengine

import (
	"context"
	pb "github.com/liangdas/lithengine/golang"
	"strings"
)

// Metadata is our way of representing request headers internally.
// They're used at the RPC level and translate back and forth
// from Transport headers.
type Metadata map[string]*pb.Struct

// New creates an MD from a given key-values map.
func New(mds ...map[string]*pb.Struct) Metadata {
	md := Metadata{}
	for _, m := range mds {
		for k, v := range m {
			md.Set(k, v)
		}
	}
	return md
}

// Get returns the value associated with the passed key.
func (m Metadata) Get(key string) *pb.Struct {
	k := strings.ToLower(key)
	return m[k]
}

// Set stores the key-value pair.
func (m Metadata) Set(key string, value *pb.Struct) {
	if key == "" || value == nil {
		return
	}
	k := strings.ToLower(key)
	m[k] = value
}

// Range iterate over element in metadata.
func (m Metadata) Range(f func(k string, v *pb.Struct) bool) {
	for k, v := range m {
		ret := f(k, v)
		if !ret {
			break
		}
	}
}

// Clone returns a deep copy of Metadata
func (m Metadata) Clone() Metadata {
	md := Metadata{}
	for k, v := range m {
		md[k] = v
	}
	return md
}

type metadataKey struct{}

// NewContext creates a new gaea with client md attached.
func NewContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metadataKey{}, md)
}

// FromContext returns the server metadata in ctx if it exists.
func FromContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(metadataKey{}).(Metadata)
	return md, ok
}

// MergeToContext merge new metadata into ctx.
func MergeToContext(ctx context.Context, cmd Metadata) context.Context {
	md, _ := FromContext(ctx)
	md = md.Clone()
	for k, v := range cmd {
		md[k] = v
	}
	return NewContext(ctx, md)
}
