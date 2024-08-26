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
    : DOMAIN
    ;

hostname
    : domain
    ;

domain
    : DOMAIN
    ;

ipAddress
    : (ipv4Address | ipv6Address)
    ;

ipv4Address
    : singleIPv4Address
        // Allow optional range to tell user ranges are not allowed
        ipRange?
    ;

singleIPv4Address
    : ipv4Digit DOT ipv4Digit DOT ipv4Digit DOT ipv4Digit
    ;

// This is not correct but fits for now
ipv6Address
    : singleIPv6Address
        // Allow optional range to tell user ranges are not allowed
        ipRange?
    ;

singleIPv6Address
    : (ipv6Octet COLON)+ ipv6Octet
    ;

ipv4Digit
    : DIGITS
    ;

ipv6Octet
    : OCTETS
    ;

ipRange
    : SLASH ipRangeBits
    ;

ipRangeBits
    : DIGITS
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

DIGITS
    : DIGIT+
    ;

fragment DIGIT
    : [0-9]
    ;

OCTETS
    : OCTET+
    ;

fragment OCTET
    : [0-9a-fA-F]
    ;

DOMAIN
    : ((STRING)+ (DOT [a-zA-Z]+)*)
    ;

fragment STRING
    : ~(' ' | '\t' | '\n' | '\r' | '#' | '.')
    ;
