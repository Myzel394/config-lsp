grammar Config;

lineStatement
    : (entry | (WHITESPACE? leadingComment) | WHITESPACE?) EOF
    ;

entry
    : WHITESPACE? key? WHITESPACE? value? WHITESPACE? leadingComment?
    ;

key
    : STRING
    ;

value
    : (STRING WHITESPACE)? STRING
    ;

leadingComment
    : HASH WHITESPACE? (STRING WHITESPACE?)+
    ;

HASH
    : '#'
    ;

WHITESPACE
    : [ \t]+
    ;

STRING
    : ~(' ' | '\t' | '\r' | '\n' | '#')+
    ;

NEWLINE
    : '\r'? '\n'
    ;
