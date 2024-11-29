grammar Fstab;

entry
    :
        WHITESPACE? spec?
        WHITESPACE? mountPoint?
        WHITESPACE? fileSystem?
        WHITESPACE? mountOptions?
        WHITESPACE? freq?
        WHITESPACE? pass? WHITESPACE?
        EOF
    ;

spec
    : QUOTED_STRING | STRING
    ;

mountPoint
    : QUOTED_STRING | STRING
    ;

fileSystem
    : STRING | QUOTED_STRING
    ;

mountOptions
    : QUOTED_STRING | STRING
    ;

freq
    : DIGITS
    ;

pass
    : DIGITS
    ;

DIGITS
    : [0-9]+
    ;

WHITESPACE
    : [ \t]+
    ;

HASH
    : '#'
    ;

STRING
    :  ~(' ' | '\t' | '#')+
    ;

QUOTED_STRING
    : '"' WHITESPACE? (STRING WHITESPACE)* STRING? ('"')?
    ;

