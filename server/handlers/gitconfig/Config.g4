grammar Config;

lineStatement
    : entry EOF
    ;

entry
    : WHITESPACE? key? separator? value? leadingComment?
    ;

leadingComment
    : HASH WHITESPACE? (string WHITESPACE?)+
    ;

key
    : string
    ;

separator
    : EQUAL
    ;

value
    : string
    ;

string
    : (QUOTED_STRING | STRING)
    ;

///////////////////////////////////////////////

EQUAL
    : '='
    ;

HASH
    : '#'
    ;

WHITESPACE
    : [ \t]+
    ;

STRING
    : ~('#' | '\r' | '\n' | '"' | ' ' | '\t')+
    ;

QUOTED_STRING
    : '"' WHITESPACE? (STRING WHITESPACE)* STRING? ('"')?
    ;
