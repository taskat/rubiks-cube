#! /bin/bash
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Create folder for java files
mkdir -p temp
# Copy grammar files
cp grammars/ConfigParser.g4 grammars/ConfigLexer.g4 temp
cd temp
# Set ANTLR folder
ANTLR="/d/antlr"
# Generate lexer files (without output)
../scripts/antlr4.sh $ANTLR ConfigLexer.g4 >/dev/null
# Run antlr to generate the java source files (without output)
../scripts/antlr4.sh $ANTLR ConfigParser.g4 >/dev/null
# Compile the generated java files
javac -cp ".:$ANTLR/antlr-4.12.0-complete.jar" *.java
echo "Start testing:"
# Start grunt
cat | ../scripts/grun.sh $ANTLR Config $1 $2
#Clean up the java files
cd ..
rm temp -r