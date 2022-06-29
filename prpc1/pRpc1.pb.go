// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: pRpc1.proto

package prpc1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{0}
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type Test1Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string          `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64           `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
	RMap *Test1CellModel `protobuf:"bytes,3,opt,name=rMap,proto3" json:"rMap,omitempty"`
}

func (x *Test1Req) Reset() {
	*x = Test1Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test1Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test1Req) ProtoMessage() {}

func (x *Test1Req) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test1Req.ProtoReflect.Descriptor instead.
func (*Test1Req) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{4}
}

func (x *Test1Req) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test1Req) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

func (x *Test1Req) GetRMap() *Test1CellModel {
	if x != nil {
		return x.RMap
	}
	return nil
}

type Test1CellModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cp string `protobuf:"bytes,1,opt,name=cp,proto3" json:"cp,omitempty"`
}

func (x *Test1CellModel) Reset() {
	*x = Test1CellModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test1CellModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test1CellModel) ProtoMessage() {}

func (x *Test1CellModel) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test1CellModel.ProtoReflect.Descriptor instead.
func (*Test1CellModel) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{5}
}

func (x *Test1CellModel) GetCp() string {
	if x != nil {
		return x.Cp
	}
	return ""
}

type Test1Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string          `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64           `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
	RMap *Test1CellModel `protobuf:"bytes,3,opt,name=rMap,proto3" json:"rMap,omitempty"`
}

func (x *Test1Resp) Reset() {
	*x = Test1Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test1Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test1Resp) ProtoMessage() {}

func (x *Test1Resp) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test1Resp.ProtoReflect.Descriptor instead.
func (*Test1Resp) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{6}
}

func (x *Test1Resp) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test1Resp) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

func (x *Test1Resp) GetRMap() *Test1CellModel {
	if x != nil {
		return x.RMap
	}
	return nil
}

type Test2Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64  `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
}

func (x *Test2Req) Reset() {
	*x = Test2Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Req) ProtoMessage() {}

func (x *Test2Req) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Req.ProtoReflect.Descriptor instead.
func (*Test2Req) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{7}
}

func (x *Test2Req) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test2Req) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

type Test2Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64  `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
}

func (x *Test2Resp) Reset() {
	*x = Test2Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Resp) ProtoMessage() {}

func (x *Test2Resp) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Resp.ProtoReflect.Descriptor instead.
func (*Test2Resp) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{8}
}

func (x *Test2Resp) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test2Resp) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

type Test3Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64  `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
}

func (x *Test3Req) Reset() {
	*x = Test3Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test3Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test3Req) ProtoMessage() {}

func (x *Test3Req) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test3Req.ProtoReflect.Descriptor instead.
func (*Test3Req) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{9}
}

func (x *Test3Req) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test3Req) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

type Test3Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RStr string `protobuf:"bytes,1,opt,name=rStr,proto3" json:"rStr,omitempty"`
	RInt int64  `protobuf:"varint,2,opt,name=rInt,proto3" json:"rInt,omitempty"`
}

func (x *Test3Resp) Reset() {
	*x = Test3Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test3Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test3Resp) ProtoMessage() {}

func (x *Test3Resp) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test3Resp.ProtoReflect.Descriptor instead.
func (*Test3Resp) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{10}
}

func (x *Test3Resp) GetRStr() string {
	if x != nil {
		return x.RStr
	}
	return ""
}

func (x *Test3Resp) GetRInt() int64 {
	if x != nil {
		return x.RInt
	}
	return 0
}

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{11}
}

func (x *AuthReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AuthReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Pic      string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pRpc1_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_pRpc1_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_pRpc1_proto_rawDescGZIP(), []int{12}
}

func (x *AuthResp) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AuthResp) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AuthResp) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

var File_pRpc1_proto protoreflect.FileDescriptor

var file_pRpc1_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x52, 0x70, 0x63, 0x31, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x5d, 0x0a, 0x08,
	0x54, 0x65, 0x73, 0x74, 0x31, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x53, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x49, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x72, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x31, 0x43, 0x65, 0x6c, 0x6c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x72, 0x4d, 0x61, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x54,
	0x65, 0x73, 0x74, 0x31, 0x43, 0x65, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x70, 0x22, 0x5e, 0x0a,
	0x09, 0x54, 0x65, 0x73, 0x74, 0x31, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x53,
	0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x53, 0x74, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x49,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x31, 0x43, 0x65,
	0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x72, 0x4d, 0x61, 0x70, 0x22, 0x32, 0x0a,
	0x08, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x53, 0x74,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x49, 0x6e,
	0x74, 0x22, 0x33, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x53,
	0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x33, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x09, 0x54, 0x65,
	0x73, 0x74, 0x33, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x53, 0x74, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x49, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x49, 0x6e, 0x74, 0x22,
	0x3b, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4a, 0x0a, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x32, 0xba, 0x02, 0x0a, 0x05, 0x50, 0x52, 0x70,
	0x63, 0x31, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x70, 0x52, 0x70,
	0x63, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x52, 0x70,
	0x63, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x45,
	0x53, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x54, 0x65, 0x73,
	0x74, 0x31, 0x12, 0x0f, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x31,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x0f,
	0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x33, 0x12, 0x0f, 0x2e, 0x70, 0x52, 0x70,
	0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x33, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x52,
	0x70, 0x63, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x33, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x2e, 0x70, 0x52, 0x70,
	0x63, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x52, 0x70,
	0x63, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x52, 0x70, 0x63, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x70, 0x63, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pRpc1_proto_rawDescOnce sync.Once
	file_pRpc1_proto_rawDescData = file_pRpc1_proto_rawDesc
)

func file_pRpc1_proto_rawDescGZIP() []byte {
	file_pRpc1_proto_rawDescOnce.Do(func() {
		file_pRpc1_proto_rawDescData = protoimpl.X.CompressGZIP(file_pRpc1_proto_rawDescData)
	})
	return file_pRpc1_proto_rawDescData
}

var file_pRpc1_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pRpc1_proto_goTypes = []interface{}{
	(*EmptyReq)(nil),       // 0: pRpc1.EmptyReq
	(*EmptyResp)(nil),      // 1: pRpc1.EmptyResp
	(*Request)(nil),        // 2: pRpc1.Request
	(*Response)(nil),       // 3: pRpc1.Response
	(*Test1Req)(nil),       // 4: pRpc1.Test1Req
	(*Test1CellModel)(nil), // 5: pRpc1.Test1CellModel
	(*Test1Resp)(nil),      // 6: pRpc1.Test1Resp
	(*Test2Req)(nil),       // 7: pRpc1.Test2Req
	(*Test2Resp)(nil),      // 8: pRpc1.Test2Resp
	(*Test3Req)(nil),       // 9: pRpc1.Test3Req
	(*Test3Resp)(nil),      // 10: pRpc1.Test3Resp
	(*AuthReq)(nil),        // 11: pRpc1.AuthReq
	(*AuthResp)(nil),       // 12: pRpc1.AuthResp
}
var file_pRpc1_proto_depIdxs = []int32{
	5,  // 0: pRpc1.Test1Req.rMap:type_name -> pRpc1.Test1CellModel
	5,  // 1: pRpc1.Test1Resp.rMap:type_name -> pRpc1.Test1CellModel
	2,  // 2: pRpc1.PRpc1.Ping:input_type -> pRpc1.Request
	0,  // 3: pRpc1.PRpc1.ESGuide:input_type -> pRpc1.EmptyReq
	4,  // 4: pRpc1.PRpc1.Test1:input_type -> pRpc1.Test1Req
	7,  // 5: pRpc1.PRpc1.Test2:input_type -> pRpc1.Test2Req
	9,  // 6: pRpc1.PRpc1.Test3:input_type -> pRpc1.Test3Req
	11, // 7: pRpc1.PRpc1.CheckAuth:input_type -> pRpc1.AuthReq
	11, // 8: pRpc1.PRpc1.login:input_type -> pRpc1.AuthReq
	3,  // 9: pRpc1.PRpc1.Ping:output_type -> pRpc1.Response
	1,  // 10: pRpc1.PRpc1.ESGuide:output_type -> pRpc1.EmptyResp
	6,  // 11: pRpc1.PRpc1.Test1:output_type -> pRpc1.Test1Resp
	8,  // 12: pRpc1.PRpc1.Test2:output_type -> pRpc1.Test2Resp
	10, // 13: pRpc1.PRpc1.Test3:output_type -> pRpc1.Test3Resp
	12, // 14: pRpc1.PRpc1.CheckAuth:output_type -> pRpc1.AuthResp
	12, // 15: pRpc1.PRpc1.login:output_type -> pRpc1.AuthResp
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pRpc1_proto_init() }
func file_pRpc1_proto_init() {
	if File_pRpc1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pRpc1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_pRpc1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResp); i {
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
		file_pRpc1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pRpc1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pRpc1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test1Req); i {
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
		file_pRpc1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test1CellModel); i {
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
		file_pRpc1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test1Resp); i {
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
		file_pRpc1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Req); i {
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
		file_pRpc1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Resp); i {
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
		file_pRpc1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test3Req); i {
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
		file_pRpc1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test3Resp); i {
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
		file_pRpc1_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_pRpc1_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
			RawDescriptor: file_pRpc1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pRpc1_proto_goTypes,
		DependencyIndexes: file_pRpc1_proto_depIdxs,
		MessageInfos:      file_pRpc1_proto_msgTypes,
	}.Build()
	File_pRpc1_proto = out.File
	file_pRpc1_proto_rawDesc = nil
	file_pRpc1_proto_goTypes = nil
	file_pRpc1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PRpc1Client is the client API for PRpc1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PRpc1Client interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	//ESGuide elastic教程
	ESGuide(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
	//Test1 测试1
	Test1(ctx context.Context, in *Test1Req, opts ...grpc.CallOption) (*Test1Resp, error)
	//Test2 测试2
	Test2(ctx context.Context, in *Test2Req, opts ...grpc.CallOption) (*Test2Resp, error)
	// Test3 测试3
	Test3(ctx context.Context, in *Test3Req, opts ...grpc.CallOption) (*Test3Resp, error)
	// CheckAuth 风控检查
	CheckAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
	// CheckAuth 风控检查
	Login(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
}

type pRpc1Client struct {
	cc grpc.ClientConnInterface
}

func NewPRpc1Client(cc grpc.ClientConnInterface) PRpc1Client {
	return &pRpc1Client{cc}
}

func (c *pRpc1Client) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) ESGuide(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/ESGuide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) Test1(ctx context.Context, in *Test1Req, opts ...grpc.CallOption) (*Test1Resp, error) {
	out := new(Test1Resp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/Test1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) Test2(ctx context.Context, in *Test2Req, opts ...grpc.CallOption) (*Test2Resp, error) {
	out := new(Test2Resp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/Test2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) Test3(ctx context.Context, in *Test3Req, opts ...grpc.CallOption) (*Test3Resp, error) {
	out := new(Test3Resp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/Test3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) CheckAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRpc1Client) Login(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/pRpc1.PRpc1/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PRpc1Server is the server API for PRpc1 service.
type PRpc1Server interface {
	Ping(context.Context, *Request) (*Response, error)
	//ESGuide elastic教程
	ESGuide(context.Context, *EmptyReq) (*EmptyResp, error)
	//Test1 测试1
	Test1(context.Context, *Test1Req) (*Test1Resp, error)
	//Test2 测试2
	Test2(context.Context, *Test2Req) (*Test2Resp, error)
	// Test3 测试3
	Test3(context.Context, *Test3Req) (*Test3Resp, error)
	// CheckAuth 风控检查
	CheckAuth(context.Context, *AuthReq) (*AuthResp, error)
	// CheckAuth 风控检查
	Login(context.Context, *AuthReq) (*AuthResp, error)
}

// UnimplementedPRpc1Server can be embedded to have forward compatible implementations.
type UnimplementedPRpc1Server struct {
}

func (*UnimplementedPRpc1Server) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedPRpc1Server) ESGuide(context.Context, *EmptyReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ESGuide not implemented")
}
func (*UnimplementedPRpc1Server) Test1(context.Context, *Test1Req) (*Test1Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test1 not implemented")
}
func (*UnimplementedPRpc1Server) Test2(context.Context, *Test2Req) (*Test2Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test2 not implemented")
}
func (*UnimplementedPRpc1Server) Test3(context.Context, *Test3Req) (*Test3Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test3 not implemented")
}
func (*UnimplementedPRpc1Server) CheckAuth(context.Context, *AuthReq) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (*UnimplementedPRpc1Server) Login(context.Context, *AuthReq) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterPRpc1Server(s *grpc.Server, srv PRpc1Server) {
	s.RegisterService(&_PRpc1_serviceDesc, srv)
}

func _PRpc1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_ESGuide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).ESGuide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/ESGuide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).ESGuide(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_Test1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test1Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).Test1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/Test1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).Test1(ctx, req.(*Test1Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_Test2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).Test2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/Test2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).Test2(ctx, req.(*Test2Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_Test3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).Test3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/Test3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).Test3(ctx, req.(*Test3Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).CheckAuth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRpc1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRpc1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pRpc1.PRpc1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRpc1Server).Login(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PRpc1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pRpc1.PRpc1",
	HandlerType: (*PRpc1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PRpc1_Ping_Handler,
		},
		{
			MethodName: "ESGuide",
			Handler:    _PRpc1_ESGuide_Handler,
		},
		{
			MethodName: "Test1",
			Handler:    _PRpc1_Test1_Handler,
		},
		{
			MethodName: "Test2",
			Handler:    _PRpc1_Test2_Handler,
		},
		{
			MethodName: "Test3",
			Handler:    _PRpc1_Test3_Handler,
		},
		{
			MethodName: "CheckAuth",
			Handler:    _PRpc1_CheckAuth_Handler,
		},
		{
			MethodName: "login",
			Handler:    _PRpc1_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pRpc1.proto",
}
