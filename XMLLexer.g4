lexer grammar XMLLexer;

OPEN : '<'  -> pushMode(INSIDE);
COMMENT : '<!--' .*? '-->' -> skip;
EntityRef : '&' [a-z]+ ';';
TEXT : ~('<'|'&')+; // 匹配任意除< &之外的字符

//---------
mode INSIDE;
CLOSE : '>' -> popMode;
SLASH_CLOSE : '/>' -> popMode;
EQUALS : '=' ;
STRING : '"' .*? '"';
SlashName : '/' Name ;
Name : ALPHA (ALPHA | DIGIT)*;
S :[\t\r\n] -> skip;

fragment
ALPHA :[a-zA-Z];

fragment
DIGIT :[0-9];
