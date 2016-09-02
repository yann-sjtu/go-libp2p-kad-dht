// Code generated by protoc-gen-gogo.
// source: dht.proto
// DO NOT EDIT!

/*
Package dht_pb is a generated protocol buffer package.

It is generated from these files:
	dht.proto

It has these top-level messages:
	Message
*/
package dht_pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Message_MessageType int32

const (
	Message_PUT_VALUE     Message_MessageType = 0
	Message_GET_VALUE     Message_MessageType = 1
	Message_ADD_PROVIDER  Message_MessageType = 2
	Message_GET_PROVIDERS Message_MessageType = 3
	Message_FIND_NODE     Message_MessageType = 4
	Message_PING          Message_MessageType = 5
)

var Message_MessageType_name = map[int32]string{
	0: "PUT_VALUE",
	1: "GET_VALUE",
	2: "ADD_PROVIDER",
	3: "GET_PROVIDERS",
	4: "FIND_NODE",
	5: "PING",
}
var Message_MessageType_value = map[string]int32{
	"PUT_VALUE":     0,
	"GET_VALUE":     1,
	"ADD_PROVIDER":  2,
	"GET_PROVIDERS": 3,
	"FIND_NODE":     4,
	"PING":          5,
}

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}
func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (x *Message_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_MessageType_value, data, "Message_MessageType")
	if err != nil {
		return err
	}
	*x = Message_MessageType(value)
	return nil
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorDht, []int{0, 0} }

type Message_ConnectionType int32

const (
	// sender does not have a connection to peer, and no extra information (default)
	Message_NOT_CONNECTED Message_ConnectionType = 0
	// sender has a live connection to peer
	Message_CONNECTED Message_ConnectionType = 1
	// sender recently connected to peer
	Message_CAN_CONNECT Message_ConnectionType = 2
	// sender recently tried to connect to peer repeatedly but failed to connect
	// ("try" here is loose, but this should signal "made strong effort, failed")
	Message_CANNOT_CONNECT Message_ConnectionType = 3
)

var Message_ConnectionType_name = map[int32]string{
	0: "NOT_CONNECTED",
	1: "CONNECTED",
	2: "CAN_CONNECT",
	3: "CANNOT_CONNECT",
}
var Message_ConnectionType_value = map[string]int32{
	"NOT_CONNECTED":  0,
	"CONNECTED":      1,
	"CAN_CONNECT":    2,
	"CANNOT_CONNECT": 3,
}

func (x Message_ConnectionType) Enum() *Message_ConnectionType {
	p := new(Message_ConnectionType)
	*p = x
	return p
}
func (x Message_ConnectionType) String() string {
	return proto.EnumName(Message_ConnectionType_name, int32(x))
}
func (x *Message_ConnectionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_ConnectionType_value, data, "Message_ConnectionType")
	if err != nil {
		return err
	}
	*x = Message_ConnectionType(value)
	return nil
}
func (Message_ConnectionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorDht, []int{0, 1} }

type Message struct {
	// defines what type of message it is.
	Type *Message_MessageType `protobuf:"varint,1,opt,name=type,enum=dht.pb.Message_MessageType" json:"type,omitempty"`
	// defines what coral cluster level this query/response belongs to.
	ClusterLevelRaw *int32 `protobuf:"varint,10,opt,name=clusterLevelRaw" json:"clusterLevelRaw,omitempty"`
	// Used to specify the key associated with this message.
	// PUT_VALUE, GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	Key *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	// Used to return a value
	// PUT_VALUE, GET_VALUE
	Record *record_pb.Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
	// Used to return peers closer to a key in a query
	// GET_VALUE, GET_PROVIDERS, FIND_NODE
	CloserPeers []*Message_Peer `protobuf:"bytes,8,rep,name=closerPeers" json:"closerPeers,omitempty"`
	// Used to return Providers
	// GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	ProviderPeers    []*Message_Peer `protobuf:"bytes,9,rep,name=providerPeers" json:"providerPeers,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorDht, []int{0} }

func (m *Message) GetType() Message_MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Message_PUT_VALUE
}

func (m *Message) GetClusterLevelRaw() int32 {
	if m != nil && m.ClusterLevelRaw != nil {
		return *m.ClusterLevelRaw
	}
	return 0
}

func (m *Message) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Message) GetRecord() *record_pb.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *Message) GetCloserPeers() []*Message_Peer {
	if m != nil {
		return m.CloserPeers
	}
	return nil
}

func (m *Message) GetProviderPeers() []*Message_Peer {
	if m != nil {
		return m.ProviderPeers
	}
	return nil
}

type Message_Peer struct {
	// ID of a given peer.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// multiaddrs for a given peer
	Addrs [][]byte `protobuf:"bytes,2,rep,name=addrs" json:"addrs,omitempty"`
	// used to signal the sender's connection capabilities to the peer
	Connection       *Message_ConnectionType `protobuf:"varint,3,opt,name=connection,enum=dht.pb.Message_ConnectionType" json:"connection,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Message_Peer) Reset()                    { *m = Message_Peer{} }
func (m *Message_Peer) String() string            { return proto.CompactTextString(m) }
func (*Message_Peer) ProtoMessage()               {}
func (*Message_Peer) Descriptor() ([]byte, []int) { return fileDescriptorDht, []int{0, 0} }

func (m *Message_Peer) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Message_Peer) GetAddrs() [][]byte {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *Message_Peer) GetConnection() Message_ConnectionType {
	if m != nil && m.Connection != nil {
		return *m.Connection
	}
	return Message_NOT_CONNECTED
}

func init() {
	proto.RegisterType((*Message)(nil), "dht.pb.Message")
	proto.RegisterType((*Message_Peer)(nil), "dht.pb.Message.Peer")
	proto.RegisterEnum("dht.pb.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("dht.pb.Message_ConnectionType", Message_ConnectionType_name, Message_ConnectionType_value)
}

var fileDescriptorDht = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x8f, 0x9a, 0x50,
	0x14, 0x86, 0xcb, 0x87, 0x56, 0x0e, 0x8a, 0x78, 0xe2, 0x82, 0xd8, 0xa4, 0x31, 0xae, 0xec, 0x86,
	0x26, 0x2e, 0xba, 0xe8, 0xa2, 0x09, 0x01, 0x6a, 0x4c, 0xec, 0xc5, 0xdc, 0xa2, 0x5d, 0x12, 0x85,
	0x9b, 0x0e, 0x19, 0x22, 0x06, 0x18, 0x27, 0xfe, 0xc2, 0xf9, 0x5b, 0x73, 0x01, 0x99, 0x41, 0x17,
	0xb3, 0xba, 0xef, 0x39, 0xf7, 0x79, 0xcf, 0x17, 0x28, 0xd1, 0x43, 0x61, 0x9e, 0xb2, 0xb4, 0x48,
	0xb1, 0x5b, 0xc9, 0xc3, 0xa4, 0x9f, 0xb1, 0x30, 0xcd, 0xa2, 0x3a, 0x3b, 0x7b, 0x91, 0xe1, 0xf3,
	0x1f, 0x96, 0xe7, 0xfb, 0xff, 0x0c, 0xbf, 0x83, 0x5c, 0x5c, 0x4e, 0xcc, 0x10, 0xa6, 0xc2, 0x5c,
	0x5b, 0x7c, 0x31, 0x6b, 0x83, 0x79, 0xfd, 0x6e, 0x5e, 0x9f, 0x23, 0xb4, 0x02, 0x71, 0x0e, 0xc3,
	0x30, 0x79, 0xca, 0x0b, 0x96, 0xad, 0xd9, 0x99, 0x25, 0x74, 0xff, 0x6c, 0x00, 0xf7, 0x76, 0xe8,
	0x7d, 0x1a, 0x75, 0x90, 0x1e, 0xd9, 0xc5, 0x10, 0xf9, 0xaf, 0x42, 0x4b, 0x89, 0xdf, 0xa0, 0x5b,
	0x0f, 0x62, 0x48, 0x3c, 0xa9, 0x2e, 0x46, 0x66, 0x33, 0xd7, 0xc1, 0xa4, 0x95, 0xa2, 0x57, 0x00,
	0x7f, 0x80, 0x1a, 0x26, 0x69, 0xce, 0xb2, 0x0d, 0x63, 0x59, 0x6e, 0xf4, 0xa6, 0x12, 0xe7, 0xc7,
	0xf7, 0xe3, 0x95, 0x9f, 0xb4, 0x0d, 0xe2, 0x4f, 0x18, 0xf0, 0x25, 0xcf, 0x71, 0xd4, 0x38, 0x95,
	0x0f, 0x9c, 0xb7, 0xe8, 0x24, 0x01, 0xb9, 0x14, 0xa8, 0x81, 0x18, 0x47, 0xd5, 0x45, 0x14, 0xca,
	0x15, 0x8e, 0xa1, 0xb3, 0x8f, 0x22, 0x5e, 0x4b, 0xe4, 0xb5, 0xfa, 0xb4, 0x0e, 0xf0, 0x17, 0x40,
	0x98, 0x1e, 0x8f, 0x2c, 0x2c, 0xe2, 0xf4, 0x58, 0x2d, 0xa4, 0x2d, 0xbe, 0xde, 0xb7, 0xb1, 0xdf,
	0x88, 0xea, 0x84, 0x2d, 0xc7, 0x2c, 0x06, 0xb5, 0x75, 0x5d, 0x1c, 0x80, 0xb2, 0xd9, 0xfa, 0xc1,
	0xce, 0x5a, 0x6f, 0x5d, 0xfd, 0x53, 0x19, 0x2e, 0xdd, 0x26, 0x14, 0xf8, 0x2d, 0xfb, 0x96, 0xe3,
	0x04, 0x1b, 0xea, 0xed, 0x56, 0x8e, 0x4b, 0x75, 0x11, 0x47, 0x30, 0x28, 0x81, 0x26, 0xf3, 0x57,
	0x97, 0x4a, 0xcf, 0xef, 0x15, 0x71, 0x02, 0xe2, 0x39, 0xae, 0x2e, 0x63, 0x8f, 0xaf, 0xb3, 0x22,
	0x4b, 0xbd, 0x33, 0xfb, 0x07, 0xda, 0xed, 0x20, 0xa5, 0x9b, 0x78, 0x7e, 0x60, 0x7b, 0x84, 0xb8,
	0xb6, 0xef, 0x3a, 0x75, 0xc7, 0xf7, 0x50, 0xc0, 0x21, 0xa8, 0xb6, 0x45, 0x1a, 0x82, 0x37, 0x44,
	0x5e, 0xc4, 0x22, 0x2d, 0x97, 0x2e, 0xbd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x4c, 0x62, 0x25,
	0x6b, 0x02, 0x00, 0x00,
}
