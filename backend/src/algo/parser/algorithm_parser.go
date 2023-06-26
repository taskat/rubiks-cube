// Code generated from grammars/AlgorithmParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package algoparser // AlgorithmParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AlgorithmParser struct {
	*antlr.BaseParser
}

var algorithmparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algorithmparserParserInit() {
	staticData := &algorithmparserParserStaticData
	staticData.literalNames = []string{
		"", "'all'", "'any'", "'branches'", "'consecutive'", "'do'", "'goal'",
		"'helpers'", "'if'", "'none'", "'piece'", "'pos'", "'position'", "'prepare'",
		"'runs'", "'step'", "'steps'", "':'", "','", "'?'", "'''", "'('", "')'",
		"'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "ALL", "ANY", "BRANCHES", "CONSECUTIVE", "DO", "GOAL", "HELPERS",
		"IF", "NONE", "PIECE", "POS", "POSITION", "PREPARE", "RUNS", "STEP",
		"STEPS", "COLON", "COMMA", "QUESTIONMARK", "PRIME", "LPAREN", "RPAREN",
		"LBRACKET", "RBRACKET", "AND", "OR", "NOT", "NUMBER", "WORD", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"algorithmFile", "helpers", "helperLine", "steps", "step", "stepLine",
		"goal", "runs", "branches", "doDef", "branch", "ifBranch", "prepareBranch",
		"consecutive", "algorithm", "turn", "boolExpr", "unaryOp", "binaryOp",
		"expr", "unaryExpr", "binaryExpr", "functionalExpr", "function", "parameter",
		"singleNode", "node", "piece", "position", "coord", "list", "sides",
		"side",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 287, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 3, 0, 68, 8, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 4, 1, 76, 8, 1, 11, 1, 12, 1, 77, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 5, 3, 87, 8, 3, 10, 3, 12, 3, 90, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5,
		4, 96, 8, 4, 10, 4, 12, 4, 99, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		106, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 4, 8, 119, 8, 8, 11, 8, 12, 8, 120, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 3, 10, 129, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 142, 8, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 4, 14, 149, 8, 14, 11, 14, 12, 14, 150, 1, 15, 1, 15, 3,
		15, 155, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 162, 8, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 173,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 179, 8, 16, 10, 16, 12, 16, 182,
		9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 191, 8,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 218, 8, 24, 1, 25, 1, 25, 3, 25,
		222, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30,
		242, 8, 30, 10, 30, 12, 30, 245, 9, 30, 1, 30, 1, 30, 1, 30, 5, 30, 250,
		8, 30, 10, 30, 12, 30, 253, 9, 30, 1, 30, 1, 30, 1, 30, 5, 30, 258, 8,
		30, 10, 30, 12, 30, 261, 9, 30, 1, 30, 1, 30, 1, 30, 5, 30, 266, 8, 30,
		10, 30, 12, 30, 269, 9, 30, 3, 30, 271, 8, 30, 3, 30, 273, 8, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 5, 31, 280, 8, 31, 10, 31, 12, 31, 283, 9,
		31, 1, 32, 1, 32, 1, 32, 0, 1, 32, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 0, 4, 2, 0, 20, 20, 28, 28, 1, 0, 25, 26, 2, 0, 1,
		2, 9, 9, 1, 0, 11, 12, 287, 0, 67, 1, 0, 0, 0, 2, 72, 1, 0, 0, 0, 4, 79,
		1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8, 91, 1, 0, 0, 0, 10, 105, 1, 0, 0, 0,
		12, 107, 1, 0, 0, 0, 14, 111, 1, 0, 0, 0, 16, 115, 1, 0, 0, 0, 18, 122,
		1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 130, 1, 0, 0, 0, 24, 135, 1, 0, 0,
		0, 26, 143, 1, 0, 0, 0, 28, 148, 1, 0, 0, 0, 30, 161, 1, 0, 0, 0, 32, 172,
		1, 0, 0, 0, 34, 183, 1, 0, 0, 0, 36, 185, 1, 0, 0, 0, 38, 190, 1, 0, 0,
		0, 40, 192, 1, 0, 0, 0, 42, 197, 1, 0, 0, 0, 44, 201, 1, 0, 0, 0, 46, 208,
		1, 0, 0, 0, 48, 217, 1, 0, 0, 0, 50, 219, 1, 0, 0, 0, 52, 223, 1, 0, 0,
		0, 54, 227, 1, 0, 0, 0, 56, 230, 1, 0, 0, 0, 58, 233, 1, 0, 0, 0, 60, 237,
		1, 0, 0, 0, 62, 276, 1, 0, 0, 0, 64, 284, 1, 0, 0, 0, 66, 68, 3, 2, 1,
		0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70,
		3, 6, 3, 0, 70, 71, 5, 0, 0, 1, 71, 1, 1, 0, 0, 0, 72, 73, 5, 7, 0, 0,
		73, 75, 5, 17, 0, 0, 74, 76, 3, 4, 2, 0, 75, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 3, 1, 0, 0, 0, 79,
		80, 5, 29, 0, 0, 80, 81, 5, 17, 0, 0, 81, 82, 3, 28, 14, 0, 82, 5, 1, 0,
		0, 0, 83, 84, 5, 16, 0, 0, 84, 88, 5, 17, 0, 0, 85, 87, 3, 8, 4, 0, 86,
		85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0,
		0, 89, 7, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 15, 0, 0, 92, 93,
		5, 29, 0, 0, 93, 97, 5, 17, 0, 0, 94, 96, 3, 10, 5, 0, 95, 94, 1, 0, 0,
		0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 9, 1,
		0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 106, 3, 12, 6, 0, 101, 106, 3, 2, 1,
		0, 102, 106, 3, 14, 7, 0, 103, 106, 3, 16, 8, 0, 104, 106, 3, 18, 9, 0,
		105, 100, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 105, 102, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 11, 1, 0, 0, 0, 107, 108, 5,
		6, 0, 0, 108, 109, 5, 17, 0, 0, 109, 110, 3, 32, 16, 0, 110, 13, 1, 0,
		0, 0, 111, 112, 5, 14, 0, 0, 112, 113, 5, 17, 0, 0, 113, 114, 5, 28, 0,
		0, 114, 15, 1, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116, 118, 5, 17, 0, 0, 117,
		119, 3, 20, 10, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 17, 1, 0, 0, 0, 122, 123, 5, 5,
		0, 0, 123, 124, 5, 17, 0, 0, 124, 125, 3, 28, 14, 0, 125, 19, 1, 0, 0,
		0, 126, 129, 3, 22, 11, 0, 127, 129, 3, 24, 12, 0, 128, 126, 1, 0, 0, 0,
		128, 127, 1, 0, 0, 0, 129, 21, 1, 0, 0, 0, 130, 131, 5, 8, 0, 0, 131, 132,
		3, 32, 16, 0, 132, 133, 5, 17, 0, 0, 133, 134, 3, 18, 9, 0, 134, 23, 1,
		0, 0, 0, 135, 136, 5, 13, 0, 0, 136, 141, 5, 17, 0, 0, 137, 138, 3, 18,
		9, 0, 138, 139, 3, 26, 13, 0, 139, 142, 1, 0, 0, 0, 140, 142, 3, 28, 14,
		0, 141, 137, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 25, 1, 0, 0, 0, 143,
		144, 5, 4, 0, 0, 144, 145, 5, 17, 0, 0, 145, 146, 5, 28, 0, 0, 146, 27,
		1, 0, 0, 0, 147, 149, 3, 30, 15, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1,
		0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 29, 1, 0, 0,
		0, 152, 154, 5, 29, 0, 0, 153, 155, 7, 0, 0, 0, 154, 153, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 162, 1, 0, 0, 0, 156, 157, 5, 28, 0, 0, 157, 158,
		5, 21, 0, 0, 158, 159, 3, 28, 14, 0, 159, 160, 5, 22, 0, 0, 160, 162, 1,
		0, 0, 0, 161, 152, 1, 0, 0, 0, 161, 156, 1, 0, 0, 0, 162, 31, 1, 0, 0,
		0, 163, 164, 6, 16, -1, 0, 164, 165, 3, 34, 17, 0, 165, 166, 3, 32, 16,
		4, 166, 173, 1, 0, 0, 0, 167, 168, 5, 21, 0, 0, 168, 169, 3, 32, 16, 0,
		169, 170, 5, 22, 0, 0, 170, 173, 1, 0, 0, 0, 171, 173, 3, 38, 19, 0, 172,
		163, 1, 0, 0, 0, 172, 167, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 180,
		1, 0, 0, 0, 174, 175, 10, 3, 0, 0, 175, 176, 3, 36, 18, 0, 176, 177, 3,
		32, 16, 4, 177, 179, 1, 0, 0, 0, 178, 174, 1, 0, 0, 0, 179, 182, 1, 0,
		0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 33, 1, 0, 0, 0,
		182, 180, 1, 0, 0, 0, 183, 184, 5, 27, 0, 0, 184, 35, 1, 0, 0, 0, 185,
		186, 7, 1, 0, 0, 186, 37, 1, 0, 0, 0, 187, 191, 3, 40, 20, 0, 188, 191,
		3, 42, 21, 0, 189, 191, 3, 44, 22, 0, 190, 187, 1, 0, 0, 0, 190, 188, 1,
		0, 0, 0, 190, 189, 1, 0, 0, 0, 191, 39, 1, 0, 0, 0, 192, 193, 5, 29, 0,
		0, 193, 194, 5, 21, 0, 0, 194, 195, 3, 48, 24, 0, 195, 196, 5, 22, 0, 0,
		196, 41, 1, 0, 0, 0, 197, 198, 3, 48, 24, 0, 198, 199, 5, 29, 0, 0, 199,
		200, 3, 48, 24, 0, 200, 43, 1, 0, 0, 0, 201, 202, 3, 46, 23, 0, 202, 203,
		5, 21, 0, 0, 203, 204, 3, 32, 16, 0, 204, 205, 5, 18, 0, 0, 205, 206, 3,
		60, 30, 0, 206, 207, 5, 22, 0, 0, 207, 45, 1, 0, 0, 0, 208, 209, 7, 2,
		0, 0, 209, 47, 1, 0, 0, 0, 210, 218, 3, 50, 25, 0, 211, 218, 3, 52, 26,
		0, 212, 218, 3, 54, 27, 0, 213, 218, 3, 56, 28, 0, 214, 218, 3, 58, 29,
		0, 215, 218, 3, 60, 30, 0, 216, 218, 5, 19, 0, 0, 217, 210, 1, 0, 0, 0,
		217, 211, 1, 0, 0, 0, 217, 212, 1, 0, 0, 0, 217, 213, 1, 0, 0, 0, 217,
		214, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 216, 1, 0, 0, 0, 218, 49, 1,
		0, 0, 0, 219, 221, 3, 62, 31, 0, 220, 222, 5, 28, 0, 0, 221, 220, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 51, 1, 0, 0, 0, 223, 224, 5, 21, 0, 0,
		224, 225, 3, 50, 25, 0, 225, 226, 5, 22, 0, 0, 226, 53, 1, 0, 0, 0, 227,
		228, 5, 10, 0, 0, 228, 229, 3, 52, 26, 0, 229, 55, 1, 0, 0, 0, 230, 231,
		7, 3, 0, 0, 231, 232, 3, 52, 26, 0, 232, 57, 1, 0, 0, 0, 233, 234, 3, 64,
		32, 0, 234, 235, 5, 28, 0, 0, 235, 236, 5, 28, 0, 0, 236, 59, 1, 0, 0,
		0, 237, 272, 5, 23, 0, 0, 238, 243, 3, 52, 26, 0, 239, 240, 5, 18, 0, 0,
		240, 242, 3, 52, 26, 0, 241, 239, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243,
		241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 273, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 246, 251, 3, 54, 27, 0, 247, 248, 5, 18, 0, 0, 248, 250, 3,
		54, 27, 0, 249, 247, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 251, 252, 1, 0, 0, 0, 252, 273, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0,
		254, 259, 3, 56, 28, 0, 255, 256, 5, 18, 0, 0, 256, 258, 3, 56, 28, 0,
		257, 255, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 271, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 267,
		3, 58, 29, 0, 263, 264, 5, 18, 0, 0, 264, 266, 3, 58, 29, 0, 265, 263,
		1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0,
		0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270, 254, 1, 0, 0, 0,
		270, 262, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 238, 1, 0, 0, 0, 272,
		246, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275,
		5, 24, 0, 0, 275, 61, 1, 0, 0, 0, 276, 281, 3, 64, 32, 0, 277, 278, 5,
		18, 0, 0, 278, 280, 3, 64, 32, 0, 279, 277, 1, 0, 0, 0, 280, 283, 1, 0,
		0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 63, 1, 0, 0, 0,
		283, 281, 1, 0, 0, 0, 284, 285, 5, 29, 0, 0, 285, 65, 1, 0, 0, 0, 23, 67,
		77, 88, 97, 105, 120, 128, 141, 150, 154, 161, 172, 180, 190, 217, 221,
		243, 251, 259, 267, 270, 272, 281,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AlgorithmParserInit initializes any static state used to implement AlgorithmParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAlgorithmParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlgorithmParserInit() {
	staticData := &algorithmparserParserStaticData
	staticData.once.Do(algorithmparserParserInit)
}

// NewAlgorithmParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAlgorithmParser(input antlr.TokenStream) *AlgorithmParser {
	AlgorithmParserInit()
	this := new(AlgorithmParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &algorithmparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AlgorithmParser.g4"

	return this
}

// AlgorithmParser tokens.
const (
	AlgorithmParserEOF           = antlr.TokenEOF
	AlgorithmParserALL           = 1
	AlgorithmParserANY           = 2
	AlgorithmParserBRANCHES      = 3
	AlgorithmParserCONSECUTIVE   = 4
	AlgorithmParserDO            = 5
	AlgorithmParserGOAL          = 6
	AlgorithmParserHELPERS       = 7
	AlgorithmParserIF            = 8
	AlgorithmParserNONE          = 9
	AlgorithmParserPIECE         = 10
	AlgorithmParserPOS           = 11
	AlgorithmParserPOSITION      = 12
	AlgorithmParserPREPARE       = 13
	AlgorithmParserRUNS          = 14
	AlgorithmParserSTEP          = 15
	AlgorithmParserSTEPS         = 16
	AlgorithmParserCOLON         = 17
	AlgorithmParserCOMMA         = 18
	AlgorithmParserQUESTIONMARK  = 19
	AlgorithmParserPRIME         = 20
	AlgorithmParserLPAREN        = 21
	AlgorithmParserRPAREN        = 22
	AlgorithmParserLBRACKET      = 23
	AlgorithmParserRBRACKET      = 24
	AlgorithmParserAND           = 25
	AlgorithmParserOR            = 26
	AlgorithmParserNOT           = 27
	AlgorithmParserNUMBER        = 28
	AlgorithmParserWORD          = 29
	AlgorithmParserLINE_COMMENT  = 30
	AlgorithmParserBLOCK_COMMENT = 31
	AlgorithmParserWS            = 32
)

// AlgorithmParser rules.
const (
	AlgorithmParserRULE_algorithmFile  = 0
	AlgorithmParserRULE_helpers        = 1
	AlgorithmParserRULE_helperLine     = 2
	AlgorithmParserRULE_steps          = 3
	AlgorithmParserRULE_step           = 4
	AlgorithmParserRULE_stepLine       = 5
	AlgorithmParserRULE_goal           = 6
	AlgorithmParserRULE_runs           = 7
	AlgorithmParserRULE_branches       = 8
	AlgorithmParserRULE_doDef          = 9
	AlgorithmParserRULE_branch         = 10
	AlgorithmParserRULE_ifBranch       = 11
	AlgorithmParserRULE_prepareBranch  = 12
	AlgorithmParserRULE_consecutive    = 13
	AlgorithmParserRULE_algorithm      = 14
	AlgorithmParserRULE_turn           = 15
	AlgorithmParserRULE_boolExpr       = 16
	AlgorithmParserRULE_unaryOp        = 17
	AlgorithmParserRULE_binaryOp       = 18
	AlgorithmParserRULE_expr           = 19
	AlgorithmParserRULE_unaryExpr      = 20
	AlgorithmParserRULE_binaryExpr     = 21
	AlgorithmParserRULE_functionalExpr = 22
	AlgorithmParserRULE_function       = 23
	AlgorithmParserRULE_parameter      = 24
	AlgorithmParserRULE_singleNode     = 25
	AlgorithmParserRULE_node           = 26
	AlgorithmParserRULE_piece          = 27
	AlgorithmParserRULE_position       = 28
	AlgorithmParserRULE_coord          = 29
	AlgorithmParserRULE_list           = 30
	AlgorithmParserRULE_sides          = 31
	AlgorithmParserRULE_side           = 32
)

// IAlgorithmFileContext is an interface to support dynamic dispatch.
type IAlgorithmFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Steps() IStepsContext
	EOF() antlr.TerminalNode
	Helpers() IHelpersContext

	// IsAlgorithmFileContext differentiates from other interfaces.
	IsAlgorithmFileContext()
}

type AlgorithmFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmFileContext() *AlgorithmFileContext {
	var p = new(AlgorithmFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_algorithmFile
	return p
}

func (*AlgorithmFileContext) IsAlgorithmFileContext() {}

func NewAlgorithmFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmFileContext {
	var p = new(AlgorithmFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_algorithmFile

	return p
}

func (s *AlgorithmFileContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmFileContext) Steps() IStepsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStepsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStepsContext)
}

func (s *AlgorithmFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserEOF, 0)
}

func (s *AlgorithmFileContext) Helpers() IHelpersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelpersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelpersContext)
}

func (s *AlgorithmFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) AlgorithmFile() (localctx IAlgorithmFileContext) {
	this := p
	_ = this

	localctx = NewAlgorithmFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AlgorithmParserRULE_algorithmFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserHELPERS {
		{
			p.SetState(66)
			p.Helpers()
		}

	}
	{
		p.SetState(69)
		p.Steps()
	}
	{
		p.SetState(70)
		p.Match(AlgorithmParserEOF)
	}

	return localctx
}

// IHelpersContext is an interface to support dynamic dispatch.
type IHelpersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HELPERS() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllHelperLine() []IHelperLineContext
	HelperLine(i int) IHelperLineContext

	// IsHelpersContext differentiates from other interfaces.
	IsHelpersContext()
}

type HelpersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelpersContext() *HelpersContext {
	var p = new(HelpersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_helpers
	return p
}

func (*HelpersContext) IsHelpersContext() {}

func NewHelpersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelpersContext {
	var p = new(HelpersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_helpers

	return p
}

func (s *HelpersContext) GetParser() antlr.Parser { return s.parser }

func (s *HelpersContext) HELPERS() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserHELPERS, 0)
}

func (s *HelpersContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *HelpersContext) AllHelperLine() []IHelperLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHelperLineContext); ok {
			len++
		}
	}

	tst := make([]IHelperLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHelperLineContext); ok {
			tst[i] = t.(IHelperLineContext)
			i++
		}
	}

	return tst
}

func (s *HelpersContext) HelperLine(i int) IHelperLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelperLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelperLineContext)
}

func (s *HelpersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelpersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Helpers() (localctx IHelpersContext) {
	this := p
	_ = this

	localctx = NewHelpersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AlgorithmParserRULE_helpers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(AlgorithmParserHELPERS)
	}
	{
		p.SetState(73)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AlgorithmParserWORD {
		{
			p.SetState(74)
			p.HelperLine()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHelperLineContext is an interface to support dynamic dispatch.
type IHelperLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Algorithm() IAlgorithmContext

	// IsHelperLineContext differentiates from other interfaces.
	IsHelperLineContext()
}

type HelperLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelperLineContext() *HelperLineContext {
	var p = new(HelperLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_helperLine
	return p
}

func (*HelperLineContext) IsHelperLineContext() {}

func NewHelperLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelperLineContext {
	var p = new(HelperLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_helperLine

	return p
}

func (s *HelperLineContext) GetParser() antlr.Parser { return s.parser }

func (s *HelperLineContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *HelperLineContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *HelperLineContext) Algorithm() IAlgorithmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmContext)
}

func (s *HelperLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelperLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) HelperLine() (localctx IHelperLineContext) {
	this := p
	_ = this

	localctx = NewHelperLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AlgorithmParserRULE_helperLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(80)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(81)
		p.Algorithm()
	}

	return localctx
}

// IStepsContext is an interface to support dynamic dispatch.
type IStepsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STEPS() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStep() []IStepContext
	Step(i int) IStepContext

	// IsStepsContext differentiates from other interfaces.
	IsStepsContext()
}

type StepsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepsContext() *StepsContext {
	var p = new(StepsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_steps
	return p
}

func (*StepsContext) IsStepsContext() {}

func NewStepsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepsContext {
	var p = new(StepsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_steps

	return p
}

func (s *StepsContext) GetParser() antlr.Parser { return s.parser }

func (s *StepsContext) STEPS() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserSTEPS, 0)
}

func (s *StepsContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *StepsContext) AllStep() []IStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStepContext); ok {
			len++
		}
	}

	tst := make([]IStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStepContext); ok {
			tst[i] = t.(IStepContext)
			i++
		}
	}

	return tst
}

func (s *StepsContext) Step(i int) IStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *StepsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Steps() (localctx IStepsContext) {
	this := p
	_ = this

	localctx = NewStepsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AlgorithmParserRULE_steps)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(AlgorithmParserSTEPS)
	}
	{
		p.SetState(84)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AlgorithmParserSTEP {
		{
			p.SetState(85)
			p.Step()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STEP() antlr.TerminalNode
	WORD() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStepLine() []IStepLineContext
	StepLine(i int) IStepLineContext

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) STEP() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserSTEP, 0)
}

func (s *StepContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *StepContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *StepContext) AllStepLine() []IStepLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStepLineContext); ok {
			len++
		}
	}

	tst := make([]IStepLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStepLineContext); ok {
			tst[i] = t.(IStepLineContext)
			i++
		}
	}

	return tst
}

func (s *StepContext) StepLine(i int) IStepLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStepLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStepLineContext)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Step() (localctx IStepContext) {
	this := p
	_ = this

	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AlgorithmParserRULE_step)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(AlgorithmParserSTEP)
	}
	{
		p.SetState(92)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(93)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16616) != 0 {
		{
			p.SetState(94)
			p.StepLine()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepLineContext is an interface to support dynamic dispatch.
type IStepLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Goal() IGoalContext
	Helpers() IHelpersContext
	Runs() IRunsContext
	Branches() IBranchesContext
	DoDef() IDoDefContext

	// IsStepLineContext differentiates from other interfaces.
	IsStepLineContext()
}

type StepLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepLineContext() *StepLineContext {
	var p = new(StepLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_stepLine
	return p
}

func (*StepLineContext) IsStepLineContext() {}

func NewStepLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepLineContext {
	var p = new(StepLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_stepLine

	return p
}

func (s *StepLineContext) GetParser() antlr.Parser { return s.parser }

func (s *StepLineContext) Goal() IGoalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoalContext)
}

func (s *StepLineContext) Helpers() IHelpersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelpersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelpersContext)
}

func (s *StepLineContext) Runs() IRunsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunsContext)
}

func (s *StepLineContext) Branches() IBranchesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranchesContext)
}

func (s *StepLineContext) DoDef() IDoDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoDefContext)
}

func (s *StepLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) StepLine() (localctx IStepLineContext) {
	this := p
	_ = this

	localctx = NewStepLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AlgorithmParserRULE_stepLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserGOAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Goal()
		}

	case AlgorithmParserHELPERS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Helpers()
		}

	case AlgorithmParserRUNS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Runs()
		}

	case AlgorithmParserBRANCHES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Branches()
		}

	case AlgorithmParserDO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.DoDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGoalContext is an interface to support dynamic dispatch.
type IGoalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GOAL() antlr.TerminalNode
	COLON() antlr.TerminalNode
	BoolExpr() IBoolExprContext

	// IsGoalContext differentiates from other interfaces.
	IsGoalContext()
}

type GoalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoalContext() *GoalContext {
	var p = new(GoalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_goal
	return p
}

func (*GoalContext) IsGoalContext() {}

func NewGoalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoalContext {
	var p = new(GoalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_goal

	return p
}

func (s *GoalContext) GetParser() antlr.Parser { return s.parser }

func (s *GoalContext) GOAL() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserGOAL, 0)
}

func (s *GoalContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *GoalContext) BoolExpr() IBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *GoalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Goal() (localctx IGoalContext) {
	this := p
	_ = this

	localctx = NewGoalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AlgorithmParserRULE_goal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(AlgorithmParserGOAL)
	}
	{
		p.SetState(108)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(109)
		p.boolExpr(0)
	}

	return localctx
}

// IRunsContext is an interface to support dynamic dispatch.
type IRunsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RUNS() antlr.TerminalNode
	COLON() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsRunsContext differentiates from other interfaces.
	IsRunsContext()
}

type RunsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunsContext() *RunsContext {
	var p = new(RunsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_runs
	return p
}

func (*RunsContext) IsRunsContext() {}

func NewRunsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunsContext {
	var p = new(RunsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_runs

	return p
}

func (s *RunsContext) GetParser() antlr.Parser { return s.parser }

func (s *RunsContext) RUNS() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRUNS, 0)
}

func (s *RunsContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *RunsContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
}

func (s *RunsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Runs() (localctx IRunsContext) {
	this := p
	_ = this

	localctx = NewRunsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AlgorithmParserRULE_runs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(AlgorithmParserRUNS)
	}
	{
		p.SetState(112)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(113)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IBranchesContext is an interface to support dynamic dispatch.
type IBranchesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BRANCHES() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext

	// IsBranchesContext differentiates from other interfaces.
	IsBranchesContext()
}

type BranchesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchesContext() *BranchesContext {
	var p = new(BranchesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_branches
	return p
}

func (*BranchesContext) IsBranchesContext() {}

func NewBranchesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchesContext {
	var p = new(BranchesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_branches

	return p
}

func (s *BranchesContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchesContext) BRANCHES() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserBRANCHES, 0)
}

func (s *BranchesContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *BranchesContext) AllBranch() []IBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBranchContext); ok {
			len++
		}
	}

	tst := make([]IBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBranchContext); ok {
			tst[i] = t.(IBranchContext)
			i++
		}
	}

	return tst
}

func (s *BranchesContext) Branch(i int) IBranchContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *BranchesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Branches() (localctx IBranchesContext) {
	this := p
	_ = this

	localctx = NewBranchesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AlgorithmParserRULE_branches)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(AlgorithmParserBRANCHES)
	}
	{
		p.SetState(116)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AlgorithmParserIF || _la == AlgorithmParserPREPARE {
		{
			p.SetState(117)
			p.Branch()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDoDefContext is an interface to support dynamic dispatch.
type IDoDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DO() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Algorithm() IAlgorithmContext

	// IsDoDefContext differentiates from other interfaces.
	IsDoDefContext()
}

type DoDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoDefContext() *DoDefContext {
	var p = new(DoDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_doDef
	return p
}

func (*DoDefContext) IsDoDefContext() {}

func NewDoDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoDefContext {
	var p = new(DoDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_doDef

	return p
}

func (s *DoDefContext) GetParser() antlr.Parser { return s.parser }

func (s *DoDefContext) DO() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserDO, 0)
}

func (s *DoDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *DoDefContext) Algorithm() IAlgorithmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmContext)
}

func (s *DoDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) DoDef() (localctx IDoDefContext) {
	this := p
	_ = this

	localctx = NewDoDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AlgorithmParserRULE_doDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(AlgorithmParserDO)
	}
	{
		p.SetState(123)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(124)
		p.Algorithm()
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfBranch() IIfBranchContext
	PrepareBranch() IPrepareBranchContext

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) IfBranch() IIfBranchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBranchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBranchContext)
}

func (s *BranchContext) PrepareBranch() IPrepareBranchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrepareBranchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrepareBranchContext)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Branch() (localctx IBranchContext) {
	this := p
	_ = this

	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AlgorithmParserRULE_branch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.IfBranch()
		}

	case AlgorithmParserPREPARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.PrepareBranch()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfBranchContext is an interface to support dynamic dispatch.
type IIfBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	BoolExpr() IBoolExprContext
	COLON() antlr.TerminalNode
	DoDef() IDoDefContext

	// IsIfBranchContext differentiates from other interfaces.
	IsIfBranchContext()
}

type IfBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBranchContext() *IfBranchContext {
	var p = new(IfBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_ifBranch
	return p
}

func (*IfBranchContext) IsIfBranchContext() {}

func NewIfBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBranchContext {
	var p = new(IfBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_ifBranch

	return p
}

func (s *IfBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBranchContext) IF() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserIF, 0)
}

func (s *IfBranchContext) BoolExpr() IBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *IfBranchContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *IfBranchContext) DoDef() IDoDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoDefContext)
}

func (s *IfBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) IfBranch() (localctx IIfBranchContext) {
	this := p
	_ = this

	localctx = NewIfBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AlgorithmParserRULE_ifBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(AlgorithmParserIF)
	}
	{
		p.SetState(131)
		p.boolExpr(0)
	}
	{
		p.SetState(132)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(133)
		p.DoDef()
	}

	return localctx
}

// IPrepareBranchContext is an interface to support dynamic dispatch.
type IPrepareBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PREPARE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Algorithm() IAlgorithmContext
	DoDef() IDoDefContext
	Consecutive() IConsecutiveContext

	// IsPrepareBranchContext differentiates from other interfaces.
	IsPrepareBranchContext()
}

type PrepareBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrepareBranchContext() *PrepareBranchContext {
	var p = new(PrepareBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_prepareBranch
	return p
}

func (*PrepareBranchContext) IsPrepareBranchContext() {}

func NewPrepareBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrepareBranchContext {
	var p = new(PrepareBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_prepareBranch

	return p
}

func (s *PrepareBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *PrepareBranchContext) PREPARE() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPREPARE, 0)
}

func (s *PrepareBranchContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *PrepareBranchContext) Algorithm() IAlgorithmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmContext)
}

func (s *PrepareBranchContext) DoDef() IDoDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoDefContext)
}

func (s *PrepareBranchContext) Consecutive() IConsecutiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsecutiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsecutiveContext)
}

func (s *PrepareBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrepareBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) PrepareBranch() (localctx IPrepareBranchContext) {
	this := p
	_ = this

	localctx = NewPrepareBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AlgorithmParserRULE_prepareBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(AlgorithmParserPREPARE)
	}
	{
		p.SetState(136)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserDO:
		{
			p.SetState(137)
			p.DoDef()
		}
		{
			p.SetState(138)
			p.Consecutive()
		}

	case AlgorithmParserNUMBER, AlgorithmParserWORD:
		{
			p.SetState(140)
			p.Algorithm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConsecutiveContext is an interface to support dynamic dispatch.
type IConsecutiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONSECUTIVE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsConsecutiveContext differentiates from other interfaces.
	IsConsecutiveContext()
}

type ConsecutiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsecutiveContext() *ConsecutiveContext {
	var p = new(ConsecutiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_consecutive
	return p
}

func (*ConsecutiveContext) IsConsecutiveContext() {}

func NewConsecutiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsecutiveContext {
	var p = new(ConsecutiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_consecutive

	return p
}

func (s *ConsecutiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsecutiveContext) CONSECUTIVE() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCONSECUTIVE, 0)
}

func (s *ConsecutiveContext) COLON() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOLON, 0)
}

func (s *ConsecutiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
}

func (s *ConsecutiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsecutiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Consecutive() (localctx IConsecutiveContext) {
	this := p
	_ = this

	localctx = NewConsecutiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AlgorithmParserRULE_consecutive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(AlgorithmParserCONSECUTIVE)
	}
	{
		p.SetState(144)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(145)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IAlgorithmContext is an interface to support dynamic dispatch.
type IAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTurn() []ITurnContext
	Turn(i int) ITurnContext

	// IsAlgorithmContext differentiates from other interfaces.
	IsAlgorithmContext()
}

type AlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmContext() *AlgorithmContext {
	var p = new(AlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_algorithm
	return p
}

func (*AlgorithmContext) IsAlgorithmContext() {}

func NewAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmContext {
	var p = new(AlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_algorithm

	return p
}

func (s *AlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmContext) AllTurn() []ITurnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITurnContext); ok {
			len++
		}
	}

	tst := make([]ITurnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITurnContext); ok {
			tst[i] = t.(ITurnContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmContext) Turn(i int) ITurnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITurnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITurnContext)
}

func (s *AlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Algorithm() (localctx IAlgorithmContext) {
	this := p
	_ = this

	localctx = NewAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AlgorithmParserRULE_algorithm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(147)
				p.Turn()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// ITurnContext is an interface to support dynamic dispatch.
type ITurnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	PRIME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Algorithm() IAlgorithmContext
	RPAREN() antlr.TerminalNode

	// IsTurnContext differentiates from other interfaces.
	IsTurnContext()
}

type TurnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTurnContext() *TurnContext {
	var p = new(TurnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_turn
	return p
}

func (*TurnContext) IsTurnContext() {}

func NewTurnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TurnContext {
	var p = new(TurnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_turn

	return p
}

func (s *TurnContext) GetParser() antlr.Parser { return s.parser }

func (s *TurnContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *TurnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
}

func (s *TurnContext) PRIME() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPRIME, 0)
}

func (s *TurnContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *TurnContext) Algorithm() IAlgorithmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmContext)
}

func (s *TurnContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *TurnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TurnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Turn() (localctx ITurnContext) {
	this := p
	_ = this

	localctx = NewTurnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AlgorithmParserRULE_turn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(AlgorithmParserWORD)
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(153)
				_la = p.GetTokenStream().LA(1)

				if !(_la == AlgorithmParserPRIME || _la == AlgorithmParserNUMBER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case AlgorithmParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(AlgorithmParserNUMBER)
		}
		{
			p.SetState(157)
			p.Match(AlgorithmParserLPAREN)
		}
		{
			p.SetState(158)
			p.Algorithm()
		}
		{
			p.SetState(159)
			p.Match(AlgorithmParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOp() IUnaryOpContext
	AllBoolExpr() []IBoolExprContext
	BoolExpr(i int) IBoolExprContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext
	BinaryOp() IBinaryOpContext

	// IsBoolExprContext differentiates from other interfaces.
	IsBoolExprContext()
}

type BoolExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprContext() *BoolExprContext {
	var p = new(BoolExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_boolExpr
	return p
}

func (*BoolExprContext) IsBoolExprContext() {}

func NewBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprContext {
	var p = new(BoolExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_boolExpr

	return p
}

func (s *BoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *BoolExprContext) AllBoolExpr() []IBoolExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExprContext); ok {
			len++
		}
	}

	tst := make([]IBoolExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExprContext); ok {
			tst[i] = t.(IBoolExprContext)
			i++
		}
	}

	return tst
}

func (s *BoolExprContext) BoolExpr(i int) IBoolExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *BoolExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *BoolExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *BoolExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BoolExprContext) BinaryOp() IBinaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) BoolExpr() (localctx IBoolExprContext) {
	return p.boolExpr(0)
}

func (p *AlgorithmParser) boolExpr(_p int) (localctx IBoolExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, AlgorithmParserRULE_boolExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(164)
			p.UnaryOp()
		}
		{
			p.SetState(165)
			p.boolExpr(4)
		}

	case 2:
		{
			p.SetState(167)
			p.Match(AlgorithmParserLPAREN)
		}
		{
			p.SetState(168)
			p.boolExpr(0)
		}
		{
			p.SetState(169)
			p.Match(AlgorithmParserRPAREN)
		}

	case 3:
		{
			p.SetState(171)
			p.Expr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBoolExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AlgorithmParserRULE_boolExpr)
			p.SetState(174)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(175)
				p.BinaryOp()
			}
			{
				p.SetState(176)
				p.boolExpr(4)
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNOT, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) UnaryOp() (localctx IUnaryOpContext) {
	this := p
	_ = this

	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AlgorithmParserRULE_unaryOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(AlgorithmParserNOT)
	}

	return localctx
}

// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsBinaryOpContext differentiates from other interfaces.
	IsBinaryOpContext()
}

type BinaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpContext() *BinaryOpContext {
	var p = new(BinaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_binaryOp
	return p
}

func (*BinaryOpContext) IsBinaryOpContext() {}

func NewBinaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_binaryOp

	return p
}

func (s *BinaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpContext) AND() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserAND, 0)
}

func (s *BinaryOpContext) OR() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserOR, 0)
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) BinaryOp() (localctx IBinaryOpContext) {
	this := p
	_ = this

	localctx = NewBinaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AlgorithmParserRULE_binaryOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AlgorithmParserAND || _la == AlgorithmParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	BinaryExpr() IBinaryExprContext
	FunctionalExpr() IFunctionalExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExprContext) BinaryExpr() IBinaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryExprContext)
}

func (s *ExprContext) FunctionalExpr() IFunctionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionalExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AlgorithmParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.UnaryExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.BinaryExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.FunctionalExpr()
		}

	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Parameter() IParameterContext
	RPAREN() antlr.TerminalNode

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *UnaryExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *UnaryExprContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *UnaryExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) UnaryExpr() (localctx IUnaryExprContext) {
	this := p
	_ = this

	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AlgorithmParserRULE_unaryExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(193)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(194)
		p.Parameter()
	}
	{
		p.SetState(195)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IBinaryExprContext is an interface to support dynamic dispatch.
type IBinaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	WORD() antlr.TerminalNode

	// IsBinaryExprContext differentiates from other interfaces.
	IsBinaryExprContext()
}

type BinaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryExprContext() *BinaryExprContext {
	var p = new(BinaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_binaryExpr
	return p
}

func (*BinaryExprContext) IsBinaryExprContext() {}

func NewBinaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryExprContext {
	var p = new(BinaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_binaryExpr

	return p
}

func (s *BinaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryExprContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExprContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *BinaryExprContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) BinaryExpr() (localctx IBinaryExprContext) {
	this := p
	_ = this

	localctx = NewBinaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AlgorithmParserRULE_binaryExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Parameter()
	}
	{
		p.SetState(198)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(199)
		p.Parameter()
	}

	return localctx
}

// IFunctionalExprContext is an interface to support dynamic dispatch.
type IFunctionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() IFunctionContext
	LPAREN() antlr.TerminalNode
	BoolExpr() IBoolExprContext
	COMMA() antlr.TerminalNode
	List() IListContext
	RPAREN() antlr.TerminalNode

	// IsFunctionalExprContext differentiates from other interfaces.
	IsFunctionalExprContext()
}

type FunctionalExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionalExprContext() *FunctionalExprContext {
	var p = new(FunctionalExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_functionalExpr
	return p
}

func (*FunctionalExprContext) IsFunctionalExprContext() {}

func NewFunctionalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionalExprContext {
	var p = new(FunctionalExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_functionalExpr

	return p
}

func (s *FunctionalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionalExprContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionalExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *FunctionalExprContext) BoolExpr() IBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *FunctionalExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOMMA, 0)
}

func (s *FunctionalExprContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *FunctionalExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *FunctionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) FunctionalExpr() (localctx IFunctionalExprContext) {
	this := p
	_ = this

	localctx = NewFunctionalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AlgorithmParserRULE_functionalExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Function()
	}
	{
		p.SetState(202)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(203)
		p.boolExpr(0)
	}
	{
		p.SetState(204)
		p.Match(AlgorithmParserCOMMA)
	}
	{
		p.SetState(205)
		p.List()
	}
	{
		p.SetState(206)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALL() antlr.TerminalNode
	ANY() antlr.TerminalNode
	NONE() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) ALL() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserALL, 0)
}

func (s *FunctionContext) ANY() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserANY, 0)
}

func (s *FunctionContext) NONE() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNONE, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AlgorithmParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&518) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleNode() ISingleNodeContext
	Node() INodeContext
	Piece() IPieceContext
	Position() IPositionContext
	Coord() ICoordContext
	List() IListContext
	QUESTIONMARK() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) SingleNode() ISingleNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleNodeContext)
}

func (s *ParameterContext) Node() INodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *ParameterContext) Piece() IPieceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPieceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPieceContext)
}

func (s *ParameterContext) Position() IPositionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPositionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *ParameterContext) Coord() ICoordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordContext)
}

func (s *ParameterContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ParameterContext) QUESTIONMARK() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserQUESTIONMARK, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AlgorithmParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.SingleNode()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Node()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)
			p.Piece()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(213)
			p.Position()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(214)
			p.Coord()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(215)
			p.List()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(216)
			p.Match(AlgorithmParserQUESTIONMARK)
		}

	}

	return localctx
}

// ISingleNodeContext is an interface to support dynamic dispatch.
type ISingleNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sides() ISidesContext
	NUMBER() antlr.TerminalNode

	// IsSingleNodeContext differentiates from other interfaces.
	IsSingleNodeContext()
}

type SingleNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleNodeContext() *SingleNodeContext {
	var p = new(SingleNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_singleNode
	return p
}

func (*SingleNodeContext) IsSingleNodeContext() {}

func NewSingleNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleNodeContext {
	var p = new(SingleNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_singleNode

	return p
}

func (s *SingleNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleNodeContext) Sides() ISidesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISidesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISidesContext)
}

func (s *SingleNodeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
}

func (s *SingleNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) SingleNode() (localctx ISingleNodeContext) {
	this := p
	_ = this

	localctx = NewSingleNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AlgorithmParserRULE_singleNode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Sides()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(220)
			p.Match(AlgorithmParserNUMBER)
		}

	}

	return localctx
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	SingleNode() ISingleNodeContext
	RPAREN() antlr.TerminalNode

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_node
	return p
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *NodeContext) SingleNode() ISingleNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleNodeContext)
}

func (s *NodeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Node() (localctx INodeContext) {
	this := p
	_ = this

	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AlgorithmParserRULE_node)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(224)
		p.SingleNode()
	}
	{
		p.SetState(225)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IPieceContext is an interface to support dynamic dispatch.
type IPieceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIECE() antlr.TerminalNode
	Node() INodeContext

	// IsPieceContext differentiates from other interfaces.
	IsPieceContext()
}

type PieceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPieceContext() *PieceContext {
	var p = new(PieceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_piece
	return p
}

func (*PieceContext) IsPieceContext() {}

func NewPieceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PieceContext {
	var p = new(PieceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_piece

	return p
}

func (s *PieceContext) GetParser() antlr.Parser { return s.parser }

func (s *PieceContext) PIECE() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPIECE, 0)
}

func (s *PieceContext) Node() INodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *PieceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PieceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Piece() (localctx IPieceContext) {
	this := p
	_ = this

	localctx = NewPieceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AlgorithmParserRULE_piece)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(AlgorithmParserPIECE)
	}
	{
		p.SetState(228)
		p.Node()
	}

	return localctx
}

// IPositionContext is an interface to support dynamic dispatch.
type IPositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Node() INodeContext
	POS() antlr.TerminalNode
	POSITION() antlr.TerminalNode

	// IsPositionContext differentiates from other interfaces.
	IsPositionContext()
}

type PositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositionContext() *PositionContext {
	var p = new(PositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_position
	return p
}

func (*PositionContext) IsPositionContext() {}

func NewPositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositionContext {
	var p = new(PositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_position

	return p
}

func (s *PositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PositionContext) Node() INodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *PositionContext) POS() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPOS, 0)
}

func (s *PositionContext) POSITION() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPOSITION, 0)
}

func (s *PositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Position() (localctx IPositionContext) {
	this := p
	_ = this

	localctx = NewPositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AlgorithmParserRULE_position)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AlgorithmParserPOS || _la == AlgorithmParserPOSITION) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(231)
		p.Node()
	}

	return localctx
}

// ICoordContext is an interface to support dynamic dispatch.
type ICoordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Side() ISideContext
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsCoordContext differentiates from other interfaces.
	IsCoordContext()
}

type CoordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordContext() *CoordContext {
	var p = new(CoordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_coord
	return p
}

func (*CoordContext) IsCoordContext() {}

func NewCoordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordContext {
	var p = new(CoordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_coord

	return p
}

func (s *CoordContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordContext) Side() ISideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISideContext)
}

func (s *CoordContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(AlgorithmParserNUMBER)
}

func (s *CoordContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, i)
}

func (s *CoordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Coord() (localctx ICoordContext) {
	this := p
	_ = this

	localctx = NewCoordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AlgorithmParserRULE_coord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Side()
	}
	{
		p.SetState(234)
		p.Match(AlgorithmParserNUMBER)
	}
	{
		p.SetState(235)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllNode() []INodeContext
	Node(i int) INodeContext
	AllPiece() []IPieceContext
	Piece(i int) IPieceContext
	AllPosition() []IPositionContext
	Position(i int) IPositionContext
	AllCoord() []ICoordContext
	Coord(i int) ICoordContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLBRACKET, 0)
}

func (s *ListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRBRACKET, 0)
}

func (s *ListContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *ListContext) AllPiece() []IPieceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPieceContext); ok {
			len++
		}
	}

	tst := make([]IPieceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPieceContext); ok {
			tst[i] = t.(IPieceContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Piece(i int) IPieceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPieceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPieceContext)
}

func (s *ListContext) AllPosition() []IPositionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPositionContext); ok {
			len++
		}
	}

	tst := make([]IPositionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPositionContext); ok {
			tst[i] = t.(IPositionContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Position(i int) IPositionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPositionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *ListContext) AllCoord() []ICoordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICoordContext); ok {
			len++
		}
	}

	tst := make([]ICoordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICoordContext); ok {
			tst[i] = t.(ICoordContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Coord(i int) ICoordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordContext)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AlgorithmParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOMMA, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AlgorithmParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(AlgorithmParserLBRACKET)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserLPAREN:
		{
			p.SetState(238)
			p.Node()
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlgorithmParserCOMMA {
			{
				p.SetState(239)
				p.Match(AlgorithmParserCOMMA)
			}
			{
				p.SetState(240)
				p.Node()
			}

			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case AlgorithmParserPIECE:
		{
			p.SetState(246)
			p.Piece()
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlgorithmParserCOMMA {
			{
				p.SetState(247)
				p.Match(AlgorithmParserCOMMA)
			}
			{
				p.SetState(248)
				p.Piece()
			}

			p.SetState(253)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case AlgorithmParserPOS, AlgorithmParserPOSITION, AlgorithmParserWORD:
		p.SetState(270)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AlgorithmParserPOS, AlgorithmParserPOSITION:
			{
				p.SetState(254)
				p.Position()
			}
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlgorithmParserCOMMA {
				{
					p.SetState(255)
					p.Match(AlgorithmParserCOMMA)
				}
				{
					p.SetState(256)
					p.Position()
				}

				p.SetState(261)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case AlgorithmParserWORD:
			{
				p.SetState(262)
				p.Coord()
			}
			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlgorithmParserCOMMA {
				{
					p.SetState(263)
					p.Match(AlgorithmParserCOMMA)
				}
				{
					p.SetState(264)
					p.Coord()
				}

				p.SetState(269)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(274)
		p.Match(AlgorithmParserRBRACKET)
	}

	return localctx
}

// ISidesContext is an interface to support dynamic dispatch.
type ISidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSide() []ISideContext
	Side(i int) ISideContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSidesContext differentiates from other interfaces.
	IsSidesContext()
}

type SidesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidesContext() *SidesContext {
	var p = new(SidesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_sides
	return p
}

func (*SidesContext) IsSidesContext() {}

func NewSidesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidesContext {
	var p = new(SidesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_sides

	return p
}

func (s *SidesContext) GetParser() antlr.Parser { return s.parser }

func (s *SidesContext) AllSide() []ISideContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISideContext); ok {
			len++
		}
	}

	tst := make([]ISideContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISideContext); ok {
			tst[i] = t.(ISideContext)
			i++
		}
	}

	return tst
}

func (s *SidesContext) Side(i int) ISideContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISideContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISideContext)
}

func (s *SidesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AlgorithmParserCOMMA)
}

func (s *SidesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AlgorithmParserCOMMA, i)
}

func (s *SidesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Sides() (localctx ISidesContext) {
	this := p
	_ = this

	localctx = NewSidesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AlgorithmParserRULE_sides)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Side()
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(277)
				p.Match(AlgorithmParserCOMMA)
			}
			{
				p.SetState(278)
				p.Side()
			}

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// ISideContext is an interface to support dynamic dispatch.
type ISideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode

	// IsSideContext differentiates from other interfaces.
	IsSideContext()
}

type SideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySideContext() *SideContext {
	var p = new(SideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlgorithmParserRULE_side
	return p
}

func (*SideContext) IsSideContext() {}

func NewSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SideContext {
	var p = new(SideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlgorithmParserRULE_side

	return p
}

func (s *SideContext) GetParser() antlr.Parser { return s.parser }

func (s *SideContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
}

func (s *SideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AlgorithmParser) Side() (localctx ISideContext) {
	this := p
	_ = this

	localctx = NewSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AlgorithmParserRULE_side)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(AlgorithmParserWORD)
	}

	return localctx
}

func (p *AlgorithmParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *BoolExprContext = nil
		if localctx != nil {
			t = localctx.(*BoolExprContext)
		}
		return p.BoolExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlgorithmParser) BoolExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
