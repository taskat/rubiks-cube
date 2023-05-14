parser grammar AlgorithmParser;

options { tokenVocab=AlgorithmLexer; }

algorithmFile: helpers steps;
helpers: HELPERS COLON helperLine+;
helperLine: WORD COLON algorithm;
steps: step*;
step: STEP WORD COLON stepLine*;
stepLine: goal | helpers | runs | branches | doDef;

goal: GOAL COLON boolExpr;
runs: RUNS COLON NUMBER;
branches: BRANCHES COLON branch+;
doDef: DO COLON algorithm;

branch: ifBranch | prepareBranch;
ifBranch: IF boolExpr COLON doDef;
prepareBranch: PREPARE COLON doDef consecutive?;
consecutive: CONSECUTIVE COLON NUMBER;

algorithm: turn+;
turn: WORD (NUMBER | PRIME)?;

boolExpr: unaryOp boolExpr | boolExpr binaryOp boolExpr | LPAREN boolExpr RPAREN | expr;
unaryOp: NOT;
binaryOp: AND | OR;
expr: unaryExpr | binaryExpr | functionalExpr;
unaryExpr: WORD LPAREN parameter RPAREN;
binaryExpr: parameter WORD parameter;
functionalExpr: function LPAREN expr COMMA list RPAREN;
function: ALL | ANY | NONE;
parameter: piece | position | coord | list | QUESTIONMARK;

piece: PIECE? LPAREN sides NUMBER? RPAREN;
position: (POS? | POSITION?) LPAREN sides NUMBER? RPAREN;
coord: WORD NUMBER NUMBER;
list: LBRACKET (piece+ | position+ | coord+) RBRACKET;

sides: side (COMMA side)*;
side: WORD;

