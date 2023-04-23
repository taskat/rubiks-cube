package cube

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/color"
)

type validator struct {
	cube *Cube
}

func newValidator(c *Cube) *validator {
	return &validator{cube: c}
}

func (v *validator) checkCornerTwists() []string {
	upColor := v.cube.sides[Up][1][1]
	downColor := v.cube.sides[Down][1][1]
	isGoodColor := func(c color.Color) bool {
		return c == upColor || c == downColor
	}
	corners := v.cube.getCorners()
	sum := 0
	for sideCoord, color := range corners {
		if !isGoodColor(color) {
			continue
		}
		if sideCoord.side == Up || sideCoord.side == Down {
			continue
		}
		coordSum := sideCoord.Row + sideCoord.Col
		if coordSum == 2 {
			sum++
		} else {
			sum += 2
		}
	}
	if sum%3 != 0 {
		return []string{fmt.Sprintf("There are some corner twists (%d)", sum%3)}
	}
	return nil
}

func (v *validator) checkEdgeFlips() []string {
	edgePieces := v.cube.getEdgePieces()
	moves := 0
	for _, piece := range edgePieces {
		moves += piece.movesToGoal()
	}
	if moves%2 != 0 {
		return []string{fmt.Sprintf("There are some edge flips")}
	}
	return nil
}

func (v *validator) checkPermutations() []string {
	fmt.Println("checkPermutations")
	return nil
}

func (v *validator) Validate() []string {
	errors := v.checkCornerTwists()
	errors = append(errors, v.checkEdgeFlips()...)
	errors = append(errors, v.checkPermutations()...)
	return errors
}
