package generatorvisitor

import (
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type advancedStateVisitor struct {
	size            int
	sides           map[parameters.Side]cube.Side
	sideConstructor func(string) parameters.Side
}

func newAdvancedStateVisitor(size int, sideConstructor func(string) parameters.Side) *advancedStateVisitor {
	return &advancedStateVisitor{size: size, sides: make(map[parameters.Side]cube.Side, 6), sideConstructor: sideConstructor}
}

func (v *advancedStateVisitor) emptySide() cube.Side {
	side := make(cube.Side, v.size)
	for i := range side {
		side[i] = make([]color.Color, v.size)
		for j := range side[i] {
			side[i][j] = color.Color("-")
		}
	}
	return side
}

func (v *advancedStateVisitor) visitAdvancedState(ctx *cp.AdvancedStateContext) models.Puzzle {
	sideToColor := map[parameters.Side]string{
		v.sideConstructor("Front"): "b",
		v.sideConstructor("Back"):  "g",
		v.sideConstructor("Up"):    "w",
		v.sideConstructor("Down"):  "y",
		v.sideConstructor("Left"):  "r",
		v.sideConstructor("Right"): "o",
	}
	for side, c := range sideToColor {
		newSide := v.emptySide()
		if v.size%2 == 1 {
			middle := v.size / 2
			newSide[middle][middle] = color.Color(c)
		}
		v.sides[side] = newSide
	}
	v.visitCorners(ctx.Corners().(*cp.CornersContext))
	if v.size > 2 {
		v.visitEdges(ctx.Edges().(*cp.EdgesContext))
	}
	return cube.NewWithSides(v.sides, v.size)
}

func (v *advancedStateVisitor) visitCornerLayer(ctx *cp.CornerLayerContext) {
	max := v.size - 1
	layerToSideCoords := map[string][][]parameters.Coord{
		"up": {
			{parameters.NewCoord(v.sideConstructor("Up"), 0, 0), parameters.NewCoord(v.sideConstructor("Left"), 0, 0), parameters.NewCoord(v.sideConstructor("Back"), 0, max)},
			{parameters.NewCoord(v.sideConstructor("Up"), 0, max), parameters.NewCoord(v.sideConstructor("Back"), 0, 0), parameters.NewCoord(v.sideConstructor("Right"), 0, max)},
			{parameters.NewCoord(v.sideConstructor("Up"), max, max), parameters.NewCoord(v.sideConstructor("Right"), 0, 0), parameters.NewCoord(v.sideConstructor("Front"), 0, max)},
			{parameters.NewCoord(v.sideConstructor("Up"), max, 0), parameters.NewCoord(v.sideConstructor("Front"), 0, 0), parameters.NewCoord(v.sideConstructor("Left"), 0, max)},
		},
		"down": {
			{parameters.NewCoord(v.sideConstructor("Down"), 0, 0), parameters.NewCoord(v.sideConstructor("Left"), max, max), parameters.NewCoord(v.sideConstructor("Front"), max, 0)},
			{parameters.NewCoord(v.sideConstructor("Down"), 0, max), parameters.NewCoord(v.sideConstructor("Front"), max, max), parameters.NewCoord(v.sideConstructor("Right"), max, 0)},
			{parameters.NewCoord(v.sideConstructor("Down"), max, max), parameters.NewCoord(v.sideConstructor("Right"), max, max), parameters.NewCoord(v.sideConstructor("Back"), max, 0)},
			{parameters.NewCoord(v.sideConstructor("Down"), max, 0), parameters.NewCoord(v.sideConstructor("Back"), max, max), parameters.NewCoord(v.sideConstructor("Left"), max, 0)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, corner := range ctx.AllCorner() {
		colors := strings.Split(corner.(*cp.CornerContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[sc.Side][sc.Row][sc.Col] = color.Color(c)
		}
	}
}

func (v *advancedStateVisitor) visitCorners(ctx *cp.CornersContext) {
	for _, cornerLayer := range ctx.AllCornerLayer() {
		v.visitCornerLayer(cornerLayer.(*cp.CornerLayerContext))
	}
}

func (v *advancedStateVisitor) visitEdgeLayer(ctx *cp.EdgeLayerContext) {
	layerToSideCoords := map[string][][]parameters.Coord{
		"up": {
			{parameters.NewCoord(v.sideConstructor("Up"), 0, 1), parameters.NewCoord(v.sideConstructor("Back"), 0, 1)},
			{parameters.NewCoord(v.sideConstructor("Up"), 1, 2), parameters.NewCoord(v.sideConstructor("Right"), 0, 1)},
			{parameters.NewCoord(v.sideConstructor("Up"), 2, 1), parameters.NewCoord(v.sideConstructor("Front"), 0, 1)},
			{parameters.NewCoord(v.sideConstructor("Up"), 1, 0), parameters.NewCoord(v.sideConstructor("Left"), 0, 1)},
		},
		"middle": {
			{parameters.NewCoord(v.sideConstructor("Left"), 1, 0), parameters.NewCoord(v.sideConstructor("Back"), 1, 2)},
			{parameters.NewCoord(v.sideConstructor("Back"), 1, 0), parameters.NewCoord(v.sideConstructor("Right"), 1, 2)},
			{parameters.NewCoord(v.sideConstructor("Right"), 1, 0), parameters.NewCoord(v.sideConstructor("Front"), 1, 2)},
			{parameters.NewCoord(v.sideConstructor("Front"), 1, 0), parameters.NewCoord(v.sideConstructor("Left"), 1, 2)},
		},
		"down": {
			{parameters.NewCoord(v.sideConstructor("Down"), 0, 1), parameters.NewCoord(v.sideConstructor("Front"), 2, 1)},
			{parameters.NewCoord(v.sideConstructor("Down"), 1, 2), parameters.NewCoord(v.sideConstructor("Right"), 2, 1)},
			{parameters.NewCoord(v.sideConstructor("Down"), 2, 1), parameters.NewCoord(v.sideConstructor("Back"), 2, 1)},
			{parameters.NewCoord(v.sideConstructor("Down"), 1, 0), parameters.NewCoord(v.sideConstructor("Left"), 2, 1)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, edge := range ctx.AllEdge() {
		colors := strings.Split(edge.(*cp.EdgeContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[sc.Side][sc.Row][sc.Col] = color.Color(c)
		}
	}
}

func (v *advancedStateVisitor) visitEdges(ctx *cp.EdgesContext) {
	for _, edgeLayer := range ctx.AllEdgeLayer() {
		v.visitEdgeLayer(edgeLayer.(*cp.EdgeLayerContext))
	}
}
