package errorvisitor

import (
	"fmt"
	"strings"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type advancedStateVisitor struct {
	baseVisitor
	size int
}

func newAdvancedStateVisitor(fileName string, errorHandler *eh.Errorhandler, size int) *advancedStateVisitor {
	return &advancedStateVisitor{baseVisitor: *newBaseVisitor(fileName, errorHandler), size: size}
}

func (v *advancedStateVisitor) checkColors(ctx *cp.AdvancedStateContext) {
	if !v.valid {
		return
	}
	for _, layer := range ctx.Corners().(*cp.CornersContext).AllCornerLayer() {
		for _, corner := range layer.(*cp.CornerLayerContext).AllCorner() {
			colors := strings.Split(corner.GetText(), "")
			for _, color := range colors {
				v.colors[color]++
			}
		}
	}
	if v.size > 2 {
		for _, layer := range ctx.Edges().(*cp.EdgesContext).AllEdgeLayer() {
			for _, edge := range layer.(*cp.EdgeLayerContext).AllEdge() {
				colors := strings.Split(edge.GetText(), "")
				for _, color := range colors {
					v.colors[color]++
				}

			}
		}
	}
	stateCtx := ctx.GetParent().(*cp.StateContext)
	stateDefCtx := stateCtx.GetParent().(*cp.StateDefContext)
	expected := v.size * v.size
	if v.size%2 == 1 {
		expected--
	}
	checkColors(stateDefCtx, v, expected)
}

func (v *advancedStateVisitor) visitAdvancedState(ctx *cp.AdvancedStateContext) {
	if v.size > 3 {
		v.valid = false
		v.Eh().AddError(ctx, "Advanced description only valid for 2x2x2 and 3x3x3", v.FileName())
		return
	}
	v.visitCorners(ctx.Corners().(*cp.CornersContext))
	if v.size > 2 {
		v.visitEdges(ctx.Edges().(*cp.EdgesContext))
	}
	v.checkColors(ctx)
}

func (v *advancedStateVisitor) visitCornerLayer(ctx *cp.CornerLayerContext) {
	corners := ctx.AllCorner()
	if len(corners) < 4 {
		v.finished = false
		warningMsg := fmt.Sprintf("layer should have 4 corners, has %d", len(corners))
		v.Eh().AddWarning(ctx, warningMsg, v.FileName())
	}
	if len(corners) > 4 {
		v.valid = false
		errorMsg := fmt.Sprintf("invalid number of corners, wanted 4, has %d", len(corners))
		v.Eh().AddError(ctx, errorMsg, v.FileName())
	}
}

func (v *advancedStateVisitor) visitCorners(ctx *cp.CornersContext) {
	layers := ctx.AllCornerLayer()
	layerDefs := make(map[string][]*cp.CornerLayerContext)
	for _, layer := range layers {
		layerDef := layer.(*cp.CornerLayerContext).LayerDef().GetText()
		layerDefs[layerDef] = append(layerDefs[layerDef], layer.(*cp.CornerLayerContext))
	}
	for layerDef, layerCtxs := range layerDefs {
		if len(layerCtxs) > 1 {
			v.valid = false
			errorMsg := layerDef + " layer is defined multiple times"
			for _, layerCtx := range layerCtxs {
				v.Eh().AddError(layerCtx, errorMsg, v.FileName())
			}
		}
	}
	if layerCtxs, ok := layerDefs["Middle"]; ok {
		for _, layerCtx := range layerCtxs {
			v.Eh().AddError(layerCtx, "middle layer is invalid for corner definitions", v.FileName())
		}
	}
	necessaryLayers := []string{"Up", "Down"}
	for _, necessaryLayer := range necessaryLayers {
		if _, ok := layerDefs[necessaryLayer]; !ok {
			v.finished = false
			v.Eh().AddWarning(ctx, necessaryLayer+" layer is missing", v.FileName())
		}
	}
	for _, layer := range ctx.AllCornerLayer() {
		v.visitCornerLayer(layer.(*cp.CornerLayerContext))
	}
}

func (v *advancedStateVisitor) visitEdgeLayer(ctx *cp.EdgeLayerContext) {
	edges := ctx.AllEdge()
	if len(edges) < 4 {
		v.finished = false
		warningMsg := fmt.Sprintf("layer should have 4 edges, has %d", len(edges))
		v.Eh().AddWarning(ctx, warningMsg, v.FileName())
	}
	if len(edges) > 4 {
		v.valid = false
		errorMsg := fmt.Sprintf("invalid number of edges, wanted 4, has %d", len(edges))
		v.Eh().AddError(ctx, errorMsg, v.FileName())
	}
}

func (v *advancedStateVisitor) visitEdges(ctx *cp.EdgesContext) {
	layers := ctx.AllEdgeLayer()
	layerDefs := make(map[string][]*cp.EdgeLayerContext)
	for _, layer := range layers {
		layerDef := layer.(*cp.EdgeLayerContext).LayerDef().GetText()
		layerDefs[layerDef] = append(layerDefs[layerDef], layer.(*cp.EdgeLayerContext))
	}
	for layerDef, layerCtxs := range layerDefs {
		if len(layerCtxs) > 1 {
			v.valid = false
			errorMsg := layerDef + " layer is defined multiple times"
			for _, layerCtx := range layerCtxs {
				v.Eh().AddError(layerCtx, errorMsg, v.FileName())
			}
		}
	}
	necessaryLayers := []string{"Up", "Middle", "Down"}
	for _, necessaryLayer := range necessaryLayers {
		if _, ok := layerDefs[necessaryLayer]; !ok {
			v.finished = false
			v.Eh().AddWarning(ctx, necessaryLayer+" layer is missing", v.FileName())
		}
	}
	for _, layer := range ctx.AllEdgeLayer() {
		v.visitEdgeLayer(layer.(*cp.EdgeLayerContext))
	}
}
