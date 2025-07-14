package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/bitcoin_conf/ast"
	"config-lsp/parsers/ini"
	"errors"
	"fmt"
)

func CreateIndexes(config *ast.BTCConfig) (*BTCIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := &BTCIndexes{
		SectionsByName:    make(map[string]*ini.Section),
		UnknownProperties: make(map[uint32]BTCIndexPropertyInfo),
	}

	for _, section := range config.Sections {
		if section.Header == nil {
			// Root section
			continue
		}

		sectionName := section.Header.Name

		if existingSection, found := indexes.SectionsByName[sectionName]; found {
			errs = append(errs, common.LSPError{
				Range: section.Header.LocationRange,
				Err:   errors.New(fmt.Sprintf("Section '%s' already defined at line %d", sectionName, existingSection.Header.Start.Line+1)),
			})

			continue
		}
		indexes.SectionsByName[sectionName] = section

		// Unknown properties will be checked later in `analyzeProperties`
	}

	return indexes, errs
}
