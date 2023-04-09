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
		"", "'advanced'", "'back'", "'beginner'", "'corners'", "'cube'", "'down'",
		"'edges'", "'front'", "'left'", "'middle'", "'puzzle'", "'random'",
		"'right'", "'size'", "'state'", "'state description'", "'up'", "'{'",
		"'}'", "'['", "']'", "'('", "')'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "ADVANCED", "BACK", "BEGINNER", "CORNERS", "CUBE", "DOWN", "EDGES",
		"FRONT", "LEFT", "MIDDLE", "PUZZLE", "RANDOM", "RIGHT", "SIZE", "STATE",
		"STATE_DESCRIPTION", "UP", "LCURLY", "RCURLY", "LBRACKET", "RBRACKET",
		"LPAREN", "RPAREN", "COLON", "NUMBER", "WORD", "LINE_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"configFile", "configLine", "puzzleTypeDef", "puzzleType", "sizeDef",
		"stateDescriptionDef", "stateDescription", "stateDef", "state", "beginnerState",
		"side", "sideDef", "row", "color", "advancedState", "corners", "cornerLayer",
		"layerDef", "corner", "edges", "edgeLayer", "edge",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 153, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 82, 8, 8, 1, 9, 1, 9, 4, 9, 86, 8,
		9, 11, 9, 12, 9, 87, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 4, 10, 95, 8, 10,
		11, 10, 12, 10, 96, 1, 11, 1, 11, 1, 12, 1, 12, 4, 12, 103, 8, 12, 11,
		12, 12, 12, 104, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 117, 8, 14, 1, 15, 1, 15, 1, 15, 4, 15, 122, 8, 15,
		11, 15, 12, 15, 123, 1, 16, 1, 16, 1, 16, 4, 16, 129, 8, 16, 11, 16, 12,
		16, 130, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 4, 19, 140, 8,
		19, 11, 19, 12, 19, 141, 1, 20, 1, 20, 1, 20, 4, 20, 147, 8, 20, 11, 20,
		12, 20, 148, 1, 21, 1, 21, 1, 21, 0, 0, 22, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 3, 2, 0, 1,
		1, 3, 3, 5, 0, 2, 2, 6, 6, 8, 9, 13, 13, 17, 17, 3, 0, 6, 6, 10, 10, 17,
		17, 144, 0, 47, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 62,
		1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 72, 1, 0, 0, 0,
		14, 74, 1, 0, 0, 0, 16, 81, 1, 0, 0, 0, 18, 83, 1, 0, 0, 0, 20, 91, 1,
		0, 0, 0, 22, 98, 1, 0, 0, 0, 24, 100, 1, 0, 0, 0, 26, 108, 1, 0, 0, 0,
		28, 116, 1, 0, 0, 0, 30, 118, 1, 0, 0, 0, 32, 125, 1, 0, 0, 0, 34, 132,
		1, 0, 0, 0, 36, 134, 1, 0, 0, 0, 38, 136, 1, 0, 0, 0, 40, 143, 1, 0, 0,
		0, 42, 150, 1, 0, 0, 0, 44, 46, 3, 2, 1, 0, 45, 44, 1, 0, 0, 0, 46, 49,
		1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0,
		49, 47, 1, 0, 0, 0, 50, 51, 5, 0, 0, 1, 51, 1, 1, 0, 0, 0, 52, 57, 3, 4,
		2, 0, 53, 57, 3, 8, 4, 0, 54, 57, 3, 10, 5, 0, 55, 57, 3, 14, 7, 0, 56,
		52, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 55, 1, 0, 0,
		0, 57, 3, 1, 0, 0, 0, 58, 59, 5, 11, 0, 0, 59, 60, 5, 24, 0, 0, 60, 61,
		3, 6, 3, 0, 61, 5, 1, 0, 0, 0, 62, 63, 5, 5, 0, 0, 63, 7, 1, 0, 0, 0, 64,
		65, 5, 14, 0, 0, 65, 66, 5, 24, 0, 0, 66, 67, 5, 25, 0, 0, 67, 9, 1, 0,
		0, 0, 68, 69, 5, 16, 0, 0, 69, 70, 5, 24, 0, 0, 70, 71, 3, 12, 6, 0, 71,
		11, 1, 0, 0, 0, 72, 73, 7, 0, 0, 0, 73, 13, 1, 0, 0, 0, 74, 75, 5, 15,
		0, 0, 75, 76, 5, 24, 0, 0, 76, 77, 3, 16, 8, 0, 77, 15, 1, 0, 0, 0, 78,
		82, 5, 12, 0, 0, 79, 82, 3, 18, 9, 0, 80, 82, 3, 28, 14, 0, 81, 78, 1,
		0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 17, 1, 0, 0, 0, 83,
		85, 5, 18, 0, 0, 84, 86, 3, 20, 10, 0, 85, 84, 1, 0, 0, 0, 86, 87, 1, 0,
		0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90,
		5, 19, 0, 0, 90, 19, 1, 0, 0, 0, 91, 92, 3, 22, 11, 0, 92, 94, 5, 24, 0,
		0, 93, 95, 3, 24, 12, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 21, 1, 0, 0, 0, 98, 99, 7, 1, 0, 0,
		99, 23, 1, 0, 0, 0, 100, 102, 5, 22, 0, 0, 101, 103, 3, 26, 13, 0, 102,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 23, 0, 0, 107, 25, 1, 0,
		0, 0, 108, 109, 5, 26, 0, 0, 109, 27, 1, 0, 0, 0, 110, 111, 3, 30, 15,
		0, 111, 112, 3, 38, 19, 0, 112, 117, 1, 0, 0, 0, 113, 114, 3, 38, 19, 0,
		114, 115, 3, 30, 15, 0, 115, 117, 1, 0, 0, 0, 116, 110, 1, 0, 0, 0, 116,
		113, 1, 0, 0, 0, 117, 29, 1, 0, 0, 0, 118, 119, 5, 4, 0, 0, 119, 121, 5,
		24, 0, 0, 120, 122, 3, 32, 16, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 31, 1, 0, 0, 0,
		125, 126, 3, 34, 17, 0, 126, 128, 5, 24, 0, 0, 127, 129, 3, 36, 18, 0,
		128, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130,
		131, 1, 0, 0, 0, 131, 33, 1, 0, 0, 0, 132, 133, 7, 2, 0, 0, 133, 35, 1,
		0, 0, 0, 134, 135, 5, 26, 0, 0, 135, 37, 1, 0, 0, 0, 136, 137, 5, 7, 0,
		0, 137, 139, 5, 24, 0, 0, 138, 140, 3, 40, 20, 0, 139, 138, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		39, 1, 0, 0, 0, 143, 144, 3, 34, 17, 0, 144, 146, 5, 24, 0, 0, 145, 147,
		3, 42, 21, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1,
		0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 41, 1, 0, 0, 0, 150, 151, 5, 26, 0,
		0, 151, 43, 1, 0, 0, 0, 11, 47, 56, 81, 87, 96, 104, 116, 123, 130, 141,
		148,
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
	ConfigParserCORNERS           = 4
	ConfigParserCUBE              = 5
	ConfigParserDOWN              = 6
	ConfigParserEDGES             = 7
	ConfigParserFRONT             = 8
	ConfigParserLEFT              = 9
	ConfigParserMIDDLE            = 10
	ConfigParserPUZZLE            = 11
	ConfigParserRANDOM            = 12
	ConfigParserRIGHT             = 13
	ConfigParserSIZE              = 14
	ConfigParserSTATE             = 15
	ConfigParserSTATE_DESCRIPTION = 16
	ConfigParserUP                = 17
	ConfigParserLCURLY            = 18
	ConfigParserRCURLY            = 19
	ConfigParserLBRACKET          = 20
	ConfigParserRBRACKET          = 21
	ConfigParserLPAREN            = 22
	ConfigParserRPAREN            = 23
	ConfigParserCOLON             = 24
	ConfigParserNUMBER            = 25
	ConfigParserWORD              = 26
	ConfigParserLINE_COMMENT      = 27
	ConfigParserWS                = 28
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
	ConfigParserRULE_row                 = 12
	ConfigParserRULE_color               = 13
	ConfigParserRULE_advancedState       = 14
	ConfigParserRULE_corners             = 15
	ConfigParserRULE_cornerLayer         = 16
	ConfigParserRULE_layerDef            = 17
	ConfigParserRULE_corner              = 18
	ConfigParserRULE_edges               = 19
	ConfigParserRULE_edgeLayer           = 20
	ConfigParserRULE_edge                = 21
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserPUZZLE)|(1<<ConfigParserSIZE)|(1<<ConfigParserSTATE)|(1<<ConfigParserSTATE_DESCRIPTION))) != 0 {
		{
			p.SetState(44)
			p.ConfigLine()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigParserPUZZLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.PuzzleTypeDef()
		}

	case ConfigParserSIZE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.SizeDef()
		}

	case ConfigParserSTATE_DESCRIPTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.StateDescriptionDef()
		}

	case ConfigParserSTATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
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
		p.SetState(58)
		p.Match(ConfigParserPUZZLE)
	}
	{
		p.SetState(59)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(60)
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
		p.SetState(62)
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
		p.SetState(64)
		p.Match(ConfigParserSIZE)
	}
	{
		p.SetState(65)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(66)
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
		p.SetState(68)
		p.Match(ConfigParserSTATE_DESCRIPTION)
	}
	{
		p.SetState(69)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(70)
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
		p.SetState(72)
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
		p.SetState(74)
		p.Match(ConfigParserSTATE)
	}
	{
		p.SetState(75)
		p.Match(ConfigParserCOLON)
	}
	{
		p.SetState(76)
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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigParserRANDOM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(ConfigParserRANDOM)
		}

	case ConfigParserLCURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.BeginnerState()
		}

	case ConfigParserCORNERS, ConfigParserEDGES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.AdvancedState()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(83)
		p.Match(ConfigParserLCURLY)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserBACK)|(1<<ConfigParserDOWN)|(1<<ConfigParserFRONT)|(1<<ConfigParserLEFT)|(1<<ConfigParserRIGHT)|(1<<ConfigParserUP))) != 0) {
		{
			p.SetState(84)
			p.Side()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
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

func (s *SideContext) AllRow() []IRowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRowContext); ok {
			len++
		}
	}

	tst := make([]IRowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRowContext); ok {
			tst[i] = t.(IRowContext)
			i++
		}
	}

	return tst
}

func (s *SideContext) Row(i int) IRowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowContext); ok {
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

	return t.(IRowContext)
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
		p.SideDef()
	}
	{
		p.SetState(92)
		p.Match(ConfigParserCOLON)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserLPAREN {
		{
			p.SetState(93)
			p.Row()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
		p.SetState(98)
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

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ConfigParserLPAREN, 0)
}

func (s *RowContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ConfigParserRPAREN, 0)
}

func (s *RowContext) AllColor() []IColorContext {
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

func (s *RowContext) Color(i int) IColorContext {
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

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Row() (localctx IRowContext) {
	this := p
	_ = this

	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ConfigParserRULE_row)
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
		p.SetState(100)
		p.Match(ConfigParserLPAREN)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserWORD {
		{
			p.SetState(101)
			p.Color()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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

func (s *ColorContext) WORD() antlr.TerminalNode {
	return s.GetToken(ConfigParserWORD, 0)
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
	p.EnterRule(localctx, 26, ConfigParserRULE_color)

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
		p.SetState(108)
		p.Match(ConfigParserWORD)
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

func (s *AdvancedStateContext) Corners() ICornersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICornersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICornersContext)
}

func (s *AdvancedStateContext) Edges() IEdgesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdgesContext)
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
	p.EnterRule(localctx, 28, ConfigParserRULE_advancedState)

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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigParserCORNERS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Corners()
		}
		{
			p.SetState(111)
			p.Edges()
		}

	case ConfigParserEDGES:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Edges()
		}
		{
			p.SetState(114)
			p.Corners()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICornersContext is an interface to support dynamic dispatch.
type ICornersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCornersContext differentiates from other interfaces.
	IsCornersContext()
}

type CornersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCornersContext() *CornersContext {
	var p = new(CornersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_corners
	return p
}

func (*CornersContext) IsCornersContext() {}

func NewCornersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CornersContext {
	var p = new(CornersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_corners

	return p
}

func (s *CornersContext) GetParser() antlr.Parser { return s.parser }

func (s *CornersContext) CORNERS() antlr.TerminalNode {
	return s.GetToken(ConfigParserCORNERS, 0)
}

func (s *CornersContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *CornersContext) AllCornerLayer() []ICornerLayerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICornerLayerContext); ok {
			len++
		}
	}

	tst := make([]ICornerLayerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICornerLayerContext); ok {
			tst[i] = t.(ICornerLayerContext)
			i++
		}
	}

	return tst
}

func (s *CornersContext) CornerLayer(i int) ICornerLayerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICornerLayerContext); ok {
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

	return t.(ICornerLayerContext)
}

func (s *CornersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CornersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Corners() (localctx ICornersContext) {
	this := p
	_ = this

	localctx = NewCornersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConfigParserRULE_corners)
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
		p.SetState(118)
		p.Match(ConfigParserCORNERS)
	}
	{
		p.SetState(119)
		p.Match(ConfigParserCOLON)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserDOWN)|(1<<ConfigParserMIDDLE)|(1<<ConfigParserUP))) != 0) {
		{
			p.SetState(120)
			p.CornerLayer()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICornerLayerContext is an interface to support dynamic dispatch.
type ICornerLayerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCornerLayerContext differentiates from other interfaces.
	IsCornerLayerContext()
}

type CornerLayerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCornerLayerContext() *CornerLayerContext {
	var p = new(CornerLayerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_cornerLayer
	return p
}

func (*CornerLayerContext) IsCornerLayerContext() {}

func NewCornerLayerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CornerLayerContext {
	var p = new(CornerLayerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_cornerLayer

	return p
}

func (s *CornerLayerContext) GetParser() antlr.Parser { return s.parser }

func (s *CornerLayerContext) LayerDef() ILayerDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILayerDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILayerDefContext)
}

func (s *CornerLayerContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *CornerLayerContext) AllCorner() []ICornerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICornerContext); ok {
			len++
		}
	}

	tst := make([]ICornerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICornerContext); ok {
			tst[i] = t.(ICornerContext)
			i++
		}
	}

	return tst
}

func (s *CornerLayerContext) Corner(i int) ICornerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICornerContext); ok {
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

	return t.(ICornerContext)
}

func (s *CornerLayerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CornerLayerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) CornerLayer() (localctx ICornerLayerContext) {
	this := p
	_ = this

	localctx = NewCornerLayerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ConfigParserRULE_cornerLayer)
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
		p.SetState(125)
		p.LayerDef()
	}
	{
		p.SetState(126)
		p.Match(ConfigParserCOLON)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserWORD {
		{
			p.SetState(127)
			p.Corner()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILayerDefContext is an interface to support dynamic dispatch.
type ILayerDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayerDefContext differentiates from other interfaces.
	IsLayerDefContext()
}

type LayerDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayerDefContext() *LayerDefContext {
	var p = new(LayerDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_layerDef
	return p
}

func (*LayerDefContext) IsLayerDefContext() {}

func NewLayerDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayerDefContext {
	var p = new(LayerDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_layerDef

	return p
}

func (s *LayerDefContext) GetParser() antlr.Parser { return s.parser }

func (s *LayerDefContext) UP() antlr.TerminalNode {
	return s.GetToken(ConfigParserUP, 0)
}

func (s *LayerDefContext) MIDDLE() antlr.TerminalNode {
	return s.GetToken(ConfigParserMIDDLE, 0)
}

func (s *LayerDefContext) DOWN() antlr.TerminalNode {
	return s.GetToken(ConfigParserDOWN, 0)
}

func (s *LayerDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayerDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) LayerDef() (localctx ILayerDefContext) {
	this := p
	_ = this

	localctx = NewLayerDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ConfigParserRULE_layerDef)
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
		p.SetState(132)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserDOWN)|(1<<ConfigParserMIDDLE)|(1<<ConfigParserUP))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICornerContext is an interface to support dynamic dispatch.
type ICornerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCornerContext differentiates from other interfaces.
	IsCornerContext()
}

type CornerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCornerContext() *CornerContext {
	var p = new(CornerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_corner
	return p
}

func (*CornerContext) IsCornerContext() {}

func NewCornerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CornerContext {
	var p = new(CornerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_corner

	return p
}

func (s *CornerContext) GetParser() antlr.Parser { return s.parser }

func (s *CornerContext) WORD() antlr.TerminalNode {
	return s.GetToken(ConfigParserWORD, 0)
}

func (s *CornerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CornerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Corner() (localctx ICornerContext) {
	this := p
	_ = this

	localctx = NewCornerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ConfigParserRULE_corner)

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
		p.SetState(134)
		p.Match(ConfigParserWORD)
	}

	return localctx
}

// IEdgesContext is an interface to support dynamic dispatch.
type IEdgesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgesContext differentiates from other interfaces.
	IsEdgesContext()
}

type EdgesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgesContext() *EdgesContext {
	var p = new(EdgesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_edges
	return p
}

func (*EdgesContext) IsEdgesContext() {}

func NewEdgesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgesContext {
	var p = new(EdgesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_edges

	return p
}

func (s *EdgesContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgesContext) EDGES() antlr.TerminalNode {
	return s.GetToken(ConfigParserEDGES, 0)
}

func (s *EdgesContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *EdgesContext) AllEdgeLayer() []IEdgeLayerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEdgeLayerContext); ok {
			len++
		}
	}

	tst := make([]IEdgeLayerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEdgeLayerContext); ok {
			tst[i] = t.(IEdgeLayerContext)
			i++
		}
	}

	return tst
}

func (s *EdgesContext) EdgeLayer(i int) IEdgeLayerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgeLayerContext); ok {
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

	return t.(IEdgeLayerContext)
}

func (s *EdgesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Edges() (localctx IEdgesContext) {
	this := p
	_ = this

	localctx = NewEdgesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ConfigParserRULE_edges)
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
		p.SetState(136)
		p.Match(ConfigParserEDGES)
	}
	{
		p.SetState(137)
		p.Match(ConfigParserCOLON)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserDOWN)|(1<<ConfigParserMIDDLE)|(1<<ConfigParserUP))) != 0) {
		{
			p.SetState(138)
			p.EdgeLayer()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdgeLayerContext is an interface to support dynamic dispatch.
type IEdgeLayerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeLayerContext differentiates from other interfaces.
	IsEdgeLayerContext()
}

type EdgeLayerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeLayerContext() *EdgeLayerContext {
	var p = new(EdgeLayerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_edgeLayer
	return p
}

func (*EdgeLayerContext) IsEdgeLayerContext() {}

func NewEdgeLayerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeLayerContext {
	var p = new(EdgeLayerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_edgeLayer

	return p
}

func (s *EdgeLayerContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeLayerContext) LayerDef() ILayerDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILayerDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILayerDefContext)
}

func (s *EdgeLayerContext) COLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserCOLON, 0)
}

func (s *EdgeLayerContext) AllEdge() []IEdgeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEdgeContext); ok {
			len++
		}
	}

	tst := make([]IEdgeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEdgeContext); ok {
			tst[i] = t.(IEdgeContext)
			i++
		}
	}

	return tst
}

func (s *EdgeLayerContext) Edge(i int) IEdgeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgeContext); ok {
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

	return t.(IEdgeContext)
}

func (s *EdgeLayerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeLayerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) EdgeLayer() (localctx IEdgeLayerContext) {
	this := p
	_ = this

	localctx = NewEdgeLayerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ConfigParserRULE_edgeLayer)
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
		p.SetState(143)
		p.LayerDef()
	}
	{
		p.SetState(144)
		p.Match(ConfigParserCOLON)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserWORD {
		{
			p.SetState(145)
			p.Edge()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdgeContext is an interface to support dynamic dispatch.
type IEdgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeContext differentiates from other interfaces.
	IsEdgeContext()
}

type EdgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeContext() *EdgeContext {
	var p = new(EdgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_edge
	return p
}

func (*EdgeContext) IsEdgeContext() {}

func NewEdgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeContext {
	var p = new(EdgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_edge

	return p
}

func (s *EdgeContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeContext) WORD() antlr.TerminalNode {
	return s.GetToken(ConfigParserWORD, 0)
}

func (s *EdgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ConfigParser) Edge() (localctx IEdgeContext) {
	this := p
	_ = this

	localctx = NewEdgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ConfigParserRULE_edge)

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
		p.SetState(150)
		p.Match(ConfigParserWORD)
	}

	return localctx
}
