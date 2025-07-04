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
	"id",
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
	//"let",
}

func MapToStruct(s *Struct, st interface{}) (*Struct, error) {
	switch st.(type) {
	case nil:
		s.StructType = StructType_nil
		break
	case string:
		s.StructType = StructType_string
		s.String_ = st.(string)
		break
	case int64:
		s.StructType = StructType_int64
		s.Int64 = st.(int64)
		break
	case bool:
		s.StructType = StructType_bool
		s.Bool = st.(bool)
		break
	case float64:
		s.StructType = StructType_double
		s.Double = st.(float64)
		break
	case []interface{}:
		s.StructType = StructType_list
		inputs, ok := st.([]interface{})
		if !ok {
			return nil, errors.New(fmt.Sprintf(`list %T not []interface{}`, st))
		}
		s.List = []*Struct{}
		for ii, ip := range inputs {
			o := new(Struct)
			ms, err := MapToStruct(o, ip)
			if err != nil {
				return nil, errors.New(fmt.Sprintf(`The value of the list[%v] %v`, ii, err))
			}
			s.List = append(s.List, ms)
		}
		break
	case map[string]interface{}:
		m := st.(map[string]interface{})
		hasType := false
		if i, ok := m["type"]; ok {
			inputs, ok := i.(float64)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`type %T not int`, i))
			}
			s.StructType = StructType(int32(inputs))
			hasType = true
		}
		if i, ok := m["int64"]; ok {
			inputs, ok := i.(float64)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`int64 %T not int64`, i))
			}
			s.Int64 = int64(inputs)
			if !hasType {
				s.StructType = StructType_int64
				hasType = true
			}
		}
		if i, ok := m["string"]; ok {
			inputs, ok := i.(string)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`string %T not string`, i))
			}
			s.String_ = inputs
			if !hasType {
				s.StructType = StructType_string
				hasType = true
			}
		}
		if i, ok := m["double"]; ok {
			inputs, ok := i.(float64)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`double %T not float64`, i))
			}
			s.Double = inputs
			if !hasType {
				s.StructType = StructType_double
				hasType = true
			}
		}
		if i, ok := m["bool"]; ok {
			inputs, ok := i.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`bool %T not bool`, i))
			}
			s.Bool = inputs
			if !hasType {
				s.StructType = StructType_bool
				hasType = true
			}
		}
		if i, ok := m["closure"]; ok {
			inputs, ok := i.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`closure %T not bool`, i))
			}
			s.Closure = inputs
		}
		if i, ok := m["nil"]; ok {
			inputs, ok := i.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`nil %T not bool`, i))
			}
			s.Bool = inputs
			if !hasType {
				s.StructType = StructType_nil
				hasType = true
			}
		}
		if i, ok := m["func"]; ok {
			inputs, ok := i.(string)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`func %T not string`, i))
			}
			s.FuncId = inputs
			if !hasType {
				s.StructType = StructType_function
				hasType = true
			}
		}
		if i, ok := m["name"]; ok {
			name, ok := i.(string)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`name %T not string`, i))
			}
			s.Name = name
		}
		if i, ok := m["optional"]; ok {
			optional, ok := i.(bool)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`optional %T not string`, i))
			}
			s.Optional = optional
		} else {
			s.Optional = true
		}
		if i, ok := m["id"]; ok {
			id, ok := i.(string)
			if !ok {
				return nil, errors.New(fmt.Sprintf(`id %T not string`, i))
			}
			s.Id = id
		}
		if i, ok := m["schema"]; ok {
			inputs, ok := i.(map[string]interface{})
			if !ok {
				return nil, errors.New(fmt.Sprintf(`schema %T not {"input":[...],"output":[...]}`, i))
			}
			fs := &FunctionSchema{}
			if input, ok := inputs["inputType"]; ok {
				i, ok := input.([]interface{})
				if !ok {
					return nil, errors.New(fmt.Sprintf(`"input" %T not [...]`, i))
				}
				var l []*Struct
				for ii, ip := range i {
					o := new(Struct)
					ms, err := MapToStruct(o, ip)
					if err != nil {
						return nil, errors.New(fmt.Sprintf(`The value of the list[%v] %v`, ii, err))
					}
					l = append(l, ms)
				}
				fs.InputType = l
			}
			if output, ok := inputs["outputType"]; ok {
				i, ok := output.([]interface{})
				if !ok {
					return nil, errors.New(fmt.Sprintf(`"output" %T not [...]`, i))
				}
				var l []*Struct
				for ii, ip := range i {
					o := new(Struct)
					ms, err := MapToStruct(o, ip)
					if err != nil {
						return nil, errors.New(fmt.Sprintf(`The value of the list[%v] %v`, ii, err))
					}
					l = append(l, ms)
				}
				fs.OutputType = l
			}
			s.Schema = fs
		}
		if i, ok := m["list"]; ok {
			inputs, ok := i.([]interface{})
			if !ok {
				return nil, errors.New(fmt.Sprintf(`list %T not []interface{}`, i))
			}
			s.List = []*Struct{}
			for ii, ip := range inputs {
				o := new(Struct)
				ms, err := MapToStruct(o, ip)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`The value of the list[%v] %v`, ii, err))
				}
				s.List = append(s.List, ms)
			}
			if !hasType {
				s.StructType = StructType_list
				hasType = true
			}
		}
		if i, ok := m["hash"]; ok {
			inputs, ok := i.(map[string]interface{})
			if !ok {
				return nil, errors.New(fmt.Sprintf(`hash %T not map[string]interface{}`, i))
			}
			s.Hash = map[string]*Struct{}
			for k, ip := range inputs {
				o := new(Struct)
				ms, err := MapToStruct(o, ip)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`hash The value of the "%v" field %v`, k, err))
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
			ms, err := MapToStruct(o, i)
			if err != nil {
				return nil, errors.New(fmt.Sprintf(`The value of the "pointer" field %v`, err))
			}
			s.Pointer = ms
			if !hasType {
				s.StructType = StructType_pointer
				hasType = true
			}
		}
		if i, ok := m["return"]; ok {
			if inputs, ok := i.([]interface{}); ok {
				s.Return = []*Struct{}
				for ii, ip := range inputs {
					o := new(Struct)
					ms, err := MapToStruct(o, ip)
					if err != nil {
						return nil, errors.New(fmt.Sprintf(`The value of the return[%v] %v`, ii, err))
					}
					s.Return = append(s.Return, ms)
				}
			} else {
				s.Return = []*Struct{}
				o := new(Struct)
				ms, err := MapToStruct(o, i)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`The value of the return %v`, err))
				}
				s.Return = append(s.Return, ms)
			}
			if !hasType {
				s.StructType = StructType_return
				hasType = true
			}
		}
		if i, ok := m["input"]; ok {
			if inputs, ok := i.([]interface{}); ok {
				s.FuncInput = []*Struct{}
				for ii, ip := range inputs {
					o := new(Struct)
					ms, err := MapToStruct(o, ip)
					if err != nil {
						return nil, errors.New(fmt.Sprintf(`The value of the input[%v] %v,array`, ii, err))
					}
					s.FuncInput = append(s.FuncInput, ms)
				}
			} else {
				s.FuncInput = []*Struct{}
				o := new(Struct)
				ms, err := MapToStruct(o, i)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`The value of the input %v`, err))
				}
				s.FuncInput = append(s.FuncInput, ms)
			}
		}
		if i, ok := m["args"]; ok {
			inputs, ok := i.(map[string]interface{})
			if !ok {
				return nil, errors.New(fmt.Sprintf(`args %T not map[string]interface{}`, i))
			}
			s.Args = map[string]*Struct{}
			for k, ip := range inputs {
				o := new(Struct)
				ms, err := MapToStruct(o, ip)
				if err != nil {
					return nil, errors.New(fmt.Sprintf(`args The value of the "%v" field %v`, k, err))
				}
				s.Args[k] = ms
			}
		}
		//if i, ok := m["let"]; ok {
		//	inputs, ok := i.(map[string]interface{})
		//	if !ok {
		//		return nil, errors.New(fmt.Sprintf(`let %T not map[string]interface{}`, i))
		//	}
		//	s.Let = map[string]*Struct{}
		//	for k, ip := range inputs {
		//		o := new(Struct)
		//		ms, err := MapToStruct(o, ip)
		//		if err != nil {
		//			return nil, errors.New(fmt.Sprintf(`let The value of the "%v" field %v`, k, err))
		//		}
		//		s.Let[k] = ms
		//	}
		//}

		for k, i := range m {
			isReservedField := false
			for _, field := range ReservedFields {
				if field == k {
					isReservedField = true
					break
				}
			}
			if !isReservedField {
				if s.FuncId != "" && s.FuncId != k {
					return nil, errors.New(fmt.Sprintf(`'%v' and '%v' cannot exist in a block`, s.FuncId, k))
				}
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
						ms, err := MapToStruct(o, ip)
						if err != nil {
							return nil, err
						}
						s.FuncInput = append(s.FuncInput, ms)
					}
				} else {
					s.FuncInput = []*Struct{}
					o := new(Struct)
					ms, err := MapToStruct(o, i)
					if err != nil {
						return nil, errors.New(fmt.Sprintf(`The value of the "%v" field %v`, k, err))
					}
					s.FuncInput = append(s.FuncInput, ms)
				}
			}
		}
		break
	default:
		return nil, errors.New(fmt.Sprintf(`must be an string,int64,bool,double,[...],{...}`))
	}
	return s, nil
}
