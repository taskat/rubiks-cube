parser grammar ConfigParser;

options { tokenVocab=ConfigLexer; }

configFile : puzzleTypeDef;
puzzleTypeDef : PUZZLE COLON puzzleType;
puzzleType: CUBE;