package lithengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
)

type Function func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error)

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
	//if input.StructType == pb.StructType_function && !input.Closure {
	//	o, err := e.FunctionOne(context, input)
	//	if err != nil {
	//		return nil, err
	//	}
	//	input = o
	//}
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
	//for i, in := range input {
	//	if in.StructType == pb.StructType_function && !in.Closure {
	//		o, err := e.FunctionOne(context, in)
	//		if err != nil {
	//			return nil, err
	//		}
	//		input[i] = o
	//	}
	//}
	output, err := function(context, e, input)
	if err != nil {
		return nil, err
	}
	if len(output) != 1 {
		return nil, errors.New(fmt.Sprintf("output(%v) not one", len(output)))
	}
	return output[0], nil
}

func (e *Engine) BaseFunctionMore(context context.Context, function Function, input []*pb.Struct) ([]*pb.Struct, error) {
	//for i, in := range input {
	//	if in.StructType == pb.StructType_function && !in.Closure {
	//		o, err := e.FunctionOne(context, in)
	//		if err != nil {
	//			return nil, err
	//		}
	//		input[i] = o
	//	}
	//}
	return function(context, e, input)
}

func (e *Engine) ExecParse(context context.Context, s []byte) (*pb.Struct, error) {
	st, err := ParseJson(s)
	if err != nil {
		return nil, err
	}
	switch st.StructType {
	case pb.StructType_function:
		if _, ok := e.LoadFunc(st.Func()); ok {
			o, err := e.FunctionOne(context, st)
			if err != nil {
				return nil, err
			}
			return o, nil
		} else if _, ok := e.LoadBlock(st.Func()); ok {
			return e.BlockOne(context, st)
		} else {
			return nil, errors.New(fmt.Sprintf("%v is not function or block ", st.Func()))
		}
	default:
		return st, nil
	}
}

func (e *Engine) Exec(context context.Context, st *pb.Struct) (*pb.Struct, error) {
	switch st.StructType {
	case pb.StructType_function:
		if _, ok := e.LoadFunc(st.Func()); ok {
			o, err := e.FunctionOne(context, st)
			if err != nil {
				return nil, err
			}
			return o, nil
		} else if _, ok := e.LoadBlock(st.Func()); ok {
			return e.BlockOne(context, st)
		} else {
			return nil, errors.New(fmt.Sprintf("%v is not function or block ", st.Func()))
		}
	default:
		return st, nil
	}
}

func (e *Engine) FunctionOne(context context.Context, function *pb.Struct) (*pb.Struct, error) {
	if function.StructType != pb.StructType_function {
		return nil, errors.New(fmt.Sprintf("%v Cannot execute", function.StructType.String()))
	}
	if f, ok := e.LoadFunc(function.Func()); !ok {
		return nil, errors.New(fmt.Sprintf("%v nofind", function.Func()))
	} else {
		//覆盖环境变量
		if function.Args != nil {
			context = MergeToContext(context, function.Args)
		}
		//初始化局部变量
		if function.Let != nil {
			let := map[string]*pb.Struct{}
			for k, v := range function.Let {
				varName := fmt.Sprintf("__%v__", k)
				let[varName] = &pb.Struct{
					StructType: pb.StructType_pointer,
					Pointer:    v,
				}
			}
			context = MergeToContext(context, let)
		}
		output, err := e.BaseFunctionMore2One(context, f, function.GetFuncInput())
		if err != nil {
			return nil, err
		}
		return output, nil
	}
}

func (e *Engine) FunctionMore(context context.Context, function *pb.Struct) ([]*pb.Struct, error) {
	if function.StructType != pb.StructType_function {
		return nil, errors.New(fmt.Sprintf("%v Cannot execute", function.StructType.String()))
	}
	if f, ok := e.LoadFunc(function.Func()); !ok {
		return nil, errors.New(fmt.Sprintf("func nofind: %v", function.Func()))
	} else {
		//覆盖环境变量
		if function.Args != nil {
			context = MergeToContext(context, function.Args)
		}
		//初始化局部变量
		if function.Let != nil {
			let := map[string]*pb.Struct{}
			for k, v := range function.Let {
				varName := fmt.Sprintf("__%v__", k)
				let[varName] = &pb.Struct{
					StructType: pb.StructType_pointer,
					Pointer:    v,
				}
			}
			context = MergeToContext(context, let)
		}
		return e.BaseFunctionMore(context, f, function.GetFuncInput())
	}
}

func (e *Engine) BlockOne(context context.Context, block *pb.Struct) (*pb.Struct, error) {
	if f, ok := e.LoadBlock(block.Func()); !ok {
		return nil, errors.New(fmt.Sprintf("%v nofind", block.Func()))
	} else {
		//覆盖环境变量
		if block.Args != nil {
			context = MergeToContext(context, block.Args)
		}
		if f.Args != nil {
			context = MergeToContext(context, f.Args)
		}
		//初始化局部变量
		if f.Let != nil {
			let := map[string]*pb.Struct{}
			for k, v := range f.Let {
				varName := fmt.Sprintf("__%v__", k)
				let[varName] = &pb.Struct{
					StructType: pb.StructType_pointer,
					Pointer:    v,
				}
			}
			context = MergeToContext(context, let)
		}
		return e.Exec(context, f)
	}
}
