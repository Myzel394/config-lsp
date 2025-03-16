package ast

func (o SSHDOption) GetType() SSHDEntryType {
	return SSHDEntryTypeOption
}

func (o SSHDOption) GetOption() SSHDOption {
	return o
}

func (m SSHDMatchBlock) GetType() SSHDEntryType {
	return SSHDEntryTypeMatchBlock
}

func (m SSHDMatchBlock) GetOption() SSHDOption {
	return *m.MatchOption
}

func (c SSHDConfig) FindMatchBlock(line uint32) *SSHDMatchBlock {
	for currentLine := line; currentLine > 0; currentLine-- {
		rawEntry, found := c.Options.Get(currentLine)

		if !found {
			continue
		}

		switch entry := rawEntry.(type) {
		case *SSHDMatchBlock:
			return entry
		}
	}

	return nil
}

func (c SSHDConfig) FindOption(line uint32) (*SSHDOption, *SSHDMatchBlock) {
	matchBlock := c.FindMatchBlock(line)

	if matchBlock != nil {
		if line == matchBlock.MatchOption.Start.Line {
			return matchBlock.MatchOption, matchBlock
		}

		rawEntry, found := matchBlock.Options.Get(line)

		if found {
			return rawEntry.(*SSHDOption), matchBlock
		} else {
			return nil, matchBlock
		}
	}

	rawEntry, found := c.Options.Get(line)

	if found {
		switch rawEntry.(type) {
		case *SSHDMatchBlock:
			return rawEntry.(*SSHDMatchBlock).MatchOption, rawEntry.(*SSHDMatchBlock)
		case *SSHDOption:
			return rawEntry.(*SSHDOption), nil
		}
	}

	return nil, nil
}

func (c SSHDConfig) GetAllOptions() []SSHDOptionInfo {
	infos := make(
		[]SSHDOptionInfo,
		0,
		// Approximation, this does not need to be exact
		c.Options.Size()+10,
	)

	var currentMatchBlock *SSHDMatchBlock = nil

	for _, rawEntry := range c.Options.Values() {
		switch entry := rawEntry.(type) {
		case *SSHDOption:
			infos = append(infos, SSHDOptionInfo{
				Option:     entry,
				MatchBlock: currentMatchBlock,
			})
		case *SSHDMatchBlock:
			currentMatchBlock = entry

			infos = append(infos, SSHDOptionInfo{
				Option:     entry.MatchOption,
				MatchBlock: currentMatchBlock,
			})
		}
	}

	return infos
}
