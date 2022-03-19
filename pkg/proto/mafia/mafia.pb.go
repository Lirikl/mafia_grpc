// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: mafia.proto

package mafia

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

type ConnectionStatus int32

const (
	ConnectionStatus_Connect    ConnectionStatus = 0
	ConnectionStatus_Disconnect ConnectionStatus = 1
	ConnectionStatus_None       ConnectionStatus = 2
	ConnectionStatus_Start      ConnectionStatus = 3
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "Connect",
		1: "Disconnect",
		2: "None",
		3: "Start",
	}
	ConnectionStatus_value = map[string]int32{
		"Connect":    0,
		"Disconnect": 1,
		"None":       2,
		"Start":      3,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_mafia_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_mafia_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[0]
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
	return file_mafia_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[1]
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
	return file_mafia_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Connect ConnectionStatus `protobuf:"varint,2,opt,name=connect,proto3,enum=mafia.ConnectionStatus" json:"connect,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectionRequest) GetConnect() ConnectionStatus {
	if x != nil {
		return x.Connect
	}
	return ConnectionStatus_Connect
}

type ConnectionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Connect   ConnectionStatus `protobuf:"varint,2,opt,name=connect,proto3,enum=mafia.ConnectionStatus" json:"connect,omitempty"`
	Users     []string         `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	SessionID int64            `protobuf:"varint,4,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Role      string           `protobuf:"bytes,5,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *ConnectionUpdate) Reset() {
	*x = ConnectionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionUpdate) ProtoMessage() {}

func (x *ConnectionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionUpdate.ProtoReflect.Descriptor instead.
func (*ConnectionUpdate) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectionUpdate) GetConnect() ConnectionStatus {
	if x != nil {
		return x.Connect
	}
	return ConnectionStatus_Connect
}

func (x *ConnectionUpdate) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ConnectionUpdate) GetSessionID() int64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *ConnectionUpdate) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GameCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Vote      string `protobuf:"bytes,2,opt,name=vote,proto3" json:"vote,omitempty"`
	SessionID int64  `protobuf:"varint,3,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
}

func (x *GameCommand) Reset() {
	*x = GameCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameCommand) ProtoMessage() {}

func (x *GameCommand) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameCommand.ProtoReflect.Descriptor instead.
func (*GameCommand) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{4}
}

func (x *GameCommand) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GameCommand) GetVote() string {
	if x != nil {
		return x.Vote
	}
	return ""
}

func (x *GameCommand) GetSessionID() int64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

type GameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Victim      string `protobuf:"bytes,2,opt,name=victim,proto3" json:"victim,omitempty"`
	Suspect     string `protobuf:"bytes,3,opt,name=suspect,proto3" json:"suspect,omitempty"`
	CheckResult bool   `protobuf:"varint,4,opt,name=CheckResult,proto3" json:"CheckResult,omitempty"`
	Winner      int64  `protobuf:"varint,5,opt,name=Winner,proto3" json:"Winner,omitempty"`
}

func (x *GameEvent) Reset() {
	*x = GameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEvent) ProtoMessage() {}

func (x *GameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEvent.ProtoReflect.Descriptor instead.
func (*GameEvent) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{5}
}

func (x *GameEvent) GetVictim() string {
	if x != nil {
		return x.Victim
	}
	return ""
}

func (x *GameEvent) GetSuspect() string {
	if x != nil {
		return x.Suspect
	}
	return ""
}

func (x *GameEvent) GetCheckResult() bool {
	if x != nil {
		return x.CheckResult
	}
	return false
}

func (x *GameEvent) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

var File_mafia_proto protoreflect.FileDescriptor

var file_mafia_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x61, 0x66, 0x69, 0x61, 0x22, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x5a, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x66, 0x69,
	0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22,
	0x53, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x22, 0x77, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2a, 0x44, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x10, 0x03, 0x32, 0xad, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x66, 0x69, 0x61, 0x12, 0x27, 0x0a,
	0x02, 0x44, 0x6f, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61,
	0x66, 0x69, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x6d, 0x61,
	0x66, 0x69, 0x61, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4c, 0x69, 0x72, 0x69, 0x6b, 0x6c, 0x2f, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mafia_proto_rawDescOnce sync.Once
	file_mafia_proto_rawDescData = file_mafia_proto_rawDesc
)

func file_mafia_proto_rawDescGZIP() []byte {
	file_mafia_proto_rawDescOnce.Do(func() {
		file_mafia_proto_rawDescData = protoimpl.X.CompressGZIP(file_mafia_proto_rawDescData)
	})
	return file_mafia_proto_rawDescData
}

var file_mafia_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mafia_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mafia_proto_goTypes = []interface{}{
	(ConnectionStatus)(0),     // 0: mafia.ConnectionStatus
	(*Request)(nil),           // 1: mafia.Request
	(*Response)(nil),          // 2: mafia.Response
	(*ConnectionRequest)(nil), // 3: mafia.ConnectionRequest
	(*ConnectionUpdate)(nil),  // 4: mafia.ConnectionUpdate
	(*GameCommand)(nil),       // 5: mafia.GameCommand
	(*GameEvent)(nil),         // 6: mafia.GameEvent
}
var file_mafia_proto_depIdxs = []int32{
	0, // 0: mafia.ConnectionRequest.connect:type_name -> mafia.ConnectionStatus
	0, // 1: mafia.ConnectionUpdate.connect:type_name -> mafia.ConnectionStatus
	1, // 2: mafia.Mafia.Do:input_type -> mafia.Request
	3, // 3: mafia.Mafia.Connect:input_type -> mafia.ConnectionRequest
	5, // 4: mafia.Mafia.GameSession:input_type -> mafia.GameCommand
	2, // 5: mafia.Mafia.Do:output_type -> mafia.Response
	4, // 6: mafia.Mafia.Connect:output_type -> mafia.ConnectionUpdate
	6, // 7: mafia.Mafia.GameSession:output_type -> mafia.GameEvent
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mafia_proto_init() }
func file_mafia_proto_init() {
	if File_mafia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mafia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_mafia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_mafia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_mafia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionUpdate); i {
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
		file_mafia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameCommand); i {
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
		file_mafia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEvent); i {
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
			RawDescriptor: file_mafia_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mafia_proto_goTypes,
		DependencyIndexes: file_mafia_proto_depIdxs,
		EnumInfos:         file_mafia_proto_enumTypes,
		MessageInfos:      file_mafia_proto_msgTypes,
	}.Build()
	File_mafia_proto = out.File
	file_mafia_proto_rawDesc = nil
	file_mafia_proto_goTypes = nil
	file_mafia_proto_depIdxs = nil
}
