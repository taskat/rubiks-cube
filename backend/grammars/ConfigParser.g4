parser grammar ConfigParser;

options { tokenVocab=ConfigLexer; }

configFile : configLine* EOF;
configLine: puzzleTypedef | sizeDef | stateDescriptionDef | stateDef;

// Configlines
puzzleTypedef: PUZZLE COLON puzzleType;
puzzleType: CUBE;

sizeDef: SIZE COLON NUMBER;

stateDescriptionDef: STATE_DESCRIPTION COLON stateDescription;
stateDescription: BEGINNER | ADVANCED;

stateDef: STATE COLON state;
state: RANDOM | beginnerState | advancedState;

beginnerState: LCURLY side+ RCURLY;
side: sideDef COLON sideState;
sideDef: FRONT | BACK | LEFT | RIGHT | UP | DOWN;
sideState: LBRACKET sideStateRow+ RBRACKET;
sideStateRow: LBRACKET color+ RBRACKET;
color: STRING;

advancedState: LCURLY RCURLY;