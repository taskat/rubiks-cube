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
	eh       eh.Errorhandler
}

func newBeginnerStateVisitor(fileName string, errorHandler eh.Errorhandler) *beginnerStateVisitor {
	return &beginnerStateVisitor{fileName: fileName, finished: true, valid: true, eh: errorHandler}
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
	colors := make(map[string][]*cp.ColorContext)
	for _, side := range ctx.AllSide() {
		for _, row := range side.(*cp.SideContext).AllRow() {
			for _, cell := range row.(*cp.RowContext).AllColor() {
				colors[cell.GetText()] = append(colors[cell.GetText()], cell.(*cp.ColorContext))
			}
		}
	}
	for color, colorCtxs := range colors {
		if len(colorCtxs) > 9 || (v.finished && len(colorCtxs) < 9) {
			errorMsg := fmt.Sprintf("color %s is defined %d times, should be 9 times", color, len(colorCtxs))
			for _, colorCtx := range colorCtxs {
				v.eh.AddWarning(colorCtx, errorMsg, v.fileName)
			}
		}
	}
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
	rows := ctx.AllRow()
	if len(rows) < 3 {
		errorMsg := fmt.Sprintf("side state should have 3 rows, have %d", len(rows))
		v.eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	} else if len(rows) > 3 {
		errorMsg := fmt.Sprintf("invalid number of rows, wanted 3, have %d", len(rows))
		v.eh.AddError(ctx, errorMsg, v.fileName)
		v.valid = false
	}
	for _, row := range rows {
		v.visitRow(row.(*cp.RowContext))
	}
}

func (v *beginnerStateVisitor) visitRow(ctx *cp.RowContext) {
	cells := ctx.AllColor()
	if len(cells) < 3 {
		errorMsg := fmt.Sprintf("side state row should have 3 cells, have %d", len(cells))
		v.eh.AddWarning(ctx, errorMsg, v.fileName)
		v.finished = false
	} else if len(cells) > 3 {
		errorMsg := fmt.Sprintf("invalid number of cells, wanted 3, have %d", len(cells))
		v.eh.AddError(ctx, errorMsg, v.fileName)
		v.valid = false
	}
}
