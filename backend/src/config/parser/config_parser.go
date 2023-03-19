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
		"", "'advanced'", "'back'", "'beginner'", "'cube'", "'down'", "'front'",
		"'left'", "'puzzle'", "'random'", "'right'", "'size'", "'state'", "'state description'",
		"'up'", "'{'", "'}'", "'['", "']'", "'('", "')'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "ADVANCED", "BACK", "BEGINNER", "CUBE", "DOWN", "FRONT", "LEFT",
		"PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE", "STATE_DESCRIPTION", "UP",
		"LCURLY", "RCURLY", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "COLON",
		"NUMBER", "STRING", "LINE_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"configFile", "configLine", "puzzleTypeDef", "puzzleType", "sizeDef",
		"stateDescriptionDef", "stateDescription", "stateDef", "state", "beginnerState",
		"side", "sideDef", "sideState", "sideStateRow", "color", "advancedState",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 108, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 45, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 3, 8, 70, 8, 8, 1, 9, 1, 9, 4, 9, 74, 8, 9, 11, 9, 12,
		9, 75, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 4, 12, 88, 8, 12, 11, 12, 12, 12, 89, 1, 12, 1, 12, 1, 13, 1, 13, 4,
		13, 96, 8, 13, 11, 13, 12, 13, 97, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 0, 2, 2, 0, 1, 1, 3, 3, 4, 0, 2, 2, 5, 7, 10, 10, 14,
		14, 100, 0, 35, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 50,
		1, 0, 0, 0, 8, 52, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 60, 1, 0, 0, 0,
		14, 62, 1, 0, 0, 0, 16, 69, 1, 0, 0, 0, 18, 71, 1, 0, 0, 0, 20, 79, 1,
		0, 0, 0, 22, 83, 1, 0, 0, 0, 24, 85, 1, 0, 0, 0, 26, 93, 1, 0, 0, 0, 28,
		101, 1, 0, 0, 0, 30, 103, 1, 0, 0, 0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0,
		0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38,
		1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1, 39, 1, 1, 0, 0, 0,
		40, 45, 3, 4, 2, 0, 41, 45, 3, 8, 4, 0, 42, 45, 3, 10, 5, 0, 43, 45, 3,
		14, 7, 0, 44, 40, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		43, 1, 0, 0, 0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 8, 0, 0, 47, 48, 5, 21, 0,
		0, 48, 49, 3, 6, 3, 0, 49, 5, 1, 0, 0, 0, 50, 51, 5, 4, 0, 0, 51, 7, 1,
		0, 0, 0, 52, 53, 5, 11, 0, 0, 53, 54, 5, 21, 0, 0, 54, 55, 5, 22, 0, 0,
		55, 9, 1, 0, 0, 0, 56, 57, 5, 13, 0, 0, 57, 58, 5, 21, 0, 0, 58, 59, 3,
		12, 6, 0, 59, 11, 1, 0, 0, 0, 60, 61, 7, 0, 0, 0, 61, 13, 1, 0, 0, 0, 62,
		63, 5, 12, 0, 0, 63, 64, 5, 21, 0, 0, 64, 65, 3, 16, 8, 0, 65, 15, 1, 0,
		0, 0, 66, 70, 5, 9, 0, 0, 67, 70, 3, 18, 9, 0, 68, 70, 3, 30, 15, 0, 69,
		66, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 17, 1, 0, 0,
		0, 71, 73, 5, 15, 0, 0, 72, 74, 3, 20, 10, 0, 73, 72, 1, 0, 0, 0, 74, 75,
		1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0,
		77, 78, 5, 16, 0, 0, 78, 19, 1, 0, 0, 0, 79, 80, 3, 22, 11, 0, 80, 81,
		5, 21, 0, 0, 81, 82, 3, 24, 12, 0, 82, 21, 1, 0, 0, 0, 83, 84, 7, 1, 0,
		0, 84, 23, 1, 0, 0, 0, 85, 87, 5, 17, 0, 0, 86, 88, 3, 26, 13, 0, 87, 86,
		1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 92, 5, 18, 0, 0, 92, 25, 1, 0, 0, 0, 93, 95, 5,
		19, 0, 0, 94, 96, 3, 28, 14, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 5,
		20, 0, 0, 100, 27, 1, 0, 0, 0, 101, 102, 5, 23, 0, 0, 102, 29, 1, 0, 0,
		0, 103, 104, 5, 15, 0, 0, 104, 105, 5, 1, 0, 0, 105, 106, 5, 16, 0, 0,
		106, 31, 1, 0, 0, 0, 6, 35, 44, 69, 75, 89, 97,
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
	ConfigParserEOF               = antlr.TokenEOF
	ConfigParserADVANCED          = 1
	ConfigParserBACK              = 2
	ConfigParserBEGINNER          = 3
	ConfigParserCUBE              = 4
	ConfigParserDOWN              = 5
	ConfigParserFRONT             = 6
	ConfigParserLEFT              = 7
	ConfigParserPUZZLE            = 8
	ConfigParserRANDOM            = 9
	ConfigParserRIGHT             = 10
	ConfigParserSIZE              = 11
	ConfigParserSTATE             = 12
	ConfigParserSTATE_DESCRIPTION = 13
	ConfigParserUP                = 14
	ConfigParserLCURLY            = 15
	ConfigParserRCURLY            = 16
	ConfigParserLBRACKET          = 17
	ConfigParserRBRACKET          = 18
	ConfigParserLPAREN            = 19
	ConfigParserRPAREN            = 20
	ConfigParserCOLON             = 21
	ConfigParserNUMBER            = 22
	ConfigParserSTRING            = 23
	ConfigParserLINE_COMMENT      = 24
	ConfigParserWS                = 25
)

// ConfigParser rules.
const (
	ConfigParserRULE_configFile          = 0
	ConfigParserRULE_configLine          = 1
	ConfigParserRULE_puzzleTypeDef       = 2
	ConfigParserRULE_puzzleType          = 3
	ConfigParserRULE_sizeDef             = 4
	ConfigParserRULE_stateDescriptionDef = 5
	ConfigParserRULE_stateDescription    = 6
	ConfigParserRULE_stateDef            = 7
	ConfigParserRULE_state               = 8
	ConfigParserRULE_beginnerState       = 9
	ConfigParserRULE_side                = 10
	ConfigParserRULE_sideDef             = 11
	ConfigParserRULE_sideState           = 12
	ConfigParserRULE_sideStateRow        = 13
	ConfigParserRULE_color               = 14
	ConfigParserRULE_advancedState       = 15
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

func (s *ConfigFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigParserEOF, 0)
}

func (s *ConfigFileContext) AllConfigLine() []IConfigLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConfigLineContext); ok {
			len++
		}
	}

	tst := make([]IConfigLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConfigLineContext); ok {
			tst[i] = t.(IConfigLineContext)
			i++
		}
	}

	return tst
}

func (s *ConfigFileContext) ConfigLine(i int) IConfigLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfigLineContext); ok {
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

	return t.(IConfigLineContext)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserPUZZLE)|(1<<ConfigParserSIZE)|(1<<ConfigParserSTATE)|(1<<ConfigParserSTATE_DESCRIPTION))) != 0 {
		{
			p.SetState(32)
			p.ConfigLine()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(ConfigParserEOF)
	}

	return localctx
}

// IConfigLineContext is an interface to support dynamic dispatch.
type IConfigLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigLineContext differentiates from other interfaces.
	IsConfigLineContext()
}

type ConfigLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigLineContext() *ConfigLineContext {
	var p = new(ConfigLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_configLine
	return p
}

func (*ConfigLineContext) IsConfigLineContext() {}

func NewConfigLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigLineContext {
	var p = new(ConfigLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_configLine

	return p
}

func (s *ConfigLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigLineContext) PuzzleTypeDef() IPuzzleTypeDefContext {
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

func (s *ConfigLineContext) SizeDef() ISizeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISizeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISizeDefContext)
}

func (s *ConfigLineContext) StateDescriptionDef() IStateDescriptionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateDescriptionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateDescriptionDefContext)
}

func (s *ConfigLineContext) StateDef() IStateDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateDefContext)
}

func (s *ConfigLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) ConfigLine() (localctx IConfigLineContext) {
	this := p
	_ = this

	localctx = NewConfigLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigParserRULE_configLine)

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

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigParserPUZZLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.PuzzleTypeDef()
		}

	case ConfigParserSIZE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.SizeDef()
		}

	case ConfigParserSTATE_DESCRIPTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.StateDescriptionDef()
		}

	case ConfigParserSTATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.StateDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 4, ConfigParserRULE_puzzleTypeDef)

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
		p.SetState(46)
		p.Match(ConfigParserPUZZLE)
	}
	{
		p.SetState(47)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(48)
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
	p.EnterRule(localctx, 6, ConfigParserRULE_puzzleType)

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
		p.SetState(50)
		p.Match(ConfigParserCUBE)
	}

	return localctx
}

// ISizeDefContext is an interface to support dynamic dispatch.
type ISizeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizeDefContext differentiates from other interfaces.
	IsSizeDefContext()
}

type SizeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeDefContext() *SizeDefContext {
	var p = new(SizeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_sizeDef
	return p
}

func (*SizeDefContext) IsSizeDefContext() {}

func NewSizeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeDefContext {
	var p = new(SizeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_sizeDef

	return p
}

func (s *SizeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeDefContext) SIZE() antlr.TerminalNode {
	return s.GetToken(ConfigParserSIZE, 0)
}

func (s *SizeDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *SizeDefContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ConfigParserNUMBER, 0)
}

func (s *SizeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) SizeDef() (localctx ISizeDefContext) {
	this := p
	_ = this

	localctx = NewSizeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigParserRULE_sizeDef)

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
		p.SetState(52)
		p.Match(ConfigParserSIZE)
	}
	{
		p.SetState(53)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(54)
		p.Match(ConfigParserNUMBER)
	}

	return localctx
}

// IStateDescriptionDefContext is an interface to support dynamic dispatch.
type IStateDescriptionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateDescriptionDefContext differentiates from other interfaces.
	IsStateDescriptionDefContext()
}

type StateDescriptionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateDescriptionDefContext() *StateDescriptionDefContext {
	var p = new(StateDescriptionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_stateDescriptionDef
	return p
}

func (*StateDescriptionDefContext) IsStateDescriptionDefContext() {}

func NewStateDescriptionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDescriptionDefContext {
	var p = new(StateDescriptionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_stateDescriptionDef

	return p
}

func (s *StateDescriptionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDescriptionDefContext) STATE_DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(ConfigParserSTATE_DESCRIPTION, 0)
}

func (s *StateDescriptionDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *StateDescriptionDefContext) StateDescription() IStateDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateDescriptionContext)
}

func (s *StateDescriptionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDescriptionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) StateDescriptionDef() (localctx IStateDescriptionDefContext) {
	this := p
	_ = this

	localctx = NewStateDescriptionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConfigParserRULE_stateDescriptionDef)

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
		p.SetState(56)
		p.Match(ConfigParserSTATE_DESCRIPTION)
	}
	{
		p.SetState(57)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(58)
		p.StateDescription()
	}

	return localctx
}

// IStateDescriptionContext is an interface to support dynamic dispatch.
type IStateDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateDescriptionContext differentiates from other interfaces.
	IsStateDescriptionContext()
}

type StateDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateDescriptionContext() *StateDescriptionContext {
	var p = new(StateDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_stateDescription
	return p
}

func (*StateDescriptionContext) IsStateDescriptionContext() {}

func NewStateDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDescriptionContext {
	var p = new(StateDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_stateDescription

	return p
}

func (s *StateDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDescriptionContext) BEGINNER() antlr.TerminalNode {
	return s.GetToken(ConfigParserBEGINNER, 0)
}

func (s *StateDescriptionContext) ADVANCED() antlr.TerminalNode {
	return s.GetToken(ConfigParserADVANCED, 0)
}

func (s *StateDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) StateDescription() (localctx IStateDescriptionContext) {
	this := p
	_ = this

	localctx = NewStateDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConfigParserRULE_stateDescription)
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
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConfigParserADVANCED || _la == ConfigParserBEGINNER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStateDefContext is an interface to support dynamic dispatch.
type IStateDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateDefContext differentiates from other interfaces.
	IsStateDefContext()
}

type StateDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateDefContext() *StateDefContext {
	var p = new(StateDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_stateDef
	return p
}

func (*StateDefContext) IsStateDefContext() {}

func NewStateDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDefContext {
	var p = new(StateDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_stateDef

	return p
}

func (s *StateDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDefContext) STATE() antlr.TerminalNode {
	return s.GetToken(ConfigParserSTATE, 0)
}

func (s *StateDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *StateDefContext) State() IStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateContext)
}

func (s *StateDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) StateDef() (localctx IStateDefContext) {
	this := p
	_ = this

	localctx = NewStateDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConfigParserRULE_stateDef)

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
		p.Match(ConfigParserSTATE)
	}
	{
		p.SetState(63)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(64)
		p.State()
	}

	return localctx
}

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateContext differentiates from other interfaces.
	IsStateContext()
}

type StateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateContext() *StateContext {
	var p = new(StateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_state
	return p
}

func (*StateContext) IsStateContext() {}

func NewStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateContext {
	var p = new(StateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_state

	return p
}

func (s *StateContext) GetParser() antlr.Parser { return s.parser }

func (s *StateContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(ConfigParserRANDOM, 0)
}

func (s *StateContext) BeginnerState() IBeginnerStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeginnerStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBeginnerStateContext)
}

func (s *StateContext) AdvancedState() IAdvancedStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdvancedStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdvancedStateContext)
}

func (s *StateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) State() (localctx IStateContext) {
	this := p
	_ = this

	localctx = NewStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConfigParserRULE_state)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(ConfigParserRANDOM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.BeginnerState()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.AdvancedState()
		}

	}

	return localctx
}

// IBeginnerStateContext is an interface to support dynamic dispatch.
type IBeginnerStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeginnerStateContext differentiates from other interfaces.
	IsBeginnerStateContext()
}

type BeginnerStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeginnerStateContext() *BeginnerStateContext {
	var p = new(BeginnerStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_beginnerState
	return p
}

func (*BeginnerStateContext) IsBeginnerStateContext() {}

func NewBeginnerStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginnerStateContext {
	var p = new(BeginnerStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_beginnerState

	return p
}

func (s *BeginnerStateContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginnerStateContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(ConfigParserLCURLY, 0)
}

func (s *BeginnerStateContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(ConfigParserRCURLY, 0)
}

func (s *BeginnerStateContext) AllSide() []ISideContext {
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

func (s *BeginnerStateContext) Side(i int) ISideContext {
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

func (s *BeginnerStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginnerStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) BeginnerState() (localctx IBeginnerStateContext) {
	this := p
	_ = this

	localctx = NewBeginnerStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConfigParserRULE_beginnerState)
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
		p.SetState(71)
		p.Match(ConfigParserLCURLY)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserBACK)|(1<<ConfigParserDOWN)|(1<<ConfigParserFRONT)|(1<<ConfigParserLEFT)|(1<<ConfigParserRIGHT)|(1<<ConfigParserUP))) != 0) {
		{
			p.SetState(72)
			p.Side()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(ConfigParserRCURLY)
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
	p.RuleIndex = ConfigParserRULE_side
	return p
}

func (*SideContext) IsSideContext() {}

func NewSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SideContext {
	var p = new(SideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_side

	return p
}

func (s *SideContext) GetParser() antlr.Parser { return s.parser }

func (s *SideContext) SideDef() ISideDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISideDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISideDefContext)
}

func (s *SideContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *SideContext) SideState() ISideStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISideStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISideStateContext)
}

func (s *SideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Side() (localctx ISideContext) {
	this := p
	_ = this

	localctx = NewSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConfigParserRULE_side)

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
		p.SideDef()
	}
	{
		p.SetState(80)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(81)
		p.SideState()
	}

	return localctx
}

// ISideDefContext is an interface to support dynamic dispatch.
type ISideDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSideDefContext differentiates from other interfaces.
	IsSideDefContext()
}

type SideDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySideDefContext() *SideDefContext {
	var p = new(SideDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_sideDef
	return p
}

func (*SideDefContext) IsSideDefContext() {}

func NewSideDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SideDefContext {
	var p = new(SideDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_sideDef

	return p
}

func (s *SideDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SideDefContext) FRONT() antlr.TerminalNode {
	return s.GetToken(ConfigParserFRONT, 0)
}

func (s *SideDefContext) BACK() antlr.TerminalNode {
	return s.GetToken(ConfigParserBACK, 0)
}

func (s *SideDefContext) LEFT() antlr.TerminalNode {
	return s.GetToken(ConfigParserLEFT, 0)
}

func (s *SideDefContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(ConfigParserRIGHT, 0)
}

func (s *SideDefContext) UP() antlr.TerminalNode {
	return s.GetToken(ConfigParserUP, 0)
}

func (s *SideDefContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ConfigParserDOWN, 0)
}

func (s *SideDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SideDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) SideDef() (localctx ISideDefContext) {
	this := p
	_ = this

	localctx = NewSideDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConfigParserRULE_sideDef)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserBACK)|(1<<ConfigParserDOWN)|(1<<ConfigParserFRONT)|(1<<ConfigParserLEFT)|(1<<ConfigParserRIGHT)|(1<<ConfigParserUP))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISideStateContext is an interface to support dynamic dispatch.
type ISideStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSideStateContext differentiates from other interfaces.
	IsSideStateContext()
}

type SideStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySideStateContext() *SideStateContext {
	var p = new(SideStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_sideState
	return p
}

func (*SideStateContext) IsSideStateContext() {}

func NewSideStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SideStateContext {
	var p = new(SideStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_sideState

	return p
}

func (s *SideStateContext) GetParser() antlr.Parser { return s.parser }

func (s *SideStateContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ConfigParserLBRACKET, 0)
}

func (s *SideStateContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ConfigParserRBRACKET, 0)
}

func (s *SideStateContext) AllSideStateRow() []ISideStateRowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISideStateRowContext); ok {
			len++
		}
	}

	tst := make([]ISideStateRowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISideStateRowContext); ok {
			tst[i] = t.(ISideStateRowContext)
			i++
		}
	}

	return tst
}

func (s *SideStateContext) SideStateRow(i int) ISideStateRowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISideStateRowContext); ok {
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

	return t.(ISideStateRowContext)
}

func (s *SideStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SideStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) SideState() (localctx ISideStateContext) {
	this := p
	_ = this

	localctx = NewSideStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ConfigParserRULE_sideState)
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
		p.SetState(85)
		p.Match(ConfigParserLBRACKET)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserLPAREN {
		{
			p.SetState(86)
			p.SideStateRow()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(ConfigParserRBRACKET)
	}

	return localctx
}

// ISideStateRowContext is an interface to support dynamic dispatch.
type ISideStateRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSideStateRowContext differentiates from other interfaces.
	IsSideStateRowContext()
}

type SideStateRowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySideStateRowContext() *SideStateRowContext {
	var p = new(SideStateRowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_sideStateRow
	return p
}

func (*SideStateRowContext) IsSideStateRowContext() {}

func NewSideStateRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SideStateRowContext {
	var p = new(SideStateRowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_sideStateRow

	return p
}

func (s *SideStateRowContext) GetParser() antlr.Parser { return s.parser }

func (s *SideStateRowContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ConfigParserLPAREN, 0)
}

func (s *SideStateRowContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ConfigParserRPAREN, 0)
}

func (s *SideStateRowContext) AllColor() []IColorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColorContext); ok {
			len++
		}
	}

	tst := make([]IColorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColorContext); ok {
			tst[i] = t.(IColorContext)
			i++
		}
	}

	return tst
}

func (s *SideStateRowContext) Color(i int) IColorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColorContext); ok {
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

	return t.(IColorContext)
}

func (s *SideStateRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SideStateRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) SideStateRow() (localctx ISideStateRowContext) {
	this := p
	_ = this

	localctx = NewSideStateRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ConfigParserRULE_sideStateRow)
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
		p.SetState(93)
		p.Match(ConfigParserLPAREN)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserSTRING {
		{
			p.SetState(94)
			p.Color()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(ConfigParserRPAREN)
	}

	return localctx
}

// IColorContext is an interface to support dynamic dispatch.
type IColorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColorContext differentiates from other interfaces.
	IsColorContext()
}

type ColorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColorContext() *ColorContext {
	var p = new(ColorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_color
	return p
}

func (*ColorContext) IsColorContext() {}

func NewColorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColorContext {
	var p = new(ColorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_color

	return p
}

func (s *ColorContext) GetParser() antlr.Parser { return s.parser }

func (s *ColorContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConfigParserSTRING, 0)
}

func (s *ColorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Color() (localctx IColorContext) {
	this := p
	_ = this

	localctx = NewColorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ConfigParserRULE_color)

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
		p.SetState(101)
		p.Match(ConfigParserSTRING)
	}

	return localctx
}

// IAdvancedStateContext is an interface to support dynamic dispatch.
type IAdvancedStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdvancedStateContext differentiates from other interfaces.
	IsAdvancedStateContext()
}

type AdvancedStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdvancedStateContext() *AdvancedStateContext {
	var p = new(AdvancedStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_advancedState
	return p
}

func (*AdvancedStateContext) IsAdvancedStateContext() {}

func NewAdvancedStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdvancedStateContext {
	var p = new(AdvancedStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_advancedState

	return p
}

func (s *AdvancedStateContext) GetParser() antlr.Parser { return s.parser }

func (s *AdvancedStateContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(ConfigParserLCURLY, 0)
}

func (s *AdvancedStateContext) ADVANCED() antlr.TerminalNode {
	return s.GetToken(ConfigParserADVANCED, 0)
}

func (s *AdvancedStateContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(ConfigParserRCURLY, 0)
}

func (s *AdvancedStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdvancedStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) AdvancedState() (localctx IAdvancedStateContext) {
	this := p
	_ = this

	localctx = NewAdvancedStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConfigParserRULE_advancedState)

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
		p.SetState(103)
		p.Match(ConfigParserLCURLY)
	}
	{
		p.SetState(104)
		p.Match(ConfigParserADVANCED)
	}
	{
		p.SetState(105)
		p.Match(ConfigParserRCURLY)
	}

	return localctx
}
