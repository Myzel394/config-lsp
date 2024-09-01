grammar Aliases;

lineStatement
    : entry EOF
    ;

entry
    : SEPARATOR? key SEPARATOR? separator SEPARATOR? values
    ;

separator
    : COLON
    ;

key
    : STRING
    ;

// // Values // //
values
    : (value COMMA SEPARATOR)* value
    ;

value
    : (user | file | command | include | email | error)
    ;

user
    : STRING
    ;

file
    : SLASH (STRING SLASH)* STRING?
    ;

command
    : VERTLINE STRING?
    ;

include
    : COLON INCLUDE COLON file?
    ;

comment
    : NUMBER_SIGN (SEPARATOR? STRING)+ SEPARATOR?
    ;

email
    : STRING AT STRING
    ;

error
    : ERROR COLON errorCode? SEPARATOR? errorMessage?
    ;

errorCode
    : DIGITS
    ;

errorMessage
    : STRING
    ;

DIGITS
    : [0-9]+
    ;

ERROR
    : 'e' 'r' 'r' 'o' 'r'
    ;

SEPARATOR
    : [ \t]+
    ;

AT
    : '@'
    ;

INCLUDE
    : 'i' 'n' 'c' 'l' 'u' 'd' 'e'
    ;

VERTLINE
    : '|'
    ;

COLON
    : ':'
    ;

COMMA
    : ','
    ;

NUMBER_SIGN
    : '#'
    ;

SLASH
    : '/'
    ;

STRING
    : ~(' ' | '\t' | '\n' | '\r' | ':' | ',' | '#' | '@' | '|' | '/')+
    ;
