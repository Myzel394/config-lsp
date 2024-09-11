grammar Config;

lineStatement
    : (entry | (WHITESPACE? leadingComment) | WHITESPACE?) EOF
    ;

entry
    : WHITESPACE? key? separator? value? WHITESPACE? leadingComment?
    ;

separator
    : WHITESPACE
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
