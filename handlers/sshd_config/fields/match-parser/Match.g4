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
    : (USER | GROUP | HOST| LOCALADDRESS | LOCALPORT | RDOMAIN | ADDRESS)
    ;

values
    : value (COMMA value?)*
    ;

value
    : STRING
    ;

USER
    : ('U'|'u') ('S'|'s') ('E'|'e') ('R'|'r')
    ;

GROUP
    : ('G'|'g') ('R'|'r') ('O'|'o') ('U'|'u') ('P'|'p')
    ;

HOST
    : ('H'|'h') ('O'|'o') ('S'|'s') ('T'|'t')
    ;

LOCALADDRESS
    : ('L'|'l') ('O'|'o') ('C'|'c') ('A'|'a') ('L'|'l') ('A'|'a') ('D'|'d') ('D'|'d') ('R'|'r') ('E'|'e') ('S'|'s') ('S'|'s')
    ;

LOCALPORT
    : ('L'|'l') ('O'|'o') ('C'|'c') ('A'|'a') ('L'|'l') ('P'|'p') ('O'|'o') ('R'|'r') ('T'|'t')
    ;

RDOMAIN
    : ('R'|'r') ('D'|'d') ('O'|'o') ('M'|'m') ('A'|'a') ('I'|'i') ('N'|'n')
    ;

ADDRESS
    : ('A'|'a') ('D'|'d') ('D'|'d') ('R'|'r') ('E'|'e') ('S'|'s') ('S'|'s')
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
