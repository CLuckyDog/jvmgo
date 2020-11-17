package main

import "fmt"
import "jvmgo/ch06/instructions"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	fmt.Println("1111111111111111111111111")
	for {
		fmt.Println("222222222222222")
		pc := frame.NextPC()
		thread.SetPC(pc)
		fmt.Println("333333333333333")
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		fmt.Println("44444444444444444")
		fmt.Printf("opcode ---%v\n",opcode)

		inst := instructions.NewInstruction(opcode)
		fmt.Println("55555555555555555")
		inst.FetchOperands(reader)
		fmt.Println("66666666666666666")
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
