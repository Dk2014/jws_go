// Code generated by protoc-gen-go.
// source: ProtobufGen_gstgoldbonus.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GSTGOLDBONUS struct {
	// * 科技id
	GSTid *uint32 `protobuf:"varint,1,req,def=0" json:"GSTid,omitempty"`
	// * 科技等级
	Level *uint32 `protobuf:"varint,2,req,def=0" json:"Level,omitempty"`
	// * 军团等级需求
	GuildLevelReq *uint32 `protobuf:"varint,3,req,def=0" json:"GuildLevelReq,omitempty"`
	// * 金币购买加成比率
	GoldPurchaseBonusRate *float32 `protobuf:"fixed32,4,req,def=0" json:"GoldPurchaseBonusRate,omitempty"`
	// * 需要多少科技点升这一级
	GSP              *uint32 `protobuf:"varint,5,req,def=0" json:"GSP,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GSTGOLDBONUS) Reset()         { *m = GSTGOLDBONUS{} }
func (m *GSTGOLDBONUS) String() string { return proto.CompactTextString(m) }
func (*GSTGOLDBONUS) ProtoMessage()    {}

const Default_GSTGOLDBONUS_GSTid uint32 = 0
const Default_GSTGOLDBONUS_Level uint32 = 0
const Default_GSTGOLDBONUS_GuildLevelReq uint32 = 0
const Default_GSTGOLDBONUS_GoldPurchaseBonusRate float32 = 0
const Default_GSTGOLDBONUS_GSP uint32 = 0

func (m *GSTGOLDBONUS) GetGSTid() uint32 {
	if m != nil && m.GSTid != nil {
		return *m.GSTid
	}
	return Default_GSTGOLDBONUS_GSTid
}

func (m *GSTGOLDBONUS) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_GSTGOLDBONUS_Level
}

func (m *GSTGOLDBONUS) GetGuildLevelReq() uint32 {
	if m != nil && m.GuildLevelReq != nil {
		return *m.GuildLevelReq
	}
	return Default_GSTGOLDBONUS_GuildLevelReq
}

func (m *GSTGOLDBONUS) GetGoldPurchaseBonusRate() float32 {
	if m != nil && m.GoldPurchaseBonusRate != nil {
		return *m.GoldPurchaseBonusRate
	}
	return Default_GSTGOLDBONUS_GoldPurchaseBonusRate
}

func (m *GSTGOLDBONUS) GetGSP() uint32 {
	if m != nil && m.GSP != nil {
		return *m.GSP
	}
	return Default_GSTGOLDBONUS_GSP
}

type GSTGOLDBONUS_ARRAY struct {
	Items            []*GSTGOLDBONUS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GSTGOLDBONUS_ARRAY) Reset()         { *m = GSTGOLDBONUS_ARRAY{} }
func (m *GSTGOLDBONUS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GSTGOLDBONUS_ARRAY) ProtoMessage()    {}

func (m *GSTGOLDBONUS_ARRAY) GetItems() []*GSTGOLDBONUS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}