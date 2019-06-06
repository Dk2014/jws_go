// Code generated by protoc-gen-go.
// source: ProtobufGen_gvgdailygift.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GVGDAILYGIFT struct {
	// * ID
	CityID           *uint32                  `protobuf:"varint,1,req,def=0" json:"CityID,omitempty"`
	Loot_Table       []*GVGDAILYGIFT_LootRule `protobuf:"bytes,2,rep" json:"Loot_Table,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *GVGDAILYGIFT) Reset()         { *m = GVGDAILYGIFT{} }
func (m *GVGDAILYGIFT) String() string { return proto.CompactTextString(m) }
func (*GVGDAILYGIFT) ProtoMessage()    {}

const Default_GVGDAILYGIFT_CityID uint32 = 0

func (m *GVGDAILYGIFT) GetCityID() uint32 {
	if m != nil && m.CityID != nil {
		return *m.CityID
	}
	return Default_GVGDAILYGIFT_CityID
}

func (m *GVGDAILYGIFT) GetLoot_Table() []*GVGDAILYGIFT_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

type GVGDAILYGIFT_LootRule struct {
	// * 奖励1
	DailyItemID *string `protobuf:"bytes,1,opt,def=" json:"DailyItemID,omitempty"`
	// * 数量1
	DailyItemNum     *uint32 `protobuf:"varint,2,opt,def=0" json:"DailyItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GVGDAILYGIFT_LootRule) Reset()         { *m = GVGDAILYGIFT_LootRule{} }
func (m *GVGDAILYGIFT_LootRule) String() string { return proto.CompactTextString(m) }
func (*GVGDAILYGIFT_LootRule) ProtoMessage()    {}

const Default_GVGDAILYGIFT_LootRule_DailyItemNum uint32 = 0

func (m *GVGDAILYGIFT_LootRule) GetDailyItemID() string {
	if m != nil && m.DailyItemID != nil {
		return *m.DailyItemID
	}
	return ""
}

func (m *GVGDAILYGIFT_LootRule) GetDailyItemNum() uint32 {
	if m != nil && m.DailyItemNum != nil {
		return *m.DailyItemNum
	}
	return Default_GVGDAILYGIFT_LootRule_DailyItemNum
}

type GVGDAILYGIFT_ARRAY struct {
	Items            []*GVGDAILYGIFT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GVGDAILYGIFT_ARRAY) Reset()         { *m = GVGDAILYGIFT_ARRAY{} }
func (m *GVGDAILYGIFT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GVGDAILYGIFT_ARRAY) ProtoMessage()    {}

func (m *GVGDAILYGIFT_ARRAY) GetItems() []*GVGDAILYGIFT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}