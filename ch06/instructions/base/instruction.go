package base

import "jvmgo/ch06/rtdata"

// 指令操作: 获取操作数 -> 执行
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtdata.Frame)
}

type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// empty
}

func (self *NoOperandsInstruction) Execute(frame *rtdata.Frame) {
	// empty
}

// 跳转指令，表示跳转Offset偏移量
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 存储和装载都是单字节指令，Index表示局部变量表的位置索引
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUInt8())
}

// 有些指令要访问运行时常量池，常量池索引是2个字节，用这个抽象,Index表示常量池索引
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUInt16())
}
