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
	var upColor, downColor color.Color
	if v.cube.size%2 == 1 {
		middle := v.cube.size / 2
		upColor = v.cube.sides[Up][middle][middle]
		downColor = v.cube.sides[Down][middle][middle]
	} else {
		upColor = color.Color("w")
		downColor = color.Color("y")
	}
	isGoodColor := func(c color.Color) bool {
		return c == upColor || c == downColor
	}
	corners := v.cube.getCornerPieces()
	sum := 0
	for _, corner := range corners {
		sum += corner.getTwist(isGoodColor)
	}
	if sum%3 != 0 {
		return []string{fmt.Sprintf("There are some corner twists (%d)", sum%3)}
	}
	return nil
}

func (v *validator) checkCornerCycles() int {
	cornerPiecesArray := v.cube.getCornerPieces()
	cornerPieces := make(map[string]cornerPiece, len(cornerPiecesArray))
	for _, piece := range cornerPiecesArray {
		cornerPieces[piece.location.getHash()] = piece
	}
	numberOfCycles := 0
	var startLocation, currentLocation string
	for len(cornerPieces) > 0 {
		for location := range cornerPieces {
			startLocation = location
			currentLocation = location
			break
		}
		for {
			nextLocation := cornerPieces[currentLocation].goalLocation.getHash()
			delete(cornerPieces, currentLocation)
			currentLocation = nextLocation
			if currentLocation == startLocation {
				break
			}
		}
		numberOfCycles++
	}
	return numberOfCycles
}

func (v *validator) checkEdgeCycles() int {
	edgePiecesArray := v.cube.getEdgePieces()
	edgePieces := make(map[string]edgePiece, len(edgePiecesArray))
	for _, piece := range edgePiecesArray {
		edgePieces[piece.location.getHash()] = piece
	}
	numberOfCycles := 0
	var startLocation, currentLocation string
	for len(edgePieces) > 0 {
		for location := range edgePieces {
			startLocation = location
			currentLocation = location
			break
		}
		for {
			nextLocation := edgePieces[currentLocation].goalLocation.getHash()
			delete(edgePieces, currentLocation)
			currentLocation = nextLocation
			if currentLocation == startLocation {
				break
			}
		}
		numberOfCycles++
	}
	return numberOfCycles
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
	cornerCycles := v.checkCornerCycles()
	edgeCycles := v.checkEdgeCycles()
	if (cornerCycles+edgeCycles)%2 != 0 {
		return []string{fmt.Sprintf("There are some impossible permutations")}
	}
	return nil
}

func (v *validator) Validate() []string {
	errors := v.checkCornerTwists()
	if v.cube.size%2 == 1 {
		errors = append(errors, v.checkEdgeFlips()...)
		errors = append(errors, v.checkPermutations()...)
	}
	return errors
}
