lexer grammar ConfigLexer;

// Keywords
ADVANCED: 'advanced';
BACK: 'back';
BEGINNER: 'beginner';
CORNERS: 'corners';
CUBE: 'cube';
DOWN: 'down';
EDGES: 'edges';
FRONT: 'front';
LEFT: 'left';
MIDDLE: 'middle';
PUZZLE: 'puzzle';
RANDOM: 'random';
RIGHT: 'right';
SIZE: 'size';
STATE: 'state';
STATE_DESCRIPTION: 'state description';
UP: 'up';

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