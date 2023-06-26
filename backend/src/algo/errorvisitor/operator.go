package errorvisitor

import "fmt"

type operator interface {
	toUnary() *unaryOperator
	toBinary() *binaryOperator
	fmt.Stringer
}
