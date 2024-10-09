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
    : ADFS | AFFS | BTRFS | EXFAT
    // Still match unknown file systems
    | STRING | QUOTED_STRING
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

// ///// Supported file systems /////

ADFS
    : ('A' | 'a') ('D' | 'd') ('F' | 'f') ('S' | 's')
    ;

AFFS
    : ('A' | 'a') ('F' | 'f') ('F' | 'f') ('S' | 's')
    ;

BTRFS
    : ('B' | 'b') ('T' | 't') ('R' | 'r') ('F' | 'f') ('S' | 's')
    ;

EXFAT
    : ('E' | 'e') ('X' | 'x') ('F' | 'f') ('A' | 'a') ('T' | 't')
    ;
