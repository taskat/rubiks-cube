lexer grammar ConfigLexer;

// Keywords
ADVANCED: 'advanced';
BACK: 'back';
BEGINNER: 'beginner';
CUBE: 'cube';
DOWN: 'down';
FRONT: 'front';
LEFT: 'left';
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

// Punctuation
COLON: ':';

// Numbers
NUMBER: [1-9][0-9]*;

// Strings
LETTER: [a-zA-Z];
STRING: LETTER+;

// Whitespaces
WS: [ \t\r\n] -> skip;