// Code generated by protoc-gen-go.
// source: ProtobufGen_guildnumbers.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GUILDNUMBERS struct {
	// * 公会等级
	GuildLevel *uint32 `protobuf:"varint,1,req,def=0" json:"GuildLevel,omitempty"`
	// * 公会人数上限
	GuildNumbersLimit *uint32 `protobuf:"varint,2,req,def=0" json:"GuildNumbersLimit,omitempty"`
	// * 公会升一级所需经验
	GuildEX *uint32 `protobuf:"varint,3,req,def=0" json:"GuildEX,omitempty"`
	// * 祈福代币加成
	GuildSignAward   *float32 `protobuf:"fixed32,4,req,def=0" json:"GuildSignAward,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GUILDNUMBERS) Reset()         { *m = GUILDNUMBERS{} }
func (m *GUILDNUMBERS) String() string { return proto.CompactTextString(m) }
func (*GUILDNUMBERS) ProtoMessage()    {}

const Default_GUILDNUMBERS_GuildLevel uint32 = 0
const Default_GUILDNUMBERS_GuildNumbersLimit uint32 = 0
const Default_GUILDNUMBERS_GuildEX uint32 = 0
const Default_GUILDNUMBERS_GuildSignAward float32 = 0

func (m *GUILDNUMBERS) GetGuildLevel() uint32 {
	if m != nil && m.GuildLevel != nil {
		return *m.GuildLevel
	}
	return Default_GUILDNUMBERS_GuildLevel
}

func (m *GUILDNUMBERS) GetGuildNumbersLimit() uint32 {
	if m != nil && m.GuildNumbersLimit != nil {
		return *m.GuildNumbersLimit
	}
	return Default_GUILDNUMBERS_GuildNumbersLimit
}

func (m *GUILDNUMBERS) GetGuildEX() uint32 {
	if m != nil && m.GuildEX != nil {
		return *m.GuildEX
	}
	return Default_GUILDNUMBERS_GuildEX
}

func (m *GUILDNUMBERS) GetGuildSignAward() float32 {
	if m != nil && m.GuildSignAward != nil {
		return *m.GuildSignAward
	}
	return Default_GUILDNUMBERS_GuildSignAward
}

type GUILDNUMBERS_ARRAY struct {
	Items            []*GUILDNUMBERS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GUILDNUMBERS_ARRAY) Reset()         { *m = GUILDNUMBERS_ARRAY{} }
func (m *GUILDNUMBERS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GUILDNUMBERS_ARRAY) ProtoMessage()    {}

func (m *GUILDNUMBERS_ARRAY) GetItems() []*GUILDNUMBERS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}