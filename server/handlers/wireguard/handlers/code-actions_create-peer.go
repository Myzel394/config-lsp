package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard"
	wgcommands "config-lsp/handlers/wireguard/commands"
	"config-lsp/parsers/ini"
	"fmt"
	"net"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionCreatePeerArgs struct {
	URI  protocol.DocumentUri
	Line uint32
}

func CodeActionCreatePeerArgsFromArguments(arguments map[string]any) CodeActionCreatePeerArgs {
	return CodeActionCreatePeerArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		Line: uint32(arguments["Line"].(float64)),
	}
}

func (args CodeActionCreatePeerArgs) RunCommand(d *wireguard.WGDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	interfaceSection := d.Indexes.SectionsByName["Interface"][0]
	section := d.Config.FindSectionByLine(args.Line)

	label := fmt.Sprintf("Add Peer based on peer on line %d", args.Line)

	newSection := section

	// IP Address
	ipAddressLine, ipAddress := newSection.FindFirstPropertyByName("AllowedIPs")
	_, address := interfaceSection.FindFirstPropertyByName("Address")

	if ipAddress != nil && address != nil {
		_, network, err := net.ParseCIDR(address.Value.Value)

		if err == nil {
			newIPAddress := createNewIP(*network, ipAddress.Value.Value)

			valueEnd := common.Location{
				Line:      ipAddress.End.Line,
				Character: ipAddress.Value.Start.Character + uint32(len(newIPAddress)) + 1,
			}

			// Create the underlying ini.Property first
			iniProperty := &ini.Property{
				LocationRange: common.LocationRange{
					Start: ipAddress.Start,
					End:   valueEnd,
				},
				Key:       ipAddress.Key,
				RawValue:  newIPAddress,
				Separator: address.Separator,
				Value: &ini.PropertyValue{
					LocationRange: common.LocationRange{
						Start: ipAddress.Value.Start,
						End:   valueEnd,
					},
					Value: newIPAddress,
				},
			}

			// Then wrap it with a WGProperty
			newSection.Properties.Put(ipAddressLine, iniProperty)
		}
	}

	// Preshared Key
	presharedKeyLine, presharedKey := newSection.FindFirstPropertyByName("PresharedKey")

	if presharedKey != nil {
		var newKey string

		if wgcommands.AreWireguardToolsAvailable() {
			createdKey, err := wgcommands.CreatePresharedKey()

			if err == nil {
				newKey = createdKey
			}
		} else {
			newKey = "[preshared key]"
		}

		valueEnd := common.Location{
			Line:      presharedKey.End.Line,
			Character: presharedKey.Value.Start.Character + uint32(len(newKey)) + 1,
		}

		// Create the underlying ini.Property first
		iniProperty := &ini.Property{
			LocationRange: common.LocationRange{
				Start: presharedKey.Start,
				End:   valueEnd,
			},
			Key:       presharedKey.Key,
			RawValue:  newKey,
			Separator: presharedKey.Separator,
			Value: &ini.PropertyValue{
				LocationRange: common.LocationRange{
					Start: presharedKey.Value.Start,
					End:   valueEnd,
				},
				Value: newKey,
			},
		}

		// Then add it to the section
		newSection.Properties.Put(presharedKeyLine, iniProperty)
	}

	lastProperty := newSection.GetLastProperty()
	newText := writeSection(*newSection)

	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      lastProperty.End.Line,
								Character: lastProperty.End.Character,
							},
							End: protocol.Position{
								Line:      lastProperty.End.Line,
								Character: lastProperty.End.Character,
							},
						},
						NewText: newText,
					},
				},
			},
		},
	}, nil
}

func writeSection(section ini.Section) string {
	text := "\n\n"

	text += fmt.Sprintf("[%s]\n", section.Header.Name)

	it := section.Properties.Iterator()
	for it.Next() {
		property := it.Value().(*ini.Property)
		text += fmt.Sprintf("%s = %s\n", property.Key.Name, property.Value.Value)
	}

	return text
}

// Try incrementing the IP address
func createNewIP(
	network net.IPNet,
	rawIP string,
) string {
	parsedIP, _, err := net.ParseCIDR(rawIP)
	parsedIP = parsedIP.To4()

	if parsedIP == nil {
		// IPv6 is not supported
		return ""
	}

	if err != nil {
		return ""
	}

	lastAddress := uint32(network.IP[0])<<24 | uint32(network.IP[1])<<16 | uint32(network.IP[2])<<8 | uint32(network.IP[3])

	networkMask, _ := network.Mask.Size()
	for index := range 32 - networkMask {
		lastAddress |= 1 << index
	}

	newIP := uint32(parsedIP[0])<<24 | uint32(parsedIP[1])<<16 | uint32(parsedIP[2])<<8 | uint32(parsedIP[3])
	newIP += 1

	if newIP >= lastAddress || newIP == 0 {
		// The IP is the last one, which can't be used
		// or even worse, it did a whole overflow
		return ""
	}

	// Here, we successfully incremented the IP correctly

	// Let's return the formatted IP now.
	return fmt.Sprintf("%d.%d.%d.%d/32", newIP>>24, newIP>>16&0xFF, newIP>>8&0xFF, newIP&0xFF)
}
