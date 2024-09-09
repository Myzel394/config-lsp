grammar Config;

lineStatement
    : (entry | (WHITESPACE? leadingComment) | WHITESPACE?) EOF
    ;

entry
    : WHITESPACE? key? WHITESPACE value? WHITESPACE? leadingComment?
    ;

key
    : STRING
    ;

value
    : STRING
    ;

leadingComment
    : HASH WHITESPACE? (STRING WHITESPACE?)+
    ;


HASH
    : '#'
    ;

MATCH
    : ('M' | 'm') ('A' | 'a') ('T' | 't') ('C' | 'c') ('H' | 'h')
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
