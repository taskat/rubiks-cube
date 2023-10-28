parser grammar ConfigParser;

options { tokenVocab=ConfigLexer; }

configFile: configLine* EOF;
configLine: puzzleTypeDef | sizeDef | stateDescriptionDef | stateDef;

// Configlines
puzzleTypeDef: PUZZLE COLON puzzleType;
puzzleType: CUBE;

sizeDef: SIZE COLON NUMBER;

stateDescriptionDef: STATE_DESCRIPTION COLON stateDescription;
stateDescription: BEGINNER | ADVANCED;

stateDef: STATE COLON state;
state: RANDOM | beginnerState | advancedState;

beginnerState: side+;
side: sideDef COLON color+;
sideDef: FRONT | BACK | LEFT | RIGHT | UP | DOWN;
color: WORD;

advancedState: corners edges | edges corners | corners;
corners: CORNERS COLON cornerLayer+;
cornerLayer: layerDef COLON corner+;
layerDef: UP | MIDDLE | DOWN;
corner: WORD;
edges: EDGES COLON edgeLayer+;
edgeLayer: layerDef COLON edge+;
edge: WORD;