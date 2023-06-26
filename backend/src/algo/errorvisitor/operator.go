package errorvisitor

type operator interface {
	toUnary() *unaryOperator
	toBinary() *binaryOperator
}
