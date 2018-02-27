package main

import (
	"fmt"
	"jvmgo/ch06/instructions"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

func interpert(method *heap.Method) {
	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
	//codeAttr := methodInfo.CodeAttribute()
	//maxLocals := codeAttr.MaxLocals()
	//maxStack := codeAttr.MaxStack()
	//bytecode := codeAttr.Code()
	//
	//thread := rtdata.NewThread()
	//frame := thread.NewFrame(maxLocals, maxStack)
	//thread.PushFrame(frame)
	//
	//defer catchErr(frame)
	//loop(thread, bytecode)
}

func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtdata.Thread, bytes []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 解码
		reader.Reset(bytes, pc)
		opcode := reader.ReadUInt8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.PC())

		// 执行
		fmt.Printf("pc: %2d, inst: %T %v \n", pc, inst, inst)
		inst.Execute(frame)
	}

}
