package generatorvisitor

import (
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
)

type advancedStateVisitor struct {
	sides map[cube.CubeSide]cube.Side
}

func newAdvancedStateVisitor() *advancedStateVisitor {
	return &advancedStateVisitor{sides: make(map[cube.CubeSide]cube.Side, 6)}
}

func (v *advancedStateVisitor) emptySide() cube.Side {
	side := make(cube.Side, 3)
	for i := range side {
		side[i] = make([]color.Color, 3)
		for j := range side[i] {
			side[i][j] = color.Color("-")
		}
	}
	return side
}

func (v *advancedStateVisitor) visitAdvancedState(ctx *cp.AdvancedStateContext) models.Puzzle {
	sideToColor := map[cube.CubeSide]string{
		cube.Front: "b",
		cube.Back:  "g",
		cube.Up:    "w",
		cube.Down:  "y",
		cube.Left:  "r",
		cube.Right: "o",
	}
	for side, c := range sideToColor {
		newSide := v.emptySide()
		newSide[1][1] = color.Color(c)
		v.sides[side] = newSide
	}
	v.visitCorners(ctx.Corners().(*cp.CornersContext))
	v.visitEdges(ctx.Edges().(*cp.EdgesContext))
	return cube.NewWithSides(v.sides)
}

func (v *advancedStateVisitor) visitCornerLayer(ctx *cp.CornerLayerContext) {
	layerToSideCoords := map[string][][]sideCoord{
		"up": {
			{newSideCoord(cube.Up, 0, 0), newSideCoord(cube.Left, 0, 0), newSideCoord(cube.Back, 0, 2)},
			{newSideCoord(cube.Up, 0, 2), newSideCoord(cube.Back, 0, 0), newSideCoord(cube.Right, 0, 2)},
			{newSideCoord(cube.Up, 2, 2), newSideCoord(cube.Right, 0, 0), newSideCoord(cube.Front, 0, 2)},
			{newSideCoord(cube.Up, 2, 0), newSideCoord(cube.Front, 0, 0), newSideCoord(cube.Left, 0, 2)},
		},
		"down": {
			{newSideCoord(cube.Down, 0, 0), newSideCoord(cube.Left, 2, 2), newSideCoord(cube.Front, 2, 0)},
			{newSideCoord(cube.Down, 0, 2), newSideCoord(cube.Front, 2, 2), newSideCoord(cube.Right, 2, 0)},
			{newSideCoord(cube.Down, 2, 2), newSideCoord(cube.Right, 2, 2), newSideCoord(cube.Back, 2, 0)},
			{newSideCoord(cube.Down, 2, 0), newSideCoord(cube.Back, 2, 2), newSideCoord(cube.Left, 2, 0)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, corner := range ctx.AllCorner() {
		colors := strings.Split(corner.(*cp.CornerContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[sc.side][sc.Row][sc.Col] = color.Color(c)
		}
	}
}

func (v *advancedStateVisitor) visitCorners(ctx *cp.CornersContext) {
	for _, cornerLayer := range ctx.AllCornerLayer() {
		v.visitCornerLayer(cornerLayer.(*cp.CornerLayerContext))
	}
}

func (v *advancedStateVisitor) visitEdgeLayer(ctx *cp.EdgeLayerContext) {
	layerToSideCoords := map[string][][]sideCoord{
		"up": {
			{newSideCoord(cube.Up, 0, 1), newSideCoord(cube.Back, 0, 1)},
			{newSideCoord(cube.Up, 1, 2), newSideCoord(cube.Right, 0, 1)},
			{newSideCoord(cube.Up, 2, 1), newSideCoord(cube.Front, 0, 1)},
			{newSideCoord(cube.Up, 1, 0), newSideCoord(cube.Left, 0, 1)},
		},
		"middle": {
			{newSideCoord(cube.Left, 1, 0), newSideCoord(cube.Back, 1, 2)},
			{newSideCoord(cube.Back, 1, 0), newSideCoord(cube.Right, 1, 2)},
			{newSideCoord(cube.Right, 1, 0), newSideCoord(cube.Front, 1, 2)},
			{newSideCoord(cube.Front, 1, 0), newSideCoord(cube.Left, 1, 2)},
		},
		"down": {
			{newSideCoord(cube.Down, 0, 1), newSideCoord(cube.Front, 2, 1)},
			{newSideCoord(cube.Down, 1, 2), newSideCoord(cube.Right, 2, 1)},
			{newSideCoord(cube.Down, 2, 1), newSideCoord(cube.Back, 2, 1)},
			{newSideCoord(cube.Down, 1, 0), newSideCoord(cube.Left, 2, 1)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, edge := range ctx.AllEdge() {
		colors := strings.Split(edge.(*cp.EdgeContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[sc.side][sc.Row][sc.Col] = color.Color(c)
		}
	}
}

func (v *advancedStateVisitor) visitEdges(ctx *cp.EdgesContext) {
	for _, edgeLayer := range ctx.AllEdgeLayer() {
		v.visitEdgeLayer(edgeLayer.(*cp.EdgeLayerContext))
	}
}
