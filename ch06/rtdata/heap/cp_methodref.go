package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
