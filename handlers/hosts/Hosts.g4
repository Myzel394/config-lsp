grammar Hosts;

lineStatement
    : SEPARATOR? entry SEPARATOR? leadingComment? EOF
    ;

entry
    : ipAddress SEPARATOR hostname (SEPARATOR aliases)?
    ;

aliases
    : (alias SEPARATOR?)+
    ;

alias
    : domain
    ;

hostname
    : domain
    ;

domain
    : (STRING)+ (DOT STRING*)*
    ;

ipAddress
    : (ipv4Address | ipv6Address)
    ;

ipv4Address
    : (STRING DOT)+ STRING (ipRange? | ipPort?)
    ;

ipv6Address
    : (((STRING COLON)+ STRING) | (STRING? COLON COLON STRING)) (ipRange? | ipPort?)
    ;

/*
ipv4Address
    : singleIPv4Address
        // Allow optional range to tell user ranges are not allowed
        (ipRange? | ipPort?)
    ;

singleIPv4Address
    : ipv4Digit DOT ipv4Digit DOT ipv4Digit DOT ipv4Digit
    ;

// This is not correct but fits for now
ipv6Address
    : singleIPv6Address
        // Allow optional range to tell user ranges are not allowed
        (ipRange? | ipPort?)
    ;

singleIPv6Address
    : (ipv6Octet COLON)+ ipv6Octet
    ;

ipv4Digit
    : STRING
    ;

ipv6Octet
    : STRING
    ;
*/

ipRange
    : SLASH ipRangeBits
    ;

ipRangeBits
    : STRING
    ;

ipPort
    : COLON STRING
    ;

comment
    : COMMENTLINE
    ;

leadingComment
    : COMMENTLINE
    ;

COMMENTLINE
    : HASHTAG ~[\r\n]+
    ;

SLASH
    : '/'
    ;

DOT
    : '.'
    ;

COLON
    : ':'
    ;

HASHTAG
    : '#'
    ;

SEPARATOR
    : [ \t]+
    ;

NEWLINE
    :  [\r]? [\n]
    ;

STRING
    : [a-zA-Z0-9_\-]+
    ;
