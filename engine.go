package lithengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
	"strings"
)

type Function func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error)

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
		"args":     Args,
	}
	_blockMap = map[string]*pb.Struct{}
}

func RegisterFunc(id string, f Function) error {
	_funcMap[id] = f
	return nil
}

func RegisterBlock(id string, block *pb.Struct) error {
	_blockMap[id] = block
	return nil
}

func RegisterBlockFromJson(id string, js string) error {
	block, err := ParseJson([]byte(js))
	if err != nil {
		return err
	}
	_blockMap[id] = block
	return nil
}

func ParseJson(s []byte) (*pb.Struct, error) {
	out := &pb.Struct{}
	err := json.Unmarshal(s, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func ToJson(s *pb.Struct) string {
	str, _ := json.Marshal(s)
	return string(str)
}

type Engine struct {
	_funcMap  map[string]Function
	_blockMap map[string]*pb.Struct
}

func NewBaseEngine() *Engine {
	return &Engine{
		_funcMap:  map[string]Function{},
		_blockMap: map[string]*pb.Struct{},
	}
}

func NewEngine(funcMap map[string]Function, blockMap map[string]*pb.Struct) *Engine {
	return &Engine{
		_funcMap:  funcMap,
		_blockMap: blockMap,
	}
}

func (e *Engine) RegisterFunc(id string, f Function) *Engine {
	e._funcMap[id] = f
	return nil
}

func (e *Engine) RegisterBlock(id string, block *pb.Struct) *Engine {
	e._blockMap[id] = block
	return nil
}

func (e *Engine) RegisterBlockFromJson(id string, js string) error {
	block, err := ParseJson([]byte(js))
	if err != nil {
		return err
	}
	e._blockMap[id] = block
	return nil
}

func (e *Engine) LoadFunc(id string) (Function, bool) {
	if f, ok := e._funcMap[id]; ok {
		return f, ok
	}
	if f, ok := _funcMap[id]; ok {
		return f, ok
	}
	return nil, false
}

func (e *Engine) LoadBlock(id string) (*pb.Struct, bool) {
	if f, ok := e._blockMap[id]; ok {
		return f, ok
	}
	if f, ok := _blockMap[id]; ok {
		return f, ok
	}
	return nil, false
}

func (e *Engine) BaseFunctionOne2One(context context.Context, function Function, input *pb.Struct) (*pb.Struct, error) {
	if input.StructType == pb.StructType_function {
		o, err := e.FunctionOne(context, input)
		if err != nil {
			return nil, err
		}
		input = o
	}
	output, err := function(context, e, []*pb.Struct{input})
	if err != nil {
		return nil, err
	}
	if len(output) != 1 {
		return nil, errors.New(fmt.Sprintf("output(%v) not one", len(output)))
	}
	return output[0], nil
}

func (e *Engine) BaseFunctionMore2One(context context.Context, function Function, input []*pb.Struct) (*pb.Struct, error) {
	for i, in := range input {
		if in.StructType == pb.StructType_function {
			o, err := e.FunctionOne(context, in)
			if err != nil {
				return nil, err
			}
			input[i] = o
		}
	}
	output, err := function(context, e, input)
	if err != nil {
		return nil, err
	}
	if len(output) != 1 {
		return nil, errors.New(fmt.Sprintf("output(%v) not one", len(output)))
	}
	return output[0], nil
}

func (e *Engine) ExecParse(context context.Context, s []byte) (*pb.Struct, error) {
	st, err := ParseJson(s)
	if err != nil {
		return nil, err
	}
	switch st.StructType {
	case pb.StructType_function, pb.StructType_closure:
		o, err := e.FunctionOne(context, st)
		if err != nil {
			return nil, err
		}
		return o, nil
	case pb.StructType_block:
		return e.BlockOne(context, st)
	default:
		return st, nil
	}
}

func (e *Engine) Exec(context context.Context, st *pb.Struct) (*pb.Struct, error) {
	switch st.StructType {
	case pb.StructType_function, pb.StructType_closure:
		o, err := e.FunctionOne(context, st)
		if err != nil {
			return nil, err
		}
		return o, nil
	case pb.StructType_block:
		return e.BlockOne(context, st)
	default:
		return st, nil
	}
}

func (e *Engine) FunctionOne(context context.Context, function *pb.Struct) (*pb.Struct, error) {
	if function.StructType != pb.StructType_function && function.StructType != pb.StructType_closure {
		return nil, errors.New(fmt.Sprintf("%v Cannot execute", function.StructType.String()))
	}
	if f, ok := e.LoadFunc(function.GetFuncId()); !ok {
		return nil, errors.New(fmt.Sprintf("%v nofind", function.GetFuncId()))
	} else {
		if function.Args != nil {
			context = MergeToContext(context, function.Args)
		}
		output, err := e.BaseFunctionMore2One(context, f, function.GetFuncInput())
		if err != nil {
			return nil, err
		}
		return output, nil
	}
}

func (e *Engine) FunctionMore(context context.Context, function *pb.Struct) ([]*pb.Struct, error) {
	if function.StructType != pb.StructType_function && function.StructType != pb.StructType_closure {
		return nil, errors.New(fmt.Sprintf("%v Cannot execute", function.StructType.String()))
	}
	if f, ok := e.LoadFunc(function.GetFuncId()); !ok {
		return nil, errors.New(fmt.Sprintf("func nofind: %v", function.GetFuncId()))
	} else {
		if function.Args != nil {
			context = MergeToContext(context, function.Args)
		}
		output, err := f(context, e, function.GetFuncInput())
		if err != nil {
			return nil, err
		}
		return output, nil
	}
}

func (e *Engine) BlockOne(context context.Context, block *pb.Struct) (*pb.Struct, error) {
	if block.StructType != pb.StructType_block {
		return nil, errors.New(fmt.Sprintf("%v Cannot execute", block.StructType.String()))
	}
	if f, ok := e.LoadBlock(block.Block); !ok {
		return nil, errors.New(fmt.Sprintf("%v nofind", block.Block))
	} else {
		if block.Args != nil {
			context = MergeToContext(context, block.Args)
		}
		if f.Args != nil {
			context = MergeToContext(context, f.Args)
		}
		return e.Exec(context, f)
	}
}

func Args(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	name := inputs[0].String_
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

func Add(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     0,
	}
	for _, input := range inputs {
		switch input.StructType {
		case pb.StructType_Int64:
			output.Double += float64(input.Int64)
		case pb.StructType_double:
			output.Double += input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't add")
		case pb.StructType_String:
			return nil, errors.New("string can't add")
		case pb.StructType_function:
			o, err := e.FunctionOne(context, input)
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
		case pb.StructType_block:
			o, err := e.BlockOne(context, input)
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
		case pb.StructType_Int64:
			output.Double -= float64(input.Int64)
		case pb.StructType_double:
			output.Double -= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't add")
		case pb.StructType_String:
			return nil, errors.New("string can't add")
		case pb.StructType_function:
			o, err := e.FunctionOne(context, input)
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
		case pb.StructType_block:
			o, err := e.BlockOne(context, input)
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
		case pb.StructType_Int64:
			output.Double *= float64(input.Int64)
		case pb.StructType_double:
			output.Double *= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't multiply")
		case pb.StructType_String:
			return nil, errors.New("string can't multiply")
		case pb.StructType_function:
			o, err := e.FunctionOne(context, input)
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
		case pb.StructType_block:
			o, err := e.BlockOne(context, input)
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
		case pb.StructType_Int64:
			output.Double /= float64(input.Int64)
		case pb.StructType_double:
			output.Double /= input.Double
		case pb.StructType_bool:
			return nil, errors.New("bool can't divide")
		case pb.StructType_String:
			return nil, errors.New("string can't divide")
		case pb.StructType_function:
			o, err := e.FunctionOne(context, input)
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
		case pb.StructType_block:
			o, err := e.BlockOne(context, input)
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
		}
	}
	return []*pb.Struct{output}, nil
}

func Int64(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	output := &pb.Struct{
		StructType: pb.StructType_Int64,
		Int64:      0,
	}
	if len(inputs) != 2 {
		return nil, errors.New("DoubleToInt64 input len  != 2")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	if a.StructType != pb.StructType_double && a.StructType != pb.StructType_Int64 {
		return nil, errors.New(fmt.Sprintf("%v not be int64 or double", a.StructType.String()))
	}
	switch a.StructType {
	case pb.StructType_double:
		output.Int64 = int64(a.Double)
		break
	case pb.StructType_Int64:
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
	case pb.StructType_Int64:
		output.Bool = a.Int64 == b.Int64
	case pb.StructType_double:
		output.Bool = a.Double == b.Double
	case pb.StructType_bool:
		output.Bool = a.Bool == b.Bool
	case pb.StructType_String:
		output.Bool = strings.Compare(a.String_, b.String_) == 0
	case pb.StructType_function:
		return nil, errors.New("Function cannot be compared")
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
	case pb.StructType_Int64:
		output.Bool = a.Int64 > b.Int64
	case pb.StructType_double:
		output.Bool = a.Double > b.Double
	case pb.StructType_bool:
		return nil, errors.New("bool cannot be compared")
	case pb.StructType_String:
		output.Bool = strings.Compare(a.String_, b.String_) == 1
	case pb.StructType_function:
		return nil, errors.New("Function cannot be compared")
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
	case pb.StructType_Int64:
		output.Bool = a.Int64 >= b.Int64
	case pb.StructType_double:
		output.Bool = a.Double >= b.Double
	case pb.StructType_bool:
		return nil, errors.New("bool cannot be compared")
	case pb.StructType_String:
		output.Bool = strings.Compare(a.String_, b.String_) == 1 || strings.Compare(a.String_, b.String_) == 0
	case pb.StructType_function:
		return nil, errors.New("Function cannot be compared")
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
	case pb.StructType_Int64:
		output.Bool = a.Int64 < b.Int64
	case pb.StructType_double:
		output.Bool = a.Double < b.Double
	case pb.StructType_bool:
		return nil, errors.New("bool cannot be compared")
	case pb.StructType_String:
		output.Bool = strings.Compare(a.String_, b.String_) == -1
	case pb.StructType_function:
		return nil, errors.New("Function cannot be compared")
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
	case pb.StructType_Int64:
		output.Bool = a.Int64 <= b.Int64
	case pb.StructType_double:
		output.Bool = a.Double <= b.Double
	case pb.StructType_bool:
		return nil, errors.New("bool cannot be compared")
	case pb.StructType_String:
		output.Bool = strings.Compare(a.String_, b.String_) == -1 || strings.Compare(a.String_, b.String_) == 0
	case pb.StructType_function:
		return nil, errors.New("Function cannot be compared")
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
	return []*pb.Struct{output}, nil
}

// If (test-clause) (action<sub>1</sub>) (action<sub>2</sub>)
func If(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) != 3 {
		return nil, errors.New("if input len  != 3")
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
	b, err := e.Exec(context, inputs[2])
	if err != nil {
		return nil, err
	}
	return []*pb.Struct{b}, nil
}

func oddNumber(n int) bool {
	if n < 2 {
		return n%2 != 0
	}
	return oddNumber(n - 2)
}

// Case keyform default key1 action1  key2 action2 ...
//(case day Sunday 1 Monday 2 Tuesday 3 Wednesday 4 Thursday 5 Friday 6 Saturday 7 Sunday
func Case(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if oddNumber(len(inputs)) {
		return nil, errors.New("case input must be odd number")
	}
	if len(inputs) < 2 {
		return nil, errors.New("case input len  < 2")
	}
	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	switch a.StructType {
	case pb.StructType_String:
		break
	case pb.StructType_Int64:
		break
	case pb.StructType_bool:
		break
	case pb.StructType_double:
		break
	default:
		return nil, errors.New(fmt.Sprintf("case keyform %v must be int64 or string or bool or double", a.StructType.String()))
	}

	keys := (len(inputs) - 2) / 2

	for i := 1; i <= keys; i++ {
		//每次for循环应该步进2
		k := inputs[i+(i)]
		v := inputs[i+(i)+1]
		ka, err := e.Exec(context, k)
		if err != nil {
			return nil, err
		}
		if ka.StructType != a.StructType {
			return nil, errors.New(fmt.Sprintf("keyform != key => %v!=%v", a.StructType.String(), ka.StructType.String()))
		}
		switch a.StructType {
		case pb.StructType_String:
			if ka.String_ != a.String_ {
				continue
			}
			break
		case pb.StructType_Int64:
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
	df, err := e.Exec(context, inputs[1])
	if err != nil {
		return nil, err
	}
	return []*pb.Struct{df}, nil
}
