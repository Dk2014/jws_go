// Code generated by protoc-gen-go.
// source: ProtobufGen_hdprewardlist.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HDPREWARDLIST struct {
	// * 等级对应档
	RewardListID *uint32 `protobuf:"varint,1,req,def=0" json:"RewardListID,omitempty"`
	// * 积分对应档
	PointRankID      *uint32                   `protobuf:"varint,4,req,def=0" json:"PointRankID,omitempty"`
	Show_Table       []*HDPREWARDLIST_ShowRule `protobuf:"bytes,2,rep" json:"Show_Table,omitempty"`
	Loot_Table       []*HDPREWARDLIST_LootRule `protobuf:"bytes,3,rep" json:"Loot_Table,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *HDPREWARDLIST) Reset()         { *m = HDPREWARDLIST{} }
func (m *HDPREWARDLIST) String() string { return proto.CompactTextString(m) }
func (*HDPREWARDLIST) ProtoMessage()    {}

const Default_HDPREWARDLIST_RewardListID uint32 = 0
const Default_HDPREWARDLIST_PointRankID uint32 = 0

func (m *HDPREWARDLIST) GetRewardListID() uint32 {
	if m != nil && m.RewardListID != nil {
		return *m.RewardListID
	}
	return Default_HDPREWARDLIST_RewardListID
}

func (m *HDPREWARDLIST) GetPointRankID() uint32 {
	if m != nil && m.PointRankID != nil {
		return *m.PointRankID
	}
	return Default_HDPREWARDLIST_PointRankID
}

func (m *HDPREWARDLIST) GetShow_Table() []*HDPREWARDLIST_ShowRule {
	if m != nil {
		return m.Show_Table
	}
	return nil
}

func (m *HDPREWARDLIST) GetLoot_Table() []*HDPREWARDLIST_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

type HDPREWARDLIST_ShowRule struct {
	// * 掉落1
	ShowItemID *string `protobuf:"bytes,1,opt,def=" json:"ShowItemID,omitempty"`
	// * 掉落数量
	ShowItemNum      *uint32 `protobuf:"varint,2,opt,def=0" json:"ShowItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HDPREWARDLIST_ShowRule) Reset()         { *m = HDPREWARDLIST_ShowRule{} }
func (m *HDPREWARDLIST_ShowRule) String() string { return proto.CompactTextString(m) }
func (*HDPREWARDLIST_ShowRule) ProtoMessage()    {}

const Default_HDPREWARDLIST_ShowRule_ShowItemNum uint32 = 0

func (m *HDPREWARDLIST_ShowRule) GetShowItemID() string {
	if m != nil && m.ShowItemID != nil {
		return *m.ShowItemID
	}
	return ""
}

func (m *HDPREWARDLIST_ShowRule) GetShowItemNum() uint32 {
	if m != nil && m.ShowItemNum != nil {
		return *m.ShowItemNum
	}
	return Default_HDPREWARDLIST_ShowRule_ShowItemNum
}

type HDPREWARDLIST_LootRule struct {
	// * 是否掉落的概率
	LootChance *float32 `protobuf:"fixed32,1,opt,def=0" json:"LootChance,omitempty"`
	// * 掉落组ID
	LootTemplateID *string `protobuf:"bytes,2,opt,def=" json:"LootTemplateID,omitempty"`
	// * 抽取次数
	LootTime         *uint32 `protobuf:"varint,3,opt,def=0" json:"LootTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HDPREWARDLIST_LootRule) Reset()         { *m = HDPREWARDLIST_LootRule{} }
func (m *HDPREWARDLIST_LootRule) String() string { return proto.CompactTextString(m) }
func (*HDPREWARDLIST_LootRule) ProtoMessage()    {}

const Default_HDPREWARDLIST_LootRule_LootChance float32 = 0
const Default_HDPREWARDLIST_LootRule_LootTime uint32 = 0

func (m *HDPREWARDLIST_LootRule) GetLootChance() float32 {
	if m != nil && m.LootChance != nil {
		return *m.LootChance
	}
	return Default_HDPREWARDLIST_LootRule_LootChance
}

func (m *HDPREWARDLIST_LootRule) GetLootTemplateID() string {
	if m != nil && m.LootTemplateID != nil {
		return *m.LootTemplateID
	}
	return ""
}

func (m *HDPREWARDLIST_LootRule) GetLootTime() uint32 {
	if m != nil && m.LootTime != nil {
		return *m.LootTime
	}
	return Default_HDPREWARDLIST_LootRule_LootTime
}

type HDPREWARDLIST_ARRAY struct {
	Items            []*HDPREWARDLIST `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HDPREWARDLIST_ARRAY) Reset()         { *m = HDPREWARDLIST_ARRAY{} }
func (m *HDPREWARDLIST_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HDPREWARDLIST_ARRAY) ProtoMessage()    {}

func (m *HDPREWARDLIST_ARRAY) GetItems() []*HDPREWARDLIST {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}