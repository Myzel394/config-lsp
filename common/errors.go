package common

type LSPError struct {
	Range LocationRange
	Err   error
}

type SyntaxError struct {
	Message string
}

func (s SyntaxError) Error() string {
	return s.Message
}
