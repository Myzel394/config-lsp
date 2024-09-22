grammar Match;

root
    : matchEntry? (WHITESPACE matchEntry?)* EOF
    ;

matchEntry
    : criteria separator? values?
    ;

separator
    : WHITESPACE
    ;

criteria
    : string
    ;

values
    : value? (COMMA value?)*
    ;

value
    : string
    ;

string
    : (QUOTED_STRING | STRING)
    ;

COMMA
    : ','
    ;

STRING
    : ~(' ' | '\t' | '\r' | '\n' | '#' | ',')+
    ;

WHITESPACE
    : [ \t]+
    ;

QUOTED_STRING
    : '"' WHITESPACE? (STRING WHITESPACE)* STRING? ('"')?
    ;
