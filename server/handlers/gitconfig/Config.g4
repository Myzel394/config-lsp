grammar Config;

lineStatement
    : WHITESPACE? entry WHITESPACE? leadingComment? EOF
    ;

entry
    : key? WHITESPACE? separator? WHITESPACE? value?
    ;

leadingComment
    : commentSymbol WHITESPACE? (string WHITESPACE?)+
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
    : (WHITESPACE? ((STRING | QUOTED_STRING) WHITESPACE)* (STRING | QUOTED_STRING))
    ;

commentSymbol
    : HASH | SEMICOLON
    ;

///////////////////////////////////////////////

EQUAL
    : '='
    ;

HASH
    : '#'
    ;

SEMICOLON
    : ';'
    ;

WHITESPACE
    : [ \t]+
    ;

STRING
    : ~('\r' | '\n' | '"' | ' ' | '\t')+
    ;

QUOTED_STRING
    : '"' WHITESPACE? (STRING WHITESPACE)* STRING? ('"')?
    ;
