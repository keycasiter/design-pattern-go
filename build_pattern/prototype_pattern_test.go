package build_pattern

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPrototypeWithShallowCopy(t *testing.T) {
	prototype := &Prototype{
		PrototypeAttr: "attr",
		PrototypeObj: &PrototypeObj{
			PrototypeObjAttr: "objAttr",
		},
	}
	shallowPrototype := prototype.Clone(Shallow_Copy)

	fmt.Println(unsafe.Pointer(prototype), unsafe.Pointer(shallowPrototype))
	//此时内嵌对象指向同一个内存地址，彼此修改受影响
	fmt.Println(unsafe.Pointer(prototype.PrototypeObj), unsafe.Pointer(shallowPrototype.PrototypeObj))
	fmt.Println(prototype.PrototypeAttr,
		shallowPrototype.PrototypeAttr,
		prototype.PrototypeObj.PrototypeObjAttr,
		shallowPrototype.PrototypeObj.PrototypeObjAttr,
	)
}

func TestPrototypeWithDeepCopy(t *testing.T) {
	prototype := &Prototype{
		PrototypeAttr: "attr",
		PrototypeObj: &PrototypeObj{
			PrototypeObjAttr: "objAttr",
		},
	}

	deepPrototype := prototype.Clone(Deep_Copy)

	fmt.Println(unsafe.Pointer(prototype), unsafe.Pointer(deepPrototype))
	//此时内嵌对象指向同一个内存地址，彼此修改受影响
	fmt.Println(unsafe.Pointer(prototype.PrototypeObj), unsafe.Pointer(deepPrototype.PrototypeObj))
	fmt.Println(prototype.PrototypeAttr,
		deepPrototype.PrototypeAttr,
		prototype.PrototypeObj.PrototypeObjAttr,
		deepPrototype.PrototypeObj.PrototypeObjAttr,
	)
}
