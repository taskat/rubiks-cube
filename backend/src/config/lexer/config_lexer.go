// Code generated from grammars/ConfigLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configlexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"", "'advanced'", "'back'", "'beginner'", "'cube'", "'down'", "'front'",
		"'left'", "'puzzle'", "'random'", "'right'", "'size'", "'state'", "'state description'",
		"'up'", "'{'", "'}'", "'['", "']'", "'('", "')'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "ADVANCED", "BACK", "BEGINNER", "CUBE", "DOWN", "FRONT", "LEFT",
		"PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE", "STATE_DESCRIPTION", "UP",
		"LCURLY", "RCURLY", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "COLON",
		"NUMBER", "WORD", "LINE_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"ADVANCED", "BACK", "BEGINNER", "CUBE", "DOWN", "FRONT", "LEFT", "PUZZLE",
		"RANDOM", "RIGHT", "SIZE", "STATE", "STATE_DESCRIPTION", "UP", "LCURLY",
		"RCURLY", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "COLON", "NUMBER",
		"WORD", "LETTER", "LINE_COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 192, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21,
		166, 8, 21, 10, 21, 12, 21, 169, 9, 21, 1, 22, 4, 22, 172, 8, 22, 11, 22,
		12, 22, 173, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 182, 8, 24,
		10, 24, 12, 24, 185, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 0,
		0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 0, 49, 24, 51, 25, 1, 0, 5, 1, 0, 49,
		57, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9,
		10, 13, 13, 32, 32, 193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0,
		0, 0, 3, 62, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 76, 1, 0, 0, 0, 9, 81, 1,
		0, 0, 0, 11, 86, 1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17,
		104, 1, 0, 0, 0, 19, 111, 1, 0, 0, 0, 21, 117, 1, 0, 0, 0, 23, 122, 1,
		0, 0, 0, 25, 128, 1, 0, 0, 0, 27, 146, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0,
		31, 151, 1, 0, 0, 0, 33, 153, 1, 0, 0, 0, 35, 155, 1, 0, 0, 0, 37, 157,
		1, 0, 0, 0, 39, 159, 1, 0, 0, 0, 41, 161, 1, 0, 0, 0, 43, 163, 1, 0, 0,
		0, 45, 171, 1, 0, 0, 0, 47, 175, 1, 0, 0, 0, 49, 177, 1, 0, 0, 0, 51, 188,
		1, 0, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 100, 0, 0, 55, 56, 5, 118,
		0, 0, 56, 57, 5, 97, 0, 0, 57, 58, 5, 110, 0, 0, 58, 59, 5, 99, 0, 0, 59,
		60, 5, 101, 0, 0, 60, 61, 5, 100, 0, 0, 61, 2, 1, 0, 0, 0, 62, 63, 5, 98,
		0, 0, 63, 64, 5, 97, 0, 0, 64, 65, 5, 99, 0, 0, 65, 66, 5, 107, 0, 0, 66,
		4, 1, 0, 0, 0, 67, 68, 5, 98, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 103,
		0, 0, 70, 71, 5, 105, 0, 0, 71, 72, 5, 110, 0, 0, 72, 73, 5, 110, 0, 0,
		73, 74, 5, 101, 0, 0, 74, 75, 5, 114, 0, 0, 75, 6, 1, 0, 0, 0, 76, 77,
		5, 99, 0, 0, 77, 78, 5, 117, 0, 0, 78, 79, 5, 98, 0, 0, 79, 80, 5, 101,
		0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 100, 0, 0, 82, 83, 5, 111, 0, 0, 83,
		84, 5, 119, 0, 0, 84, 85, 5, 110, 0, 0, 85, 10, 1, 0, 0, 0, 86, 87, 5,
		102, 0, 0, 87, 88, 5, 114, 0, 0, 88, 89, 5, 111, 0, 0, 89, 90, 5, 110,
		0, 0, 90, 91, 5, 116, 0, 0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 108, 0, 0, 93,
		94, 5, 101, 0, 0, 94, 95, 5, 102, 0, 0, 95, 96, 5, 116, 0, 0, 96, 14, 1,
		0, 0, 0, 97, 98, 5, 112, 0, 0, 98, 99, 5, 117, 0, 0, 99, 100, 5, 122, 0,
		0, 100, 101, 5, 122, 0, 0, 101, 102, 5, 108, 0, 0, 102, 103, 5, 101, 0,
		0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 97, 0, 0,
		106, 107, 5, 110, 0, 0, 107, 108, 5, 100, 0, 0, 108, 109, 5, 111, 0, 0,
		109, 110, 5, 109, 0, 0, 110, 18, 1, 0, 0, 0, 111, 112, 5, 114, 0, 0, 112,
		113, 5, 105, 0, 0, 113, 114, 5, 103, 0, 0, 114, 115, 5, 104, 0, 0, 115,
		116, 5, 116, 0, 0, 116, 20, 1, 0, 0, 0, 117, 118, 5, 115, 0, 0, 118, 119,
		5, 105, 0, 0, 119, 120, 5, 122, 0, 0, 120, 121, 5, 101, 0, 0, 121, 22,
		1, 0, 0, 0, 122, 123, 5, 115, 0, 0, 123, 124, 5, 116, 0, 0, 124, 125, 5,
		97, 0, 0, 125, 126, 5, 116, 0, 0, 126, 127, 5, 101, 0, 0, 127, 24, 1, 0,
		0, 0, 128, 129, 5, 115, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 97,
		0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 101, 0, 0, 133, 134, 5, 32,
		0, 0, 134, 135, 5, 100, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 115,
		0, 0, 137, 138, 5, 99, 0, 0, 138, 139, 5, 114, 0, 0, 139, 140, 5, 105,
		0, 0, 140, 141, 5, 112, 0, 0, 141, 142, 5, 116, 0, 0, 142, 143, 5, 105,
		0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5, 110, 0, 0, 145, 26, 1, 0, 0,
		0, 146, 147, 5, 117, 0, 0, 147, 148, 5, 112, 0, 0, 148, 28, 1, 0, 0, 0,
		149, 150, 5, 123, 0, 0, 150, 30, 1, 0, 0, 0, 151, 152, 5, 125, 0, 0, 152,
		32, 1, 0, 0, 0, 153, 154, 5, 91, 0, 0, 154, 34, 1, 0, 0, 0, 155, 156, 5,
		93, 0, 0, 156, 36, 1, 0, 0, 0, 157, 158, 5, 40, 0, 0, 158, 38, 1, 0, 0,
		0, 159, 160, 5, 41, 0, 0, 160, 40, 1, 0, 0, 0, 161, 162, 5, 58, 0, 0, 162,
		42, 1, 0, 0, 0, 163, 167, 7, 0, 0, 0, 164, 166, 7, 1, 0, 0, 165, 164, 1,
		0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0,
		0, 168, 44, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 172, 3, 47, 23, 0, 171,
		170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 46, 1, 0, 0, 0, 175, 176, 7, 2, 0, 0, 176, 48, 1, 0, 0,
		0, 177, 178, 5, 47, 0, 0, 178, 179, 5, 47, 0, 0, 179, 183, 1, 0, 0, 0,
		180, 182, 8, 3, 0, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183,
		181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183,
		1, 0, 0, 0, 186, 187, 6, 24, 0, 0, 187, 50, 1, 0, 0, 0, 188, 189, 7, 4,
		0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 6, 25, 0, 0, 191, 52, 1, 0, 0, 0,
		4, 0, 167, 173, 183, 1, 6, 0, 0,
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
	ConfigLexerCUBE              = 4
	ConfigLexerDOWN              = 5
	ConfigLexerFRONT             = 6
	ConfigLexerLEFT              = 7
	ConfigLexerPUZZLE            = 8
	ConfigLexerRANDOM            = 9
	ConfigLexerRIGHT             = 10
	ConfigLexerSIZE              = 11
	ConfigLexerSTATE             = 12
	ConfigLexerSTATE_DESCRIPTION = 13
	ConfigLexerUP                = 14
	ConfigLexerLCURLY            = 15
	ConfigLexerRCURLY            = 16
	ConfigLexerLBRACKET          = 17
	ConfigLexerRBRACKET          = 18
	ConfigLexerLPAREN            = 19
	ConfigLexerRPAREN            = 20
	ConfigLexerCOLON             = 21
	ConfigLexerNUMBER            = 22
	ConfigLexerWORD              = 23
	ConfigLexerLINE_COMMENT      = 24
	ConfigLexerWS                = 25
)
