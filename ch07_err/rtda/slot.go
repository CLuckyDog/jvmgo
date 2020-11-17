package rtda

import "jvmgo/ch07_err/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
