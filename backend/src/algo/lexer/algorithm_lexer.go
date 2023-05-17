// Code generated from grammars/AlgorithmLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package algolexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AlgorithmLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var algorithmlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func algorithmlexerLexerInit() {
	staticData := &algorithmlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"ALL", "ANY", "BRANCHES", "CONSECUTIVE", "DO", "GOAL", "HELPERS", "IF",
		"NONE", "PIECE", "POS", "POSITION", "PREPARE", "RUNS", "STEP", "STEPS",
		"COLON", "COMMA", "QUESTIONMARK", "PRIME", "LPAREN", "RPAREN", "LBRACKET",
		"RBRACKET", "AND", "OR", "NOT", "NUMBER", "DIGIT", "NONZERO", "WORD",
		"LETTER", "LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 253, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 3, 24, 188, 8, 24, 1, 25, 1, 25, 1, 25, 3, 25, 193, 8, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 3, 26, 199, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 204,
		8, 27, 10, 27, 12, 27, 207, 9, 27, 3, 27, 209, 8, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 5, 30, 218, 8, 30, 10, 30, 12, 30, 221,
		9, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 229, 8, 32, 10,
		32, 12, 32, 232, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33,
		240, 8, 33, 10, 33, 12, 33, 243, 9, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 241, 0, 35, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 0, 59, 0, 61, 29, 63, 0, 65,
		30, 67, 31, 69, 32, 1, 0, 5, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 65, 90,
		97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 258, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0,
		7, 88, 1, 0, 0, 0, 9, 100, 1, 0, 0, 0, 11, 103, 1, 0, 0, 0, 13, 108, 1,
		0, 0, 0, 15, 116, 1, 0, 0, 0, 17, 119, 1, 0, 0, 0, 19, 124, 1, 0, 0, 0,
		21, 130, 1, 0, 0, 0, 23, 134, 1, 0, 0, 0, 25, 143, 1, 0, 0, 0, 27, 151,
		1, 0, 0, 0, 29, 156, 1, 0, 0, 0, 31, 161, 1, 0, 0, 0, 33, 167, 1, 0, 0,
		0, 35, 169, 1, 0, 0, 0, 37, 171, 1, 0, 0, 0, 39, 173, 1, 0, 0, 0, 41, 175,
		1, 0, 0, 0, 43, 177, 1, 0, 0, 0, 45, 179, 1, 0, 0, 0, 47, 181, 1, 0, 0,
		0, 49, 187, 1, 0, 0, 0, 51, 192, 1, 0, 0, 0, 53, 198, 1, 0, 0, 0, 55, 208,
		1, 0, 0, 0, 57, 210, 1, 0, 0, 0, 59, 212, 1, 0, 0, 0, 61, 214, 1, 0, 0,
		0, 63, 222, 1, 0, 0, 0, 65, 224, 1, 0, 0, 0, 67, 235, 1, 0, 0, 0, 69, 249,
		1, 0, 0, 0, 71, 72, 5, 97, 0, 0, 72, 73, 5, 108, 0, 0, 73, 74, 5, 108,
		0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 5, 97, 0, 0, 76, 77, 5, 110, 0, 0, 77,
		78, 5, 121, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 98, 0, 0, 80, 81, 5, 114,
		0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 110, 0, 0, 83, 84, 5, 99, 0, 0, 84,
		85, 5, 104, 0, 0, 85, 86, 5, 101, 0, 0, 86, 87, 5, 115, 0, 0, 87, 6, 1,
		0, 0, 0, 88, 89, 5, 99, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5, 110, 0,
		0, 91, 92, 5, 115, 0, 0, 92, 93, 5, 101, 0, 0, 93, 94, 5, 99, 0, 0, 94,
		95, 5, 117, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 105, 0, 0, 97, 98, 5,
		118, 0, 0, 98, 99, 5, 101, 0, 0, 99, 8, 1, 0, 0, 0, 100, 101, 5, 100, 0,
		0, 101, 102, 5, 111, 0, 0, 102, 10, 1, 0, 0, 0, 103, 104, 5, 103, 0, 0,
		104, 105, 5, 111, 0, 0, 105, 106, 5, 97, 0, 0, 106, 107, 5, 108, 0, 0,
		107, 12, 1, 0, 0, 0, 108, 109, 5, 104, 0, 0, 109, 110, 5, 101, 0, 0, 110,
		111, 5, 108, 0, 0, 111, 112, 5, 112, 0, 0, 112, 113, 5, 101, 0, 0, 113,
		114, 5, 114, 0, 0, 114, 115, 5, 115, 0, 0, 115, 14, 1, 0, 0, 0, 116, 117,
		5, 105, 0, 0, 117, 118, 5, 102, 0, 0, 118, 16, 1, 0, 0, 0, 119, 120, 5,
		110, 0, 0, 120, 121, 5, 111, 0, 0, 121, 122, 5, 110, 0, 0, 122, 123, 5,
		101, 0, 0, 123, 18, 1, 0, 0, 0, 124, 125, 5, 112, 0, 0, 125, 126, 5, 105,
		0, 0, 126, 127, 5, 101, 0, 0, 127, 128, 5, 99, 0, 0, 128, 129, 5, 101,
		0, 0, 129, 20, 1, 0, 0, 0, 130, 131, 5, 112, 0, 0, 131, 132, 5, 111, 0,
		0, 132, 133, 5, 115, 0, 0, 133, 22, 1, 0, 0, 0, 134, 135, 5, 112, 0, 0,
		135, 136, 5, 111, 0, 0, 136, 137, 5, 115, 0, 0, 137, 138, 5, 105, 0, 0,
		138, 139, 5, 116, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5, 111, 0, 0,
		141, 142, 5, 110, 0, 0, 142, 24, 1, 0, 0, 0, 143, 144, 5, 112, 0, 0, 144,
		145, 5, 114, 0, 0, 145, 146, 5, 101, 0, 0, 146, 147, 5, 112, 0, 0, 147,
		148, 5, 97, 0, 0, 148, 149, 5, 114, 0, 0, 149, 150, 5, 101, 0, 0, 150,
		26, 1, 0, 0, 0, 151, 152, 5, 114, 0, 0, 152, 153, 5, 117, 0, 0, 153, 154,
		5, 110, 0, 0, 154, 155, 5, 115, 0, 0, 155, 28, 1, 0, 0, 0, 156, 157, 5,
		115, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 101, 0, 0, 159, 160, 5,
		112, 0, 0, 160, 30, 1, 0, 0, 0, 161, 162, 5, 115, 0, 0, 162, 163, 5, 116,
		0, 0, 163, 164, 5, 101, 0, 0, 164, 165, 5, 112, 0, 0, 165, 166, 5, 115,
		0, 0, 166, 32, 1, 0, 0, 0, 167, 168, 5, 58, 0, 0, 168, 34, 1, 0, 0, 0,
		169, 170, 5, 44, 0, 0, 170, 36, 1, 0, 0, 0, 171, 172, 5, 63, 0, 0, 172,
		38, 1, 0, 0, 0, 173, 174, 5, 39, 0, 0, 174, 40, 1, 0, 0, 0, 175, 176, 5,
		40, 0, 0, 176, 42, 1, 0, 0, 0, 177, 178, 5, 41, 0, 0, 178, 44, 1, 0, 0,
		0, 179, 180, 5, 91, 0, 0, 180, 46, 1, 0, 0, 0, 181, 182, 5, 93, 0, 0, 182,
		48, 1, 0, 0, 0, 183, 188, 5, 38, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186,
		5, 110, 0, 0, 186, 188, 5, 100, 0, 0, 187, 183, 1, 0, 0, 0, 187, 184, 1,
		0, 0, 0, 188, 50, 1, 0, 0, 0, 189, 193, 5, 124, 0, 0, 190, 191, 5, 111,
		0, 0, 191, 193, 5, 114, 0, 0, 192, 189, 1, 0, 0, 0, 192, 190, 1, 0, 0,
		0, 193, 52, 1, 0, 0, 0, 194, 199, 5, 33, 0, 0, 195, 196, 5, 110, 0, 0,
		196, 197, 5, 111, 0, 0, 197, 199, 5, 116, 0, 0, 198, 194, 1, 0, 0, 0, 198,
		195, 1, 0, 0, 0, 199, 54, 1, 0, 0, 0, 200, 209, 5, 48, 0, 0, 201, 205,
		3, 59, 29, 0, 202, 204, 3, 57, 28, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1,
		0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 209, 1, 0, 0,
		0, 207, 205, 1, 0, 0, 0, 208, 200, 1, 0, 0, 0, 208, 201, 1, 0, 0, 0, 209,
		56, 1, 0, 0, 0, 210, 211, 7, 0, 0, 0, 211, 58, 1, 0, 0, 0, 212, 213, 7,
		1, 0, 0, 213, 60, 1, 0, 0, 0, 214, 219, 3, 63, 31, 0, 215, 218, 3, 63,
		31, 0, 216, 218, 5, 95, 0, 0, 217, 215, 1, 0, 0, 0, 217, 216, 1, 0, 0,
		0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220,
		62, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 7, 2, 0, 0, 223, 64, 1,
		0, 0, 0, 224, 225, 5, 47, 0, 0, 225, 226, 5, 47, 0, 0, 226, 230, 1, 0,
		0, 0, 227, 229, 8, 3, 0, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0,
		230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232,
		230, 1, 0, 0, 0, 233, 234, 6, 32, 0, 0, 234, 66, 1, 0, 0, 0, 235, 236,
		5, 47, 0, 0, 236, 237, 5, 42, 0, 0, 237, 241, 1, 0, 0, 0, 238, 240, 9,
		0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 242, 1, 0, 0,
		0, 241, 239, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244,
		245, 5, 42, 0, 0, 245, 246, 5, 47, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248,
		6, 33, 0, 0, 248, 68, 1, 0, 0, 0, 249, 250, 7, 4, 0, 0, 250, 251, 1, 0,
		0, 0, 251, 252, 6, 34, 0, 0, 252, 70, 1, 0, 0, 0, 10, 0, 187, 192, 198,
		205, 208, 217, 219, 230, 241, 1, 6, 0, 0,
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

// AlgorithmLexerInit initializes any static state used to implement AlgorithmLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAlgorithmLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlgorithmLexerInit() {
	staticData := &algorithmlexerLexerStaticData
	staticData.once.Do(algorithmlexerLexerInit)
}

// NewAlgorithmLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAlgorithmLexer(input antlr.CharStream) *AlgorithmLexer {
	AlgorithmLexerInit()
	l := new(AlgorithmLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &algorithmlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "AlgorithmLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlgorithmLexer tokens.
const (
	AlgorithmLexerALL           = 1
	AlgorithmLexerANY           = 2
	AlgorithmLexerBRANCHES      = 3
	AlgorithmLexerCONSECUTIVE   = 4
	AlgorithmLexerDO            = 5
	AlgorithmLexerGOAL          = 6
	AlgorithmLexerHELPERS       = 7
	AlgorithmLexerIF            = 8
	AlgorithmLexerNONE          = 9
	AlgorithmLexerPIECE         = 10
	AlgorithmLexerPOS           = 11
	AlgorithmLexerPOSITION      = 12
	AlgorithmLexerPREPARE       = 13
	AlgorithmLexerRUNS          = 14
	AlgorithmLexerSTEP          = 15
	AlgorithmLexerSTEPS         = 16
	AlgorithmLexerCOLON         = 17
	AlgorithmLexerCOMMA         = 18
	AlgorithmLexerQUESTIONMARK  = 19
	AlgorithmLexerPRIME         = 20
	AlgorithmLexerLPAREN        = 21
	AlgorithmLexerRPAREN        = 22
	AlgorithmLexerLBRACKET      = 23
	AlgorithmLexerRBRACKET      = 24
	AlgorithmLexerAND           = 25
	AlgorithmLexerOR            = 26
	AlgorithmLexerNOT           = 27
	AlgorithmLexerNUMBER        = 28
	AlgorithmLexerWORD          = 29
	AlgorithmLexerLINE_COMMENT  = 30
	AlgorithmLexerBLOCK_COMMENT = 31
	AlgorithmLexerWS            = 32
)