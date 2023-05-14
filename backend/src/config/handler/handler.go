package confighandler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ev "github.com/taskat/rubiks-cube/src/config/errorvisitor"
	gv "github.com/taskat/rubiks-cube/src/config/generatorvisitor"
	cl "github.com/taskat/rubiks-cube/src/config/lexer"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	el "github.com/taskat/rubiks-cube/src/errorlistener"
	"github.com/taskat/rubiks-cube/src/models"
)

type Handler struct {
	fileName     string
	content      string
	tree         antlr.ParseTree
	errorHandler eh.Errorhandler
}

func NewHandler(fileName, content string, errorHandler eh.Errorhandler) *Handler {
	return &Handler{fileName: fileName, content: content, errorHandler: errorHandler}
}

func (h *Handler) buildTree(stream *antlr.CommonTokenStream) antlr.ParseTree {
	parser := cp.NewConfigParser(stream)
	el.AddCustomErrorListener(parser, h.fileName, h.errorHandler)
	parser.BuildParseTrees = true
	return parser.ConfigFile()
}

func (h *Handler) checkTree() {
	visitor := ev.NewVisitor(h.fileName, h.errorHandler)
	visitor.Visit(h.tree)
}

func (h Handler) createCube() models.Puzzle {
	visitor := gv.NewVisitor()
	return visitor.Visit(h.tree)
}

func (h *Handler) getTokens(input *antlr.InputStream) *antlr.CommonTokenStream {
	lexer := cl.NewConfigLexer(input)
	el.AddCustomErrorListener(lexer, h.fileName, h.errorHandler)
	return antlr.NewCommonTokenStream(lexer, 0)
}

func Handle(fileName, content string, errorHandler eh.Errorhandler) models.Puzzle {
	handler := NewHandler(fileName, content, errorHandler)
	handler.readConfig()
	if handler.errorHandler.HasErrors() {
		return nil
	}
	cube := handler.createCube()
	validator := cube.GetValidator()
	errors := validator.Validate()
	for _, err := range errors {
		errorHandler.AddError(eh.NewContext(-1, -1), err, fileName)
	}
	return cube
}

func (h *Handler) readConfig() {
	input := h.getFileStream()
	stream := h.getTokens(input)
	h.tree = h.buildTree(stream)
	h.checkTree()
}

func (h *Handler) getFileStream() *antlr.InputStream {
	return antlr.NewInputStream(h.content)
}
