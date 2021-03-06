// Code generated by protoc-gen-go.
// source: ProtobufGen_servergroup.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SERVERGROUP struct {
	// * 服务器组ID
	ServerGroupID *uint32 `protobuf:"varint,1,req,def=0" json:"ServerGroupID,omitempty"`
	// * 服务器分组类型
	ServerGroupType *uint32 `protobuf:"varint,2,req,def=0" json:"ServerGroupType,omitempty"`
	// * 服务器分组子类型
	ServerGroupSubType *uint32                        `protobuf:"varint,4,req,def=0" json:"ServerGroupSubType,omitempty"`
	AccCon_Table       []*SERVERGROUP_AcceptCondition `protobuf:"bytes,3,rep" json:"AccCon_Table,omitempty"`
	XXX_unrecognized   []byte                         `json:"-"`
}

func (m *SERVERGROUP) Reset()         { *m = SERVERGROUP{} }
func (m *SERVERGROUP) String() string { return proto.CompactTextString(m) }
func (*SERVERGROUP) ProtoMessage()    {}

const Default_SERVERGROUP_ServerGroupID uint32 = 0
const Default_SERVERGROUP_ServerGroupType uint32 = 0
const Default_SERVERGROUP_ServerGroupSubType uint32 = 0

func (m *SERVERGROUP) GetServerGroupID() uint32 {
	if m != nil && m.ServerGroupID != nil {
		return *m.ServerGroupID
	}
	return Default_SERVERGROUP_ServerGroupID
}

func (m *SERVERGROUP) GetServerGroupType() uint32 {
	if m != nil && m.ServerGroupType != nil {
		return *m.ServerGroupType
	}
	return Default_SERVERGROUP_ServerGroupType
}

func (m *SERVERGROUP) GetServerGroupSubType() uint32 {
	if m != nil && m.ServerGroupSubType != nil {
		return *m.ServerGroupSubType
	}
	return Default_SERVERGROUP_ServerGroupSubType
}

func (m *SERVERGROUP) GetAccCon_Table() []*SERVERGROUP_AcceptCondition {
	if m != nil {
		return m.AccCon_Table
	}
	return nil
}

type SERVERGROUP_AcceptCondition struct {
	// * 服务器分组参数1
	ServerGroupValue1 *uint32 `protobuf:"varint,1,opt,def=0" json:"ServerGroupValue1,omitempty"`
	// * 服务器分组参数2
	ServerGroupValue2 *uint32 `protobuf:"varint,2,opt,def=0" json:"ServerGroupValue2,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *SERVERGROUP_AcceptCondition) Reset()         { *m = SERVERGROUP_AcceptCondition{} }
func (m *SERVERGROUP_AcceptCondition) String() string { return proto.CompactTextString(m) }
func (*SERVERGROUP_AcceptCondition) ProtoMessage()    {}

const Default_SERVERGROUP_AcceptCondition_ServerGroupValue1 uint32 = 0
const Default_SERVERGROUP_AcceptCondition_ServerGroupValue2 uint32 = 0

func (m *SERVERGROUP_AcceptCondition) GetServerGroupValue1() uint32 {
	if m != nil && m.ServerGroupValue1 != nil {
		return *m.ServerGroupValue1
	}
	return Default_SERVERGROUP_AcceptCondition_ServerGroupValue1
}

func (m *SERVERGROUP_AcceptCondition) GetServerGroupValue2() uint32 {
	if m != nil && m.ServerGroupValue2 != nil {
		return *m.ServerGroupValue2
	}
	return Default_SERVERGROUP_AcceptCondition_ServerGroupValue2
}

type SERVERGROUP_ARRAY struct {
	Items            []*SERVERGROUP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SERVERGROUP_ARRAY) Reset()         { *m = SERVERGROUP_ARRAY{} }
func (m *SERVERGROUP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SERVERGROUP_ARRAY) ProtoMessage()    {}

func (m *SERVERGROUP_ARRAY) GetItems() []*SERVERGROUP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
