package util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Visitor interface {
	Eh() *eh.Errorhandler
	FileName() string
}

func CheckZeroDef[def antlr.ParserRuleContext](defs []def, defType string, v Visitor, parentCtx eh.IContext) *def {
	if len(defs) != 0 {
		for _, d := range defs {
			v.Eh().AddWarning(d, defType+" will be ignored", v.FileName())
		}
	}
	return nil
}

func CheckOneDef[def antlr.ParserRuleContext](defs []def, defType string, v Visitor, parentCtx eh.IContext) *def {
	switch {
	case len(defs) > 1:
		for _, d := range defs {
			v.Eh().AddError(d, "Multiple "+defType+" found", v.FileName())
		}
		return nil
	case len(defs) == 0:
		v.Eh().AddError(parentCtx, "No "+defType+" found", v.FileName())
		return nil
	}
	return &defs[0]
}

func CheckOptionalDef[def antlr.ParserRuleContext](defs []def, defType string, v Visitor, parentCtx eh.IContext) *def {
	if len(defs) > 1 {
		for _, d := range defs {
			v.Eh().AddError(d, "Multiple "+defType+" found", v.FileName())
		}
		return nil
	}
	if len(defs) == 1 {
		return &defs[0]
	}
	return nil
}

func GetLines[def, lineType antlr.ParserRuleContext](lines []lineType) []def {
	result := make([]def, 0, 1)
	for _, l := range lines {
		line := l.GetChild(0)
		if line, ok := line.(def); ok {
			result = append(result, line)
		}
	}
	return result
}
