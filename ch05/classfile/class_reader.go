package classfile

//即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
import (
	"encoding/binary"
)

/*
	相同类型的多条数据一般按表（table）的形式存储在class文件
	中。表由表头和表项（item）构成，表头是u2或u4整数。假设表头是
	n，后面就紧跟着n个表项数据。
*/

type ClassReader struct {
	data []byte
}

//readUint8（）读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//readUint16（）读取u2类型
//Go标准库encoding/binary包中定义了一个变量BigEndian，正好
//可以从[]byte中解码多字节数据
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//eadUint32（）读取u4类型数据
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
//readUint64（）读取uint64（Java虚拟机规范并没有定义u8）类型
//数据
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//readUint16s（）读取uint16表，表的大小由开头的uint16数据指
//出
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//是readBytes（），用于读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}




