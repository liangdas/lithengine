package lithengine

import "encoding/json"

func (s *Struct) UnmarshalJSON(b []byte) (err error) {
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	hasType := false
	if st, ok := m["type"]; ok {
		s.StructType = StructType(int32(st.(float64)))
		hasType = true
	}
	if i, ok := m["int64"]; ok {
		s.Int64 = int64(i.(float64))
		if !hasType {
			s.StructType = StructType_Int64
			hasType = true
		}
	}
	if i, ok := m["string"]; ok {
		s.String_ = i.(string)
		if !hasType {
			s.StructType = StructType_String
			hasType = true
		}
	}
	if i, ok := m["double"]; ok {
		s.Double = i.(float64)
		if !hasType {
			s.StructType = StructType_double
			hasType = true
		}
	}
	if i, ok := m["bool"]; ok {
		s.Bool = i.(bool)
		if !hasType {
			s.StructType = StructType_bool
			hasType = true
		}
	}
	if i, ok := m["block"]; ok {
		s.Block = i.(string)
		if !hasType {
			s.StructType = StructType_block
			hasType = true
		}
	}
	if i, ok := m["func"]; ok {
		s.FuncId = i.(string)
		if !hasType {
			s.StructType = StructType_function
			hasType = true
		}
	}
	if i, ok := m["closure"]; ok {
		s.ClosureId = i.(string)
		if !hasType {
			s.StructType = StructType_closure
			hasType = true
		}
	}
	if i, ok := m["name"]; ok {
		s.Name = i.(string)
	}
	if i, ok := m["schema"]; ok {
		s.Schema = i.(string)
	}
	if i, ok := m["input"]; ok {
		intputs := i.([]interface{})
		s.FuncInput = []*Struct{}
		for _, ip := range intputs {
			s.FuncInput = append(s.FuncInput, MapToStruct(ip.(map[string]interface{})))
		}
	}
	if i, ok := m["args"]; ok {
		intputs := i.(map[string]map[string]interface{})
		s.Args = map[string]*Struct{}
		for k, ip := range intputs {
			s.Args[k] = MapToStruct(ip)
		}
	}
	return nil
}

func MapToStruct(m map[string]interface{}) *Struct {
	s := new(Struct)
	hasType := false
	if st, ok := m["type"]; ok {
		s.StructType = StructType(int32(st.(float64)))
		hasType = true
	}
	if i, ok := m["int64"]; ok {
		s.Int64 = int64(i.(float64))
		if !hasType {
			s.StructType = StructType_Int64
			hasType = true
		}
	}
	if i, ok := m["string"]; ok {
		s.String_ = i.(string)
		if !hasType {
			s.StructType = StructType_String
			hasType = true
		}
	}
	if i, ok := m["double"]; ok {
		s.Double = i.(float64)
		if !hasType {
			s.StructType = StructType_double
			hasType = true
		}
	}
	if i, ok := m["bool"]; ok {
		s.Bool = i.(bool)
		if !hasType {
			s.StructType = StructType_bool
			hasType = true
		}
	}
	if i, ok := m["block"]; ok {
		s.Block = i.(string)
		if !hasType {
			s.StructType = StructType_block
			hasType = true
		}
	}
	if i, ok := m["func"]; ok {
		s.FuncId = i.(string)
		if !hasType {
			s.StructType = StructType_function
			hasType = true
		}
	}
	if i, ok := m["closure"]; ok {
		s.ClosureId = i.(string)
		if !hasType {
			s.StructType = StructType_closure
			hasType = true
		}
	}
	if i, ok := m["name"]; ok {
		s.Name = i.(string)
	}
	if i, ok := m["schema"]; ok {
		s.Schema = i.(string)
	}
	if i, ok := m["input"]; ok {
		intputs := i.([]interface{})
		s.FuncInput = []*Struct{}
		for _, ip := range intputs {
			s.FuncInput = append(s.FuncInput, MapToStruct(ip.(map[string]interface{})))
		}
	}
	if i, ok := m["args"]; ok {
		intputs := i.(map[string]interface{})
		s.Args = map[string]*Struct{}
		for k, ip := range intputs {
			s.Args[k] = MapToStruct(ip.(map[string]interface{}))
		}
	}
	return s
}
