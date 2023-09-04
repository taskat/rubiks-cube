package algohandler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ev "github.com/taskat/rubiks-cube/src/algo/errorvisitor"
	al "github.com/taskat/rubiks-cube/src/algo/lexer"
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	el "github.com/taskat/rubiks-cube/src/errorlistener"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
)

type Handler struct {
	fileName     string
	content      string
	tree         antlr.ParseTree
	errorHandler *eh.Errorhandler
	constraint   models.Constraint
}

func NewHandler(fileName, content string, errorHandler *eh.Errorhandler, constraint models.Constraint) *Handler {
	return &Handler{fileName: fileName, content: content, errorHandler: errorHandler, constraint: constraint}
}

func (h *Handler) buildTree(stream *antlr.CommonTokenStream) antlr.ParseTree {
	parser := ap.NewAlgorithmParser(stream)
	el.AddCustomErrorListener(parser, h.fileName, h.errorHandler)
	parser.BuildParseTrees = true
	return parser.AlgorithmFile()
}

func (h *Handler) checkTree() {
	visitor := ev.NewVisitor(h.fileName, h.errorHandler, h.constraint)
	visitor.Visit(h.tree)
}

func (h *Handler) getFileStream() *antlr.InputStream {
	return antlr.NewInputStream(h.content)
}

func (h *Handler) getTokens(input *antlr.InputStream) *antlr.CommonTokenStream {
	lexer := al.NewAlgorithmLexer(input)
	el.AddCustomErrorListener(lexer, h.fileName, h.errorHandler)
	return antlr.NewCommonTokenStream(lexer, 0)
}

func Handle(fileName, content string, errorHandler *eh.Errorhandler, cosntraint models.Constraint) algorithm.Algorithm {
	handler := NewHandler(fileName, content, errorHandler, cosntraint)
	handler.readConfig()
	return nil
}

func (h *Handler) readConfig() {
	input := h.getFileStream()
	stream := h.getTokens(input)
	h.tree = h.buildTree(stream)
	h.checkTree()
}
