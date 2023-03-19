// Code generated from grammars/ConfigParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configparser // ConfigParser
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

type ConfigParser struct {
	*antlr.BaseParser
}

var configparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func configparserParserInit() {
	staticData := &configparserParserStaticData
	staticData.literalNames = []string{
		"", "'puzzle'", "'cube'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "PUZZLE", "CUBE", "COLON", "WS",
	}
	staticData.ruleNames = []string{
		"configFile", "puzzleTypeDef", "puzzleType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 15, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0, 11, 0, 6, 1, 0, 0,
		0, 2, 8, 1, 0, 0, 0, 4, 12, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 1, 1, 0, 0,
		0, 8, 9, 5, 1, 0, 0, 9, 10, 5, 3, 0, 0, 10, 11, 3, 4, 2, 0, 11, 3, 1, 0,
		0, 0, 12, 13, 5, 2, 0, 0, 13, 5, 1, 0, 0, 0, 0,
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

// ConfigParserInit initializes any static state used to implement ConfigParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConfigParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfigParserInit() {
	staticData := &configparserParserStaticData
	staticData.once.Do(configparserParserInit)
}

// NewConfigParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConfigParser(input antlr.TokenStream) *ConfigParser {
	ConfigParserInit()
	this := new(ConfigParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &configparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ConfigParser.g4"

	return this
}

// ConfigParser tokens.
const (
	ConfigParserEOF    = antlr.TokenEOF
	ConfigParserPUZZLE = 1
	ConfigParserCUBE   = 2
	ConfigParserCOLON  = 3
	ConfigParserWS     = 4
)

// ConfigParser rules.
const (
	ConfigParserRULE_configFile    = 0
	ConfigParserRULE_puzzleTypeDef = 1
	ConfigParserRULE_puzzleType    = 2
)

// IConfigFileContext is an interface to support dynamic dispatch.
type IConfigFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigFileContext differentiates from other interfaces.
	IsConfigFileContext()
}

type ConfigFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigFileContext() *ConfigFileContext {
	var p = new(ConfigFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_configFile
	return p
}

func (*ConfigFileContext) IsConfigFileContext() {}

func NewConfigFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigFileContext {
	var p = new(ConfigFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_configFile

	return p
}

func (s *ConfigFileContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigFileContext) PuzzleTypeDef() IPuzzleTypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPuzzleTypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPuzzleTypeDefContext)
}

func (s *ConfigFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) ConfigFile() (localctx IConfigFileContext) {
	this := p
	_ = this

	localctx = NewConfigFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigParserRULE_configFile)

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
		p.SetState(6)
		p.PuzzleTypeDef()
	}

	return localctx
}

// IPuzzleTypeDefContext is an interface to support dynamic dispatch.
type IPuzzleTypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPuzzleTypeDefContext differentiates from other interfaces.
	IsPuzzleTypeDefContext()
}

type PuzzleTypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPuzzleTypeDefContext() *PuzzleTypeDefContext {
	var p = new(PuzzleTypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_puzzleTypeDef
	return p
}

func (*PuzzleTypeDefContext) IsPuzzleTypeDefContext() {}

func NewPuzzleTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PuzzleTypeDefContext {
	var p = new(PuzzleTypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_puzzleTypeDef

	return p
}

func (s *PuzzleTypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PuzzleTypeDefContext) PUZZLE() antlr.TerminalNode {
	return s.GetToken(ConfigParserPUZZLE, 0)
}

func (s *PuzzleTypeDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *PuzzleTypeDefContext) PuzzleType() IPuzzleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPuzzleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPuzzleTypeContext)
}

func (s *PuzzleTypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PuzzleTypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) PuzzleTypeDef() (localctx IPuzzleTypeDefContext) {
	this := p
	_ = this

	localctx = NewPuzzleTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigParserRULE_puzzleTypeDef)

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
		p.SetState(8)
		p.Match(ConfigParserPUZZLE)
	}
	{
		p.SetState(9)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(10)
		p.PuzzleType()
	}

	return localctx
}

// IPuzzleTypeContext is an interface to support dynamic dispatch.
type IPuzzleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPuzzleTypeContext differentiates from other interfaces.
	IsPuzzleTypeContext()
}

type PuzzleTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPuzzleTypeContext() *PuzzleTypeContext {
	var p = new(PuzzleTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_puzzleType
	return p
}

func (*PuzzleTypeContext) IsPuzzleTypeContext() {}

func NewPuzzleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PuzzleTypeContext {
	var p = new(PuzzleTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_puzzleType

	return p
}

func (s *PuzzleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PuzzleTypeContext) CUBE() antlr.TerminalNode {
	return s.GetToken(ConfigParserCUBE, 0)
}

func (s *PuzzleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PuzzleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) PuzzleType() (localctx IPuzzleTypeContext) {
	this := p
	_ = this

	localctx = NewPuzzleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConfigParserRULE_puzzleType)

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
		p.SetState(12)
		p.Match(ConfigParserCUBE)
	}

	return localctx
}
