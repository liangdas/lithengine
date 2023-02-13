package lithengine

import (
	"context"
	pb "github.com/liangdas/lithengine/golang"
	"sync"
)

// Metadata is our way of representing request headers internally.
// They're used at the RPC level and translate back and forth
// from Transport headers.
type Metadata struct {
	engine sync.Map
	golang sync.Map
}

// New creates an MD from a given key-values map.
func New(mds map[string]*pb.Struct) *Metadata {
	md := &Metadata{
		golang: sync.Map{},
		engine: sync.Map{},
	}
	for k, v := range mds {
		md.Set(k, v)
	}
	return md
}

// Get returns the value associated with the passed key.
func (m *Metadata) Get(key string) (*pb.Struct, bool) {
	k := key

	if v, ok := m.engine.Load(k); ok {
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
func (m *Metadata) Set(key string, value *pb.Struct) {
	if key == "" || value == nil {
		return
	}
	k := key
	m.engine.Store(k, value)
}

// Remove stores the key-value pair.
func (m *Metadata) Remove(key string) {
	if key == "" {
		return
	}
	k := key
	m.engine.Delete(k)
}

// GetExtra returns the value associated with the passed key.
func (m *Metadata) GetExtra(key string) (interface{}, bool) {
	k := key
	return m.golang.Load(k)
}

// SetExtra stores the key-value pair.
func (m *Metadata) SetExtra(key string, value interface{}) {
	if key == "" || value == nil {
		return
	}
	k := key
	m.golang.Store(k, value)
}

// RemoveExtra stores the key-value pair.
func (m *Metadata) RemoveExtra(key string) {
	if key == "" {
		return
	}
	k := key
	m.golang.Delete(k)
}

// Range iterate over element in metadata.
func (m *Metadata) Range(f func(k string, v *pb.Struct) bool) {
	m.engine.Range(func(key, value interface{}) bool {
		if okV, ok := value.(*pb.Struct); ok {
			return f(key.(string), okV)
		}
		return true
	})
}

// RangeExtra iterate over element in metadata.
func (m *Metadata) RangeExtra(f func(k string, v interface{}) bool) {
	m.golang.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

// Clone returns a deep copy of Metadata
func (m *Metadata) Clone() *Metadata {
	md := &Metadata{
		golang: sync.Map{},
		engine: sync.Map{},
	}
	m.engine.Range(func(key, value interface{}) bool {
		md.engine.Store(key, value)
		return true
	})
	m.golang.Range(func(key, value interface{}) bool {
		md.golang.Store(key, value)
		return true
	})
	return md
}

type metadataKey struct{}

// NewContext creates a new gaea with client md attached.
func NewContext(ctx context.Context, md *Metadata) context.Context {
	return context.WithValue(ctx, metadataKey{}, md)
}

// FromContext returns the server metadata in ctx if it exists.
func FromContext(ctx context.Context) (*Metadata, bool) {
	md, ok := ctx.Value(metadataKey{}).(*Metadata)
	return md, ok
}

// MergeToContext merge new metadata into ctx.
func MergeToContext(ctx context.Context, cmd map[string]*pb.Struct) context.Context {
	md, ok := FromContext(ctx)
	if ok {
		md = md.Clone()
		for k, v := range cmd {
			md.engine.Store(k, v)
		}
	} else {
		md = New(cmd)
	}
	return NewContext(ctx, md)
}

// MergeForExtra merge new metadata into ctx.
func MergeForExtra(ctx context.Context, cmd map[string]interface{}) context.Context {
	md, ok := FromContext(ctx)
	if ok {
		md = md.Clone()
	} else {
		md = New(map[string]*pb.Struct{})
	}
	for k, v := range cmd {
		md.golang.Store(k, v)
	}
	return NewContext(ctx, md)
}
