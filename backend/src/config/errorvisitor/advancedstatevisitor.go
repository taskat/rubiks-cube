package errorvisitor

import (
	"fmt"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type advancedStateVisitor struct {
	fileName string
	finished bool
	valid    bool
	eh       eh.Errorhandler
}

func newAdvancedStateVisitor(fileName string, errorHandler eh.Errorhandler) *advancedStateVisitor {
	return &advancedStateVisitor{fileName: fileName, finished: true, valid: true, eh: errorHandler}
}

func (v *advancedStateVisitor) checkCorners(layers []cp.ICornerLayerContext) {
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
}

func (v *advancedStateVisitor) checkColors(ctx *cp.BeginnerStateContext) {
	if !v.valid {
		return
	}
	colors := make(map[string]int)
	for _, side := range ctx.AllSide() {
		for _, row := range side.(*cp.SideContext).AllRow() {
			for _, cell := range row.(*cp.RowContext).AllColor() {
				colors[cell.GetText()] += 1
			}
		}
	}
	stateCtx := ctx.GetParent().(*cp.StateContext)
	stateDefCtx := stateCtx.GetParent().(*cp.StateDefContext)
	for color, amount := range colors {
		if amount > 9 || (v.finished && amount < 9) {
			errorMsg := fmt.Sprintf("color %s is defined %d times, should be 9 times", color, amount)
			v.eh.AddWarning(stateDefCtx, errorMsg, v.fileName)
		}
	}
}

func (v *advancedStateVisitor) visitAdvancedState(ctx *cp.AdvancedStateContext) {
	v.checkCorners(ctx.Corners().(*cp.CornersContext).AllCornerLayer())
	v.visitCorners(ctx.Corners().(*cp.CornersContext))
	v.visitEdges(ctx.Edges().(*cp.EdgesContext))
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
	for _, layer := range ctx.AllCornerLayer() {
		v.visitCornerLayer(layer.(*cp.CornerLayerContext))
	}
}

func (v *advancedStateVisitor) visitEdges(ctx *cp.EdgesContext) {

}
