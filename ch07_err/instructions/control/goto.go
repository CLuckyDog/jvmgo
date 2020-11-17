package control

import "jvmgo/ch07_err/instructions/base"
import "jvmgo/ch07_err/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
