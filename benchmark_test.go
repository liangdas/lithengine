package lithengine

import (
	"context"
	"errors"
	pb "github.com/liangdas/lithengine/golang"
	"testing"
)

var e *Engine
var p *pb.Struct

func init() {
	e = NewBaseEngine()
	e.RegisterFunc("IsIPBlack", IsIPBlack)
	p, _ = ParseJson([]byte(
		`{
            "||": [
                {
                    "&&": [
                        {
                            "IsIPBlack": "127.0.0.11"
                        },
                        {
                            "=": [
                                {
                                    "+": [
                                        10,
                                        15,
                                        5
                                    ]
                                },
                                30.0
                            ]
                        }
                    ]
                },
                {
                    "IsIPBlack": "127.0.0.1"
                }
            ]
        }`,
	))
}

func BenchmarkRiskLogic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = e.Exec(context.Background(), p)
	}
}

func IsIPBlack(context context.Context, e *Engine, inputs []*pb.Struct) ([]*pb.Struct, error) {
	if len(inputs) < 1 {
		return nil, errors.New("in input len  < 1")
	}

	a, err := e.Exec(context, inputs[0])
	if err != nil {
		return nil, err
	}

	in := false
	ins := []string{"127.0.0.1", "192.168.20.119"}
	for _, v := range ins {
		if v == a.GetString_() {
			in = true
			break
		}
	}

	output := &pb.Struct{
		StructType: pb.StructType_bool,
		Bool:       in,
	}
	return []*pb.Struct{output}, nil
}
