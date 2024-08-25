package analyzer

import "fmt"

type DuplicateHostEntry struct {
	AlreadyFoundAt uint32
	Hostname       string
}

func (d DuplicateHostEntry) Error() string {
	return fmt.Sprintf("'%s' already defined on line %d", d.Hostname, d.AlreadyFoundAt)
}

type DuplicateIPDeclaration struct {
	AlreadyFoundAt uint32
}

func (d DuplicateIPDeclaration) Error() string {
	return fmt.Sprintf("This IP address is already defined on line %d", d.AlreadyFoundAt)
}
