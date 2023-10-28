package errorvisitor

import "fmt"

type iType interface {
	canUseAs(iType) bool
	fmt.Stringer
}

type basicType struct {
	name    string
	parents []iType
}

func newBasicType(name string, parents ...iType) basicType {
	return basicType{name: name, parents: parents}
}

func (t basicType) addParents(parents ...iType) {
	t.parents = append(t.parents, parents...)
}

func (t basicType) canUseAs(other iType) bool {
	otherBasicType, ok := other.(basicType)
	if !ok {
		return false
	}
	if t.name == otherBasicType.name {
		return true
	}
	for _, parent := range t.parents {
		if parent.canUseAs(other) {
			return true
		}
	}
	return false
}

func (t basicType) String() string {
	return t.name
}

type listType struct {
	elemType iType
}

func newListType(elemType iType) listType {
	return listType{elemType: elemType}
}

func (t listType) canUseAs(other iType) bool {
	otherListType, ok := other.(listType)
	if !ok {
		return false
	}
	return t.elemType.canUseAs(otherListType.elemType)
}

func (t listType) String() string {
	return fmt.Sprintf("[%s]", t.elemType)
}
