package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
