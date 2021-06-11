// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: sc.proto

package proto

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

type Response int32

const (
	Response_ResponseBeginIndex Response = 0
)

// Enum value maps for Response.
var (
	Response_name = map[int32]string{
		0: "ResponseBeginIndex",
	}
	Response_value = map[string]int32{
		"ResponseBeginIndex": 0,
	}
)

func (x Response) Enum() *Response {
	p := new(Response)
	*p = x
	return p
}

func (x Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response) Descriptor() protoreflect.EnumDescriptor {
	return file_sc_proto_enumTypes[0].Descriptor()
}

func (Response) Type() protoreflect.EnumType {
	return &file_sc_proto_enumTypes[0]
}

func (x Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response.Descriptor instead.
func (Response) EnumDescriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{0}
}

type ResultCode int32

const (
	ResultCode_Success   ResultCode = 0 //协议请求成功，其余失败
	ResultCode_Fail      ResultCode = 1
	ResultCode_SuccessOK ResultCode = 10000
)

// Enum value maps for ResultCode.
var (
	ResultCode_name = map[int32]string{
		0:     "Success",
		1:     "Fail",
		10000: "SuccessOK",
	}
	ResultCode_value = map[string]int32{
		"Success":   0,
		"Fail":      1,
		"SuccessOK": 10000,
	}
)

func (x ResultCode) Enum() *ResultCode {
	p := new(ResultCode)
	*p = x
	return p
}

func (x ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sc_proto_enumTypes[1].Descriptor()
}

func (ResultCode) Type() protoreflect.EnumType {
	return &file_sc_proto_enumTypes[1]
}

func (x ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultCode.Descriptor instead.
func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{1}
}

type ResponsePkg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdId        Response      `protobuf:"varint,1,opt,name=cmdId,proto3,enum=proto.Response" json:"cmdId,omitempty"` // 协议ID
	Code         ResultCode    `protobuf:"varint,2,opt,name=code,proto3,enum=proto.ResultCode" json:"code,omitempty"` //返回码
	ErrMsg       string        `protobuf:"bytes,10,opt,name=errMsg,proto3" json:"errMsg,omitempty"`                   //消息
	ServerTime   int64         `protobuf:"varint,11,opt,name=serverTime,proto3" json:"serverTime,omitempty"`          //服务器时间
	ArticleList  []*Article    `protobuf:"bytes,12,rep,name=articleList,proto3" json:"articleList,omitempty"`
	BlogHomeInfo *BlogHomeInfo `protobuf:"bytes,13,opt,name=blogHomeInfo,proto3" json:"blogHomeInfo,omitempty"`
}

func (x *ResponsePkg) Reset() {
	*x = ResponsePkg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePkg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePkg) ProtoMessage() {}

func (x *ResponsePkg) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePkg.ProtoReflect.Descriptor instead.
func (*ResponsePkg) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{0}
}

func (x *ResponsePkg) GetCmdId() Response {
	if x != nil {
		return x.CmdId
	}
	return Response_ResponseBeginIndex
}

func (x *ResponsePkg) GetCode() ResultCode {
	if x != nil {
		return x.Code
	}
	return ResultCode_Success
}

func (x *ResponsePkg) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *ResponsePkg) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *ResponsePkg) GetArticleList() []*Article {
	if x != nil {
		return x.ArticleList
	}
	return nil
}

func (x *ResponsePkg) GetBlogHomeInfo() *BlogHomeInfo {
	if x != nil {
		return x.BlogHomeInfo
	}
	return nil
}

var File_sc_proto protoreflect.FileDescriptor

var file_sc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x6b, 0x67, 0x12, 0x25, 0x0a,
	0x05, 0x63, 0x6d, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x63,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x22,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x10, 0x00, 0x2a, 0x33, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4f, 0x4b, 0x10, 0x90, 0x4e, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2e, 0x2f, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sc_proto_rawDescOnce sync.Once
	file_sc_proto_rawDescData = file_sc_proto_rawDesc
)

func file_sc_proto_rawDescGZIP() []byte {
	file_sc_proto_rawDescOnce.Do(func() {
		file_sc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sc_proto_rawDescData)
	})
	return file_sc_proto_rawDescData
}

var file_sc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sc_proto_goTypes = []interface{}{
	(Response)(0),        // 0: proto.Response
	(ResultCode)(0),      // 1: proto.ResultCode
	(*ResponsePkg)(nil),  // 2: proto.ResponsePkg
	(*Article)(nil),      // 3: proto.Article
	(*BlogHomeInfo)(nil), // 4: proto.BlogHomeInfo
}
var file_sc_proto_depIdxs = []int32{
	0, // 0: proto.ResponsePkg.cmdId:type_name -> proto.Response
	1, // 1: proto.ResponsePkg.code:type_name -> proto.ResultCode
	3, // 2: proto.ResponsePkg.articleList:type_name -> proto.Article
	4, // 3: proto.ResponsePkg.blogHomeInfo:type_name -> proto.BlogHomeInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sc_proto_init() }
func file_sc_proto_init() {
	if File_sc_proto != nil {
		return
	}
	file_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePkg); i {
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
			RawDescriptor: file_sc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sc_proto_goTypes,
		DependencyIndexes: file_sc_proto_depIdxs,
		EnumInfos:         file_sc_proto_enumTypes,
		MessageInfos:      file_sc_proto_msgTypes,
	}.Build()
	File_sc_proto = out.File
	file_sc_proto_rawDesc = nil
	file_sc_proto_goTypes = nil
	file_sc_proto_depIdxs = nil
}
