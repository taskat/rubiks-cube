parser grammar AlgorithmParser;

options { tokenVocab=AlgorithmLexer; }

algorithmFile: helpers? steps EOF;
helpers: HELPERS COLON helperLine+;
helperLine: WORD COLON algorithm;
steps: STEPS COLON step*;
step: STEP WORD COLON stepLine*;
stepLine: goal | helpers | runs | branches | doDef;

goal: GOAL COLON boolExpr;
runs: RUNS COLON NUMBER;
branches: BRANCHES COLON branch+;
doDef: DO COLON algorithm;

branch: ifBranch | prepareBranch;
ifBranch: IF boolExpr COLON doDef;
prepareBranch: PREPARE COLON ((doDef consecutive) | algorithm);
consecutive: CONSECUTIVE COLON NUMBER;

algorithm: turn+;
turn: WORD (NUMBER | PRIME)? | NUMBER LPAREN algorithm RPAREN;

boolExpr: unaryOp boolExpr | boolExpr binaryOp boolExpr | LPAREN boolExpr RPAREN | expr;
unaryOp: NOT;
binaryOp: AND | OR;
expr: unaryExpr | binaryExpr | functionalExpr;
unaryExpr: WORD LPAREN parameter RPAREN;
binaryExpr: parameter WORD parameter;
functionalExpr: function LPAREN expr COMMA list RPAREN;
function: ALL | ANY | NONE;
parameter: singleNode | node | piece | position | coord | list | QUESTIONMARK;

singleNode: sides NUMBER?;
node: LPAREN singleNode RPAREN;
piece: PIECE node;
position: (POS | POSITION) node;
coord: WORD NUMBER NUMBER;
list: LBRACKET ((node (COMMA node)*) | (piece (COMMA piece)*) | (position (COMMA position)* | (coord (COMMA coord)*)) ) RBRACKET;

sides: side (COMMA side)*;
side: WORD;

