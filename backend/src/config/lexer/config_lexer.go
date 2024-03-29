// Code generated from grammars/ConfigLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package configlexer

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

type ConfigLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var configlexerLexerStaticData struct {
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

func configlexerLexerInit() {
	staticData := &configlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'advanced'", "'Back'", "'beginner'", "'corners'", "'cube'", "'Down'",
		"'edges'", "'Front'", "'Left'", "'Middle'", "'puzzle'", "'random'",
		"'Right'", "'size'", "'state'", "'state description'", "'Up'", "'{'",
		"'}'", "'['", "']'", "'('", "')'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "ADVANCED", "BACK", "BEGINNER", "CORNERS", "CUBE", "DOWN", "EDGES",
		"FRONT", "LEFT", "MIDDLE", "PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE",
		"STATE_DESCRIPTION", "UP", "LCURLY", "RCURLY", "LBRACKET", "RBRACKET",
		"LPAREN", "RPAREN", "COLON", "NUMBER", "WORD", "LINE_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"ADVANCED", "BACK", "BEGINNER", "CORNERS", "CUBE", "DOWN", "EDGES",
		"FRONT", "LEFT", "MIDDLE", "PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE",
		"STATE_DESCRIPTION", "UP", "LCURLY", "RCURLY", "LBRACKET", "RBRACKET",
		"LPAREN", "RPAREN", "COLON", "NUMBER", "WORD", "LETTER", "LINE_COMMENT",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 219, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 193, 8, 24, 10, 24, 12, 24, 196,
		9, 24, 1, 25, 4, 25, 199, 8, 25, 11, 25, 12, 25, 200, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 5, 27, 209, 8, 27, 10, 27, 12, 27, 212, 9, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 0, 55, 27, 57, 28, 1, 0, 5, 1, 0, 49, 57,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 220, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 68, 1, 0,
		0, 0, 5, 73, 1, 0, 0, 0, 7, 82, 1, 0, 0, 0, 9, 90, 1, 0, 0, 0, 11, 95,
		1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 112, 1, 0, 0,
		0, 19, 117, 1, 0, 0, 0, 21, 124, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 138,
		1, 0, 0, 0, 27, 144, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0, 31, 155, 1, 0, 0,
		0, 33, 173, 1, 0, 0, 0, 35, 176, 1, 0, 0, 0, 37, 178, 1, 0, 0, 0, 39, 180,
		1, 0, 0, 0, 41, 182, 1, 0, 0, 0, 43, 184, 1, 0, 0, 0, 45, 186, 1, 0, 0,
		0, 47, 188, 1, 0, 0, 0, 49, 190, 1, 0, 0, 0, 51, 198, 1, 0, 0, 0, 53, 202,
		1, 0, 0, 0, 55, 204, 1, 0, 0, 0, 57, 215, 1, 0, 0, 0, 59, 60, 5, 97, 0,
		0, 60, 61, 5, 100, 0, 0, 61, 62, 5, 118, 0, 0, 62, 63, 5, 97, 0, 0, 63,
		64, 5, 110, 0, 0, 64, 65, 5, 99, 0, 0, 65, 66, 5, 101, 0, 0, 66, 67, 5,
		100, 0, 0, 67, 2, 1, 0, 0, 0, 68, 69, 5, 66, 0, 0, 69, 70, 5, 97, 0, 0,
		70, 71, 5, 99, 0, 0, 71, 72, 5, 107, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5,
		98, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 103, 0, 0, 76, 77, 5, 105, 0,
		0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 110, 0, 0, 79, 80, 5, 101, 0, 0, 80,
		81, 5, 114, 0, 0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 99, 0, 0, 83, 84, 5, 111,
		0, 0, 84, 85, 5, 114, 0, 0, 85, 86, 5, 110, 0, 0, 86, 87, 5, 101, 0, 0,
		87, 88, 5, 114, 0, 0, 88, 89, 5, 115, 0, 0, 89, 8, 1, 0, 0, 0, 90, 91,
		5, 99, 0, 0, 91, 92, 5, 117, 0, 0, 92, 93, 5, 98, 0, 0, 93, 94, 5, 101,
		0, 0, 94, 10, 1, 0, 0, 0, 95, 96, 5, 68, 0, 0, 96, 97, 5, 111, 0, 0, 97,
		98, 5, 119, 0, 0, 98, 99, 5, 110, 0, 0, 99, 12, 1, 0, 0, 0, 100, 101, 5,
		101, 0, 0, 101, 102, 5, 100, 0, 0, 102, 103, 5, 103, 0, 0, 103, 104, 5,
		101, 0, 0, 104, 105, 5, 115, 0, 0, 105, 14, 1, 0, 0, 0, 106, 107, 5, 70,
		0, 0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 111, 0, 0, 109, 110, 5, 110,
		0, 0, 110, 111, 5, 116, 0, 0, 111, 16, 1, 0, 0, 0, 112, 113, 5, 76, 0,
		0, 113, 114, 5, 101, 0, 0, 114, 115, 5, 102, 0, 0, 115, 116, 5, 116, 0,
		0, 116, 18, 1, 0, 0, 0, 117, 118, 5, 77, 0, 0, 118, 119, 5, 105, 0, 0,
		119, 120, 5, 100, 0, 0, 120, 121, 5, 100, 0, 0, 121, 122, 5, 108, 0, 0,
		122, 123, 5, 101, 0, 0, 123, 20, 1, 0, 0, 0, 124, 125, 5, 112, 0, 0, 125,
		126, 5, 117, 0, 0, 126, 127, 5, 122, 0, 0, 127, 128, 5, 122, 0, 0, 128,
		129, 5, 108, 0, 0, 129, 130, 5, 101, 0, 0, 130, 22, 1, 0, 0, 0, 131, 132,
		5, 114, 0, 0, 132, 133, 5, 97, 0, 0, 133, 134, 5, 110, 0, 0, 134, 135,
		5, 100, 0, 0, 135, 136, 5, 111, 0, 0, 136, 137, 5, 109, 0, 0, 137, 24,
		1, 0, 0, 0, 138, 139, 5, 82, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5,
		103, 0, 0, 141, 142, 5, 104, 0, 0, 142, 143, 5, 116, 0, 0, 143, 26, 1,
		0, 0, 0, 144, 145, 5, 115, 0, 0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 122,
		0, 0, 147, 148, 5, 101, 0, 0, 148, 28, 1, 0, 0, 0, 149, 150, 5, 115, 0,
		0, 150, 151, 5, 116, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153, 5, 116, 0,
		0, 153, 154, 5, 101, 0, 0, 154, 30, 1, 0, 0, 0, 155, 156, 5, 115, 0, 0,
		156, 157, 5, 116, 0, 0, 157, 158, 5, 97, 0, 0, 158, 159, 5, 116, 0, 0,
		159, 160, 5, 101, 0, 0, 160, 161, 5, 32, 0, 0, 161, 162, 5, 100, 0, 0,
		162, 163, 5, 101, 0, 0, 163, 164, 5, 115, 0, 0, 164, 165, 5, 99, 0, 0,
		165, 166, 5, 114, 0, 0, 166, 167, 5, 105, 0, 0, 167, 168, 5, 112, 0, 0,
		168, 169, 5, 116, 0, 0, 169, 170, 5, 105, 0, 0, 170, 171, 5, 111, 0, 0,
		171, 172, 5, 110, 0, 0, 172, 32, 1, 0, 0, 0, 173, 174, 5, 85, 0, 0, 174,
		175, 5, 112, 0, 0, 175, 34, 1, 0, 0, 0, 176, 177, 5, 123, 0, 0, 177, 36,
		1, 0, 0, 0, 178, 179, 5, 125, 0, 0, 179, 38, 1, 0, 0, 0, 180, 181, 5, 91,
		0, 0, 181, 40, 1, 0, 0, 0, 182, 183, 5, 93, 0, 0, 183, 42, 1, 0, 0, 0,
		184, 185, 5, 40, 0, 0, 185, 44, 1, 0, 0, 0, 186, 187, 5, 41, 0, 0, 187,
		46, 1, 0, 0, 0, 188, 189, 5, 58, 0, 0, 189, 48, 1, 0, 0, 0, 190, 194, 7,
		0, 0, 0, 191, 193, 7, 1, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 50, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 197, 199, 3, 53, 26, 0, 198, 197, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 52, 1, 0,
		0, 0, 202, 203, 7, 2, 0, 0, 203, 54, 1, 0, 0, 0, 204, 205, 5, 47, 0, 0,
		205, 206, 5, 47, 0, 0, 206, 210, 1, 0, 0, 0, 207, 209, 8, 3, 0, 0, 208,
		207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211,
		1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 6, 27,
		0, 0, 214, 56, 1, 0, 0, 0, 215, 216, 7, 4, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 218, 6, 28, 0, 0, 218, 58, 1, 0, 0, 0, 4, 0, 194, 200, 210, 1, 6,
		0, 0,
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

// ConfigLexerInit initializes any static state used to implement ConfigLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewConfigLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfigLexerInit() {
	staticData := &configlexerLexerStaticData
	staticData.once.Do(configlexerLexerInit)
}

// NewConfigLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewConfigLexer(input antlr.CharStream) *ConfigLexer {
	ConfigLexerInit()
	l := new(ConfigLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &configlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ConfigLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ConfigLexer tokens.
const (
	ConfigLexerADVANCED          = 1
	ConfigLexerBACK              = 2
	ConfigLexerBEGINNER          = 3
	ConfigLexerCORNERS           = 4
	ConfigLexerCUBE              = 5
	ConfigLexerDOWN              = 6
	ConfigLexerEDGES             = 7
	ConfigLexerFRONT             = 8
	ConfigLexerLEFT              = 9
	ConfigLexerMIDDLE            = 10
	ConfigLexerPUZZLE            = 11
	ConfigLexerRANDOM            = 12
	ConfigLexerRIGHT             = 13
	ConfigLexerSIZE              = 14
	ConfigLexerSTATE             = 15
	ConfigLexerSTATE_DESCRIPTION = 16
	ConfigLexerUP                = 17
	ConfigLexerLCURLY            = 18
	ConfigLexerRCURLY            = 19
	ConfigLexerLBRACKET          = 20
	ConfigLexerRBRACKET          = 21
	ConfigLexerLPAREN            = 22
	ConfigLexerRPAREN            = 23
	ConfigLexerCOLON             = 24
	ConfigLexerNUMBER            = 25
	ConfigLexerWORD              = 26
	ConfigLexerLINE_COMMENT      = 27
	ConfigLexerWS                = 28
)
