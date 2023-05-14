// Code generated from grammars/AlgorithmParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package algoparser // AlgorithmParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"piece", "position", "coord", "list", "sides", "side",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 259, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 4, 1, 69, 8, 1, 11, 1, 12, 1, 70, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 5, 3, 78, 8, 3, 10, 3, 12, 3, 81, 9, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 87, 8, 4, 10, 4, 12, 4, 90, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 97, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 4, 8, 110, 8, 8, 11, 8, 12, 8, 111, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 3, 10, 120, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 131, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 4, 14, 138, 8, 14, 11, 14, 12, 14, 139, 1, 15, 1, 15, 3, 15,
		144, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 155, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 161, 8, 16, 10,
		16, 12, 16, 164, 9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		3, 19, 173, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 198, 8, 24, 1, 25, 3, 25, 201,
		8, 25, 1, 25, 1, 25, 1, 25, 3, 25, 206, 8, 25, 1, 25, 1, 25, 1, 26, 3,
		26, 211, 8, 26, 1, 26, 3, 26, 214, 8, 26, 3, 26, 216, 8, 26, 1, 26, 1,
		26, 1, 26, 3, 26, 221, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 4, 28, 231, 8, 28, 11, 28, 12, 28, 232, 1, 28, 4, 28, 236,
		8, 28, 11, 28, 12, 28, 237, 1, 28, 4, 28, 241, 8, 28, 11, 28, 12, 28, 242,
		3, 28, 245, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 5, 29, 252, 8, 29,
		10, 29, 12, 29, 255, 9, 29, 1, 30, 1, 30, 1, 30, 0, 1, 32, 31, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 2, 1, 0, 25, 26, 2, 0, 1, 2, 9,
		9, 260, 0, 62, 1, 0, 0, 0, 2, 65, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 79,
		1, 0, 0, 0, 8, 82, 1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12, 98, 1, 0, 0, 0,
		14, 102, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 113, 1, 0, 0, 0, 20, 119,
		1, 0, 0, 0, 22, 121, 1, 0, 0, 0, 24, 126, 1, 0, 0, 0, 26, 132, 1, 0, 0,
		0, 28, 137, 1, 0, 0, 0, 30, 141, 1, 0, 0, 0, 32, 154, 1, 0, 0, 0, 34, 165,
		1, 0, 0, 0, 36, 167, 1, 0, 0, 0, 38, 172, 1, 0, 0, 0, 40, 174, 1, 0, 0,
		0, 42, 179, 1, 0, 0, 0, 44, 183, 1, 0, 0, 0, 46, 190, 1, 0, 0, 0, 48, 197,
		1, 0, 0, 0, 50, 200, 1, 0, 0, 0, 52, 215, 1, 0, 0, 0, 54, 224, 1, 0, 0,
		0, 56, 228, 1, 0, 0, 0, 58, 248, 1, 0, 0, 0, 60, 256, 1, 0, 0, 0, 62, 63,
		3, 2, 1, 0, 63, 64, 3, 6, 3, 0, 64, 1, 1, 0, 0, 0, 65, 66, 5, 7, 0, 0,
		66, 68, 5, 17, 0, 0, 67, 69, 3, 4, 2, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 3, 1, 0, 0, 0, 72,
		73, 5, 29, 0, 0, 73, 74, 5, 17, 0, 0, 74, 75, 3, 28, 14, 0, 75, 5, 1, 0,
		0, 0, 76, 78, 3, 8, 4, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 7, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 83, 5, 15, 0, 0, 83, 84, 5, 29, 0, 0, 84, 88, 5, 17, 0, 0, 85, 87,
		3, 10, 5, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0,
		88, 89, 1, 0, 0, 0, 89, 9, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 97, 3, 12,
		6, 0, 92, 97, 3, 2, 1, 0, 93, 97, 3, 14, 7, 0, 94, 97, 3, 16, 8, 0, 95,
		97, 3, 18, 9, 0, 96, 91, 1, 0, 0, 0, 96, 92, 1, 0, 0, 0, 96, 93, 1, 0,
		0, 0, 96, 94, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 11, 1, 0, 0, 0, 98, 99,
		5, 6, 0, 0, 99, 100, 5, 17, 0, 0, 100, 101, 3, 32, 16, 0, 101, 13, 1, 0,
		0, 0, 102, 103, 5, 14, 0, 0, 103, 104, 5, 17, 0, 0, 104, 105, 5, 28, 0,
		0, 105, 15, 1, 0, 0, 0, 106, 107, 5, 3, 0, 0, 107, 109, 5, 17, 0, 0, 108,
		110, 3, 20, 10, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 17, 1, 0, 0, 0, 113, 114, 5, 5,
		0, 0, 114, 115, 5, 17, 0, 0, 115, 116, 3, 28, 14, 0, 116, 19, 1, 0, 0,
		0, 117, 120, 3, 22, 11, 0, 118, 120, 3, 24, 12, 0, 119, 117, 1, 0, 0, 0,
		119, 118, 1, 0, 0, 0, 120, 21, 1, 0, 0, 0, 121, 122, 5, 8, 0, 0, 122, 123,
		3, 32, 16, 0, 123, 124, 5, 17, 0, 0, 124, 125, 3, 18, 9, 0, 125, 23, 1,
		0, 0, 0, 126, 127, 5, 13, 0, 0, 127, 128, 5, 17, 0, 0, 128, 130, 3, 18,
		9, 0, 129, 131, 3, 26, 13, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0,
		0, 131, 25, 1, 0, 0, 0, 132, 133, 5, 4, 0, 0, 133, 134, 5, 17, 0, 0, 134,
		135, 5, 28, 0, 0, 135, 27, 1, 0, 0, 0, 136, 138, 3, 30, 15, 0, 137, 136,
		1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 29, 1, 0, 0, 0, 141, 143, 5, 29, 0, 0, 142, 144, 5, 20, 0, 0,
		143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 31, 1, 0, 0, 0, 145, 146,
		6, 16, -1, 0, 146, 147, 3, 34, 17, 0, 147, 148, 3, 32, 16, 4, 148, 155,
		1, 0, 0, 0, 149, 150, 5, 21, 0, 0, 150, 151, 3, 32, 16, 0, 151, 152, 5,
		22, 0, 0, 152, 155, 1, 0, 0, 0, 153, 155, 3, 38, 19, 0, 154, 145, 1, 0,
		0, 0, 154, 149, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 162, 1, 0, 0, 0,
		156, 157, 10, 3, 0, 0, 157, 158, 3, 36, 18, 0, 158, 159, 3, 32, 16, 4,
		159, 161, 1, 0, 0, 0, 160, 156, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162,
		160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 33, 1, 0, 0, 0, 164, 162, 1,
		0, 0, 0, 165, 166, 5, 27, 0, 0, 166, 35, 1, 0, 0, 0, 167, 168, 7, 0, 0,
		0, 168, 37, 1, 0, 0, 0, 169, 173, 3, 40, 20, 0, 170, 173, 3, 42, 21, 0,
		171, 173, 3, 44, 22, 0, 172, 169, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172,
		171, 1, 0, 0, 0, 173, 39, 1, 0, 0, 0, 174, 175, 5, 29, 0, 0, 175, 176,
		5, 21, 0, 0, 176, 177, 3, 48, 24, 0, 177, 178, 5, 22, 0, 0, 178, 41, 1,
		0, 0, 0, 179, 180, 3, 48, 24, 0, 180, 181, 5, 29, 0, 0, 181, 182, 3, 48,
		24, 0, 182, 43, 1, 0, 0, 0, 183, 184, 3, 46, 23, 0, 184, 185, 5, 21, 0,
		0, 185, 186, 3, 38, 19, 0, 186, 187, 5, 18, 0, 0, 187, 188, 3, 56, 28,
		0, 188, 189, 5, 22, 0, 0, 189, 45, 1, 0, 0, 0, 190, 191, 7, 1, 0, 0, 191,
		47, 1, 0, 0, 0, 192, 198, 3, 50, 25, 0, 193, 198, 3, 52, 26, 0, 194, 198,
		3, 54, 27, 0, 195, 198, 3, 56, 28, 0, 196, 198, 5, 19, 0, 0, 197, 192,
		1, 0, 0, 0, 197, 193, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 197, 196, 1, 0, 0, 0, 198, 49, 1, 0, 0, 0, 199, 201, 5, 10, 0, 0,
		200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202,
		203, 5, 21, 0, 0, 203, 205, 3, 58, 29, 0, 204, 206, 5, 28, 0, 0, 205, 204,
		1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 5, 22,
		0, 0, 208, 51, 1, 0, 0, 0, 209, 211, 5, 11, 0, 0, 210, 209, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 216, 1, 0, 0, 0, 212, 214, 5, 12, 0, 0, 213,
		212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 210,
		1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 5, 21,
		0, 0, 218, 220, 3, 58, 29, 0, 219, 221, 5, 28, 0, 0, 220, 219, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 22, 0, 0, 223,
		53, 1, 0, 0, 0, 224, 225, 5, 29, 0, 0, 225, 226, 5, 28, 0, 0, 226, 227,
		5, 28, 0, 0, 227, 55, 1, 0, 0, 0, 228, 244, 5, 23, 0, 0, 229, 231, 3, 50,
		25, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 245, 1, 0, 0, 0, 234, 236, 3, 52, 26, 0, 235,
		234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238,
		1, 0, 0, 0, 238, 245, 1, 0, 0, 0, 239, 241, 3, 54, 27, 0, 240, 239, 1,
		0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0,
		0, 243, 245, 1, 0, 0, 0, 244, 230, 1, 0, 0, 0, 244, 235, 1, 0, 0, 0, 244,
		240, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 24, 0, 0, 247, 57,
		1, 0, 0, 0, 248, 253, 3, 60, 30, 0, 249, 250, 5, 18, 0, 0, 250, 252, 3,
		60, 30, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0,
		0, 0, 253, 254, 1, 0, 0, 0, 254, 59, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0,
		256, 257, 5, 29, 0, 0, 257, 61, 1, 0, 0, 0, 24, 70, 79, 88, 96, 111, 119,
		130, 139, 143, 154, 162, 172, 197, 200, 205, 210, 213, 215, 220, 232, 237,
		242, 244, 253,
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
	AlgorithmParserRULE_piece          = 25
	AlgorithmParserRULE_position       = 26
	AlgorithmParserRULE_coord          = 27
	AlgorithmParserRULE_list           = 28
	AlgorithmParserRULE_sides          = 29
	AlgorithmParserRULE_side           = 30
)

// IAlgorithmFileContext is an interface to support dynamic dispatch.
type IAlgorithmFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(62)
		p.Helpers()
	}
	{
		p.SetState(63)
		p.Steps()
	}

	return localctx
}

// IHelpersContext is an interface to support dynamic dispatch.
type IHelpersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(65)
		p.Match(AlgorithmParserHELPERS)
	}
	{
		p.SetState(66)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AlgorithmParserWORD {
		{
			p.SetState(67)
			p.HelperLine()
		}

		p.SetState(70)
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
		p.SetState(72)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(73)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(74)
		p.Algorithm()
	}

	return localctx
}

// IStepsContext is an interface to support dynamic dispatch.
type IStepsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AlgorithmParserSTEP {
		{
			p.SetState(76)
			p.Step()
		}

		p.SetState(81)
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
		p.SetState(82)
		p.Match(AlgorithmParserSTEP)
	}
	{
		p.SetState(83)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(84)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlgorithmParserBRANCHES)|(1<<AlgorithmParserDO)|(1<<AlgorithmParserGOAL)|(1<<AlgorithmParserHELPERS)|(1<<AlgorithmParserRUNS))) != 0 {
		{
			p.SetState(85)
			p.StepLine()
		}

		p.SetState(90)
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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserGOAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Goal()
		}

	case AlgorithmParserHELPERS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Helpers()
		}

	case AlgorithmParserRUNS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Runs()
		}

	case AlgorithmParserBRANCHES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Branches()
		}

	case AlgorithmParserDO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
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
		p.SetState(98)
		p.Match(AlgorithmParserGOAL)
	}
	{
		p.SetState(99)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(100)
		p.boolExpr(0)
	}

	return localctx
}

// IRunsContext is an interface to support dynamic dispatch.
type IRunsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(102)
		p.Match(AlgorithmParserRUNS)
	}
	{
		p.SetState(103)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(104)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IBranchesContext is an interface to support dynamic dispatch.
type IBranchesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(106)
		p.Match(AlgorithmParserBRANCHES)
	}
	{
		p.SetState(107)
		p.Match(AlgorithmParserCOLON)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AlgorithmParserIF || _la == AlgorithmParserPREPARE {
		{
			p.SetState(108)
			p.Branch()
		}

		p.SetState(111)
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
		p.SetState(113)
		p.Match(AlgorithmParserDO)
	}
	{
		p.SetState(114)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(115)
		p.Algorithm()
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlgorithmParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.IfBranch()
		}

	case AlgorithmParserPREPARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
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
		p.SetState(121)
		p.Match(AlgorithmParserIF)
	}
	{
		p.SetState(122)
		p.boolExpr(0)
	}
	{
		p.SetState(123)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(124)
		p.DoDef()
	}

	return localctx
}

// IPrepareBranchContext is an interface to support dynamic dispatch.
type IPrepareBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(126)
		p.Match(AlgorithmParserPREPARE)
	}
	{
		p.SetState(127)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(128)
		p.DoDef()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserCONSECUTIVE {
		{
			p.SetState(129)
			p.Consecutive()
		}

	}

	return localctx
}

// IConsecutiveContext is an interface to support dynamic dispatch.
type IConsecutiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(132)
		p.Match(AlgorithmParserCONSECUTIVE)
	}
	{
		p.SetState(133)
		p.Match(AlgorithmParserCOLON)
	}
	{
		p.SetState(134)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IAlgorithmContext is an interface to support dynamic dispatch.
type IAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(136)
				p.Turn()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// ITurnContext is an interface to support dynamic dispatch.
type ITurnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *TurnContext) PRIME() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPRIME, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(AlgorithmParserWORD)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserPRIME {
		{
			p.SetState(142)
			p.Match(AlgorithmParserPRIME)
		}

	}

	return localctx
}

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(146)
			p.UnaryOp()
		}
		{
			p.SetState(147)
			p.boolExpr(4)
		}

	case 2:
		{
			p.SetState(149)
			p.Match(AlgorithmParserLPAREN)
		}
		{
			p.SetState(150)
			p.boolExpr(0)
		}
		{
			p.SetState(151)
			p.Match(AlgorithmParserRPAREN)
		}

	case 3:
		{
			p.SetState(153)
			p.Expr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBoolExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AlgorithmParserRULE_boolExpr)
			p.SetState(156)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(157)
				p.BinaryOp()
			}
			{
				p.SetState(158)
				p.boolExpr(4)
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(165)
		p.Match(AlgorithmParserNOT)
	}

	return localctx
}

// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(167)
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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.UnaryExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.BinaryExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
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
		p.SetState(174)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(175)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(176)
		p.Parameter()
	}
	{
		p.SetState(177)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IBinaryExprContext is an interface to support dynamic dispatch.
type IBinaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(179)
		p.Parameter()
	}
	{
		p.SetState(180)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(181)
		p.Parameter()
	}

	return localctx
}

// IFunctionalExprContext is an interface to support dynamic dispatch.
type IFunctionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *FunctionalExprContext) Expr() IExprContext {
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
		p.SetState(183)
		p.Function()
	}
	{
		p.SetState(184)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(185)
		p.Expr()
	}
	{
		p.SetState(186)
		p.Match(AlgorithmParserCOMMA)
	}
	{
		p.SetState(187)
		p.List()
	}
	{
		p.SetState(188)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlgorithmParserALL)|(1<<AlgorithmParserANY)|(1<<AlgorithmParserNONE))) != 0) {
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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Piece()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Position()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.Coord()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.List()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.Match(AlgorithmParserQUESTIONMARK)
		}

	}

	return localctx
}

// IPieceContext is an interface to support dynamic dispatch.
type IPieceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *PieceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *PieceContext) Sides() ISidesContext {
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

func (s *PieceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *PieceContext) PIECE() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserPIECE, 0)
}

func (s *PieceContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
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
	p.EnterRule(localctx, 50, AlgorithmParserRULE_piece)
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
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserPIECE {
		{
			p.SetState(199)
			p.Match(AlgorithmParserPIECE)
		}

	}
	{
		p.SetState(202)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(203)
		p.Sides()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserNUMBER {
		{
			p.SetState(204)
			p.Match(AlgorithmParserNUMBER)
		}

	}
	{
		p.SetState(207)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// IPositionContext is an interface to support dynamic dispatch.
type IPositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *PositionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserLPAREN, 0)
}

func (s *PositionContext) Sides() ISidesContext {
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

func (s *PositionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserRPAREN, 0)
}

func (s *PositionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserNUMBER, 0)
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
	p.EnterRule(localctx, 52, AlgorithmParserRULE_position)
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
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlgorithmParserPOS {
			{
				p.SetState(209)
				p.Match(AlgorithmParserPOS)
			}

		}

	case 2:
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlgorithmParserPOSITION {
			{
				p.SetState(212)
				p.Match(AlgorithmParserPOSITION)
			}

		}

	}
	{
		p.SetState(217)
		p.Match(AlgorithmParserLPAREN)
	}
	{
		p.SetState(218)
		p.Sides()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AlgorithmParserNUMBER {
		{
			p.SetState(219)
			p.Match(AlgorithmParserNUMBER)
		}

	}
	{
		p.SetState(222)
		p.Match(AlgorithmParserRPAREN)
	}

	return localctx
}

// ICoordContext is an interface to support dynamic dispatch.
type ICoordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *CoordContext) WORD() antlr.TerminalNode {
	return s.GetToken(AlgorithmParserWORD, 0)
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
	p.EnterRule(localctx, 54, AlgorithmParserRULE_coord)

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
		p.SetState(224)
		p.Match(AlgorithmParserWORD)
	}
	{
		p.SetState(225)
		p.Match(AlgorithmParserNUMBER)
	}
	{
		p.SetState(226)
		p.Match(AlgorithmParserNUMBER)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.EnterRule(localctx, 56, AlgorithmParserRULE_list)
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
		p.SetState(228)
		p.Match(AlgorithmParserLBRACKET)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlgorithmParserPIECE || _la == AlgorithmParserLPAREN {
			{
				p.SetState(229)
				p.Piece()
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlgorithmParserPOS)|(1<<AlgorithmParserPOSITION)|(1<<AlgorithmParserLPAREN))) != 0) {
			{
				p.SetState(234)
				p.Position()
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlgorithmParserWORD {
			{
				p.SetState(239)
				p.Coord()
			}

			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(246)
		p.Match(AlgorithmParserRBRACKET)
	}

	return localctx
}

// ISidesContext is an interface to support dynamic dispatch.
type ISidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.EnterRule(localctx, 58, AlgorithmParserRULE_sides)
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
		p.SetState(248)
		p.Side()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AlgorithmParserCOMMA {
		{
			p.SetState(249)
			p.Match(AlgorithmParserCOMMA)
		}
		{
			p.SetState(250)
			p.Side()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISideContext is an interface to support dynamic dispatch.
type ISideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.EnterRule(localctx, 60, AlgorithmParserRULE_side)

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
		p.SetState(256)
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
