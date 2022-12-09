package lithengine

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (s *Struct) Func() string {
	switch s.StructType {
	case StructType_function:
		return s.FuncId
	}
	return ""
}

func (s *Struct) UnmarshalJSON(b []byte) (err error) {
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	_, err = MapToStruct(s, m)
	if err != nil {
		return err
	}
	return nil
}

var ReservedFields = []string{
	"type",
	"int64",
	"string",
	"double",
	"bool",
	"nil",
	"func",
	"name",
	"schema",
	"list",
	"hash",
	"pointer",
	"return",
	"input",
	"args",
	"closure",
}

func MapToStruct(s *Struct, m map[string]interface{}) (*Struct, error) {
	hasType := false
	if st, ok := m["type"]; ok {
		s.StructType = StructType(int32(st.(float64)))
		hasType = true
	}
	if i, ok := m["int64"]; ok {
		s.Int64 = int64(i.(float64))
		if !hasType {
			s.StructType = StructType_int64
			hasType = true
		}
	}
	if i, ok := m["string"]; ok {
		s.String_ = i.(string)
		if !hasType {
			s.StructType = StructType_string
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
	if i, ok := m["nil"]; ok {
		s.Bool = i.(bool)
		if !hasType {
			s.StructType = StructType_nil
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
	if i, ok := m["name"]; ok {
		s.Name = i.(string)
	}
	if i, ok := m["schema"]; ok {
		s.Schema = i.(string)
	}
	if i, ok := m["list"]; ok {
		inputs := i.([]interface{})
		s.List = []*Struct{}
		for _, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			s.List = append(s.List, ms)
		}
		if !hasType {
			s.StructType = StructType_list
			hasType = true
		}
	}
	if i, ok := m["hash"]; ok {
		inputs := i.(map[string]interface{})
		s.Hash = map[string]*Struct{}
		for k, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			s.Hash[k] = ms
		}
		if !hasType {
			s.StructType = StructType_hash
			hasType = true
		}
	}
	if i, ok := m["pointer"]; ok {
		o := new(Struct)
		ms, err := MapToStruct(o, i.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		s.Pointer = ms
		if !hasType {
			s.StructType = StructType_pointer
			hasType = true
		}
	}
	if i, ok := m["return"]; ok {
		inputs := i.([]interface{})
		s.Return = []*Struct{}
		for _, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			s.Return = append(s.Return, ms)
		}
		if !hasType {
			s.StructType = StructType_return
			hasType = true
		}
	}
	if i, ok := m["input"]; ok {
		inputs := i.([]interface{})
		s.FuncInput = []*Struct{}
		for _, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			s.FuncInput = append(s.FuncInput, ms)
		}
	}
	if i, ok := m["args"]; ok {
		inputs := i.(map[string]interface{})
		s.Args = map[string]*Struct{}
		for k, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			s.Args[k] = ms
		}
	}

	for k, i := range m {
		isReservedField := false
		for _, field := range ReservedFields {
			if field == k {
				isReservedField = true
				break
			}
		}
		if !isReservedField {
			if s.FuncInput != nil {
				return nil, errors.New(fmt.Sprintf(`The "input" field cannot be used, please use {"%v":[...]}`, k))
			}
			s.FuncId = k
			if !hasType {
				s.StructType = StructType_function
				hasType = true
			}
			//不是保留字段统一设置为func类型
			if inputs, ok := i.([]interface{}); ok {
				s.FuncInput = []*Struct{}
				for _, ip := range inputs {
					o := new(Struct)
					ms, err := MapToStruct(o, ip.(map[string]interface{}))
					if err != nil {
						return nil, err
					}
					s.FuncInput = append(s.FuncInput, ms)
				}
			} else {
				return nil, errors.New(fmt.Sprintf(`The value of the "%v" field must be an array, eg {"%v":[...]}`, k, k))
			}
		}
	}
	return s, nil
}
