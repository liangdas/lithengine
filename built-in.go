package lithengine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
	"strings"
)

var _funcMap map[string]Function
var _blockMap map[string]*pb.Struct

func init() {
	_funcMap = map[string]Function{
		"add":      Add,
		"reduce":   Reduce,
		"multiply": Multiply,
		"divide":   Divide,
		"eq":       Eq,
		"gt":       Gt,
		"gte":      Gte,
		"lt":       Lt,
		"lte":      Lte,
		"and":      And,
		"or":       OR,
		"+":        Add,
		"-":        Reduce,
		"*":        Multiply,
		"/":        Divide,
		"=":        Eq,
		">":        Gt,
		">=":       Gte,
		"<":        Lt,
		"<=":       Lte,
		"&&":       And,
		"||":       OR,
		"not":      Not,
		"if":       If,
		"case":     Case,
		"int64":    Int64,
		"getArgs":  Args,
		"isType":   IsType,
		"in":       In,
		"chain":    Chain,
		"getHash":  GetHash,
	}
	_blockMap = map[string]*pb.Struct{}
}

func Args(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_string {
		return nil, errors.New(fmt.Sprintf("%v not be string", a.StructType.String()))
	}
	name := a.String_
	m, ok := FromContext(context)
	if !ok {
		return nil, errors.New("no args")
	}
	if r, ok := m[name]; !ok {
		return nil, errors.New(fmt.Sprintf("args no '%v' variables", name))
	} else {
		return []*pb.Struct{
			r,
		}, nil
	}
}

func IsType(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) != 2 {
		return nil, errors.New("is input len  != 2")
	}

	a := inputs[0]
	b := inputs[1]

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}
	return []*pb.Struct{&pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       true,
	}}, nil
}

func Add(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     0,
	}
	for _, input := range inputs {
		switch input.StructType {
		case pb.StructType_int64:
			output.Double += float64(input.Int64)
		case pb.StructType_double:
			output.Double += input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't add")
		case pb.StructType_string:
			return nil, errors.New("string can't add")
		case pb.StructType_function:
			o, err := e.Exec(context, input)
			if err != nil {
				return nil, err
			}
			oo, err := e.BaseFunctionMore2One(context, Add, []*pb.Struct{
				output, o,
			})
			if err != nil {
				return nil, err
			}
			output = oo
		default:
			return nil, errors.New(fmt.Sprintf("%v can't add", input.StructType.String()))
		}
	}
	return []*pb.Struct{output}, nil
}

func Reduce(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     0,
	}
	for _, input := range inputs {
		switch input.StructType {
		case pb.StructType_int64:
			output.Double -= float64(input.Int64)
		case pb.StructType_double:
			output.Double -= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't reduce")
		case pb.StructType_string:
			return nil, errors.New("string can't reduce")
		case pb.StructType_function:
			o, err := e.Exec(context, input)
			if err != nil {
				return nil, err
			}
			oo, err := e.BaseFunctionMore2One(context, Reduce, []*pb.Struct{
				output, o,
			})
			if err != nil {
				return nil, err
			}
			output = oo
		default:
			return nil, errors.New(fmt.Sprintf("%v can't reduce", input.StructType.String()))
		}
	}
	return []*pb.Struct{output}, nil
}

func Multiply(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     0,
	}
	for _, input := range inputs {
		switch input.StructType {
		case pb.StructType_int64:
			output.Double *= float64(input.Int64)
		case pb.StructType_double:
			output.Double *= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't multiply")
		case pb.StructType_string:
			return nil, errors.New("string can't multiply")
		case pb.StructType_function:
			o, err := e.Exec(context, input)
			if err != nil {
				return nil, err
			}
			oo, err := e.BaseFunctionMore2One(context, Multiply, []*pb.Struct{
				output, o,
			})
			if err != nil {
				return nil, err
			}
			output = oo
		default:
			return nil, errors.New(fmt.Sprintf("%v can't multiply", input.StructType.String()))
		}
	}
	return []*pb.Struct{output}, nil
}

func Divide(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     0,
	}
	for _, input := range inputs {
		switch input.StructType {
		case pb.StructType_int64:
			output.Double /= float64(input.Int64)
		case pb.StructType_double:
			output.Double /= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't divide")
		case pb.StructType_string:
			return nil, errors.New("string can't divide")
		case pb.StructType_function:
			o, err := e.Exec(context, input)
			if err != nil {
				return nil, err
			}
			oo, err := e.BaseFunctionMore2One(context, Divide, []*pb.Struct{
				output, o,
			})
			if err != nil {
				return nil, err
			}
			output = oo
		default:
			return nil, errors.New(fmt.Sprintf("%v can't divide", input.StructType.String()))
		}
	}
	return []*pb.Struct{output}, nil
}

func Int64(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_int64,
		Int64:      0,
	}
	if len(inputs) != 2 {
		return nil, errors.New("DoubleToInt64 input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_double && a.StructType != pb.StructType_int64 {
		return nil, errors.New(fmt.Sprintf("%v not be int64 or double", a.StructType.String()))
	}
	switch a.StructType {
	case pb.StructType_double:
		output.Int64 = int64(a.Double)
		break
	case pb.StructType_int64:
		output.Int64 = a.Int64
	}

	return []*pb.Struct{output}, nil
}

func Eq(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 2 {
		return nil, errors.New("eq input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}

	switch a.StructType {
	case pb.StructType_int64:
		output.Bool = a.Int64 == b.Int64
	case pb.StructType_double:
		output.Bool = a.Double == b.Double
	case pb.StructType_bool:
		output.Bool = a.Bool == b.Bool
	case pb.StructType_string:
		output.Bool = strings.Compare(a.String_, b.String_) == 0
	case pb.StructType_nil:
		output.Bool = true
	default:
		return nil, errors.New(fmt.Sprintf("%v cannot be eq", a.StructType.String()))
	}

	return []*pb.Struct{output}, nil
}

func Gt(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 2 {
		return nil, errors.New("gt input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}

	switch a.StructType {
	case pb.StructType_int64:
		output.Bool = a.Int64 > b.Int64
	case pb.StructType_double:
		output.Bool = a.Double > b.Double
	case pb.StructType_string:
		output.Bool = strings.Compare(a.String_, b.String_) == 1
	default:
		return nil, errors.New(fmt.Sprintf("%v cannot be gt", a.StructType.String()))
	}

	return []*pb.Struct{output}, nil
}

func Gte(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 2 {
		return nil, errors.New("gte input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}

	switch a.StructType {
	case pb.StructType_int64:
		output.Bool = a.Int64 >= b.Int64
	case pb.StructType_double:
		output.Bool = a.Double >= b.Double
	case pb.StructType_string:
		output.Bool = strings.Compare(a.String_, b.String_) == 1 || strings.Compare(a.String_, b.String_) == 0
	default:
		return nil, errors.New(fmt.Sprintf("%v cannot be gte", a.StructType.String()))
	}

	return []*pb.Struct{output}, nil
}

func Lt(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 2 {
		return nil, errors.New("lt input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}

	switch a.StructType {
	case pb.StructType_int64:
		output.Bool = a.Int64 < b.Int64
	case pb.StructType_double:
		output.Bool = a.Double < b.Double
	case pb.StructType_string:
		output.Bool = strings.Compare(a.String_, b.String_) == -1
	default:
		return nil, errors.New(fmt.Sprintf("%v cannot be lt", a.StructType.String()))
	}

	return []*pb.Struct{output}, nil
}

func Lte(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 2 {
		return nil, errors.New("lte input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}
	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != b.StructType {
		return nil, errors.New(fmt.Sprintf("%v %v cannot be compared", a.StructType.String(), b.StructType.String()))
	}

	switch a.StructType {
	case pb.StructType_int64:
		output.Bool = a.Int64 <= b.Int64
	case pb.StructType_double:
		output.Bool = a.Double <= b.Double
	case pb.StructType_string:
		output.Bool = strings.Compare(a.String_, b.String_) == -1 || strings.Compare(a.String_, b.String_) == 0
	default:
		return nil, errors.New(fmt.Sprintf("%v cannot be lte", a.StructType.String()))
	}

	return []*pb.Struct{output}, nil
}

func And(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	r := true
	for _, input := range inputs {
		a, err := e.Exec(context, input)
		if err != nil {
			return nil, err
		}
		if a.StructType != pb.StructType_bool {
			return nil, errors.New(fmt.Sprintf("%v cannot be bool", a.StructType.String()))
		}
		if !a.Bool {
			r = a.Bool
			break
		}
	}
	output.Bool = r
	return []*pb.Struct{output}, nil
}

func OR(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	r := false
	for _, input := range inputs {
		a, err := e.Exec(context, input)
		if err != nil {
			return nil, err
		}
		if a.StructType != pb.StructType_bool {
			return nil, errors.New(fmt.Sprintf("%v cannot be bool", a.StructType.String()))
		}
		if a.Bool {
			r = a.Bool
			break
		}
	}
	output.Bool = r
	return []*pb.Struct{output}, nil
}

func Not(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	if len(inputs) != 1 {
		return nil, errors.New("not input len  != 1")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_bool {
		return nil, errors.New(fmt.Sprintf("%v not bool", a.StructType.String()))
	}
	output.Bool = !a.Bool
	return []*pb.Struct{output}, nil
}

func In(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) < 1 {
		return nil, errors.New("in input len  < 1")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	in := false
	ins := inputs[1:]
	for _, input := range ins {
		switch input.StructType {
		case pb.StructType_int64:
			if a.StructType != input.StructType {
				continue
			}
			if a.Int64 == input.Int64 {
				in = true
				break
			}
		case pb.StructType_double:
			if a.StructType != input.StructType {
				continue
			}
			if a.Double == input.Double {
				in = true
				break
			}
		case pb.StructType_bool:
			if a.StructType != input.StructType {
				continue
			}
			if a.Bool == input.Bool {
				in = true
				break
			}
		case pb.StructType_string:
			if a.StructType != input.StructType {
				continue
			}
			if a.String_ == input.String_ {
				in = true
				break
			}
		case pb.StructType_list:
			list := []*pb.Struct{
				a,
			}
			list = append(list, input.List...)
			oo, err := e.BaseFunctionMore2One(context, In, list)
			if err != nil {
				return nil, err
			}
			if oo.Bool {
				in = true
				break
			}
		default:
			return nil, errors.New(fmt.Sprintf("%v can't in", input.StructType.String()))
		}
	}
	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       false,
	}
	output.Bool = in
	return []*pb.Struct{output}, nil
}

// If (test-clause) (action<sub>1</sub>) (action<sub>2</sub>)
func If(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) < 2 {
		return nil, errors.New("if input len  < 2")
	}
	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_bool {
		return nil, errors.New(fmt.Sprintf("%v not bool", a.StructType.String()))
	}

	if a.Bool {
		b, err := e.Exec(context, inputs[1])
		if err != nil {
			return nil, err
		}
		return []*pb.Struct{b}, nil
	}
	if len(inputs) > 2 {
		b, err := e.Exec(context, inputs[2])
		if err != nil {
			return nil, err
		}
		return []*pb.Struct{b}, nil
	}
	return []*pb.Struct{&pb.Struct{
		StructType: pb.StructType_nil,
	}}, nil
}

func oddNumber(n int) bool {
	if n < 2 {
		return n%2 != 0
	}
	return oddNumber(n - 2)
}

// Case keyform key1 action1  key2 action2 ...
//(case day Sunday 1 Monday 2 Tuesday 3 Wednesday 4 Thursday 5 Friday 6 Saturday 7 Sunday
func Case(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) < 1 {
		return nil, errors.New("case input len  < 1")
	}
	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	switch a.StructType {
	case pb.StructType_string:
		break
	case pb.StructType_int64:
		break
	case pb.StructType_bool:
		break
	case pb.StructType_double:
		break
	default:
		return nil, errors.New(fmt.Sprintf("case keyform %v must be int64 or string or bool or double", a.StructType.String()))
	}

	keys := len(inputs) - 1

	for i := 1; i <= keys; i++ {
		kv := inputs[i]
		kvv, err := e.Exec(context, kv)
		if err != nil {
			return nil, err
		}
		if kvv.StructType != pb.StructType_list {
			return nil, errors.New(fmt.Sprintf("case %v must be {'list':[key action]}}", kvv.StructType.String()))
		}
		if len(kvv.GetList()) < 2 {
			return nil, errors.New(fmt.Sprintf("case kv len=%v must be {'list':[key action]}}", len(kvv.GetList())))
		}
		//每次for循环应该步进2
		k := kvv.List[0]
		v := kvv.List[1]
		ka, err := e.Exec(context, k)
		if err != nil {
			return nil, err
		}
		if ka.StructType != a.StructType {
			return nil, errors.New(fmt.Sprintf("keyform != key => %v!=%v", a.StructType.String(), ka.StructType.String()))
		}
		switch a.StructType {
		case pb.StructType_string:
			if ka.String_ != a.String_ {
				continue
			}
			break
		case pb.StructType_int64:
			if ka.Int64 != a.Int64 {
				continue
			}
			break
		case pb.StructType_bool:
			if ka.Bool != a.Bool {
				continue
			}
			break
		case pb.StructType_double:
			if ka.Double != a.Double {
				continue
			}
			break
		default:
			return nil, errors.New(fmt.Sprintf(" %v must be int64 or string or bool or double", a.StructType.String()))
		}
		va, err := e.Exec(context, v)
		if err != nil {
			return nil, err
		}
		return []*pb.Struct{va}, nil
	}
	return []*pb.Struct{&pb.Struct{
		StructType: pb.StructType_nil,
	}}, nil
}

// Chain action1 action2 action3 顺序执行表达式，如果其中一个表达式结果为Return类型则直接返回Return的结果不再继续执行后续表达式，如果所有表达式均无return则最终返回nil
func Chain(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	for _, input := range inputs {
		kvv, err := e.Exec(context, input)
		if err != nil {
			return nil, err
		}
		if kvv.StructType == pb.StructType_return {
			return kvv.Return, nil
		}
	}
	return []*pb.Struct{&pb.Struct{
		StructType: pb.StructType_nil,
	}}, nil
}

func GetHash(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) != 2 {
		return nil, errors.New("gethash input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	b, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_hash {
		return nil, errors.New(fmt.Sprintf("%v not hash", a.StructType.String()))
	}

	if b.StructType != pb.StructType_string {
		return nil, errors.New(fmt.Sprintf("%v not string", a.StructType.String()))
	}
	if a.Hash == nil {
		return nil, errors.New("hash is nil")
	}
	if v, ok := a.Hash[b.String_]; ok {
		return []*pb.Struct{v}, nil
	}
	return []*pb.Struct{&pb.Struct{
		StructType: pb.StructType_nil,
	}}, nil
}
