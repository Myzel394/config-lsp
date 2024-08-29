grammar Aliases;

lineStatement
    : entry SEPARATOR? comment? EOF
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
    : (user | file | command | include | email)
    ;

user
    : STRING
    ;

file
    : (SLASH STRING)+ SLASH?
    ;

command
    : VERTLINE STRING
    ;

include
    : COLON INCLUDE COLON file
    ;

comment
    : NUMBER_SIGN (SEPARATOR? STRING)+ SEPARATOR?
    ;

email
    : STRING AT STRING
    ;

error
    : errorStatus COLON errorCode SEPARATOR errorMessage
    ;

errorStatus
    : STRING
    ;

errorCode
    : STRING
    ;

errorMessage
    : STRING
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
