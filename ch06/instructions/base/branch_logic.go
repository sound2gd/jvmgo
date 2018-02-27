package base

import "jvmgo/ch06/rtdata"

func Branch(frame *rtdata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPc(nextPC)
}
