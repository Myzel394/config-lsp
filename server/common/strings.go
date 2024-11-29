package common

var UnicodeWhitespace = map[rune]struct{}{
	'\u0020': {}, // Space
	'\u0009': {}, // Horizontal tab
	'\u000A': {}, // Line feed
	'\u000B': {}, // Vertical tab
	'\u000C': {}, // Form feed
	'\u000D': {}, // Carriage return
	'\u0085': {}, // Next line
	'\u00A0': {}, // No-break space
	'\u1680': {}, // Ogham space mark
	'\u2000': {}, // En quad
	'\u2001': {}, // Em quad
	'\u2002': {}, // En space
	'\u2003': {}, // Em space
	'\u2004': {}, // Three-per-em space
	'\u2005': {}, // Four-per-em space
	'\u2006': {}, // Six-per-em space
	'\u2007': {}, // Figure space
	'\u2008': {}, // Punctuation space
	'\u2009': {}, // Thin space
}
