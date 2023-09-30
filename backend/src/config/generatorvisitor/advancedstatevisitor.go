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
	sides map[parameters.Side]cube.Side
}

func newAdvancedStateVisitor() *advancedStateVisitor {
	return &advancedStateVisitor{sides: make(map[parameters.Side]cube.Side, 6)}
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
	sideToColor := map[parameters.Side]string{
		"Front": "b",
		"Back":  "g",
		"Up":    "w",
		"Down":  "y",
		"Left":  "r",
		"Right": "o",
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
			{newSideCoord("Up", 0, 0), newSideCoord("Left", 0, 0), newSideCoord("Back", 0, 2)},
			{newSideCoord("Up", 0, 2), newSideCoord("Back", 0, 0), newSideCoord("Right", 0, 2)},
			{newSideCoord("Up", 2, 2), newSideCoord("Right", 0, 0), newSideCoord("Front", 0, 2)},
			{newSideCoord("Up", 2, 0), newSideCoord("Front", 0, 0), newSideCoord("Left", 0, 2)},
		},
		"down": {
			{newSideCoord("Down", 0, 0), newSideCoord("Left", 2, 2), newSideCoord("Front", 2, 0)},
			{newSideCoord("Down", 0, 2), newSideCoord("Front", 2, 2), newSideCoord("Right", 2, 0)},
			{newSideCoord("Down", 2, 2), newSideCoord("Right", 2, 2), newSideCoord("Back", 2, 0)},
			{newSideCoord("Down", 2, 0), newSideCoord("Back", 2, 2), newSideCoord("Left", 2, 0)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, corner := range ctx.AllCorner() {
		colors := strings.Split(corner.(*cp.CornerContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[parameters.Side(sc.side.String())][sc.Row][sc.Col] = color.Color(c)
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
			{newSideCoord("Up", 0, 1), newSideCoord("Back", 0, 1)},
			{newSideCoord("Up", 1, 2), newSideCoord("Right", 0, 1)},
			{newSideCoord("Up", 2, 1), newSideCoord("Front", 0, 1)},
			{newSideCoord("Up", 1, 0), newSideCoord("Left", 0, 1)},
		},
		"middle": {
			{newSideCoord("Left", 1, 0), newSideCoord("Back", 1, 2)},
			{newSideCoord("Back", 1, 0), newSideCoord("Right", 1, 2)},
			{newSideCoord("Right", 1, 0), newSideCoord("Front", 1, 2)},
			{newSideCoord("Front", 1, 0), newSideCoord("Left", 1, 2)},
		},
		"down": {
			{newSideCoord("Down", 0, 1), newSideCoord("Front", 2, 1)},
			{newSideCoord("Down", 1, 2), newSideCoord("Right", 2, 1)},
			{newSideCoord("Down", 2, 1), newSideCoord("Back", 2, 1)},
			{newSideCoord("Down", 1, 0), newSideCoord("Left", 2, 1)},
		},
	}
	layer := ctx.LayerDef().GetText()
	for i, edge := range ctx.AllEdge() {
		colors := strings.Split(edge.(*cp.EdgeContext).GetText(), "")
		for j, c := range colors {
			sc := layerToSideCoords[layer][i][j]
			v.sides[parameters.Side(sc.side.String())][sc.Row][sc.Col] = color.Color(c)
		}
	}
}

func (v *advancedStateVisitor) visitEdges(ctx *cp.EdgesContext) {
	for _, edgeLayer := range ctx.AllEdgeLayer() {
		v.visitEdgeLayer(edgeLayer.(*cp.EdgeLayerContext))
	}
}
