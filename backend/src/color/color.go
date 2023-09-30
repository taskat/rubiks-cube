package color

type Color string

func (c Color) Equals(other Color) bool {
	return c[0] == other[0]
}

func (c Color) String() string {
	return string(c)
}
