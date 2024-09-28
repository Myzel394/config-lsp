grammar Match;

root
    : matchEntry? (WHITESPACE matchEntry?)* EOF
    ;

matchEntry
    : entrySingle | entryWithValue
    ;

entrySingle
    : criteriaSingle
    ;

entryWithValue
    : criteriaWithValue separator? values?
    ;

separator
    : WHITESPACE
    ;

values
    : value? (COMMA value?)*
    ;

value
    : string
    ;

criteriaSingle
    : QUOTE? (ALL | CANONICAL | FINAL) QUOTE?
    ;

criteriaWithValue
    : QUOTE? (EXEC | LOCALNETWORK | HOST | ORIGINALHOST | TAGGED | USER | LOCALUSER) QUOTE?
    ;

string
    : (QUOTED_STRING | STRING)
    ;

COMMA
    : ','
    ;

ALL
    : ('a' | 'A') ('l' | 'L') ('l' | 'L')
    ;

CANONICAL
    : ('c' | 'C') ('a' | 'A') ('n' | 'N') ('o' | 'O') ('n' | 'N') ('i' | 'I') ('c' | 'C') ('a' | 'A') ('l' | 'L')
    ;

FINAL
    : ('f' | 'F') ('i' | 'I') ('n' | 'N') ('a' | 'A') ('l' | 'L')
    ;

EXEC
    : ('e' | 'E') ('x' | 'X') ('e' | 'E') ('c' | 'C')
    ;

LOCALNETWORK
    : ('l' | 'L') ('o' | 'O') ('c' | 'C') ('a' | 'A') ('l' | 'L') ('n' | 'N') ('e' | 'E') ('t' | 'T') ('w' | 'W') ('o' | 'O') ('r' | 'R') ('k' | 'K')
    ;

HOST
    : ('h' | 'H') ('o' | 'O') ('s' | 'S') ('t' | 'T')
    ;

ORIGINALHOST
    : ('o' | 'O') ('r' | 'R') ('i' | 'I') ('g' | 'G') ('i' | 'I') ('n' | 'N') ('a' | 'A') ('l' | 'L') ('h' | 'H') ('o' | 'O') ('s' | 'S') ('t' | 'T')
    ;

TAGGED
    : ('t' | 'T') ('a' | 'A') ('g' | 'G') ('g' | 'G') ('e' | 'E') ('d' | 'D')
    ;

USER
    : ('u' | 'U') ('s' | 'S') ('e' | 'E') ('r' | 'R')
    ;

LOCALUSER
    : ('l' | 'L') ('o' | 'O') ('c' | 'C') ('a' | 'A') ('l' | 'L') ('u' | 'U') ('s' | 'S') ('e' | 'E') ('r' | 'R')
    ;

STRING
    : ~(' ' | '\t' | '\r' | '\n' | '#' | ',')+
    ;

WHITESPACE
    : [ \t]+
    ;

QUOTED_STRING
    : QUOTE WHITESPACE? (STRING WHITESPACE)* STRING? QUOTE?
    ;

QUOTE
    : '"'
    ;
