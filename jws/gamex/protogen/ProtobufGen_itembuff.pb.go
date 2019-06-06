// Code generated by protoc-gen-go.
// source: ProtobufGen_itembuff.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ITEMBUFF struct {
	// * Buff的ID
	ID *string `protobuf:"bytes,1,req,def=" json:"ID,omitempty"`
	// * 持续时间
	Time *float32 `protobuf:"fixed32,2,opt,name=time,def=0" json:"time,omitempty"`
	// * 角色属性枚举
	PropertyEnum *int32 `protobuf:"varint,3,opt,name=propertyEnum,def=0" json:"propertyEnum,omitempty"`
	// * 影响类别
	BonusType *int32 `protobuf:"varint,4,opt,name=bonusType,def=0" json:"bonusType,omitempty"`
	// * 影响数值
	BonusValue *float32 `protobuf:"fixed32,5,opt,name=bonusValue,def=0" json:"bonusValue,omitempty"`
	// * 光环特效
	AuraEffect *string `protobuf:"bytes,6,opt,name=auraEffect,def=" json:"auraEffect,omitempty"`
	// * 图标
	BuffIconSp *string `protobuf:"bytes,7,opt,name=buffIconSp,def=" json:"buffIconSp,omitempty"`
	// * BUFF特效
	BuffEffect *string `protobuf:"bytes,8,opt,def=" json:"BuffEffect,omitempty"`
	// * BUFF替换材质枚举
	MaterialEffectIndex *int32 `protobuf:"varint,9,opt,def=-1" json:"MaterialEffectIndex,omitempty"`
	// * 生效间隔时间
	CycleTime *float32 `protobuf:"fixed32,10,opt,def=1" json:"CycleTime,omitempty"`
	// * BUFF半径
	Radius *uint32 `protobuf:"varint,11,opt,def=0" json:"Radius,omitempty"`
	// * BUFF转化
	Bufftrans *string `protobuf:"bytes,12,opt,name=bufftrans,def=" json:"bufftrans,omitempty"`
	// * BUFF叠加
	Buffover *int32 `protobuf:"varint,13,opt,name=buffover,def=1" json:"buffover,omitempty"`
	// * BUFF开启条件
	BuffConditionType *uint32 `protobuf:"varint,14,opt,def=0" json:"BuffConditionType,omitempty"`
	// * BUFF开启参数
	BuffConditionValue *string `protobuf:"bytes,15,opt,def=" json:"BuffConditionValue,omitempty"`
	// * BUFF触发概率
	BuffTriggerRate  *float32 `protobuf:"fixed32,16,opt,def=1" json:"BuffTriggerRate,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ITEMBUFF) Reset()         { *m = ITEMBUFF{} }
func (m *ITEMBUFF) String() string { return proto.CompactTextString(m) }
func (*ITEMBUFF) ProtoMessage()    {}

const Default_ITEMBUFF_Time float32 = 0
const Default_ITEMBUFF_PropertyEnum int32 = 0
const Default_ITEMBUFF_BonusType int32 = 0
const Default_ITEMBUFF_BonusValue float32 = 0
const Default_ITEMBUFF_MaterialEffectIndex int32 = -1
const Default_ITEMBUFF_CycleTime float32 = 1
const Default_ITEMBUFF_Radius uint32 = 0
const Default_ITEMBUFF_Buffover int32 = 1
const Default_ITEMBUFF_BuffConditionType uint32 = 0
const Default_ITEMBUFF_BuffTriggerRate float32 = 1

func (m *ITEMBUFF) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *ITEMBUFF) GetTime() float32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return Default_ITEMBUFF_Time
}

func (m *ITEMBUFF) GetPropertyEnum() int32 {
	if m != nil && m.PropertyEnum != nil {
		return *m.PropertyEnum
	}
	return Default_ITEMBUFF_PropertyEnum
}

func (m *ITEMBUFF) GetBonusType() int32 {
	if m != nil && m.BonusType != nil {
		return *m.BonusType
	}
	return Default_ITEMBUFF_BonusType
}

func (m *ITEMBUFF) GetBonusValue() float32 {
	if m != nil && m.BonusValue != nil {
		return *m.BonusValue
	}
	return Default_ITEMBUFF_BonusValue
}

func (m *ITEMBUFF) GetAuraEffect() string {
	if m != nil && m.AuraEffect != nil {
		return *m.AuraEffect
	}
	return ""
}

func (m *ITEMBUFF) GetBuffIconSp() string {
	if m != nil && m.BuffIconSp != nil {
		return *m.BuffIconSp
	}
	return ""
}

func (m *ITEMBUFF) GetBuffEffect() string {
	if m != nil && m.BuffEffect != nil {
		return *m.BuffEffect
	}
	return ""
}

func (m *ITEMBUFF) GetMaterialEffectIndex() int32 {
	if m != nil && m.MaterialEffectIndex != nil {
		return *m.MaterialEffectIndex
	}
	return Default_ITEMBUFF_MaterialEffectIndex
}

func (m *ITEMBUFF) GetCycleTime() float32 {
	if m != nil && m.CycleTime != nil {
		return *m.CycleTime
	}
	return Default_ITEMBUFF_CycleTime
}

func (m *ITEMBUFF) GetRadius() uint32 {
	if m != nil && m.Radius != nil {
		return *m.Radius
	}
	return Default_ITEMBUFF_Radius
}

func (m *ITEMBUFF) GetBufftrans() string {
	if m != nil && m.Bufftrans != nil {
		return *m.Bufftrans
	}
	return ""
}

func (m *ITEMBUFF) GetBuffover() int32 {
	if m != nil && m.Buffover != nil {
		return *m.Buffover
	}
	return Default_ITEMBUFF_Buffover
}

func (m *ITEMBUFF) GetBuffConditionType() uint32 {
	if m != nil && m.BuffConditionType != nil {
		return *m.BuffConditionType
	}
	return Default_ITEMBUFF_BuffConditionType
}

func (m *ITEMBUFF) GetBuffConditionValue() string {
	if m != nil && m.BuffConditionValue != nil {
		return *m.BuffConditionValue
	}
	return ""
}

func (m *ITEMBUFF) GetBuffTriggerRate() float32 {
	if m != nil && m.BuffTriggerRate != nil {
		return *m.BuffTriggerRate
	}
	return Default_ITEMBUFF_BuffTriggerRate
}

type ITEMBUFF_ARRAY struct {
	Items            []*ITEMBUFF `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ITEMBUFF_ARRAY) Reset()         { *m = ITEMBUFF_ARRAY{} }
func (m *ITEMBUFF_ARRAY) String() string { return proto.CompactTextString(m) }
func (*ITEMBUFF_ARRAY) ProtoMessage()    {}

func (m *ITEMBUFF_ARRAY) GetItems() []*ITEMBUFF {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}