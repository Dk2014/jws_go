// Code generated by protoc-gen-go.
// source: ProtobufGen_wheelgacha.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WHEELGACHA struct {
	// * 抽取ID
	GachaID *uint32 `protobuf:"varint,1,opt,def=0" json:"GachaID,omitempty"`
	// * 物品ID
	ItemID *string `protobuf:"bytes,2,opt,def=" json:"ItemID,omitempty"`
	// * 物品数量
	ItemCount *uint32 `protobuf:"varint,3,opt,def=0" json:"ItemCount,omitempty"`
	// * 权重
	Weight *uint32 `protobuf:"varint,4,opt,def=0" json:"Weight,omitempty"`
	// * 是否稀有
	Unusual          *uint32 `protobuf:"varint,5,opt,def=0" json:"Unusual,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WHEELGACHA) Reset()         { *m = WHEELGACHA{} }
func (m *WHEELGACHA) String() string { return proto.CompactTextString(m) }
func (*WHEELGACHA) ProtoMessage()    {}

const Default_WHEELGACHA_GachaID uint32 = 0
const Default_WHEELGACHA_ItemCount uint32 = 0
const Default_WHEELGACHA_Weight uint32 = 0
const Default_WHEELGACHA_Unusual uint32 = 0

func (m *WHEELGACHA) GetGachaID() uint32 {
	if m != nil && m.GachaID != nil {
		return *m.GachaID
	}
	return Default_WHEELGACHA_GachaID
}

func (m *WHEELGACHA) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *WHEELGACHA) GetItemCount() uint32 {
	if m != nil && m.ItemCount != nil {
		return *m.ItemCount
	}
	return Default_WHEELGACHA_ItemCount
}

func (m *WHEELGACHA) GetWeight() uint32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return Default_WHEELGACHA_Weight
}

func (m *WHEELGACHA) GetUnusual() uint32 {
	if m != nil && m.Unusual != nil {
		return *m.Unusual
	}
	return Default_WHEELGACHA_Unusual
}

type WHEELGACHA_ARRAY struct {
	Items            []*WHEELGACHA `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *WHEELGACHA_ARRAY) Reset()         { *m = WHEELGACHA_ARRAY{} }
func (m *WHEELGACHA_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WHEELGACHA_ARRAY) ProtoMessage()    {}

func (m *WHEELGACHA_ARRAY) GetItems() []*WHEELGACHA {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}