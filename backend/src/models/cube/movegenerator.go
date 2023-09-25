package cube

type moveGenerator struct {
	size  int
	moves map[string]move
}

func newMoveGenerator(size int) moveGenerator {
	return moveGenerator{size: size, moves: make(map[string]move)}
}

func (mg *moveGenerator) addMove(m move, name string) {
	mg.moves[name] = m
	doubleMove := move{cycles: m.cycles, steps: 2}
	mg.moves[name+"2"] = doubleMove
	inverseMove := move{cycles: mg.inverseCycles(m.cycles), steps: 1}
	mg.moves[name+"'"] = inverseMove
}

func (mg *moveGenerator) generateAllMoves() map[string]move {
	mg.generateXMoves()
	mg.generateYMoves()
	mg.generateZMoves()
	return mg.moves
}

func (mg *moveGenerator) generateXMoves() {
	getSlices := func(col int) [4][]sideCoord {
		slices := [4][]sideCoord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], newSideCoord(Front, i, col))
			slices[1] = append(slices[1], newSideCoord(Up, i, col))
			slices[2] = append(slices[2], newSideCoord(Back, mg.size-i-1, mg.size-col-1))
			slices[3] = append(slices[3], newSideCoord(Down, i, col))
		}
		return slices
	}
	names := []string{"L", "M", "R"}
	for i := 0; i < mg.size; i++ {
		slices := getSlices(i)
		cycles := mg.getSliceCycles(slices)
		if i <= mg.size/2 {
			cycles = mg.inverseCycles(cycles)
		}
		if i == 0 {
			cycles = append(cycles, mg.getSideCycles(Left)...)
		}
		if i == mg.size-1 {
			cycles = append(cycles, mg.getSideCycles(Right)...)
		}
		mg.addMove(move{cycles: cycles, steps: 1}, names[i])
	}
	fullCubeCycles := make([]cycle, 0)
	for i := 0; i < mg.size; i++ {
		fullCubeCycles = append(fullCubeCycles, mg.getSliceCycles(getSlices(i))...)
	}
	fullCubeCycles = append(fullCubeCycles, mg.inverseCycles(mg.getSideCycles(Left))...)
	fullCubeCycles = append(fullCubeCycles, mg.getSideCycles(Right)...)
	mg.addMove(move{cycles: fullCubeCycles, steps: 1}, "x")
}

func (mg *moveGenerator) generateYMoves() {
	getSlices := func(row int) [4][]sideCoord {
		slices := [4][]sideCoord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], newSideCoord(Front, row, i))
			slices[1] = append(slices[1], newSideCoord(Left, row, i))
			slices[2] = append(slices[2], newSideCoord(Back, row, i))
			slices[3] = append(slices[3], newSideCoord(Right, row, i))
		}
		return slices
	}
	names := []string{"U", "E", "D"}
	for i := 0; i < mg.size; i++ {
		slices := getSlices(i)
		cycles := mg.getSliceCycles(slices)
		if i == 0 {
			cycles = append(cycles, mg.getSideCycles(Up)...)
		}
		if i >= mg.size/2 {
			cycles = mg.inverseCycles(cycles)
		}
		if i == mg.size-1 {
			cycles = append(cycles, mg.getSideCycles(Down)...)
		}
		mg.addMove(move{cycles: cycles, steps: 1}, names[i])
	}
	fullCubeCycles := make([]cycle, 0)
	for i := 0; i < mg.size; i++ {
		fullCubeCycles = append(fullCubeCycles, mg.getSliceCycles(getSlices(i))...)
	}
	fullCubeCycles = append(fullCubeCycles, mg.getSideCycles(Up)...)
	fullCubeCycles = append(fullCubeCycles, mg.inverseCycles(mg.getSideCycles(Down))...)
	mg.addMove(move{cycles: fullCubeCycles, steps: 1}, "y")
}

func (mg *moveGenerator) generateZMoves() {
	getSlices := func(diffFromFront int) [4][]sideCoord {
		slices := [4][]sideCoord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], newSideCoord(Up, mg.size-diffFromFront-1, i))
			slices[1] = append(slices[1], newSideCoord(Right, i, diffFromFront))
			slices[2] = append(slices[2], newSideCoord(Down, diffFromFront, mg.size-i-1))
			slices[3] = append(slices[3], newSideCoord(Left, mg.size-i-1, mg.size-diffFromFront-1))
		}
		return slices
	}
	names := []string{"F", "S", "B"}
	for i := 0; i < mg.size; i++ {
		slices := getSlices(i)
		cycles := mg.getSliceCycles(slices)
		if i == 0 {
			cycles = append(cycles, mg.getSideCycles(Front)...)
		}
		if i > mg.size/2 {
			cycles = mg.inverseCycles(cycles)
		}
		if i == mg.size-1 {
			cycles = append(cycles, mg.getSideCycles(Back)...)
		}
		mg.addMove(move{cycles: cycles, steps: 1}, names[i])
	}
	fullCubeCycles := make([]cycle, 0)
	for i := 0; i < mg.size; i++ {
		fullCubeCycles = append(fullCubeCycles, mg.getSliceCycles(getSlices(i))...)
	}
	fullCubeCycles = append(fullCubeCycles, mg.getSideCycles(Front)...)
	fullCubeCycles = append(fullCubeCycles, mg.inverseCycles(mg.getSideCycles(Back))...)
	mg.addMove(move{cycles: fullCubeCycles, steps: 1}, "z")
}

func (mg *moveGenerator) getSideCycles(side CubeSide) []cycle {
	cycles := make([]cycle, 3)
	cycles[0] = cycle{
		newSideCoord(side, 0, 0), newSideCoord(side, 0, 2),
		newSideCoord(side, 2, 2), newSideCoord(side, 2, 0),
	}
	cycles[1] = cycle{
		newSideCoord(side, 0, 1), newSideCoord(side, 1, 2),
		newSideCoord(side, 2, 1), newSideCoord(side, 1, 0),
	}
	cycles[2] = cycle{newSideCoord(side, 1, 1)}
	return cycles
}

func (mg *moveGenerator) getSliceCycles(slices [4][]sideCoord) []cycle {
	cycles := make([]cycle, len(slices[0]))
	for i := range slices[0] {
		cycles[i] = cycle{
			slices[0][i], slices[1][i], slices[2][i], slices[3][i],
		}
	}
	return cycles
}

func (moveGenerator) inverseCycles(cycles []cycle) []cycle {
	inversedCycles := make([]cycle, len(cycles))
	for i, c := range cycles {
		inversedCycles[i] = c.inverse()
	}
	return inversedCycles
}
