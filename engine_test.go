package lithengine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rFuncMap map[string]Function
var rBlockMap map[string]*pb.Struct

func init() {
	rFuncMap = map[string]Function{
		"isPay": func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
			if len(inputs) < 1 {
				if len(inputs) != 1 {
					return nil, errors.New("not input len  != 1")
				}
			}
			a, err := e.Exec(context, inputs[0])
			if err != nil {
				return nil, err
			}
			if a.StructType != pb.StructType_string {
				return nil, errors.New(fmt.Sprintf("%v not string", inputs[0].StructType.String()))
			}
			userId := a.String_
			isPay := false
			if userId == "111" {
				isPay = true
			}
			return []*pb.Struct{
				&pb.Struct{
					StructType: pb.StructType_bool,
					Bool:       isPay,
				},
			}, nil
		},
		"userRiskLevel": func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
			return []*pb.Struct{
				&pb.Struct{
					StructType: pb.StructType_int64,
					Int64:      2,
				},
			}, nil
		},
	}
	rBlockMap = map[string]*pb.Struct{}
}

func TestAdd(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	outputAddToInt64, err := engine.Exec(context.Background(), &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "+",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      10,
			},
			&pb.Struct{
				StructType: pb.StructType_double,
				Double:     15,
			},
		},
	})
	assert.Empty(t, err)
	assert.Equal(t, outputAddToInt64.Double, 25.0)

	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"+":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 10.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"+":[5,5,2,3]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 15.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"+":[5,5,2,"3"]
		 }`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"+":["5",5,2,3]
		 }`,
	))
	assert.NotEmpty(t, err)
}

func TestReduce(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"-":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 0.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"-":[5,5,2,3]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, -5.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"-":[15,5,2,3]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 5.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"-":[15,5,2,"3"]
		 }`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"-":["15",5,2,3]
		 }`,
	))
	assert.NotEmpty(t, err)
}

func TestMultiply(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 25.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,-5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, -25.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,5,2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 50.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,5,2,2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 100.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,5,2,-2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, -100.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,5,-2,-2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 100.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":["5",5,-2,-2]
		 }`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"*":[5,"5"",-2,-2]
		 }`,
	))
	assert.NotEmpty(t, err)
}

func TestDivide(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 1.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,-5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, -1.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,5,2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 0.5)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,5,2,2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 0.25)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,5,2,-2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, -0.25)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,5,-2,-2]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 0.25)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":["5",5,-2,-2]
		 }`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"/":[5,"5"",-2,-2]
		 }`,
	))
	assert.NotEmpty(t, err)
}

func TestEq(t *testing.T) {
	input0 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "+",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      10,
			},
			&pb.Struct{
				StructType: pb.StructType_double,
				Double:     15,
			},
		},
	}
	input1 := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     25,
	}
	eq := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "=",
		FuncInput: []*pb.Struct{
			input0,
			input1,
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), eq)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
        "=": [
          {"+": [10,15,5]},
          30
        ]
	}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestFunction(t *testing.T) {
	isPay := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "isPay",
		Name:       "是否充值",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    "111",
			},
		},
	}
	input0 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "+",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      10,
			},
			&pb.Struct{
				StructType: pb.StructType_double,
				Double:     15,
			},
		},
	}
	input1 := &pb.Struct{
		StructType: pb.StructType_double,
		Double:     25,
	}
	isAge25 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "=",
		FuncInput: []*pb.Struct{
			input0,
			input1,
		},
	}

	and := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "&&",
		FuncInput: []*pb.Struct{
			isPay,
			isAge25,
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), and)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestAnd(t *testing.T) {
	and := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "&&",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       false,
			},
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), and)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	and = &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "&&",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
		},
	}

	output, err = engine.Exec(context.Background(), and)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestOR(t *testing.T) {
	or := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "||",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       false,
			},
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), or)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestNot(t *testing.T) {

	not := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "not",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), not)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	not = &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "not",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       false,
			},
		},
	}
	output, err = engine.Exec(context.Background(), not)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestNotEq(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{"!=":[3,4]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":[4,4]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":["4",4]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":["4","4"]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":[true,true]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":[true,false]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":["true",{"nil":true}]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"!=":[{"nil":true},{"nil":true}]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
}

func TestIf(t *testing.T) {
	isPay := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "isPay",
		Name:       "是否充值",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    "111",
			},
		},
	}
	If := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "if",
		FuncInput: []*pb.Struct{
			isPay,
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    "高价值用户",
			},
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    "低价值用户",
			},
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), If)
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "高价值用户")

	If = &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "if",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_function,
				FuncId:     "not",
				FuncInput: []*pb.Struct{
					isPay,
				},
			},
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    "高价值用户",
			},
		},
	}
	engine = NewEngine(rFuncMap, rBlockMap)
	output, err = engine.Exec(context.Background(), If)
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_nil)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set": ["a","a"]},
				{"set": ["a","b"]},
				{"if":[
					{
						"=":[
								{
									"chain": {"return": {"get": "a"}}
								},
								"b"
						]
					},
					{"set": ["a","setToc"]}
				]},
				{"return": {
						"chain": {"return": {"get": "a"}}
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "setToc")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set": ["a","a"]},
				{"if":[
					{
						"=":[
								{
									"chain": {"return": {"get": "a"}}
								},
								"c"
						]
					},
					{"set": ["a","setToc"]}
				]},
				{"return": {
						"chain": {"return": {"get": "a"}}
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")
}

func TestCase(t *testing.T) {
	Case := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "case",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_function,
				FuncId:     "userRiskLevel",
				Name:       "用户风险等级",
				FuncInput: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "10001",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      1,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "高风险用户",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      2,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "低风险用户",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      3,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "正常用户",
					},
				},
			},
		},
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), Case)
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "低风险用户")

	Case = &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "case",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      4,
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      1,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "高风险用户",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      2,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "低风险用户",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_list,
				List: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_int64,
						Int64:      3,
					},
					&pb.Struct{
						StructType: pb.StructType_string,
						String_:    "正常用户",
					},
				},
			},
		},
	}
	engine = NewEngine(rFuncMap, rBlockMap)
	output, err = engine.Exec(context.Background(), Case)
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_nil)
}

// TestBlock 代码块注册和使用
func TestBlock(t *testing.T) {
	//block
	engine := NewEngine(rFuncMap, rBlockMap)
	err := engine.RegisterBlockFromJson("a+b",
		`{
				"schema":{"inputType":[{"name":"a"},{"name":"b"}]},
				"+": [
					{"exec":{"get": "a"}},
					{"exec":{"get": "b"}}
				]
			}`)
	assert.Empty(t, err)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"a+b":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 10.0)

	err = engine.RegisterBlockFromJson("a=b",
		`{
				"schema":{"inputType":[{"name":"a"},{"name":"b"}]},
				"=": [
					{"exec":{"get": "a"}},
					{"exec":{"get": "b"}}
				]
			}`)
	assert.Empty(t, err)
	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"a=b":[5,5]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain":[
				{"set":["blockFunc",{"a+b":[5,3]}]},
				{"return":{
						"a=b":[{"get":"blockFunc"},8]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain":[
				{"set":["a",{"a+b":[5,3]}]},
				{"return":{
						"a=b":[{"get":"a"},8]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain":[
				{"return":{
						"a=b":[{"a+b":[5,3]},8]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain":[
				{"set":["a",{"a+b":[5,3]}]},
				{"set":["c",{"get":"a"}]},
				{"return":{
						"a=b":[{"closure":true,"get":"c"},8]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

// TestArgs 环境变量和变量传递示例
func TestArgs(t *testing.T) {
	//	//block
	err := RegisterBlockFromJson("PayAndAge25",
		`{
				"&&":[
					{
						"name": "是否充值",
						"isPay": [
							{
								"getArgs": [{"string": "uid"}]
							}
						]
					},
					{
						"=": [
							{
								"+": [{"int64": 10},{"double": 15}]
							},
							{
								"double": 25
							}
						]
					}
				]
			}`)
	assert.Empty(t, err)
	engine := NewEngine(rFuncMap, rBlockMap)
	ctx := MergeToContext(context.Background(), map[string]*pb.Struct{
		"uid": &pb.Struct{
			StructType: pb.StructType_string,
			String_:    "111",
		},
	})
	output, err := engine.ExecParse(ctx, []byte(
		`{"func":"PayAndAge25"}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(ctx, []byte(
		`{"getArgs":[{"string":"appid"}]}`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(ctx, []byte(
		`{"getArgs":[{"string":"appid"},{"string":"9999"}]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "9999")

	ctx = MergeToContext(context.Background(), map[string]*pb.Struct{
		"a": &pb.Struct{
			StructType: pb.StructType_double,
			Double:     10,
		},
	})

	output, err = engine.ExecParse(ctx, []byte(
		`{">":[{"getArgs":"a"},5]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestEngine_ExecParse(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"func": "=",
			"input": [
				{
					"func": "+",
					"input": [
						{"int64": 10},
						{"double": 15},
						{"double": 5}
					]
				},
				{"double": 30}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestIsType(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"func": "isType",
			"input": [
				{"nil": true},
				{"nil": true}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestIn(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				"a",
				"a",
				"b",
				"c"
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"string": "a"},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"string": "a"},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": ["a","d",3]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"int64": 3},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": [{"string": "a"},{"string": "d"},{"int64": 3}]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"bool": false},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": ["a",false,{"int64": 3}]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				3,
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": ["a",false,3]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"int64":3},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": ["a",false,{"int64":3}]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`
		{
			"chain":[
				{"set":["name","e"]},
				{
					"return":{
						"in": [
							{"int64":3},
							{"string": "d"},
							{"string": "b"},
							{"string": "c"},
							{"get": "name"},
							{"list": ["a",false,{"int64":3}]}
						]
					}
				}
			]
		}
		`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

// TestArgs 环境变量和变量传递示例
func TestGetHash(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"getHash": [
				{"hash": {"a":{"string": "good"},"b":"b"}},
				{"string": "a"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "good")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"getHash": [
				{"hash": {"a":{"string": "good"},"b":{"string": "b"}}},
				{"string": "c"}
			]
		}`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"getHash": [
				{"hash": {"a":{"string": "good"},"b":{"string": "b"}}},
				{"string": "c"},
				{"string": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "c")
}

// TestClosure 闭包函数测试
func TestClosure(t *testing.T) {
	//	//block
	err := RegisterBlockFromJson("PayAndAge25",
		`{
				"func": "&&",
				"closure":true,
				"input": [
					{
						"func": "isPay",
						"closure":true,
						"name": "是否充值",
						"input": [
							{
								"func": "getArgs",
								"closure":true,
								"input": [
									{
										"string": "uid"
									}
								]
							}
						]
					},
					{
						"func": "=",
						"closure":true,
						"input": [
							{
								"func": "+",
								"input": [
									{
										"int64": 10
									},
									{
										"double": 15
									}
								]
							},
							{
								"double": 25
							}
						]
					}
				]
			}`)
	assert.Empty(t, err)
	args := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "PayAndAge25",
	}
	engine := NewEngine(rFuncMap, rBlockMap)
	ctx := MergeToContext(context.Background(), map[string]*pb.Struct{
		"uid": &pb.Struct{
			StructType: pb.StructType_string,
			String_:    "111",
		},
	})
	output, err := engine.Exec(ctx, args)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

// TestChain 顺序执行器
func TestChain(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"func": "chain",
			"input": [
				{"string": "a"},
				{"string": "a"},
				{"string": "b"},
				{"string": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_nil)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"func": "chain",
			"input": [
				{
					"func":"if",
					"input":[
						true,
						{"return": ["a"]}
					]
				},
				{"string": "a"},
				{"string": "b"},
				{"string": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"string": "a"},
				{"return": "b"},
				{"string": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "b")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"func": "chain",
			"input": [
				{"string": "a"},
				{"return": {"string": "b"}},
				{"return": {"string": "c"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "b")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"if":[false,{"return":  "a"}]
				},
				{
					"if":[false,{"return": "b"}]
				},
				{"return": "c"}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "c")
}

// TestExec 函数执行器
func TestExec(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"exec": {"string": "a"}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"exec": {"+":[{"double": 6},{"double": 4}]}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 10.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set": ["execFunc",{"set":["a","aa"]}]},
				{"exec": {"get":"execFunc"}},
				{"return":{"get":"a"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "aa")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a",{"nil":true}]},
				{"set": ["execFunc",{"closure":true,"set":["a","aa"]}]},
				{"return":{"get":"a"}},
				{"exec": {"get":"execFunc"}},
				{"return":{"get":"a"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_nil)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"args":{"execFunc":{"closure":true,"set":["a","aa"]}},
			"chain": [
				{"exec": {"getArgs":"execFunc"}},
				{"return":{"get":"a"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "aa")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"args":{"execFunc":{"closure":true,"set":["a","aa"]}},
			"exec":
				{
					"chain":[
						{"exec": {"getArgs":"execFunc"}},
						{"return":{"get":"a"}}
					]
				}
		}`,
	))
	fmt.Println(output)
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "aa")
}

func TestPointer(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{"pointer": {"string":"a"}}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_pointer)
	assert.Equal(t, output.Pointer.String_, "a")
}

func TestFunc(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{"+": [{"int64":2},{"int64":2}]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.StructType, pb.StructType_double)
	assert.Equal(t, output.Double, 4.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"+": {"double":2}}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 2.0)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{"jkf": [{"int64":2},{"int64":2}]}`,
	))
	assert.NotEmpty(t, err)
}

func TestSet(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set": ["a","b"]},
				{"return": {"get": "a"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "b")
}

func TestGet(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{"return": {"get": "a"}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{"return": {"get": "b"}}
			]
		}`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{"return": {"get": ["b","a"]}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{"return": {
						"chain": [
							{"return": {"get": "a"}}
						]
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{"return": {
						"chain": [
							{"set":["a","b"]},
							{"return": {"get": "a"}}
						]
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "b")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","b"]},
				{"return": {
						"chain": [
							{"set": ["a","c"]},
							{"return": {"get": "a"}}
						]
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "c")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"set":["a","a"]},
				{
					"chain": [
						{"set":["a","b"]},
						{"set": ["a","c"]}
					]
				},
				{"return": {
						"chain": [
							{"return": {"get": {"string": "a"}}}
						]
					}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "c")

}

func TestSetBlock(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"setBlock":[
						"a+b",
						{
							"schema":{
								"inputType":[{"name":"a"},{"name":"b"}]
							},
							"+": [
								{"exec":{"get": "a"}},
								{"exec":{"get": "b"}}
							]
						}
					]
				},
				{
					"setBlock":[
						"a=b",
						{
							"schema":{
								"inputType":[{"name":"a"},{"name":"b"}]
							},
							"=": [
								{"exec":{"get": "a"}},
								{"exec":{"get": "b"}}
							]
						}
					]
				},
				{
					"return":{
						"a=b":[{"a+b":[5,3]},8]
						}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"setBlock":[
						"notEq",
						{
							"schema":{
								"inputType":[{"name":"a"},{"name":"b"}]
							},
							"not": {"eq":[{"get":"a"},{"get":"b"}]}
						}
					]
				},
				{
					"return":{
						"notEq":[5,3]
						}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestDefun(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"defun":[
						"notEq",
						[{"name":"a"},{"name":"b"}],
						{"not":{"eq":[{"get":"a"},{"get":"b"}]}}
					]
				},
				{
					"return":{"notEq":[3,4]}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"defun":[
						"notEq",
						[{"name":"a"},{"name":"b"}],
						{"not":{"eq":[{"get":"a"},{"get":"b"}]}}
					]
				},
				{
					"return":{"notEq":[4,4]}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
}

func TestToInt64(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toInt64":100.0
				},
				{
					"int64":100
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toInt64":"100.0"
				},
				{
					"int64":100
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toInt64":"100"
				},
				{
					"int64":100
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestToDouble(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toDouble":100.0
				},
				100
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toDouble":"100.0"
				},
				100
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"=": [
				{
					"toDouble":"100"
				},
				100
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestList(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"int64":3},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				{"list": ["a",false,{"int64":3}]}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"in": [
				{"int64":3},
				{"string": "d"},
				{"string": "b"},
				{"string": "c"},
				["a",false,{"int64":3}]
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestMetadata_GetInterface(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	engine.RegisterFunc("getMetadata", func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
		md, ok := FromContext(context)
		if !ok {
			return nil, errors.New("no args")
		}
		a, err := e.Exec(context, inputs[0])
		if err != nil {
			return nil, err
		}
		if a.StructType != pb.StructType_string {
			return nil, errors.New(fmt.Sprintf("%v not be string", a.StructType.String()))
		}
		str, ok := md.GetExtra(a.String_)
		if !ok {
			if len(inputs) >= 2 {
				return []*pb.Struct{
					inputs[1],
				}, nil
			}
			return nil, errors.New("no args interface")
		}
		return []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_string,
				String_:    str.(string),
			},
		}, nil
	})
	engine.RegisterFunc("removeMetadata", func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
		md, ok := FromContext(context)
		if !ok {
			return nil, errors.New("no args")
		}
		a, err := e.Exec(context, inputs[0])
		if err != nil {
			return nil, err
		}
		if a.StructType != pb.StructType_string {
			return nil, errors.New(fmt.Sprintf("%v not be string", a.StructType.String()))
		}
		md.RemoveExtra(a.String_)
		return []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_bool,
				Bool:       true,
			},
		}, nil
	})
	md := New(map[string]*pb.Struct{})
	md.SetExtra("interface", "this is interface string")
	output, err := engine.ExecParse(NewContext(context.Background(), md), []byte(
		`{"getMetadata": ["interface"]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "this is interface string")
	//------测试引擎变量和原生变量重名时能否区分存储-------------
	md = New(map[string]*pb.Struct{
		"interface": &pb.Struct{
			StructType: pb.StructType_string,
			String_:    "this is engine string",
		},
	})
	md.SetExtra("interface", "this is interface string")
	output, err = engine.ExecParse(NewContext(context.Background(), md), []byte(
		`{"getMetadata": ["interface"]}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "this is interface string")
	//------测试原生变量删除接口-------------
	md = New(map[string]*pb.Struct{
		"interface": &pb.Struct{
			StructType: pb.StructType_string,
			String_:    "this is engine string",
		},
	})
	md.SetExtra("interface", "this is interface string")
	output, err = engine.ExecParse(NewContext(context.Background(), md), []byte(
		`{
			"chain": [
				{"removeMetadata": ["interface"]},
				{"return": {"getMetadata": ["interface","no args interface"]}}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "no args interface")
}

func TestHookFunc(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	getArgs, ok := engine.LoadFunc("getArgs")
	assert.Equal(t, ok, true)
	engine.RegisterFunc("getArgs", func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
		r, err := getArgs(context, e, inputs)
		assert.Empty(t, err)
		assert.Equal(t, r[0].String_, "this is engine string")
		r[0].String_ = "this is hook getArgs"
		return r, err
	})
	md := New(map[string]*pb.Struct{
		"interface": &pb.Struct{
			StructType: pb.StructType_string,
			String_:    "this is engine string",
		},
	})
	output, err := engine.ExecParse(NewContext(context.Background(), md), []byte(
		`{"getArgs": "interface"}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "this is hook getArgs")
}

func TestStruct(t *testing.T) {
	output, err := ParseJson([]byte(`{"func":"=","id":"1"}`))
	assert.Empty(t, err)
	assert.Equal(t, output.Id, "1")
}

func TestEngine_RegisterBlockFromJson(t *testing.T) {
	engine := NewEngine(rFuncMap, rBlockMap)
	_ = engine.RegisterBlockFromJson("a+b",
		`{
          "schema":{"inputType":[
			{"name": "a"},
			{"name": "b"},
			{"name": "c", "optional": false}
			]},
		  "+": [
			{"get": "a"},
			{"get": "b"}
		  ]
		}`)
	_, err := engine.ExecParse(context.Background(), []byte(
		`{"a+b": [1]}`,
	))
	assert.NotEmpty(t, err)
	_, err = engine.ExecParse(context.Background(), []byte(
		`{"a+b": [1,2]}`,
	))
	assert.Empty(t, err)
	_, err = engine.ExecParse(context.Background(), []byte(
		`{"a+b": [1,2,3]}`,
	))
	assert.Empty(t, err)
	_ = engine.RegisterBlockFromJson("a+b",
		`{
          "schema":{"inputType":[
			{"name": "a"},
			{"name": "b", "optional": false},
			{"name": "c", "optional": false}
			]},
		  "+": [
			{"get": "a"}
		  ]
		}`)
	_, err = engine.ExecParse(context.Background(), []byte(
		`{"a+b": [1]}`,
	))
	assert.Empty(t, err)
}
