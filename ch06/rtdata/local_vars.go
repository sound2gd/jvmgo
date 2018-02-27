package rtdata

import (
	"jvmgo/ch06/rtdata/heap"
	"math"
)

// 局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// float32类型也是32位，先保存成字节，存为int32，取出的时候转成float32
func (self LocalVars) SetFloat(index uint, val float32) {
	self[index].num = int32(math.Float32bits(val))
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// Long类型要拆成2个int存储
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	// 转成无符号32位
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high<<32) | int64(low)
}

// Double拆分成long型存储就可以
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}

// boolean,char,short都可以使用int处理
