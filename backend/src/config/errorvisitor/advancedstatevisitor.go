package errorvisitor

import (
	"fmt"
	"strings"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type advancedStateVisitor struct {
	baseVisitor
}

func newAdvancedStateVisitor(fileName string, errorHandler *eh.Errorhandler) *advancedStateVisitor {
	return &advancedStateVisitor{baseVisitor: *newBaseVisitor(fileName, errorHandler)}
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
	for _, layer := range ctx.Edges().(*cp.EdgesContext).AllEdgeLayer() {
		for _, edge := range layer.(*cp.EdgeLayerContext).AllEdge() {
			colors := strings.Split(edge.GetText(), "")
			for _, color := range colors {
				v.colors[color]++
			}

		}
	}
	stateCtx := ctx.GetParent().(*cp.StateContext)
	stateDefCtx := stateCtx.GetParent().(*cp.StateDefContext)
	checkColors(stateDefCtx, v, 8)
}

func (v *advancedStateVisitor) visitAdvancedState(ctx *cp.AdvancedStateContext) {
	v.visitCorners(ctx.Corners().(*cp.CornersContext))
	v.visitEdges(ctx.Edges().(*cp.EdgesContext))
	v.checkColors(ctx)
}

func (v *advancedStateVisitor) visitCornerLayer(ctx *cp.CornerLayerContext) {
	corners := ctx.AllCorner()
	if len(corners) < 4 {
		v.finished = false
		warningMsg := fmt.Sprintf("layer should have 4 corners, has %d", len(corners))
		v.eh.AddWarning(ctx, warningMsg, v.fileName)
	}
	if len(corners) > 4 {
		v.valid = false
		errorMsg := fmt.Sprintf("invalid number of corners, wanted 4, has %d", len(corners))
		v.eh.AddError(ctx, errorMsg, v.fileName)
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
				v.eh.AddError(layerCtx, errorMsg, v.fileName)
			}
		}
	}
	if layerCtxs, ok := layerDefs["middle"]; ok {
		for _, layerCtx := range layerCtxs {
			v.eh.AddError(layerCtx, "middle layer is invalid for corner definitions", v.fileName)
		}
	}
	necessaryLayers := []string{"up", "down"}
	for _, necessaryLayer := range necessaryLayers {
		if _, ok := layerDefs[necessaryLayer]; !ok {
			v.finished = false
			v.eh.AddWarning(ctx, necessaryLayer+" layer is missing", v.fileName)
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
		v.eh.AddWarning(ctx, warningMsg, v.fileName)
	}
	if len(edges) > 4 {
		v.valid = false
		errorMsg := fmt.Sprintf("invalid number of edges, wanted 4, has %d", len(edges))
		v.eh.AddError(ctx, errorMsg, v.fileName)
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
				v.eh.AddError(layerCtx, errorMsg, v.fileName)
			}
		}
	}
	necessaryLayers := []string{"up", "middle", "down"}
	for _, necessaryLayer := range necessaryLayers {
		if _, ok := layerDefs[necessaryLayer]; !ok {
			v.finished = false
			v.eh.AddWarning(ctx, necessaryLayer+" layer is missing", v.fileName)
		}
	}
	for _, layer := range ctx.AllEdgeLayer() {
		v.visitEdgeLayer(layer.(*cp.EdgeLayerContext))
	}
}
