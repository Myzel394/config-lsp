package ast

func (o SSHOption) GetType() SSHEntryType {
	return SSHEntryTypeOption
}

func (o SSHOption) GetOption() SSHOption {
	return o
}

func (m SSHMatchBlock) GetType() SSHEntryType {
	return SSHEntryTypeMatchBlock
}

func (m SSHMatchBlock) GetOption() SSHOption {
	return *m.MatchEntry
}

func (c SSHConfig) FindMatchBlock(line uint32) *SSHMatchBlock {
	for currentLine := line; currentLine > 0; currentLine-- {
		rawEntry, found := c.Options.Get(currentLine)

		if !found {
			continue
		}

		switch entry := rawEntry.(type) {
		case *SSHMatchBlock:
			return entry
		}
	}

	return nil
}

func (c SSHConfig) FindOption(line uint32) (*SSHOption, *SSHMatchBlock) {
	matchBlock := c.FindMatchBlock(line)

	if matchBlock != nil {
		if line == matchBlock.MatchEntry.Start.Line {
			return matchBlock.MatchEntry, matchBlock
		}

		rawEntry, found := matchBlock.Options.Get(line)

		if found {
			return rawEntry.(*SSHOption), matchBlock
		} else {
			return nil, matchBlock
		}
	}

	rawEntry, found := c.Options.Get(line)

	if found {
		switch rawEntry.(type) {
		case *SSHMatchBlock:
			return rawEntry.(*SSHMatchBlock).MatchEntry, rawEntry.(*SSHMatchBlock)
		case *SSHOption:
			return rawEntry.(*SSHOption), nil
		}
	}

	return nil, nil

}
