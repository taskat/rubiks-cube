package cube

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/models/parameters"
)

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

func (mg *moveGenerator) generateNames(basics []string) []string {
	half := mg.size / 2
	names := make([]string, 0, mg.size)
	names = append(names, basics[0])
	for i := 2; i <= half; i++ {
		newName := fmt.Sprintf("%d%s", i, basics[0])
		names = append(names, newName)
	}
	if mg.size%2 == 1 {
		names = append(names, basics[1])
	}
	for i := half; i >= 2; i-- {
		newName := fmt.Sprintf("%d%s", i, basics[2])
		names = append(names, newName)
	}
	names = append(names, basics[2])
	return names
}

func (mg *moveGenerator) generateXMoves() {
	getSlices := func(col int) [4][]parameters.Coord {
		slices := [4][]parameters.Coord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], parameters.NewCoord(cubeSide("Front"), i, col))
			slices[1] = append(slices[1], parameters.NewCoord(cubeSide("Up"), i, col))
			slices[2] = append(slices[2], parameters.NewCoord(cubeSide("Back"), mg.size-i-1, mg.size-col-1))
			slices[3] = append(slices[3], parameters.NewCoord(cubeSide("Down"), i, col))
		}
		return slices
	}
	names := mg.generateNames([]string{"L", "M", "R"})
	for i := 0; i < mg.size; i++ {
		slices := getSlices(i)
		cycles := mg.getSliceCycles(slices)
		if i < mg.size/2 || (mg.size%2 == 1 && i == mg.size/2) {
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
	getSlices := func(row int) [4][]parameters.Coord {
		slices := [4][]parameters.Coord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], parameters.NewCoord(cubeSide("Front"), row, i))
			slices[1] = append(slices[1], parameters.NewCoord(cubeSide("Left"), row, i))
			slices[2] = append(slices[2], parameters.NewCoord(cubeSide("Back"), row, i))
			slices[3] = append(slices[3], parameters.NewCoord(cubeSide("Right"), row, i))
		}
		return slices
	}
	names := mg.generateNames([]string{"U", "E", "D"})
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
	getSlices := func(diffFromFront int) [4][]parameters.Coord {
		slices := [4][]parameters.Coord{}
		for i := 0; i < mg.size; i++ {
			slices[0] = append(slices[0], parameters.NewCoord(cubeSide("Up"), mg.size-diffFromFront-1, i))
			slices[1] = append(slices[1], parameters.NewCoord(cubeSide("Right"), i, diffFromFront))
			slices[2] = append(slices[2], parameters.NewCoord(cubeSide("Down"), diffFromFront, mg.size-i-1))
			slices[3] = append(slices[3], parameters.NewCoord(cubeSide("Left"), mg.size-i-1, mg.size-diffFromFront-1))
		}
		return slices
	}
	names := mg.generateNames([]string{"F", "S", "B"})
	for i := 0; i < mg.size; i++ {
		slices := getSlices(i)
		cycles := mg.getSliceCycles(slices)
		if i == 0 {
			cycles = append(cycles, mg.getSideCycles(Front)...)
		}
		if i > mg.size/2 || (mg.size%2 == 0 && i == mg.size/2) {
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

func (mg *moveGenerator) getSideCycles(side cubeSide) []cycle {
	max := mg.size - 1
	cycles := make([]cycle, 0, 3)
	cycles = append(cycles, cycle{
		parameters.NewCoord(side, 0, 0), parameters.NewCoord(side, 0, max),
		parameters.NewCoord(side, max, max), parameters.NewCoord(side, max, 0),
	})
	if mg.size == 3 {
		cycles = append(cycles, cycle{
			parameters.NewCoord(side, 0, 1), parameters.NewCoord(side, 1, 2),
			parameters.NewCoord(side, 2, 1), parameters.NewCoord(side, 1, 0),
		})
	}
	if mg.size%2 == 1 && mg.size > 2 {
		cycles = append(cycles, cycle{parameters.NewCoord(side, 1, 1)})
	}
	return cycles
}

func (mg *moveGenerator) getSliceCycles(slices [4][]parameters.Coord) []cycle {
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
