package constants

import "jvmgo/ch07_err/instructions/base"
import "jvmgo/ch07_err/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
