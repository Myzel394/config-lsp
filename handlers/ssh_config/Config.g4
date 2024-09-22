grammar Config;

lineStatement
    : (entry | leadingComment | WHITESPACE?) EOF
    ;

entry
    : WHITESPACE? key? separator? value? leadingComment?
    ;

separator
    : WHITESPACE
    ;

key
    : string
    ;

value
    : (string WHITESPACE)* string? WHITESPACE?
    ;

leadingComment
    : HASH WHITESPACE? (string WHITESPACE?)+
    ;

string
    : (QUOTED_STRING | STRING)
    ;

///////////////////////////////////////////////

HASH
    : '#'
    ;

WHITESPACE
    : [ \t]+
    ;

STRING
    : ~('#' | '\r' | '\n' | '"' | ' ' | '\t')+
    ;

NEWLINE
    : '\r'? '\n'
    ;

QUOTED_STRING
    : '"' WHITESPACE? (STRING WHITESPACE)* STRING? ('"')?
    ;
