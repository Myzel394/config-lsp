package ast

func (c *GitConfig) Clear() {
	c.Sections = []*GitSection{}
	c.CommentLines = map[uint32]struct{}{}
}
