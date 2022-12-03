# lithengine
一款可以使用json编码的代码执行器，可以用来实现规则引擎

# 使用示例
```
伪代码
10.0+15.0+5.0=30.0
```
```
engine := NewBaseEngine()
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
```
更多示例请见 engine_test.go 
#支持数据类型
+ nil    
  + {"nil":true}
+ string 
  + {"string":"string"}
+ int64  
  + {"int64":666}
+ double
  + {"double":6.6}
+ bool
  + {"bool":true}
+ list
  + {"list":[{"string":"string"},{"int64":666}]}
+ hash
  + {"hash":{"a":{"string":"string"},"b":{"int64":666}}}
+ 函数
  + {
    "func": "in",
    "input": [。。。]
    }
+ 闭包函数
  + {
    "closure": "in",
    "input": [。。。]
    }
+ 代码块
# 支持特性
+ 内置支持
  + 加，
  + 减，
  + 乘，
  + 除，
  + =，
  + \>,<,>=,<=,&&,||,
  + not,
  + if,
  + case,
  + int64(浮点数转int64),
  + args(获取传参),
  + in(包含检查),
  + getHash,
  + isType(类型判断)
+ 支持添加自定义函数
