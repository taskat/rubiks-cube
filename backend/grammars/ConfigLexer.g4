lexer grammar ConfigLexer;

// Keywords
ADVANCED: 'advanced';
BACK: 'Back';
BEGINNER: 'beginner';
CORNERS: 'corners';
CUBE: 'cube';
DOWN: 'Down';
EDGES: 'edges';
FRONT: 'Front';
LEFT: 'Left';
MIDDLE: 'Middle';
PUZZLE: 'puzzle';
RANDOM: 'random';
RIGHT: 'Right';
SIZE: 'size';
STATE: 'state';
STATE_DESCRIPTION: 'state description';
UP: 'Up';

// Parenthesis
LCURLY: '{';
RCURLY: '}';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';

// Punctuation
COLON: ':';

// Numbers
NUMBER: [1-9][0-9]*;

// Strings
WORD: LETTER+;
fragment LETTER: [a-zA-Z];


// Comments
LINE_COMMENT: '//' ~[\r\n]* -> skip;

// Whitespaces
WS: [ \t\r\n] -> skip;
