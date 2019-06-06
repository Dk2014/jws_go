// Code generated by protoc-gen-go.
// source: ProtobufGen_wheelcost.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WHEELCOST struct {
	// * 活动ID
	ActivityID *uint32 `protobuf:"varint,1,opt,def=0" json:"ActivityID,omitempty"`
	// * 兑换次数
	Index *uint32 `protobuf:"varint,2,opt,def=0" json:"Index,omitempty"`
	// * 消耗代币
	Cost             *uint32 `protobuf:"varint,3,opt,def=0" json:"Cost,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WHEELCOST) Reset()         { *m = WHEELCOST{} }
func (m *WHEELCOST) String() string { return proto.CompactTextString(m) }
func (*WHEELCOST) ProtoMessage()    {}

const Default_WHEELCOST_ActivityID uint32 = 0
const Default_WHEELCOST_Index uint32 = 0
const Default_WHEELCOST_Cost uint32 = 0

func (m *WHEELCOST) GetActivityID() uint32 {
	if m != nil && m.ActivityID != nil {
		return *m.ActivityID
	}
	return Default_WHEELCOST_ActivityID
}

func (m *WHEELCOST) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_WHEELCOST_Index
}

func (m *WHEELCOST) GetCost() uint32 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return Default_WHEELCOST_Cost
}

type WHEELCOST_ARRAY struct {
	Items            []*WHEELCOST `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WHEELCOST_ARRAY) Reset()         { *m = WHEELCOST_ARRAY{} }
func (m *WHEELCOST_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WHEELCOST_ARRAY) ProtoMessage()    {}

func (m *WHEELCOST_ARRAY) GetItems() []*WHEELCOST {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}