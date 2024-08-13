package wireguard

type malformedLineError struct{}

func (e *malformedLineError) Error() string {
	return "Malformed line"
}

type lineError struct {
	Line uint32
	Err  error
}
