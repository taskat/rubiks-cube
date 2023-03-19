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
		"", "'puzzle'", "'cube'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "PUZZLE", "CUBE", "COLON", "WS",
	}
	staticData.ruleNames = []string{
		"PUZZLE", "CUBE", "COLON", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 27, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0,
		1, 3, 0, 9, 10, 13, 13, 32, 32, 26, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 16, 1, 0, 0, 0,
		5, 21, 1, 0, 0, 0, 7, 23, 1, 0, 0, 0, 9, 10, 5, 112, 0, 0, 10, 11, 5, 117,
		0, 0, 11, 12, 5, 122, 0, 0, 12, 13, 5, 122, 0, 0, 13, 14, 5, 108, 0, 0,
		14, 15, 5, 101, 0, 0, 15, 2, 1, 0, 0, 0, 16, 17, 5, 99, 0, 0, 17, 18, 5,
		117, 0, 0, 18, 19, 5, 98, 0, 0, 19, 20, 5, 101, 0, 0, 20, 4, 1, 0, 0, 0,
		21, 22, 5, 58, 0, 0, 22, 6, 1, 0, 0, 0, 23, 24, 7, 0, 0, 0, 24, 25, 1,
		0, 0, 0, 25, 26, 6, 3, 0, 0, 26, 8, 1, 0, 0, 0, 1, 0, 1, 6, 0, 0,
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
	ConfigLexerPUZZLE = 1
	ConfigLexerCUBE   = 2
	ConfigLexerCOLON  = 3
	ConfigLexerWS     = 4
)
