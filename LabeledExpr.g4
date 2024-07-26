grammar LabeledExpr;


@header{
import "strconv"
}
@members {
  cols int
}

prog : stat+ ;

stat: expr NEWLINE  # printExpr
      | ID ASSIGN expr NEWLINE # assign
      | CLEAR  NEWLINE #clear
      | NEWLINE      # blank;
expr : expr op= (MUL|DIV) expr #MulDiv
     | expr op= (ADD|SUB) expr #AddSub
     | INT                 #int
     | ID                  #id
     | '('  expr  ')'      #parens
     ;

MUL : '*';
DIV : '/';
ADD : '+';
SUB : '-';
ASSIGN:'=';
CLEAR: 'clear';




fragment A  : [aA] ;
fragment B  : [bB] ;
fragment C  : [cC] ;
fragment D  : [dD] ;
fragment E  : [eE] ;
fragment F  : [fF] ;
fragment G  : [gG] ;
fragment H  : [hH] ;
fragment I  : [iI] ;
fragment J  : [jJ] ;
fragment K  : [kK] ;
fragment L  : [lL] ;
fragment M  : [mM] ;
fragment N  : [nN] ;
fragment O  : [oO] ;
fragment P  : [pP] ;
fragment Q  : [qQ] ;
fragment R  : [rR] ;
fragment S  : [sS] ;
fragment T  : [tT] ;
fragment U  : [uU] ;
fragment V  : [vV] ;
fragment W  : [wW] ;
fragment X  : [xX] ;
fragment Y  : [yY] ;
fragment Z  : [zZ] ;

OPERATOR_LOGIC_STR
    :   A N D
    |   O R
    |   N O T
    ;

operator
    :   EQUALS
    |   NOT_EQUALS
    |   GT
    |   GTE
    |   LT
    |   LTE
    |   'IS_NULL'
    |   'STRING_START_WITH'
    |   'STRING_END_WITH'
    |   'STRING_CONTAINS'
    |   'LIST_IN'
    |   'LIST_CONTAINS'
    |   'LIST_RETAIN'
    |   'CONTAIN_REGULAR'
    |   'SUB_LIST_IN'
    |   'SUB_LIST_CONTAINS'
    |   'CONTAINS_KEYWORDS'
    |   'BETWEEN_ALL_CLOSE'
    |   'BETWEEN_ALL_OPEN'
    |   'BETWEEN_LEFT_OPEN_RIGHT_CLOSE'
    |   'BETWEEN_LEFT_CLOSE_RIGHT_OPEN'
    |   'SUB_COND'
    ;

NULL : 'null' ;

DOT         : '.' ;
COMMA       : ',' ;
COLON       : ':' ;
QUOTATION   : '"' ;
L_BRACKET   : '[' ;
R_BRACKET   : ']' ;
L_BRACE     : '{' ;
R_BRACE     : '}' ;

INT : [0-9]+ ;
EXP : [Ee] [+\-]? INT;
NAME_KEY_TYPE : [a-zA-Z][a-zA-Z0-9_#]* ;
SP_UNICODE : '\u0080'..'\uFFFF' ;

AND : '&&'  ;
OR  : '||'  ;
NOT : '!'   ;
MOD     : '%' ;

EQUALS      : '==' ;
NOT_EQUALS  : '!=' ;
GT          : '>'  ;
LT          : '<'  ;
GTE         : '>=' ;
LTE         : '<=' ;

// skip tab, newline, return
WS  :   [\t\n\r]+ -> skip ;


ID : [a-zA-Z]+;
NEWLINE: '\r' ? '\n';
