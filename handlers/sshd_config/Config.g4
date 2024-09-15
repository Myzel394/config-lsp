grammar Config;

lineStatement
    : (entry | (leadingComment) | WHITESPACE?) EOF
    ;

entry
    : WHITESPACE? key? separator? value? leadingComment?
    ;

separator
    : WHITESPACE
    ;

key
    : STRING
    ;

value
    : (STRING WHITESPACE)* STRING? WHITESPACE?
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
