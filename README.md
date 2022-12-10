# lithengine
一款可以使用json编码的代码执行器，可以用来实现规则引擎

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

# 串行执行
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

更多示例请见 engine_test.go 

