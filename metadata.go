package lithengine

import (
	"context"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
)

// Metadata is our way of representing request headers internally.
// They're used at the RPC level and translate back and forth
// from Transport headers.
type Metadata map[string]interface{}

// New creates an MD from a given key-values map.
func New(mds map[string]*pb.Struct) Metadata {
	md := Metadata{}
	for k, v := range mds {
		md.Set(k, v)
	}
	return md
}

// Get returns the value associated with the passed key.
func (m Metadata) Get(key string) (*pb.Struct, bool) {
	k := key
	if v, ok := m[k]; ok {
		if okV, k := v.(*pb.Struct); k {
			return okV, true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

// Set stores the key-value pair.
func (m Metadata) Set(key string, value *pb.Struct) {
	if key == "" || value == nil {
		return
	}
	k := key
	m[k] = value
}

// Remove stores the key-value pair.
func (m Metadata) Remove(key string) {
	if key == "" {
		return
	}
	k := key
	delete(m, k)
}

// GetInterface returns the value associated with the passed key.
func (m Metadata) GetInterface(key string) (interface{}, bool) {
	k := key
	k = fmt.Sprintf("__%v__", k)
	v, ok := m[k]
	return v, ok
}

// SetInterface stores the key-value pair.
func (m Metadata) SetInterface(key string, value interface{}) {
	if key == "" || value == nil {
		return
	}
	k := key
	k = fmt.Sprintf("__%v__", k)
	m[k] = value
}

// RemoveInterface stores the key-value pair.
func (m Metadata) RemoveInterface(key string) {
	if key == "" {
		return
	}
	k := key
	k = fmt.Sprintf("__%v__", k)
	delete(m, k)
}

// Range iterate over element in metadata.
func (m Metadata) Range(f func(k string, v *pb.Struct) bool) {
	for k, v := range m {
		if okV, ok := v.(*pb.Struct); ok {
			ret := f(k, okV)
			if !ret {
				break
			}
		} else {
			continue
		}
	}
}

// RangeInterface iterate over element in metadata.
func (m Metadata) RangeInterface(f func(k string, v interface{}) bool) {
	for k, v := range m {
		if okV, ok := v.(*pb.Struct); !ok {
			continue
		} else {
			ret := f(k, okV)
			if !ret {
				break
			}
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
func MergeToContext(ctx context.Context, cmd map[string]*pb.Struct) context.Context {
	md, _ := FromContext(ctx)
	md = md.Clone()
	for k, v := range cmd {
		md[k] = v
	}
	return NewContext(ctx, md)
}

// MergeForInterface merge new metadata into ctx.
func MergeForInterface(ctx context.Context, cmd map[string]interface{}) context.Context {
	md, _ := FromContext(ctx)
	md = md.Clone()
	for k, v := range cmd {
		k = fmt.Sprintf("__%v__", k)
		md[k] = v
	}
	return NewContext(ctx, md)
}
