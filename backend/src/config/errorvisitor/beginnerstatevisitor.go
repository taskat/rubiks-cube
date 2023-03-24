package errorvisitor

import (
	"fmt"

	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type beginnerStateVisitor struct {
	fileName string
	finished bool
	valid    bool
}

func newBeginnerStateVisitor(fileName string) *beginnerStateVisitor {
	return &beginnerStateVisitor{fileName: fileName, finished: true, valid: true}
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
				eh.AddError(sideCtx, errorMsg, v.fileName)
			}
		}
	}
}

func (v *beginnerStateVisitor) checkColors(ctx *cp.BeginnerStateContext) {
	fmt.Println("checkColors")
	if !v.valid {
		return
	}
	colors := make(map[string][]*cp.ColorContext)
	for _, side := range ctx.AllSide() {
		for _, row := range side.(*cp.SideContext).SideState().(*cp.SideStateContext).AllSideStateRow() {
			for _, cell := range row.(*cp.SideStateRowContext).AllColor() {
				colors[cell.GetText()] = append(colors[cell.GetText()], cell.(*cp.ColorContext))
			}
		}
	}
	for color, colorCtxs := range colors {
		if len(colorCtxs) > 9 || (v.finished && len(colorCtxs) < 9) {
			errorMsg := fmt.Sprintf("color %s is defined %d times, should be 9 times", color, len(colorCtxs))
			for _, colorCtx := range colorCtxs {
				eh.AddError(colorCtx, errorMsg, v.fileName)
			}
		}
	}
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) {
	sides := ctx.AllSide()
	if len(sides) != 6 {
		errorMsg := fmt.Sprintf("state should have 6 sides, have %d", len(sides))
		eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	}
	v.checkSideNames(sides)
	for _, side := range sides {
		v.visitSide(side.(*cp.SideContext))
	}
	v.checkColors(ctx)
}

func (v *beginnerStateVisitor) visitSide(ctx *cp.SideContext) {
	v.visitSideState(ctx.SideState().(*cp.SideStateContext))
}

func (v *beginnerStateVisitor) visitSideState(ctx *cp.SideStateContext) {
	rows := ctx.AllSideStateRow()
	if len(rows) < 3 {
		errorMsg := fmt.Sprintf("side state should have 3 rows, have %d", len(rows))
		eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	} else if len(rows) > 3 {
		errorMsg := fmt.Sprintf("invalid number of rows, wanted 3, have %d", len(rows))
		eh.AddError(ctx, errorMsg, v.fileName)
		v.valid = false
	}
	for _, row := range rows {
		v.visitSideStateRow(row.(*cp.SideStateRowContext))
	}
}

func (v *beginnerStateVisitor) visitSideStateRow(ctx *cp.SideStateRowContext) {
	cells := ctx.AllColor()
	if len(cells) < 3 {
		errorMsg := fmt.Sprintf("side state row should have 3 cells, have %d", len(cells))
		eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	} else if len(cells) > 3 {
		errorMsg := fmt.Sprintf("invalid number of cells, wanted 3, have %d", len(cells))
		eh.AddError(ctx, errorMsg, v.fileName)
		v.valid = false
	}
}
