grammar UserProfileLang;

// Entry point
script: statement+ ;

// Statements
statement
    : printStatement
    | userProfileDeclaration
    ;

// Print statement
printStatement: 'Print' '(' expression ')' ;

// User profile declaration
userProfileDeclaration
    : 'Controller.UserProfile' '->' userProfileBody
    ;

// Body of user profile controller
userProfileBody: userProfileMethod+ ;

// Methods inside user profile
userProfileMethod
    : 'BeforeAction' '->' '[' actionList ']'
    | methodDeclaration '->' methodBody
    ;

// List of actions
actionList: action (',' action)* ;

// Action
action: IDENTIFIER ;

// Method declaration
methodDeclaration: IDENTIFIER methodParams ;

// Method parameters
methodParams: '(' param (',' param)* ')' ;

// Method parameter
param: IDENTIFIER ':' type ;

// Method body
methodBody
    : '{' expression '}' // single expression
    | conditionalExpression  // conditional expression
    ;

// Conditional expression
conditionalExpression
    : expression '?' expression ':' expression
    ;

// Type
type: 'int' | 'Post[]' ;

// Expressions
expression
    : IDENTIFIER
    | string
    | number
    | expression '+' expression
    | expression '=' expression
    | expression '/' expression
    | expression '*' expression
    | '(' expression ')'
    ;

// Identifiers
IDENTIFIER: [a-zA-Z_] [a-zA-Z_0-9]* ;

// String literals
string: '"' .*? '"' ;

// Number literals
number: INT | FLOAT ;

// Tokens
INT: [0-9]+ ;
FLOAT: [0-9]+ '.' [0-9]+ ;
WS: [ \t\r\n]+ -> skip;