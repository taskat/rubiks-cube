package parameters

type Side string

func NewSide(s string) Side {
	return Side(s)
}

func (s Side) String() string {
	return string(s)
}
