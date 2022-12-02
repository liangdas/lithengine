// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: lithengine.proto

package lithengine

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 结构体类型，函数，代码块也被认为是一个特殊的结构体
type StructType int32

const (
	StructType_Int64    StructType = 0
	StructType_String   StructType = 1
	StructType_double   StructType = 2
	StructType_bool     StructType = 3
	StructType_nil      StructType = 4
	StructType_function StructType = 5 //函数，会在【传参前先执行】得到结果
	StructType_closure  StructType = 6 //闭包函数，会作为子针参数传递【传参前不执行】
	StructType_block    StructType = 7 //代码块，本质是一系列函数逻辑组合 使用者可以定义一些通用的代码块方便后续复用
)

// Enum value maps for StructType.
var (
	StructType_name = map[int32]string{
		0: "Int64",
		1: "String",
		2: "double",
		3: "bool",
		4: "nil",
		5: "function",
		6: "closure",
		7: "block",
	}
	StructType_value = map[string]int32{
		"Int64":    0,
		"String":   1,
		"double":   2,
		"bool":     3,
		"nil":      4,
		"function": 5,
		"closure":  6,
		"block":    7,
	}
)

func (x StructType) Enum() *StructType {
	p := new(StructType)
	*p = x
	return p
}

func (x StructType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StructType) Descriptor() protoreflect.EnumDescriptor {
	return file_lithengine_proto_enumTypes[0].Descriptor()
}

func (StructType) Type() protoreflect.EnumType {
	return &file_lithengine_proto_enumTypes[0]
}

func (x StructType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StructType.Descriptor instead.
func (StructType) EnumDescriptor() ([]byte, []int) {
	return file_lithengine_proto_rawDescGZIP(), []int{0}
}

// 结构体定义
type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"type,omitempty"
	StructType StructType `protobuf:"varint,1,opt,name=structType,proto3,enum=lithengine.StructType" json:"type,omitempty"`
	Int64      int64      `protobuf:"varint,2,opt,name=int64,proto3" json:"int64,omitempty"`
	String_    string     `protobuf:"bytes,3,opt,name=string,proto3" json:"string,omitempty"`
	Double     float64    `protobuf:"fixed64,4,opt,name=double,proto3" json:"double,omitempty"`
	Bool       bool       `protobuf:"varint,5,opt,name=bool,proto3" json:"bool,omitempty"`
	Block      string     `protobuf:"bytes,7,opt,name=block,proto3" json:"block,omitempty"`
	//Function function=6;
	// @inject_tag: json:"fid,omitempty"
	FuncId string `protobuf:"bytes,10,opt,name=funcId,proto3" json:"fid,omitempty"` //函数ID
	// @inject_tag: json:"cid,omitempty"
	ClosureId string `protobuf:"bytes,11,opt,name=closureId,proto3" json:"cid,omitempty"` //closure函数ID
	// @inject_tag: json:"name,omitempty"
	FuncName string `protobuf:"bytes,12,opt,name=funcName,proto3" json:"name,omitempty"` //函数名称
	// @inject_tag: json:"schema,omitempty"
	FuncSchema string `protobuf:"bytes,13,opt,name=funcSchema,proto3" json:"schema,omitempty"` //函数定义
	// @inject_tag: json:"input,omitempty"
	FuncInput []*Struct          `protobuf:"bytes,14,rep,name=funcInput,proto3" json:"input,omitempty"`                                                                               //函数的输入
	Args      map[string]*Struct `protobuf:"bytes,15,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //环境变量
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lithengine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file_lithengine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file_lithengine_proto_rawDescGZIP(), []int{0}
}

func (x *Struct) GetStructType() StructType {
	if x != nil {
		return x.StructType
	}
	return StructType_Int64
}

func (x *Struct) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Struct) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Struct) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Struct) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Struct) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *Struct) GetFuncId() string {
	if x != nil {
		return x.FuncId
	}
	return ""
}

func (x *Struct) GetClosureId() string {
	if x != nil {
		return x.ClosureId
	}
	return ""
}

func (x *Struct) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *Struct) GetFuncSchema() string {
	if x != nil {
		return x.FuncSchema
	}
	return ""
}

func (x *Struct) GetFuncInput() []*Struct {
	if x != nil {
		return x.FuncInput
	}
	return nil
}

func (x *Struct) GetArgs() map[string]*Struct {
	if x != nil {
		return x.Args
	}
	return nil
}

// 函数定义Schema，可以定义函数的输入/输出数据类型，用于后续参数校验
type FunctionSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Input           []*Struct `protobuf:"bytes,3,rep,name=input,proto3" json:"input,omitempty"`
	Output          []*Struct `protobuf:"bytes,4,rep,name=output,proto3" json:"output,omitempty"`
	NumberOfInputs  int64     `protobuf:"varint,5,opt,name=number_of_inputs,json=numberOfInputs,proto3" json:"number_of_inputs,omitempty"`
	NumberOfOutputs int64     `protobuf:"varint,6,opt,name=number_of_outputs,json=numberOfOutputs,proto3" json:"number_of_outputs,omitempty"`
}

func (x *FunctionSchema) Reset() {
	*x = FunctionSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lithengine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionSchema) ProtoMessage() {}

func (x *FunctionSchema) ProtoReflect() protoreflect.Message {
	mi := &file_lithengine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionSchema.ProtoReflect.Descriptor instead.
func (*FunctionSchema) Descriptor() ([]byte, []int) {
	return file_lithengine_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionSchema) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FunctionSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FunctionSchema) GetInput() []*Struct {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *FunctionSchema) GetOutput() []*Struct {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *FunctionSchema) GetNumberOfInputs() int64 {
	if x != nil {
		return x.NumberOfInputs
	}
	return 0
}

func (x *FunctionSchema) GetNumberOfOutputs() int64 {
	if x != nil {
		return x.NumberOfOutputs
	}
	return 0
}

// 函数结构体定义
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`         //函数ID
	Name   string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`     //函数名称
	Schema string    `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"` //函数定义
	Input  []*Struct `protobuf:"bytes,4,rep,name=input,proto3" json:"input,omitempty"`   //函数的输入
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lithengine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_lithengine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_lithengine_proto_rawDescGZIP(), []int{2}
}

func (x *Function) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Function) GetInput() []*Struct {
	if x != nil {
		return x.Input
	}
	return nil
}

// 规则存储结构定义
type Engine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Block      map[string]*Struct `protobuf:"bytes,2,rep,name=block,proto3" json:"block,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //代码块，使用者可以定义一些通用的代码块方便后续复用
	RunExpr    *Struct            `protobuf:"bytes,3,opt,name=runExpr,proto3" json:"runExpr,omitempty"`                                                                                     //运算表达式
	ResultExpr *Struct            `protobuf:"bytes,4,opt,name=resultExpr,proto3" json:"resultExpr,omitempty"`                                                                               //通过运算表达式得出的结果，判断执行后续函数【可选】
}

func (x *Engine) Reset() {
	*x = Engine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lithengine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engine) ProtoMessage() {}

func (x *Engine) ProtoReflect() protoreflect.Message {
	mi := &file_lithengine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engine.ProtoReflect.Descriptor instead.
func (*Engine) Descriptor() ([]byte, []int) {
	return file_lithengine_proto_rawDescGZIP(), []int{3}
}

func (x *Engine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Engine) GetBlock() map[string]*Struct {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Engine) GetRunExpr() *Struct {
	if x != nil {
		return x.RunExpr
	}
	return nil
}

func (x *Engine) GetResultExpr() *Struct {
	if x != nil {
		return x.ResultExpr
	}
	return nil
}

var File_lithengine_proto protoreflect.FileDescriptor

var file_lithengine_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0xd3,
	0x03, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x4b, 0x0a, 0x09, 0x41, 0x72, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x74,
	0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x28, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x81, 0x02, 0x0a, 0x06, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x70, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x70, 0x72, 0x1a,
	0x4c, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x68, 0x0a,
	0x0a, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x69, 0x6c, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x05, 0x12,
	0x0b, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x07, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61, 0x6e, 0x67, 0x64, 0x61, 0x73, 0x2f, 0x6c,
	0x69, 0x74, 0x68, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6c, 0x69, 0x74, 0x68, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lithengine_proto_rawDescOnce sync.Once
	file_lithengine_proto_rawDescData = file_lithengine_proto_rawDesc
)

func file_lithengine_proto_rawDescGZIP() []byte {
	file_lithengine_proto_rawDescOnce.Do(func() {
		file_lithengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_lithengine_proto_rawDescData)
	})
	return file_lithengine_proto_rawDescData
}

var file_lithengine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lithengine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lithengine_proto_goTypes = []interface{}{
	(StructType)(0),        // 0: lithengine.StructType
	(*Struct)(nil),         // 1: lithengine.Struct
	(*FunctionSchema)(nil), // 2: lithengine.FunctionSchema
	(*Function)(nil),       // 3: lithengine.Function
	(*Engine)(nil),         // 4: lithengine.Engine
	nil,                    // 5: lithengine.Struct.ArgsEntry
	nil,                    // 6: lithengine.Engine.BlockEntry
}
var file_lithengine_proto_depIdxs = []int32{
	0,  // 0: lithengine.Struct.structType:type_name -> lithengine.StructType
	1,  // 1: lithengine.Struct.funcInput:type_name -> lithengine.Struct
	5,  // 2: lithengine.Struct.args:type_name -> lithengine.Struct.ArgsEntry
	1,  // 3: lithengine.FunctionSchema.input:type_name -> lithengine.Struct
	1,  // 4: lithengine.FunctionSchema.output:type_name -> lithengine.Struct
	1,  // 5: lithengine.Function.input:type_name -> lithengine.Struct
	6,  // 6: lithengine.Engine.block:type_name -> lithengine.Engine.BlockEntry
	1,  // 7: lithengine.Engine.runExpr:type_name -> lithengine.Struct
	1,  // 8: lithengine.Engine.resultExpr:type_name -> lithengine.Struct
	1,  // 9: lithengine.Struct.ArgsEntry.value:type_name -> lithengine.Struct
	1,  // 10: lithengine.Engine.BlockEntry.value:type_name -> lithengine.Struct
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_lithengine_proto_init() }
func file_lithengine_proto_init() {
	if File_lithengine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lithengine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Struct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lithengine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionSchema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lithengine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lithengine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Engine); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lithengine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lithengine_proto_goTypes,
		DependencyIndexes: file_lithengine_proto_depIdxs,
		EnumInfos:         file_lithengine_proto_enumTypes,
		MessageInfos:      file_lithengine_proto_msgTypes,
	}.Build()
	File_lithengine_proto = out.File
	file_lithengine_proto_rawDesc = nil
	file_lithengine_proto_goTypes = nil
	file_lithengine_proto_depIdxs = nil
}
