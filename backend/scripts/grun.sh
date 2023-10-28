#! /bin/bash
# Get the first parameter
ANTLR=$1
# Shift all other parameters down 1
shift
# Call antlr
java -Xmx500M -cp ".:$ANTLR/antlr-4.12.0-complete.jar" org.antlr.v4.gui.TestRig $@