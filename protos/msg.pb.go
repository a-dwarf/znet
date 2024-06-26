// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.26.1
// source: protos/msg.proto

package __

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

type ZIdentity int32

const (
	ZIdentity_U_TYPE_CLI ZIdentity = 0 // client
	ZIdentity_U_TYPE_SER ZIdentity = 1 // server
)

// Enum value maps for ZIdentity.
var (
	ZIdentity_name = map[int32]string{
		0: "U_TYPE_CLI",
		1: "U_TYPE_SER",
	}
	ZIdentity_value = map[string]int32{
		"U_TYPE_CLI": 0,
		"U_TYPE_SER": 1,
	}
)

func (x ZIdentity) Enum() *ZIdentity {
	p := new(ZIdentity)
	*p = x
	return p
}

func (x ZIdentity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZIdentity) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_msg_proto_enumTypes[0].Descriptor()
}

func (ZIdentity) Type() protoreflect.EnumType {
	return &file_protos_msg_proto_enumTypes[0]
}

func (x ZIdentity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZIdentity.Descriptor instead.
func (ZIdentity) EnumDescriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{0}
}

type ZAction int32

const (
	ZAction_Z_TYPE_READ  ZAction = 0 // read
	ZAction_Z_TYPE_WRITE ZAction = 1 // write
)

// Enum value maps for ZAction.
var (
	ZAction_name = map[int32]string{
		0: "Z_TYPE_READ",
		1: "Z_TYPE_WRITE",
	}
	ZAction_value = map[string]int32{
		"Z_TYPE_READ":  0,
		"Z_TYPE_WRITE": 1,
	}
)

func (x ZAction) Enum() *ZAction {
	p := new(ZAction)
	*p = x
	return p
}

func (x ZAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZAction) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_msg_proto_enumTypes[1].Descriptor()
}

func (ZAction) Type() protoreflect.EnumType {
	return &file_protos_msg_proto_enumTypes[1]
}

func (x ZAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZAction.Descriptor instead.
func (ZAction) EnumDescriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{1}
}

type ZPushType int32

const (
	ZPushType_Z_TYPE_DM ZPushType = 0 // direct msg
	ZPushType_Z_TYPE_BC ZPushType = 1 // broadcast
)

// Enum value maps for ZPushType.
var (
	ZPushType_name = map[int32]string{
		0: "Z_TYPE_DM",
		1: "Z_TYPE_BC",
	}
	ZPushType_value = map[string]int32{
		"Z_TYPE_DM": 0,
		"Z_TYPE_BC": 1,
	}
)

func (x ZPushType) Enum() *ZPushType {
	p := new(ZPushType)
	*p = x
	return p
}

func (x ZPushType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZPushType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_msg_proto_enumTypes[2].Descriptor()
}

func (ZPushType) Type() protoreflect.EnumType {
	return &file_protos_msg_proto_enumTypes[2]
}

func (x ZPushType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZPushType.Descriptor instead.
func (ZPushType) EnumDescriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{2}
}

type ZType int32

const (
	ZType_Z_TYPE_RNG     ZType = 0
	ZType_Z_TYPE_EVENT   ZType = 1
	ZType_Z_TYPE_CLOCK   ZType = 2
	ZType_Z_TYPE_GATEWAY ZType = 3
	ZType_Z_TYPE_ZCHAT   ZType = 4
)

// Enum value maps for ZType.
var (
	ZType_name = map[int32]string{
		0: "Z_TYPE_RNG",
		1: "Z_TYPE_EVENT",
		2: "Z_TYPE_CLOCK",
		3: "Z_TYPE_GATEWAY",
		4: "Z_TYPE_ZCHAT",
	}
	ZType_value = map[string]int32{
		"Z_TYPE_RNG":     0,
		"Z_TYPE_EVENT":   1,
		"Z_TYPE_CLOCK":   2,
		"Z_TYPE_GATEWAY": 3,
		"Z_TYPE_ZCHAT":   4,
	}
)

func (x ZType) Enum() *ZType {
	p := new(ZType)
	*p = x
	return p
}

func (x ZType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_msg_proto_enumTypes[3].Descriptor()
}

func (ZType) Type() protoreflect.EnumType {
	return &file_protos_msg_proto_enumTypes[3]
}

func (x ZType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZType.Descriptor instead.
func (ZType) EnumDescriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{3}
}

type Zp2P struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type      ZIdentity `protobuf:"varint,2,opt,name=type,proto3,enum=protos.ZIdentity" json:"type,omitempty"`                         // for p2p
	Action    ZAction   `protobuf:"varint,3,opt,name=action,proto3,enum=protos.ZAction" json:"action,omitempty"`                       // for p2p
	PushType  ZPushType `protobuf:"varint,4,opt,name=push_type,json=pushType,proto3,enum=protos.ZPushType" json:"push_type,omitempty"` // for vlc
	Message   *ZMessage `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	PublicKey []byte    `protobuf:"bytes,6,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature []byte    `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"` // for verifying
}

func (x *Zp2P) Reset() {
	*x = Zp2P{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zp2P) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zp2P) ProtoMessage() {}

func (x *Zp2P) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zp2P.ProtoReflect.Descriptor instead.
func (*Zp2P) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Zp2P) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Zp2P) GetType() ZIdentity {
	if x != nil {
		return x.Type
	}
	return ZIdentity_U_TYPE_CLI
}

func (x *Zp2P) GetAction() ZAction {
	if x != nil {
		return x.Action
	}
	return ZAction_Z_TYPE_READ
}

func (x *Zp2P) GetPushType() ZPushType {
	if x != nil {
		return x.PushType
	}
	return ZPushType_Z_TYPE_DM
}

func (x *Zp2P) GetMessage() *ZMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Zp2P) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Zp2P) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// ZMessage
type ZMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []byte    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version   uint32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Type      ZType     `protobuf:"varint,3,opt,name=type,proto3,enum=protos.ZType" json:"type,omitempty"`
	Action    ZAction   `protobuf:"varint,4,opt,name=action,proto3,enum=protos.ZAction" json:"action,omitempty"`
	Identity  ZIdentity `protobuf:"varint,5,opt,name=identity,proto3,enum=protos.ZIdentity" json:"identity,omitempty"`
	PublicKey []byte    `protobuf:"bytes,6,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Data      []byte    `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte    `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	From      []byte    `protobuf:"bytes,9,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte    `protobuf:"bytes,10,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ZMessage) Reset() {
	*x = ZMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZMessage) ProtoMessage() {}

func (x *ZMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZMessage.ProtoReflect.Descriptor instead.
func (*ZMessage) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{1}
}

func (x *ZMessage) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ZMessage) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ZMessage) GetType() ZType {
	if x != nil {
		return x.Type
	}
	return ZType_Z_TYPE_RNG
}

func (x *ZMessage) GetAction() ZAction {
	if x != nil {
		return x.Action
	}
	return ZAction_Z_TYPE_READ
}

func (x *ZMessage) GetIdentity() ZIdentity {
	if x != nil {
		return x.Identity
	}
	return ZIdentity_U_TYPE_CLI
}

func (x *ZMessage) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *ZMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ZMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ZMessage) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ZMessage) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

type ZChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageData string     `protobuf:"bytes,1,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	Clock       *ClockInfo `protobuf:"bytes,2,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (x *ZChat) Reset() {
	*x = ZChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZChat) ProtoMessage() {}

func (x *ZChat) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZChat.ProtoReflect.Descriptor instead.
func (*ZChat) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{2}
}

func (x *ZChat) GetMessageData() string {
	if x != nil {
		return x.MessageData
	}
	return ""
}

func (x *ZChat) GetClock() *ClockInfo {
	if x != nil {
		return x.Clock
	}
	return nil
}

type MergeLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId     []byte `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToId       []byte `protobuf:"bytes,2,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	StartCount uint64 `protobuf:"varint,3,opt,name=start_count,json=startCount,proto3" json:"start_count,omitempty"`
	EndCount   uint64 `protobuf:"varint,4,opt,name=end_count,json=endCount,proto3" json:"end_count,omitempty"`
	SClock     *Clock `protobuf:"bytes,5,opt,name=s_clock,json=sClock,proto3" json:"s_clock,omitempty"`
	EClock     *Clock `protobuf:"bytes,6,opt,name=e_clock,json=eClock,proto3" json:"e_clock,omitempty"`
	MergeAt    uint64 `protobuf:"varint,7,opt,name=merge_at,json=mergeAt,proto3" json:"merge_at,omitempty"`
}

func (x *MergeLog) Reset() {
	*x = MergeLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeLog) ProtoMessage() {}

func (x *MergeLog) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeLog.ProtoReflect.Descriptor instead.
func (*MergeLog) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{3}
}

func (x *MergeLog) GetFromId() []byte {
	if x != nil {
		return x.FromId
	}
	return nil
}

func (x *MergeLog) GetToId() []byte {
	if x != nil {
		return x.ToId
	}
	return nil
}

func (x *MergeLog) GetStartCount() uint64 {
	if x != nil {
		return x.StartCount
	}
	return 0
}

func (x *MergeLog) GetEndCount() uint64 {
	if x != nil {
		return x.EndCount
	}
	return 0
}

func (x *MergeLog) GetSClock() *Clock {
	if x != nil {
		return x.SClock
	}
	return nil
}

func (x *MergeLog) GetEClock() *Clock {
	if x != nil {
		return x.EClock
	}
	return nil
}

func (x *MergeLog) GetMergeAt() uint64 {
	if x != nil {
		return x.MergeAt
	}
	return 0
}

type Clock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]uint64 `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Clock) Reset() {
	*x = Clock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clock) ProtoMessage() {}

func (x *Clock) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clock.ProtoReflect.Descriptor instead.
func (*Clock) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{4}
}

func (x *Clock) GetValues() map[string]uint64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type ClockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock     *Clock `protobuf:"bytes,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id        []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"` //id 为node id
	MessageId []byte `protobuf:"bytes,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Count     uint64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CreateAt  uint64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *ClockInfo) Reset() {
	*x = ClockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockInfo) ProtoMessage() {}

func (x *ClockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockInfo.ProtoReflect.Descriptor instead.
func (*ClockInfo) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{5}
}

func (x *ClockInfo) GetClock() *Clock {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *ClockInfo) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ClockInfo) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *ClockInfo) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ClockInfo) GetCreateAt() uint64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

// type = TYPE_CLOCK_NODE
type ClockNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock      *Clock `protobuf:"bytes,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id         []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MessageId  []byte `protobuf:"bytes,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Count      uint64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CreateAt   uint64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	RawMessage []byte `protobuf:"bytes,6,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
}

func (x *ClockNode) Reset() {
	*x = ClockNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockNode) ProtoMessage() {}

func (x *ClockNode) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockNode.ProtoReflect.Descriptor instead.
func (*ClockNode) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{6}
}

func (x *ClockNode) GetClock() *Clock {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *ClockNode) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ClockNode) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *ClockNode) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ClockNode) GetCreateAt() uint64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *ClockNode) GetRawMessage() []byte {
	if x != nil {
		return x.RawMessage
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIds []string `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_protos_msg_proto_rawDescGZIP(), []int{7}
}

func (x *NodeInfo) GetNodeIds() []string {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

var File_protos_msg_proto protoreflect.FileDescriptor

var file_protos_msg_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x5a,
	0x70, 0x32, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x09, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x50, 0x75, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x5a, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x5a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x53, 0x0a,
	0x05, 0x5a, 0x43, 0x68, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x22, 0xe1, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x73,
	0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x06, 0x65, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x31, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x92, 0x01,
	0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x22, 0xb3, 0x01, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61,
	0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x2a,
	0x2b, 0x0a, 0x09, 0x5a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x0a,
	0x55, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x55, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x2c, 0x0a, 0x07,
	0x5a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x5a, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x5a, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01, 0x2a, 0x29, 0x0a, 0x09, 0x5a, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x5a, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x4d, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x42, 0x43, 0x10, 0x01, 0x2a, 0x61, 0x0a, 0x05, 0x5a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x0a, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x43, 0x4b,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x54,
	0x45, 0x57, 0x41, 0x59, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x5a, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x5a, 0x43, 0x48, 0x41, 0x54, 0x10, 0x04, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_msg_proto_rawDescOnce sync.Once
	file_protos_msg_proto_rawDescData = file_protos_msg_proto_rawDesc
)

func file_protos_msg_proto_rawDescGZIP() []byte {
	file_protos_msg_proto_rawDescOnce.Do(func() {
		file_protos_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_msg_proto_rawDescData)
	})
	return file_protos_msg_proto_rawDescData
}

var file_protos_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_protos_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_msg_proto_goTypes = []interface{}{
	(ZIdentity)(0),    // 0: protos.ZIdentity
	(ZAction)(0),      // 1: protos.ZAction
	(ZPushType)(0),    // 2: protos.ZPushType
	(ZType)(0),        // 3: protos.ZType
	(*Zp2P)(nil),      // 4: protos.Zp2p
	(*ZMessage)(nil),  // 5: protos.ZMessage
	(*ZChat)(nil),     // 6: protos.ZChat
	(*MergeLog)(nil),  // 7: protos.MergeLog
	(*Clock)(nil),     // 8: protos.Clock
	(*ClockInfo)(nil), // 9: protos.ClockInfo
	(*ClockNode)(nil), // 10: protos.ClockNode
	(*NodeInfo)(nil),  // 11: protos.NodeInfo
	nil,               // 12: protos.Clock.ValuesEntry
}
var file_protos_msg_proto_depIdxs = []int32{
	0,  // 0: protos.Zp2p.type:type_name -> protos.ZIdentity
	1,  // 1: protos.Zp2p.action:type_name -> protos.ZAction
	2,  // 2: protos.Zp2p.push_type:type_name -> protos.ZPushType
	5,  // 3: protos.Zp2p.message:type_name -> protos.ZMessage
	3,  // 4: protos.ZMessage.type:type_name -> protos.ZType
	1,  // 5: protos.ZMessage.action:type_name -> protos.ZAction
	0,  // 6: protos.ZMessage.identity:type_name -> protos.ZIdentity
	9,  // 7: protos.ZChat.clock:type_name -> protos.ClockInfo
	8,  // 8: protos.MergeLog.s_clock:type_name -> protos.Clock
	8,  // 9: protos.MergeLog.e_clock:type_name -> protos.Clock
	12, // 10: protos.Clock.values:type_name -> protos.Clock.ValuesEntry
	8,  // 11: protos.ClockInfo.clock:type_name -> protos.Clock
	8,  // 12: protos.ClockNode.clock:type_name -> protos.Clock
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_protos_msg_proto_init() }
func file_protos_msg_proto_init() {
	if File_protos_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zp2P); i {
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
		file_protos_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZMessage); i {
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
		file_protos_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZChat); i {
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
		file_protos_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeLog); i {
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
		file_protos_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clock); i {
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
		file_protos_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockInfo); i {
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
		file_protos_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockNode); i {
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
		file_protos_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
			RawDescriptor: file_protos_msg_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_msg_proto_goTypes,
		DependencyIndexes: file_protos_msg_proto_depIdxs,
		EnumInfos:         file_protos_msg_proto_enumTypes,
		MessageInfos:      file_protos_msg_proto_msgTypes,
	}.Build()
	File_protos_msg_proto = out.File
	file_protos_msg_proto_rawDesc = nil
	file_protos_msg_proto_goTypes = nil
	file_protos_msg_proto_depIdxs = nil
}
