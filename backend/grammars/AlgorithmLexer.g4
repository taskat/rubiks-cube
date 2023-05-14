lexer grammar AlgorithmLexer;

// Keywords
ALL: 'all';
ANY: 'any';
BRANCHES: 'branches';
CONSECUTIVE: 'consecutive';
DO: 'do';
GOAL: 'goal';
HELPERS: 'helpers';
IF: 'if';
NONE: 'none';
PIECE: 'piece';
POS: 'pos';
POSITION: 'position';
PREPARE: 'prepare';
RUNS: 'runs';
STEP: 'step';
STEPS: 'steps';

// Punctuation
COLON: ':';
COMMA: ',';
QUESTIONMARK: '?';
PRIME: '\'';

// Parentheses
LPAREN: '(';
RPAREN: ')';
LBRACKET: '[';
RBRACKET: ']';

// OPERATORS
AND: '&' | 'and';
OR: '|' | 'or';
NOT: '!' | 'not';

// Numbers
NUMBER: '0' | NONZERO DIGIT*;
fragment DIGIT: [0-9];
fragment NONZERO: [1-9];

// Words
WORD: LETTER (LETTER | '_')*;
fragment LETTER: [a-zA-Z];

// Comments
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Whitespaces
WS: [ \t\r\n] -> skip;

