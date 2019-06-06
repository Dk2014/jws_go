// automatically generated by the FlatBuffers compiler, do not modify

package multiplayMsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///玩家状态(数组中以玩家数据中的idx为索引)
type PlayerState struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerState(buf []byte, offset flatbuffers.UOffsetT) *PlayerState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerState{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerState) AccountID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///玩家状态 1-掉线 2-已退出 3-未准备 4-已准备 5-已死亡 6-战斗中
func (rcv *PlayerState) State() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

///玩家状态 1-掉线 2-已退出 3-未准备 4-已准备 5-已死亡 6-战斗中
func (rcv *PlayerState) MutateState(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

///玩家hp
func (rcv *PlayerState) Hp(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PlayerState) HpLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///玩家hp
func (rcv *PlayerState) Pos() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerState) MutatePos(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

/// avatar id
func (rcv *PlayerState) AvatarID(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PlayerState) AvatarIDLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// avatar id
/// cur avatar
func (rcv *PlayerState) CurAvatar() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// cur avatar
func (rcv *PlayerState) MutateCurAvatar(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

func PlayerStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func PlayerStateAddAccountID(builder *flatbuffers.Builder, AccountID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(AccountID), 0)
}
func PlayerStateAddState(builder *flatbuffers.Builder, state int32) {
	builder.PrependInt32Slot(1, state, 0)
}
func PlayerStateAddHp(builder *flatbuffers.Builder, hp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(hp), 0)
}
func PlayerStateStartHpVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PlayerStateAddPos(builder *flatbuffers.Builder, Pos int32) {
	builder.PrependInt32Slot(3, Pos, 0)
}
func PlayerStateAddAvatarID(builder *flatbuffers.Builder, avatarID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(avatarID), 0)
}
func PlayerStateStartAvatarIDVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PlayerStateAddCurAvatar(builder *flatbuffers.Builder, curAvatar int32) {
	builder.PrependInt32Slot(5, curAvatar, 0)
}
func PlayerStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}