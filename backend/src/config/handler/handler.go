package confighandler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	ev "github.com/taskat/rubiks-cube/src/config/errorvisitor"
	gv "github.com/taskat/rubiks-cube/src/config/generatorvisitor"
	cl "github.com/taskat/rubiks-cube/src/config/lexer"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	el "github.com/taskat/rubiks-cube/src/errorlistener"
	"github.com/taskat/rubiks-cube/src/models"
)

type Handler struct {
	fileName string
	content  string
	tree     antlr.ParseTree
}

func NewHandler(fileName, content string) *Handler {
	return &Handler{fileName: fileName, content: content}
}

func (h *Handler) buildTree(stream *antlr.CommonTokenStream) antlr.ParseTree {
	parser := cp.NewConfigParser(stream)
	el.AddCustomErrorListener(parser, h.fileName)
	parser.BuildParseTrees = true
	return parser.ConfigFile()
}

func (h *Handler) checkTree() {
	visitor := ev.NewVisitor(h.fileName)
	visitor.Visit(h.tree)
}

func (h Handler) createCube() models.Puzzle {
	visitor := gv.NewVisitor()
	return visitor.Visit(h.tree)
}

func (h *Handler) getTokens(input *antlr.InputStream) *antlr.CommonTokenStream {
	lexer := cl.NewConfigLexer(input)
	el.AddCustomErrorListener(lexer, h.fileName)
	return antlr.NewCommonTokenStream(lexer, 0)
}

func Handle(fileName, content string) models.Puzzle {
	handler := NewHandler(fileName, content)
	handler.readConfig()
	if eh.HasErrors() {
		return nil
	}
	return handler.createCube()
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
