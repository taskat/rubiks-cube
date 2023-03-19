#! /bin/bash
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Set ANTLR folder
ANTLR="/d/antlr"
# Generate lexer
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package lexer -o compiler/lexer compiler/GoLexer.g4
# Generate parser
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package parser -o compiler/parser -lib compiler/lexer -visitor compiler/GoParser.g4
