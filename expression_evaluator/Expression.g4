grammar Expression;

OR: '||';
AND: '&&';

GT: '>';
GTE: '>=';
LT: '<';
LTE: '<=';
EQ: '==';
NEQ: '!=';
IN: 'in';
NOTIN: 'not in';

NOT: '!';


MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

LPAREN: '(';
RPAREN: ')';


NUMBER: [0-9]+;
TRUE: 'true';
FALSE: 'false';



WS: [ \r\n\t]+ -> skip;



IDENTIFIER: [a-zA-Z]+[a-zA-Z0-9]*;

start: expression EOF;

expression
        :mathExpression op=('>' | '<' | '<=' | '>=' | '!=' | '==') mathExpression # MathExpressionCompare
        | NOT expression                    # Not
        | expression AND expression             # And
        | expression OR expression              # Or
        | TRUE                           # True
        | FALSE                        # False
        | expression op=(EQ | NEQ) expression # Eq
        | LPAREN expression RPAREN                 # Bracket
        | IDENTIFIER                         # Identifier
        ;

mathExpression
        : mathExpression op=(MUL | DIV) mathExpression # MulDiv
        | mathExpression op=(ADD | SUB) mathExpression # AddSub
        | SUB mathExpression # Minus
        | NUMBER                             # Number
        | LPAREN mathExpression RPAREN                 # MathBracket
        | IDENTIFIER                         # MathIdentifier
        ;


