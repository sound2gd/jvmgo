package base

// 读取字节码，保存pc
type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUInt8())
}

func (self *BytecodeReader) ReadUInt8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadUInt16() uint16 {
	high := uint16(self.ReadUInt8())
	low := uint16(self.ReadUInt8())
	return (high<<8) | low
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUInt16())
}

func (self *BytecodeReader) ReadUInt32() uint32 {
	b4 := uint32(self.ReadUInt8())
	b3 := uint32(self.ReadUInt8())
	b2 := uint32(self.ReadUInt8())
	b1 := uint32(self.ReadUInt8())
	return b4<<24 | b3<<16 | b2<<8 | b1
}

func (self *BytecodeReader) ReadInt32() int32 {
	return int32(self.ReadUInt32())
}

// tableswitch和lookupswitch指令需要
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// tableswitch和lookupswitch指令需要
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUInt8()
	}
}
