grammar Calculator;

// 顶层规则：包含服务调用链和条件表达式
rule
    : serviceChain ARROW expr EOF    # CompleteRule
    ;

// 服务调用链定义
serviceChain
    : service (ARROW service)*       # ServiceSequence
    ;

// 单个服务定义
service
    : serviceType '(' STRING (',' parameter)* ')'  # ServiceCall
    ;

// 服务类型
serviceType
    : 'http'
    | 'rpc'
    | 'db'
    | 'cache'
    ;

// 参数定义，支持引用之前服务的结果
parameter
    : IDENTIFIER '=' (
        STRING                           # StringParam
        | NUMBER                         # NumberParam
        | BOOLEAN                        # BooleanParam
        | array                          # ArrayParam
        | '$' IDENTIFIER ('.' IDENTIFIER)*  # RefParam
    )
    ;

// 表达式规则
expr
    : expr op=('*'|'/') expr                      # MulDiv
    | expr op=('+'|'-') expr                      # AddSub
    | expr op=('&&'|'||') expr                    # LogicalOp
    | expr op=('>'|'<'|'>='|'<='|'=='|'!=') expr # CompareOp
    | '$' IDENTIFIER ('.' IDENTIFIER)*            # ResultRef
    | NUMBER                                      # Number
    | STRING                                      # String
    | BOOLEAN                                     # Boolean
    | array                                      # ArrayExpr
    | '(' expr ')'                               # Parens
    ;

// 数组规则
array
    : '[' (expr (',' expr)*)? ']'
    ;

// 词法规则
NUMBER: '-'? [0-9]+ ('.' [0-9]+)?;
STRING: '"' (ESC|.)*? '"';
fragment ESC: '\\"' | '\\\\';
BOOLEAN: 'true' | 'false';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
ARROW: '->';

// 运算符和分隔符
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
AND: '&&';
OR: '||';
GT: '>';
LT: '<';
GE: '>=';
LE: '<=';
EQ: '==';
NEQ: '!=';
LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';
COMMA: ',';
DOT: '.';
EQUALS: '=';
DOLLAR: '$';

// 忽略空白字符和注释
WS: [ \t\r\n]+ -> skip;
COMMENT: '//' .*? '\r'? '\n' -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;