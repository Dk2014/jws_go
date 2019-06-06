// Code generated by protoc-gen-go.
// source: ProtobufGen_herowinglevel.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HEROWINGLEVEL struct {
	// * 翅膀星级
	HWLevel *uint32 `protobuf:"varint,1,req,def=0" json:"HWLevel,omitempty"`
	// * 攻击
	ATK *float32 `protobuf:"fixed32,2,opt,def=0" json:"ATK,omitempty"`
	// * 防御
	DEF *float32 `protobuf:"fixed32,3,opt,def=0" json:"DEF,omitempty"`
	// * 生命
	HP *float32 `protobuf:"fixed32,4,opt,def=0" json:"HP,omitempty"`
	// * 回退所需的钻石数量
	RebornCost *uint32 `protobuf:"varint,7,opt,def=0" json:"RebornCost,omitempty"`
	// * 升级材料
	HWLevelupMaterial *string `protobuf:"bytes,5,opt,def=" json:"HWLevelupMaterial,omitempty"`
	// * 升级材料数量
	HWLevelupMaterialCount *uint32 `protobuf:"varint,6,opt,def=0" json:"HWLevelupMaterialCount,omitempty"`
	// * 幻甲到达此星级后，可以继续升级
	HWLevelupJudge   *uint32 `protobuf:"varint,8,opt,def=0" json:"HWLevelupJudge,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HEROWINGLEVEL) Reset()         { *m = HEROWINGLEVEL{} }
func (m *HEROWINGLEVEL) String() string { return proto.CompactTextString(m) }
func (*HEROWINGLEVEL) ProtoMessage()    {}

const Default_HEROWINGLEVEL_HWLevel uint32 = 0
const Default_HEROWINGLEVEL_ATK float32 = 0
const Default_HEROWINGLEVEL_DEF float32 = 0
const Default_HEROWINGLEVEL_HP float32 = 0
const Default_HEROWINGLEVEL_RebornCost uint32 = 0
const Default_HEROWINGLEVEL_HWLevelupMaterialCount uint32 = 0
const Default_HEROWINGLEVEL_HWLevelupJudge uint32 = 0

func (m *HEROWINGLEVEL) GetHWLevel() uint32 {
	if m != nil && m.HWLevel != nil {
		return *m.HWLevel
	}
	return Default_HEROWINGLEVEL_HWLevel
}

func (m *HEROWINGLEVEL) GetATK() float32 {
	if m != nil && m.ATK != nil {
		return *m.ATK
	}
	return Default_HEROWINGLEVEL_ATK
}

func (m *HEROWINGLEVEL) GetDEF() float32 {
	if m != nil && m.DEF != nil {
		return *m.DEF
	}
	return Default_HEROWINGLEVEL_DEF
}

func (m *HEROWINGLEVEL) GetHP() float32 {
	if m != nil && m.HP != nil {
		return *m.HP
	}
	return Default_HEROWINGLEVEL_HP
}

func (m *HEROWINGLEVEL) GetRebornCost() uint32 {
	if m != nil && m.RebornCost != nil {
		return *m.RebornCost
	}
	return Default_HEROWINGLEVEL_RebornCost
}

func (m *HEROWINGLEVEL) GetHWLevelupMaterial() string {
	if m != nil && m.HWLevelupMaterial != nil {
		return *m.HWLevelupMaterial
	}
	return ""
}

func (m *HEROWINGLEVEL) GetHWLevelupMaterialCount() uint32 {
	if m != nil && m.HWLevelupMaterialCount != nil {
		return *m.HWLevelupMaterialCount
	}
	return Default_HEROWINGLEVEL_HWLevelupMaterialCount
}

func (m *HEROWINGLEVEL) GetHWLevelupJudge() uint32 {
	if m != nil && m.HWLevelupJudge != nil {
		return *m.HWLevelupJudge
	}
	return Default_HEROWINGLEVEL_HWLevelupJudge
}

type HEROWINGLEVEL_ARRAY struct {
	Items            []*HEROWINGLEVEL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HEROWINGLEVEL_ARRAY) Reset()         { *m = HEROWINGLEVEL_ARRAY{} }
func (m *HEROWINGLEVEL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HEROWINGLEVEL_ARRAY) ProtoMessage()    {}

func (m *HEROWINGLEVEL_ARRAY) GetItems() []*HEROWINGLEVEL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}