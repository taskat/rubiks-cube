package errorvisitor

import "fmt"

type unaryOperator struct {
	name  string
	param iType
}

func newUnaryOperator(name string, param iType) unaryOperator {
	return unaryOperator{name: name, param: param}
}

func (o *unaryOperator) acceptParam(param iType) error {
	if param.canUseAs(o.param) {
		return nil
	}
	return fmt.Errorf("%s operator can only be used with %s as parameter, got: %s", o.name, o.param, param)
}

func (o *unaryOperator) toBinary() *binaryOperator {
	return nil
}

func (o *unaryOperator) toUnary() *unaryOperator {
	return o
}
