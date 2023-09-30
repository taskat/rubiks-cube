package cube

import "github.com/taskat/rubiks-cube/src/models/parameters"

type cycle []parameters.Coord

func (c cycle) inverse() cycle {
	inversed := make(cycle, len(c))
	for i, coord := range c {
		inversed[len(c)-i-1] = coord
	}
	return inversed
}
