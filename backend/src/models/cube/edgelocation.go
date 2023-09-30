package cube

import (
	"sort"
	"strings"

	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type edgeLocation struct {
	sides [2]cubeSide
	hash  string
}

func newEdgeLocation(s1, s2 parameters.Side) edgeLocation {
	return edgeLocation{sides: [2]cubeSide{cubeSide(s1), cubeSide(s2)}, hash: ""}
}

func newEdgeLocationFromCubeSides(s1, s2 cubeSide) edgeLocation {
	return edgeLocation{sides: [2]cubeSide{s1, s2}, hash: ""}
}

func (e edgeLocation) distance(other edgeLocation) int {
	if e.same(other) {
		return 0
	}
	if !e.hasSameSide(other) {
		e.sides = e.getMiddleLocation(other).sides
	}
	diffSides := e.getDifferentSides(other)
	if diffSides[0].isOpposite(diffSides[1]) {
		return 2
	}
	return 1
}

func (e edgeLocation) getDifferentSides(other edgeLocation) [2]cubeSide {
	switch {
	case e.sides[0] == other.sides[0]:
		return [2]cubeSide{e.sides[1], other.sides[1]}
	case e.sides[0] == other.sides[1]:
		return [2]cubeSide{e.sides[1], other.sides[0]}
	case e.sides[1] == other.sides[0]:
		return [2]cubeSide{e.sides[0], other.sides[1]}
	case e.sides[1] == other.sides[1]:
		return [2]cubeSide{e.sides[0], other.sides[0]}
	}
	panic("There is more than two different sides")
}

func (e edgeLocation) getSameSide(other edgeLocation) cubeSide {
	switch {
	case e.sides[0] == other.sides[0]:
		return e.sides[0]
	case e.sides[0] == other.sides[1]:
		return e.sides[0]
	case e.sides[1] == other.sides[0]:
		return e.sides[1]
	case e.sides[1] == other.sides[1]:
		return e.sides[1]
	}
	panic("There is no same side")
}

func (e edgeLocation) getMiddleLocation(other edgeLocation) edgeLocation {
	newLocation := newEdgeLocationFromCubeSides(e.sides[0], e.sides[1])
	newLocation.sides[0] = newLocation.sides[0].getOpposite()
	if newLocation.hasSameSide(other) {
		return newLocation
	}
	newLocation.sides[0] = newLocation.sides[0].getOpposite()
	newLocation.sides[1] = newLocation.sides[1].getOpposite()
	if newLocation.hasSameSide(other) {
		return newLocation
	}
	panic("There is no middle location")
}

func (e edgeLocation) getHash() string {
	if e.hash != "" {
		return e.hash
	}
	hashes := []string{e.sides[0].getHash(), e.sides[1].getHash()}
	sort.Strings(hashes)
	return strings.Join(hashes, "")
}

func (e edgeLocation) hasSameSide(other edgeLocation) bool {
	return e.sides[0] == other.sides[0] || e.sides[0] == other.sides[1] || e.sides[1] == other.sides[0] || e.sides[1] == other.sides[1]
}

func (e edgeLocation) flipCost(goal edgeLocation) int {
	if e.sides[0] == goal.sides[0] || e.sides[1] == goal.sides[1] {
		return 0
	}
	return 3
}

func (e edgeLocation) same(other edgeLocation) bool {
	return (e.sides[0] == other.sides[0] && e.sides[1] == other.sides[1]) || (e.sides[0] == other.sides[1] && e.sides[1] == other.sides[0])
}
