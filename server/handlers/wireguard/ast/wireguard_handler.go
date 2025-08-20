package ast

func (c *WGConfig) IncludesHeader(headerName string) bool {
	for _, section := range c.Sections {
		if section.Header.Name == headerName {
			return true
		}
	}

	return false
}
