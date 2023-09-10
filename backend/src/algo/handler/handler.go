package algohandler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ev "github.com/taskat/rubiks-cube/src/algo/errorvisitor"
	"github.com/taskat/rubiks-cube/src/algo/generatorvisitor"
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
	puzzle       models.Puzzle
}

func NewHandler(fileName, content string, errorHandler *eh.Errorhandler, puzzle models.Puzzle) *Handler {
	return &Handler{fileName: fileName, content: content, errorHandler: errorHandler, puzzle: puzzle}
}

func (h *Handler) buildTree(stream *antlr.CommonTokenStream) antlr.ParseTree {
	parser := ap.NewAlgorithmParser(stream)
	el.AddCustomErrorListener(parser, h.fileName, h.errorHandler)
	parser.BuildParseTrees = true
	return parser.AlgorithmFile()
}

func (h *Handler) checkTree() {
	visitor := ev.NewVisitor(h.fileName, h.errorHandler, h.puzzle.GetConstraint())
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

func Handle(fileName, content string, errorHandler *eh.Errorhandler, puzzle models.Puzzle) algorithm.Algorithm {
	handler := NewHandler(fileName, content, errorHandler, puzzle)
	handler.readConfig()
	v := generatorvisitor.NewVisitor(puzzle)
	return v.Visit(handler.tree)
}

func (h *Handler) readConfig() {
	input := h.getFileStream()
	stream := h.getTokens(input)
	h.tree = h.buildTree(stream)
	h.checkTree()
}
