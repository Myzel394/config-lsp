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
    : STRING
    ;

values
    : value? (COMMA value?)*
    ;

value
    : STRING
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
