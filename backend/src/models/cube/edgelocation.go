package cube

type edgeLocation [2]CubeSide

func newEdgeLocation(s1, s2 CubeSide) edgeLocation {
	return edgeLocation{s1, s2}
}

func (e edgeLocation) distance(other edgeLocation) int {
	if e.same(other) {
		return 0
	}
	if !e.hasSameSide(other) {
		middle := e.getMiddleLocation(other)
		e[0] = middle[0]
		e[1] = middle[1]
	}
	diffSides := e.getDifferentSides(other)
	if diffSides[0].isOpposite(diffSides[1]) {
		return 2
	}
	return 1
}

func (e edgeLocation) getDifferentSides(other edgeLocation) [2]CubeSide {
	switch {
	case e[0] == other[0]:
		return [2]CubeSide{e[1], other[1]}
	case e[0] == other[1]:
		return [2]CubeSide{e[1], other[0]}
	case e[1] == other[0]:
		return [2]CubeSide{e[0], other[1]}
	case e[1] == other[1]:
		return [2]CubeSide{e[0], other[0]}
	}
	panic("There is more than two different sides")
}

func (e edgeLocation) getSameSide(other edgeLocation) CubeSide {
	switch {
	case e[0] == other[0]:
		return e[0]
	case e[0] == other[1]:
		return e[0]
	case e[1] == other[0]:
		return e[1]
	case e[1] == other[1]:
		return e[1]
	}
	panic("There is no same side")
}

func (e edgeLocation) getMiddleLocation(other edgeLocation) edgeLocation {
	newLocation := newEdgeLocation(e[0], e[1])
	newLocation[0] = newLocation[0].getOpposite()
	if newLocation.hasSameSide(other) {
		return newLocation
	}
	newLocation[0] = newLocation[0].getOpposite()
	newLocation[1] = newLocation[1].getOpposite()
	if newLocation.hasSameSide(other) {
		return newLocation
	}
	panic("There is no middle location")

}

func (e edgeLocation) hasSameSide(other edgeLocation) bool {
	return e[0] == other[0] || e[0] == other[1] || e[1] == other[0] || e[1] == other[1]
}

func (e edgeLocation) flipCost(goal edgeLocation) int {
	if e[0] == goal[0] && e[1] == goal[1] {
		return 0
	}
	return 3
}

func (e edgeLocation) same(other edgeLocation) bool {
	return (e[0] == other[0] && e[1] == other[1]) || (e[0] == other[1] && e[1] == other[0])
}
