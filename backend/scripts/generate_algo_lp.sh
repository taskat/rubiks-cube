#! /bin/bash
# Stop on error
set -e
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Set ANTLR folder
ANTLR="/d/antlr"
# Generate lexer
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package algolexer -o src/algo/lexer -no-listener -no-visitor grammars/AlgorithmLexer.g4
# Generate parser
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package algoparser -o src/algo/parser -lib src/algo/lexer -no-listener -no-visitor grammars/AlgorithmParser.g4
# Print success message
echo "Algorithm lexer and parser generated successfully"
# Unset stopping on error
set +e
