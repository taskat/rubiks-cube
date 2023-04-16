package errorvisitor

import (
	"fmt"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type beginnerStateVisitor struct {
	baseVisitor
}

func newBeginnerStateVisitor(fileName string, errorHandler *eh.Errorhandler) *beginnerStateVisitor {
	return &beginnerStateVisitor{baseVisitor: *newBaseVisitor(fileName, errorHandler)}
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
				v.eh.AddError(sideCtx, errorMsg, v.fileName)
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
	checkColors(stateDefCtx, v, 9)
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) {
	sides := ctx.AllSide()
	if len(sides) < 6 {
		errorMsg := fmt.Sprintf("state should have 6 sides, have %d", len(sides))
		v.eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	} else if len(sides) > 6 {
		errorMsg := fmt.Sprintf("state should have 6 sides, have %d", len(sides))
		v.eh.AddError(ctx, errorMsg, v.fileName)
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
	if len(colors) < 9 {
		warningMsg := fmt.Sprintf("side state should have 9 cells, has %d", len(colors))
		v.eh.AddWarning(ctx, warningMsg, v.fileName)
		v.finished = false
	} else if len(colors) > 9 {
		errorMsg := fmt.Sprintf("invalid number of cell, wanted 9, has %d", len(colors))
		v.eh.AddError(ctx, errorMsg, v.fileName)
		v.valid = false
	}
}
