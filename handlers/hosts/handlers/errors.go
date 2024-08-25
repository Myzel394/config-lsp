package handlers

import "fmt"

type DuplicateHostEntry struct {
	AlreadyFoundAt uint32
	Hostname       string
}

func (d DuplicateHostEntry) Error() string {
	return fmt.Sprintf("'%s' already defined on line %d", d.Hostname, d.AlreadyFoundAt)
}
