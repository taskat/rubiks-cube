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
		"'up'", "'{'", "'}'", "'['", "']'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "ADVANCED", "BACK", "BEGINNER", "CUBE", "DOWN", "FRONT", "LEFT",
		"PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE", "STATE_DESCRIPTION", "UP",
		"LCURLY", "RCURLY", "LBRACKET", "RBRACKET", "COLON", "NUMBER", "LETTER",
		"STRING", "WS",
	}
	staticData.ruleNames = []string{
		"ADVANCED", "BACK", "BEGINNER", "CUBE", "DOWN", "FRONT", "LEFT", "PUZZLE",
		"RANDOM", "RIGHT", "SIZE", "STATE", "STATE_DESCRIPTION", "UP", "LCURLY",
		"RCURLY", "LBRACKET", "RBRACKET", "COLON", "NUMBER", "LETTER", "STRING",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 171, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		5, 19, 156, 8, 19, 10, 19, 12, 19, 159, 9, 19, 1, 20, 1, 20, 1, 21, 4,
		21, 164, 8, 21, 11, 21, 12, 21, 165, 1, 22, 1, 22, 1, 22, 1, 22, 0, 0,
		23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 1, 0, 4, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0,
		65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 172, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 56,
		1, 0, 0, 0, 5, 61, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11,
		80, 1, 0, 0, 0, 13, 86, 1, 0, 0, 0, 15, 91, 1, 0, 0, 0, 17, 98, 1, 0, 0,
		0, 19, 105, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 116, 1, 0, 0, 0, 25, 122,
		1, 0, 0, 0, 27, 140, 1, 0, 0, 0, 29, 143, 1, 0, 0, 0, 31, 145, 1, 0, 0,
		0, 33, 147, 1, 0, 0, 0, 35, 149, 1, 0, 0, 0, 37, 151, 1, 0, 0, 0, 39, 153,
		1, 0, 0, 0, 41, 160, 1, 0, 0, 0, 43, 163, 1, 0, 0, 0, 45, 167, 1, 0, 0,
		0, 47, 48, 5, 97, 0, 0, 48, 49, 5, 100, 0, 0, 49, 50, 5, 118, 0, 0, 50,
		51, 5, 97, 0, 0, 51, 52, 5, 110, 0, 0, 52, 53, 5, 99, 0, 0, 53, 54, 5,
		101, 0, 0, 54, 55, 5, 100, 0, 0, 55, 2, 1, 0, 0, 0, 56, 57, 5, 98, 0, 0,
		57, 58, 5, 97, 0, 0, 58, 59, 5, 99, 0, 0, 59, 60, 5, 107, 0, 0, 60, 4,
		1, 0, 0, 0, 61, 62, 5, 98, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5, 103,
		0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5, 110, 0, 0, 66, 67, 5, 110, 0, 0,
		67, 68, 5, 101, 0, 0, 68, 69, 5, 114, 0, 0, 69, 6, 1, 0, 0, 0, 70, 71,
		5, 99, 0, 0, 71, 72, 5, 117, 0, 0, 72, 73, 5, 98, 0, 0, 73, 74, 5, 101,
		0, 0, 74, 8, 1, 0, 0, 0, 75, 76, 5, 100, 0, 0, 76, 77, 5, 111, 0, 0, 77,
		78, 5, 119, 0, 0, 78, 79, 5, 110, 0, 0, 79, 10, 1, 0, 0, 0, 80, 81, 5,
		102, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 110,
		0, 0, 84, 85, 5, 116, 0, 0, 85, 12, 1, 0, 0, 0, 86, 87, 5, 108, 0, 0, 87,
		88, 5, 101, 0, 0, 88, 89, 5, 102, 0, 0, 89, 90, 5, 116, 0, 0, 90, 14, 1,
		0, 0, 0, 91, 92, 5, 112, 0, 0, 92, 93, 5, 117, 0, 0, 93, 94, 5, 122, 0,
		0, 94, 95, 5, 122, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 101, 0, 0, 97,
		16, 1, 0, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 97, 0, 0, 100, 101, 5,
		110, 0, 0, 101, 102, 5, 100, 0, 0, 102, 103, 5, 111, 0, 0, 103, 104, 5,
		109, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5, 114, 0, 0, 106, 107, 5, 105,
		0, 0, 107, 108, 5, 103, 0, 0, 108, 109, 5, 104, 0, 0, 109, 110, 5, 116,
		0, 0, 110, 20, 1, 0, 0, 0, 111, 112, 5, 115, 0, 0, 112, 113, 5, 105, 0,
		0, 113, 114, 5, 122, 0, 0, 114, 115, 5, 101, 0, 0, 115, 22, 1, 0, 0, 0,
		116, 117, 5, 115, 0, 0, 117, 118, 5, 116, 0, 0, 118, 119, 5, 97, 0, 0,
		119, 120, 5, 116, 0, 0, 120, 121, 5, 101, 0, 0, 121, 24, 1, 0, 0, 0, 122,
		123, 5, 115, 0, 0, 123, 124, 5, 116, 0, 0, 124, 125, 5, 97, 0, 0, 125,
		126, 5, 116, 0, 0, 126, 127, 5, 101, 0, 0, 127, 128, 5, 32, 0, 0, 128,
		129, 5, 100, 0, 0, 129, 130, 5, 101, 0, 0, 130, 131, 5, 115, 0, 0, 131,
		132, 5, 99, 0, 0, 132, 133, 5, 114, 0, 0, 133, 134, 5, 105, 0, 0, 134,
		135, 5, 112, 0, 0, 135, 136, 5, 116, 0, 0, 136, 137, 5, 105, 0, 0, 137,
		138, 5, 111, 0, 0, 138, 139, 5, 110, 0, 0, 139, 26, 1, 0, 0, 0, 140, 141,
		5, 117, 0, 0, 141, 142, 5, 112, 0, 0, 142, 28, 1, 0, 0, 0, 143, 144, 5,
		123, 0, 0, 144, 30, 1, 0, 0, 0, 145, 146, 5, 125, 0, 0, 146, 32, 1, 0,
		0, 0, 147, 148, 5, 91, 0, 0, 148, 34, 1, 0, 0, 0, 149, 150, 5, 93, 0, 0,
		150, 36, 1, 0, 0, 0, 151, 152, 5, 58, 0, 0, 152, 38, 1, 0, 0, 0, 153, 157,
		7, 0, 0, 0, 154, 156, 7, 1, 0, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0,
		0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 40, 1, 0, 0, 0,
		159, 157, 1, 0, 0, 0, 160, 161, 7, 2, 0, 0, 161, 42, 1, 0, 0, 0, 162, 164,
		3, 41, 20, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1,
		0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 44, 1, 0, 0, 0, 167, 168, 7, 3, 0,
		0, 168, 169, 1, 0, 0, 0, 169, 170, 6, 22, 0, 0, 170, 46, 1, 0, 0, 0, 3,
		0, 157, 165, 1, 6, 0, 0,
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
	ConfigLexerCOLON             = 19
	ConfigLexerNUMBER            = 20
	ConfigLexerLETTER            = 21
	ConfigLexerSTRING            = 22
	ConfigLexerWS                = 23
)
