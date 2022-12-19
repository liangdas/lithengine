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
			"let":{"a":"a"},
			"chain": [
				{"if":[
					{
						"=":[
								{
									"let":{"a":"b"},
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
			"let":{"a":"a"},
			"chain": [
				{"if":[
					{
						"=":[
								{
									"let":{"a":"b"},
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
				"let":{"a":5,"b":5},
				"+": [
					{"exec":{"get": "a"}},
					{"exec":{"get": "b"}}
				]
			}`)
	assert.Empty(t, err)
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"a":5,"b":3},
			"a+b":[]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 10.0)

	err = engine.RegisterBlockFromJson("a+b",
		`{
				"+": [
					{"exec":{"get": "a"}},
					{"exec":{"get": "b"}}
				]
			}`)
	assert.Empty(t, err)
	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"a":5,"b":3},
			"a+b":[]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Double, 8.0)

	err = engine.RegisterBlockFromJson("a=b",
		`{
				"=": [
					{"exec":{"get": "a"}},
					{"exec":{"get": "b"}}
				]
			}`)
	assert.Empty(t, err)
	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"a":5,"b":5},
			"a=b":[]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"blockFunc":0},
			"chain":[
				{"set":["blockFunc",{"let":{"a":5,"b":3},"a+b":[]}]},
				{"return":{
						"let":{"a":{"get":"blockFunc"},"b":8},
						"a=b":[]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"a":0},
			"chain":[
				{"set":["a",{"let":{"a":5,"b":3},"a+b":[]}]},
				{"return":{
						"let":{"a":{"get":"a"},"b":8},
						"a=b":[]
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
						"let":{"a":{"let":{"a":5,"b":3},"a+b":[]},"b":8},
						"a=b":[]
						}
				}
			]
		 }`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"let":{"a":0},
			"chain":[
				{"set":["a",{"let":{"a":5,"b":3},"a+b":[]}]},
				{"return":{
						"let":{"a":{"closure":true,"get":"c"},"b":8,"c":{"get":"a"}},
						"a=b":[]
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
			"let":{"execFunc":{"nil":true},"a":{"nil":true}},
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
			"let":{"execFunc":{"nil":true},"a":{"nil":true}},
			"chain": [
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
			"let":{"a":{"nil":true}},
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
					"let":{"a":{"nil":true}},
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
			],
			"let":{"a":{"nil":true}}
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
				{"return": {"get": {"string": "a"}}}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"return": {"get": {"string": "b"}}}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.NotEmpty(t, err)

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"return": {"get": [{"string": "b"},{"string": "a"}]}}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"return": {
						"chain": [
							{"return": {"get": {"string": "a"}}}
						]
					}
				}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"return": {
						"chain": [
							{"return": {"get": {"string": "a"}}}
						],
						"let":{"a":{"string":"b"}}
					}
				}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "b")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{"return": {
						"chain": [
							{"set": ["a","c"]},
							{"return": {"get": "a"}}
						],
						"let":{"a":"b"}
					}
				}
			],
			"let":{"a":"a"}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "c")

	output, err = engine.ExecParse(context.Background(), []byte(
		`{
			"chain": [
				{
					"chain": [
						{"set": [{"string": "a"},{"string": "c"}]}
					],
					"let":{"a":{"string":"b"}}
				},
				{"return": {
						"chain": [
							{"return": {"get": {"string": "a"}}}
						]
					}
				}
			],
			"let":{"a":{"string":"a"}}
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "a")

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
							"=": [
								{"exec":{"get": "a"}},
								{"exec":{"get": "b"}}
							]
						}
					]
				},
				{
					"return":{
						"let":{"a":{"let":{"a":5,"b":3},"a+b":[]},"b":8},
						"a=b":[]
						}
				}
			]
		}`,
	))
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
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
