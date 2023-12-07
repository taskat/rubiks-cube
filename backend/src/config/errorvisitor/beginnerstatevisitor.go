package errorvisitor

import (
	"fmt"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type beginnerStateVisitor struct {
	baseVisitor
	size int
}

func newBeginnerStateVisitor(filename string, errorHandler *eh.Errorhandler, size int) *beginnerStateVisitor {
	return &beginnerStateVisitor{baseVisitor: *newBaseVisitor(filename, errorHandler), size: size}
}

func (v *beginnerStateVisitor) checkSideNames(sides []cp.ISideContext) {
	sideNames := make(map[string][]*cp.SideContext)
	for _, side := range sides {
		sideName := side.(*cp.SideContext).SideDef().GetText()
		sideNames[sideName] = append(sideNames[sideName], side.(*cp.SideContext))
	}
	for sideName, sideCtxs := range sideNames {
		if len(sideCtxs) > 1 {
			v.valid = false
			errorMsg := "side name " + sideName + " is defined multiple times"
			for _, sideCtx := range sideCtxs {
				v.Eh().AddError(sideCtx, errorMsg, v.FileName())
			}
		}
	}
}

func (v *beginnerStateVisitor) checkColors(ctx *cp.BeginnerStateContext) {
	if !v.valid {
		return
	}
	for _, side := range ctx.AllSide() {
		for _, cell := range side.(*cp.SideContext).AllColor() {
			v.colors[cell.GetText()]++
		}
	}
	stateCtx := ctx.GetParent().(*cp.StateContext)
	stateDefCtx := stateCtx.GetParent().(*cp.StateDefContext)
	expectedColors := v.size * v.size
	checkColors(stateDefCtx, v, expectedColors)
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) {
	sides := ctx.AllSide()
	if len(sides) < 6 {
		errorMsg := fmt.Sprintf("state should have 6 sides, have %d", len(sides))
		v.Eh().AddWarning(ctx, errorMsg, v.FileName())
		v.finished = false
	} else if len(sides) > 6 {
		errorMsg := fmt.Sprintf("state should have 6 sides, have %d", len(sides))
		v.Eh().AddError(ctx, errorMsg, v.FileName())
		v.valid = false
	}
	v.checkSideNames(sides)
	for _, side := range sides {
		v.visitSide(side.(*cp.SideContext))
	}
	v.checkColors(ctx)
}

func (v *beginnerStateVisitor) visitSide(ctx *cp.SideContext) {
	colors := ctx.AllColor()
	cells := v.size * v.size
	if len(colors) < cells {
		warningMsg := fmt.Sprintf("side state should have %d cells, has %d", cells, len(colors))
		v.Eh().AddWarning(ctx, warningMsg, v.FileName())
		v.finished = false
	} else if len(colors) > cells {
		errorMsg := fmt.Sprintf("invalid number of cell, wanted %d, has %d", cells, len(colors))
		v.Eh().AddError(ctx, errorMsg, v.FileName())
		v.valid = false
	}
}
