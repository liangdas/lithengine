package lithengine

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/liangdas/lithengine/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rfuncMap map[string]Function
var rblockMap map[string]*pb.Struct

func init() {
	rfuncMap = map[string]Function{
		"isPay": func(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
			userId := inputs[0].String_
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
					StructType: pb.StructType_Int64,
					Int64:      2,
				},
			}, nil
		},
	}
	rblockMap = map[string]*pb.Struct{}
}

func TestAdd(t *testing.T) {
	engine := NewEngine(rfuncMap, rblockMap)
	outputAddToInt64, err := engine.Exec(context.Background(), &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "+",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_Int64,
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
				StructType: pb.StructType_Int64,
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
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), eq)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
}

func TestFunction(t *testing.T) {
	isPay := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "isPay",
		FuncName:   "是否充值",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "111",
			},
		},
	}
	input0 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "+",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_Int64,
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
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), and)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
	b, _ := json.MarshalIndent(and, "", "    ")
	fmt.Println(string(b))
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
	engine := NewEngine(rfuncMap, rblockMap)
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
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), or)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
	fmt.Println("---or test ", output)
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
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), not)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
	fmt.Println("---not test ", output)
}

func TestIf(t *testing.T) {
	isPay := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "isPay",
		FuncName:   "是否充值",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_String,
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
				StructType: pb.StructType_String,
				String_:    "高价值用户",
			},
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "低价值用户",
			},
		},
	}
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), If)
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "高价值用户")
	fmt.Println("---if test ", output)
	b, _ := json.MarshalIndent(If, "", "    ")
	fmt.Println("---if json = ", string(b))
}

func TestCase(t *testing.T) {
	Case := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "case",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_function,
				FuncId:     "userRiskLevel",
				FuncName:   "用户风险等级",
				FuncInput: []*pb.Struct{
					&pb.Struct{
						StructType: pb.StructType_String,
						String_:    "10001",
					},
				},
			},
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "正常用户",
			},
			&pb.Struct{
				StructType: pb.StructType_Int64,
				Int64:      1,
			},
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "高风险用户",
			},
			&pb.Struct{
				StructType: pb.StructType_Int64,
				Int64:      2,
			},
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "低风险用户",
			},
			&pb.Struct{
				StructType: pb.StructType_Int64,
				Int64:      3,
			},
			&pb.Struct{
				StructType: pb.StructType_String,
				String_:    "正常用户",
			},
		},
	}
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), Case)
	assert.Empty(t, err)
	assert.Equal(t, output.String_, "低风险用户")
	fmt.Println("---case test ", output)
	b, _ := json.MarshalIndent(Case, "", "    ")
	fmt.Println("---case json = ", string(b))
}

// TestBlock 代码块注册和使用
func TestBlock(t *testing.T) {
	//	//block
	err := RegisterBlockFromJson("PayAndAge25",
		`{
				"fid": "&&",
				"input": [
					{
						"fid": "isPay",
						"name": "是否充值",
						"input": [
							{
								"string": "111"
							}
						]
					},
					{
						"fid": "=",
						"input": [
							{
								"fid": "+",
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
	engine := NewEngine(rfuncMap, rblockMap)
	output, err := engine.Exec(context.Background(), &pb.Struct{
		StructType: pb.StructType_block,
		Block:      "PayAndAge25",
	})
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)
	fmt.Println("---block test ", output)
}

// TestArgs 环境变量和变量传递示例
func TestArgs(t *testing.T) {
	//	//block
	err := RegisterBlockFromJson("PayAndAge25",
		`{
				"fid": "&&",
				"input": [
					{
						"fid": "isPay",
						"name": "是否充值",
						"input": [
							{
								"fid": "args",
								"input": [
									{
										"string": "uid"
									}
								]
							}
						]
					},
					{
						"fid": "=",
						"input": [
							{
								"fid": "+",
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
		StructType: pb.StructType_block,
		Block:      "PayAndAge25",
	}
	engine := NewEngine(rfuncMap, rblockMap)
	ctx := MergeToContext(context.Background(), map[string]*pb.Struct{
		"uid": &pb.Struct{
			StructType: pb.StructType_String,
			String_:    "111",
		},
	})
	output, err := engine.Exec(ctx, args)
	assert.Empty(t, err)
	fmt.Println("---args test ", output)
	b, _ := json.MarshalIndent(args, "", "    ")
	fmt.Println(string(b))
}

func TestEngine_ExecParse(t *testing.T) {
	engine := NewBaseEngine()
	output, err := engine.ExecParse(context.Background(), []byte(
		`{
			"fid": "=",
			"input": [
				{
					"fid": "+",
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
