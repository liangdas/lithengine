package lithengine

import (
	"context"
	pb "github.com/liangdas/lithengine/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGt(t *testing.T) {
	input0 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     ">",
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
	engine := NewEngine(rFuncMap, rBlockMap)
	output, err := engine.Exec(context.Background(), input0)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
	input1 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "<",
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
	output, err = engine.Exec(context.Background(), input1)
	assert.Empty(t, err)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

	input2 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     ">=",
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
	output, err = engine.Exec(context.Background(), input2)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
	input3 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "<=",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_double,
				Double:     15,
			},
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      10,
			},
		},
	}
	output, err = engine.Exec(context.Background(), input3)
	assert.Empty(t, err)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)

	input5 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "=",
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
	output, err = engine.Exec(context.Background(), input5)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, false)
	input4 := &pb.Struct{
		StructType: pb.StructType_function,
		FuncId:     "<=",
		FuncInput: []*pb.Struct{
			&pb.Struct{
				StructType: pb.StructType_double,
				Double:     10,
			},
			&pb.Struct{
				StructType: pb.StructType_int64,
				Int64:      10,
			},
		},
	}
	output, err = engine.Exec(context.Background(), input4)
	assert.Empty(t, err)
	assert.Empty(t, err)
	assert.Equal(t, output.Bool, true)

}
