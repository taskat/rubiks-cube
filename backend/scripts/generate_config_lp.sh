#! /bin/bash
# Stop on error
set -e
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Set ANTLR folder
ANTLR="/d/antlr"
# Generate lexer
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package configlexer -o src/config/lexer -no-listener -no-visitor grammars/ConfigLexer.g4
# Generate parser
scripts/antlr4.sh $ANTLR -Dlanguage=Go -package configparser -o src/config/parser -lib src/config/lexer -no-listener -no-visitor grammars/ConfigParser.g4
# Print success message
echo "Config lexer and parser generated successfully"
# Unset stopping on error
set +e
