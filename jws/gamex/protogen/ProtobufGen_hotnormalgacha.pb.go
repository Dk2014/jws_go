// Code generated by protoc-gen-go.
// source: ProtobufGen_hotnormalgacha.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HOTNORMALGACHA struct {
	// * 抽取ID
	GachaID *uint32 `protobuf:"varint,1,opt,def=0" json:"GachaID,omitempty"`
	// * 物品ID
	ItemID *string `protobuf:"bytes,8,opt,def=" json:"ItemID,omitempty"`
	// * 物品数量
	ItemCount *uint32 `protobuf:"varint,9,opt,def=0" json:"ItemCount,omitempty"`
	// * 权重
	Weight           *uint32 `protobuf:"varint,10,opt,def=0" json:"Weight,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HOTNORMALGACHA) Reset()         { *m = HOTNORMALGACHA{} }
func (m *HOTNORMALGACHA) String() string { return proto.CompactTextString(m) }
func (*HOTNORMALGACHA) ProtoMessage()    {}

const Default_HOTNORMALGACHA_GachaID uint32 = 0
const Default_HOTNORMALGACHA_ItemCount uint32 = 0
const Default_HOTNORMALGACHA_Weight uint32 = 0

func (m *HOTNORMALGACHA) GetGachaID() uint32 {
	if m != nil && m.GachaID != nil {
		return *m.GachaID
	}
	return Default_HOTNORMALGACHA_GachaID
}

func (m *HOTNORMALGACHA) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *HOTNORMALGACHA) GetItemCount() uint32 {
	if m != nil && m.ItemCount != nil {
		return *m.ItemCount
	}
	return Default_HOTNORMALGACHA_ItemCount
}

func (m *HOTNORMALGACHA) GetWeight() uint32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return Default_HOTNORMALGACHA_Weight
}

type HOTNORMALGACHA_ARRAY struct {
	Items            []*HOTNORMALGACHA `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HOTNORMALGACHA_ARRAY) Reset()         { *m = HOTNORMALGACHA_ARRAY{} }
func (m *HOTNORMALGACHA_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HOTNORMALGACHA_ARRAY) ProtoMessage()    {}

func (m *HOTNORMALGACHA_ARRAY) GetItems() []*HOTNORMALGACHA {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
