# lithengine
一款可以使用json编码的代码执行器，可以用来实现规则引擎

特点：
1. 简单
2. 灵活
3. 高效

goos: darwin
goarch: amd64
pkg: github.com/liangdas/lithengine
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkRiskLogic
BenchmarkRiskLogic-8   	 1768824	       611.0 ns/op
PASS

#支持数据类型
+ nil    
  + {"nil":true}
+ string 
  + {"string":"string"}
  + "string"
+ int64  
  + {"int64":666}
+ double
  + {"double":6.6}
  + 6.6
+ bool
  + {"bool":true}
  + true
+ list
  + {"list":["string",666]}
+ hash
  + {"hash":{"a":"string","b":666}}
+ 函数
  + {
    "in": [。。。]
    }
+ 延迟函数
  + 延迟函数可以作为参数传递（不会提前执行）
  + {
    "closure": true,
    "in": [。。。]
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
  + getArgs(获取传参),
  + in(包含检查),
  + getHash,
  + isType(类型判断)
  + chain （串行执行多个表达式）
  + exec   （执行表达式）
  + set （设置变量）
  + get （读取变量）
  + setBlock (添加代码块)
+ 支持添加自定义函数

# 使用示例
```
伪代码
10.0+15.0+5.0=30.0
```
```
engine := NewBaseEngine()
output, err := engine.ExecParse(context.Background(), []byte(
    `{
        "=": [
          {"+": [10,15,5]},
          30
        ]
	}`,
))
```
# 函数调用的语法
## 1. 无入参调用
```
语法1:

{"+":[]}

语法2:

{"func":"+",input:[]}
```

## 2. 单个入参
```
语法1:

{"getArgs":"clientId"}

语法2:

{"getArgs":["clientId"]}

语法3:

{"func":"getArgs",input:["clientId"]}
```

## 3. 多个入参
```
语法1:

{"+":[1,1]}

语法2:

{"func":"+",input:[1,1]}

```


# 串行执行
> chain 会按顺序执行传入的表达式
> 
>当遇到表达式返回{"return":[结果]}类型时终止后续表达式执行且使用return的结果作为chain的输出

```
伪代码
func chain(){
  if false {
    return "a"
  }
  if false {
    return "b"
  }
  return: "c"
}
```

```
engine := NewBaseEngine()
output, err := engine.ExecParse(context.Background(), []byte(
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

结果 "c"
```

# 变量操作
```
伪代码
func chain(){
  a=nil
  a="b"
  return a
}
```

```
engine := NewBaseEngine()
output, err := engine.ExecParse(context.Background(), []byte(
    `{
          "let":{"a":{"nil":true}}
          "chain": [
              {"set": ["a","b"]},
              {"return": {"get": "a"}}
          ]
      }`,
))

结果 "b"
```

# 执行闭包表达式
```
伪代码
func execFunc{
  a="aa"
}

func chain(){
  a=nil
  getArgs("execFunc")()
  return a
}

args["execFunc"]=execFunc
chain()
```

```
engine := NewBaseEngine()
output, err := engine.ExecParse(context.Background(), []byte(
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

结果 "aa"
```

# 代码块
> 代码块是可复用的表达式，可以在golang环境中注册也可以在脚本中注册
```go
engine := NewBaseEngine()
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
                      "let":{
                              "a":{
                                  "let":{"a":5,"b":3},
                                  "a+b":[]
                              },
                              "b":8
                          },
                      "a=b":[]
                    }
            }
        ]
    }`,
))
assert.Empty(t, err)
assert.Equal(t, output.Bool, true)
```

> 代码块传参可以使用环境变量(args)或变量(let)，无法使用input，如下
```
//注入add代码块，它会读取变量a和b进行+操作
{
  "setBlock":[
    "add",
    {
      "+": [
        {"exec":{"get": "a"}},
        {"exec":{"get": "b"}}
      ]
    }
  ]
}

exec函数的作用在于如果入参是闭包函数(closure)，则能够正常执行此闭包函数，并且用得到的结果做+操作


//调用a+b函数，传入参a=5 b=3
{"add":[],"let":{"a":5,"b":3}}

伪代码 
add(a,b) ==> 8
```

更多示例请见 engine_test.go 

