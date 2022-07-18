package build_pattern

const (
	Shallow_Copy = iota
	Deep_Copy
)

type PrototypeObj struct {
	PrototypeObjAttr string
}

type Prototype struct {
	PrototypeAttr string
	PrototypeObj  *PrototypeObj
}

func (p *Prototype) Clone(copyMode int) *Prototype {
	switch copyMode {
	case Deep_Copy:
		obj := *p
		obj.PrototypeObj = &PrototypeObj{PrototypeObjAttr: p.PrototypeAttr}
		return &obj
	case Shallow_Copy:
		obj := *p
		return &obj
	default:
		panic("unsupport copy mode")
	}
}
