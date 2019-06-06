// Code generated by protoc-gen-go.
// source: ProtobufGen_item.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Item struct {
	// * 商品ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// *
	Type *string `protobuf:"bytes,2,opt,def=" json:"Type,omitempty"`
	// *
	Part *string `protobuf:"bytes,3,opt,def=" json:"Part,omitempty"`
	// *
	EnableLevel *int32 `protobuf:"varint,4,opt,def=0" json:"EnableLevel,omitempty"`
	// *
	Tier *int32 `protobuf:"varint,5,opt,def=0" json:"Tier,omitempty"`
	// *
	RareLevel *int32 `protobuf:"varint,6,opt,def=0" json:"RareLevel,omitempty"`
	// *
	FuseLevelLimit *int32 `protobuf:"varint,7,opt,def=0" json:"FuseLevelLimit,omitempty"`
	// *
	LootScore *int32 `protobuf:"varint,8,opt,def=0" json:"LootScore,omitempty"`
	// *
	SetGear *int32 `protobuf:"varint,9,opt,def=0" json:"SetGear,omitempty"`
	// *
	NameIDS *string `protobuf:"bytes,10,opt,def=" json:"NameIDS,omitempty"`
	// *
	DescriptionIDS *string `protobuf:"bytes,11,opt,def=" json:"DescriptionIDS,omitempty"`
	// *
	Icon *string `protobuf:"bytes,12,opt,def=" json:"Icon,omitempty"`
	// *
	Instance *string `protobuf:"bytes,13,opt,def=" json:"Instance,omitempty"`
	// *
	Attack *float32 `protobuf:"fixed32,14,opt,def=0" json:"Attack,omitempty"`
	// *
	Defense *float32 `protobuf:"fixed32,15,opt,def=0" json:"Defense,omitempty"`
	// *
	HP               *float32 `protobuf:"fixed32,16,opt,def=0" json:"HP,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}

const Default_Item_EnableLevel int32 = 0
const Default_Item_Tier int32 = 0
const Default_Item_RareLevel int32 = 0
const Default_Item_FuseLevelLimit int32 = 0
const Default_Item_LootScore int32 = 0
const Default_Item_SetGear int32 = 0
const Default_Item_Attack float32 = 0
const Default_Item_Defense float32 = 0
const Default_Item_HP float32 = 0

func (m *Item) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *Item) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Item) GetPart() string {
	if m != nil && m.Part != nil {
		return *m.Part
	}
	return ""
}

func (m *Item) GetEnableLevel() int32 {
	if m != nil && m.EnableLevel != nil {
		return *m.EnableLevel
	}
	return Default_Item_EnableLevel
}

func (m *Item) GetTier() int32 {
	if m != nil && m.Tier != nil {
		return *m.Tier
	}
	return Default_Item_Tier
}

func (m *Item) GetRareLevel() int32 {
	if m != nil && m.RareLevel != nil {
		return *m.RareLevel
	}
	return Default_Item_RareLevel
}

func (m *Item) GetFuseLevelLimit() int32 {
	if m != nil && m.FuseLevelLimit != nil {
		return *m.FuseLevelLimit
	}
	return Default_Item_FuseLevelLimit
}

func (m *Item) GetLootScore() int32 {
	if m != nil && m.LootScore != nil {
		return *m.LootScore
	}
	return Default_Item_LootScore
}

func (m *Item) GetSetGear() int32 {
	if m != nil && m.SetGear != nil {
		return *m.SetGear
	}
	return Default_Item_SetGear
}

func (m *Item) GetNameIDS() string {
	if m != nil && m.NameIDS != nil {
		return *m.NameIDS
	}
	return ""
}

func (m *Item) GetDescriptionIDS() string {
	if m != nil && m.DescriptionIDS != nil {
		return *m.DescriptionIDS
	}
	return ""
}

func (m *Item) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *Item) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

func (m *Item) GetAttack() float32 {
	if m != nil && m.Attack != nil {
		return *m.Attack
	}
	return Default_Item_Attack
}

func (m *Item) GetDefense() float32 {
	if m != nil && m.Defense != nil {
		return *m.Defense
	}
	return Default_Item_Defense
}

func (m *Item) GetHP() float32 {
	if m != nil && m.HP != nil {
		return *m.HP
	}
	return Default_Item_HP
}

type Item_ARRAY struct {
	Items            []*Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Item_ARRAY) Reset()         { *m = Item_ARRAY{} }
func (m *Item_ARRAY) String() string { return proto.CompactTextString(m) }
func (*Item_ARRAY) ProtoMessage()    {}

func (m *Item_ARRAY) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}