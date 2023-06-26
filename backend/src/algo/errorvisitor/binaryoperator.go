package errorvisitor

import "fmt"

type binaryOperator struct {
	name       string
	leftParam  iType
	rightParam iType
}

func newBinaryOperator(name string, leftParam, rightParam iType) binaryOperator {
	return binaryOperator{name: name, leftParam: leftParam, rightParam: rightParam}
}

func (b *binaryOperator) acceptLeftParam(param iType) error {
	if param.canUseAs(b.leftParam) {
		return nil
	}
	return fmt.Errorf("%s operator can only be used with %s as left parameter, got: %s", b.name, b.leftParam, param)
}

func (b *binaryOperator) acceptRightParam(param iType) error {
	if param.canUseAs(b.rightParam) {
		return nil
	}
	return fmt.Errorf("%s operator can only be used with %s as right parameter, got: %s", b.name, b.rightParam, param)
}

func (b *binaryOperator) String() string {
	return b.name
}

func (b *binaryOperator) toBinary() *binaryOperator {
	return b
}

func (b *binaryOperator) toUnary() *unaryOperator {
	return nil
}
