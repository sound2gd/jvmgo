package rtdata

import "jvmgo/ch06/rtdata/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
