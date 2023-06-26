package errorvisitor

import "fmt"

type iType interface {
	canUseAs(iType) bool
	fmt.Stringer
}

type basicType int

const (
	node basicType = iota
	piece
	position
	coord
)

func (t basicType) canUseAs(other iType) bool {
	otherBasicType, ok := other.(basicType)
	if !ok {
		return false
	}
	if t == otherBasicType {
		return true
	}
	if t == node && (otherBasicType == piece || otherBasicType == position) {
		return true
	}
	return false
}

func (t basicType) String() string {
	switch t {
	case node:
		return "node"
	case piece:
		return "piece"
	case position:
		return "position"
	case coord:
		return "coord"
	default:
		return "unknown"
	}
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
	return fmt.Sprintf("list of %s", t.elemType)
}
