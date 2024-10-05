package shared

import "fmt"

type DuplicateKeyEntry struct {
	AlreadyFoundAt uint32
	Key            string
}

func (d DuplicateKeyEntry) Error() string {
	return fmt.Sprintf("Alias '%s' already defined on line %d", d.Key, d.AlreadyFoundAt+1)
}
