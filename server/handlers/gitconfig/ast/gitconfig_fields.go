package ast

func (c *GitConfig) Clear() {
	c.Sections = []*GitSection{}
}

