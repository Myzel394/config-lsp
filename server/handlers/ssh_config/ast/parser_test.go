package ast

import (
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"config-lsp/utils"
	"testing"
)

func TestSSHConfigParserExample(
	t *testing.T,
) {
	input := utils.Dedent(`
HostName 1.2.3.4
User root
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)

	if !(firstEntry.Value.Value == "HostName 1.2.3.4" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 16 &&
		firstEntry.Key.Value.Value == "HostName" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 8 &&
		firstEntry.OptionValue.Value.Value == "1.2.3.4" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 9 &&
		firstEntry.OptionValue.LocationRange.End.Character == 16) {
		t.Errorf("Expected first entry to be HostName 1.2.3.4, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(1))
	secondEntry := rawSecondEntry.(*SSHOption)

	if !(secondEntry.Value.Value == "User root" &&
		secondEntry.LocationRange.Start.Line == 1 &&
		secondEntry.LocationRange.End.Line == 1 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 9 &&
		secondEntry.Key.Value.Value == "User" &&
		secondEntry.Key.LocationRange.Start.Character == 0 &&
		secondEntry.Key.LocationRange.End.Character == 4 &&
		secondEntry.OptionValue.Value.Value == "root" &&
		secondEntry.OptionValue.LocationRange.Start.Character == 5 &&
		secondEntry.OptionValue.LocationRange.End.Character == 9) {
		t.Errorf("Expected second entry to be User root, but got: %v", secondEntry)
	}
}

func TestMatchSimpleBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
Hostname 1.2.3.4

Match originalhost "192.168.0.1"
	User root
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 option and no comment lines, but got: %v entries, %v comment lines", p.Options.Size(), len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)

	if !(firstEntry.Value.Value == "Hostname 1.2.3.4" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 16 &&
		firstEntry.Key.Value.Value == "Hostname" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 8 &&
		firstEntry.OptionValue.Value.Value == "1.2.3.4" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 9 &&
		firstEntry.OptionValue.LocationRange.End.Character == 16) {
		t.Errorf("Expected first entry to be Hostname 1.2.3.4, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHMatchBlock)

	if !(secondEntry.Options.Size() == 1 &&
		secondEntry.LocationRange.Start.Line == 2 &&
		secondEntry.LocationRange.End.Line == 3 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 10 &&
		secondEntry.MatchOption.OptionValue.Value.Raw == "originalhost \"192.168.0.1\"" &&
		secondEntry.MatchOption.OptionValue.LocationRange.Start.Character == 6 &&
		secondEntry.MatchOption.OptionValue.LocationRange.End.Character == 32) {
		t.Errorf("Expected second entry to be Match originalhost \"192.168.0.1\", but got: %v; options amount: %d", secondEntry, secondEntry.Options.Size())
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Value.Raw == "\tUser root" &&
		thirdEntry.LocationRange.Start.Line == 3 &&
		thirdEntry.LocationRange.End.Line == 3 &&
		thirdEntry.LocationRange.Start.Character == 0 &&
		thirdEntry.LocationRange.End.Character == 10 &&
		thirdEntry.Key.Value.Value == "User" &&
		thirdEntry.Key.LocationRange.Start.Character == 1 &&
		thirdEntry.Key.LocationRange.End.Character == 5 &&
		thirdEntry.OptionValue.Value.Value == "root" &&
		thirdEntry.OptionValue.LocationRange.Start.Character == 6 &&
		thirdEntry.OptionValue.LocationRange.End.Character == 10) {
		t.Errorf("Expected third entry to be User root, but got: %v", thirdEntry)
	}
}

func TestSimpleHostBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
Ciphers 3des-cbc

Host example.com
	Port 22
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 option and no comment lines, but got: %v entries, %v comment lines", p.Options.Size(), len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)
	if !(firstEntry.Value.Value == "Ciphers 3des-cbc") {
		t.Errorf("Expected first entry to be Ciphers 3des-cbc, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHHostBlock)
	if !(secondEntry.Options.Size() == 1 &&
		secondEntry.LocationRange.Start.Line == 2 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Line == 3 &&
		secondEntry.LocationRange.End.Character == 8) {
		t.Errorf("Expected second entry to be Host example.com, but got: %v", secondEntry)
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Value.Raw == "\tPort 22" &&
		thirdEntry.Key.Value.Raw == "Port" &&
		thirdEntry.OptionValue.Value.Raw == "22" &&
		thirdEntry.LocationRange.Start.Line == 3 &&
		thirdEntry.LocationRange.Start.Character == 0 &&
		thirdEntry.LocationRange.End.Line == 3 &&
		thirdEntry.LocationRange.End.Character == 8) {
		t.Errorf("Expected third entry to be Port 22, but got: %v", thirdEntry)
	}

	rawFourthEntry, _ := p.Options.Get(uint32(3))

	if !(rawFourthEntry == nil) {
		t.Errorf("Expected fourth entry to be nil, but got: %v", rawFourthEntry)
	}
}

func TestComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
Host laptop
    HostName laptop.lan

Match originalhost laptop exec "[[ $(/usr/bin/dig +short laptop.lan) == '' ]]"
    HostName laptop.sdn
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 option and no comment lines, but got: %v entries, %v comment lines", p.Options.Size(), len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstBlock := rawFirstEntry.(*SSHHostBlock)
	if !(firstBlock.Options.Size() == 1 &&
		firstBlock.LocationRange.Start.Line == 0 &&
		firstBlock.LocationRange.Start.Character == 0 &&
		firstBlock.LocationRange.End.Line == 1 &&
		firstBlock.LocationRange.End.Character == 23) {
		t.Errorf("Expected first entry to be Host laptop, but got: %v", firstBlock)
	}

	rawSecondEntry, _ := firstBlock.Options.Get(uint32(1))
	secondOption := rawSecondEntry.(*SSHOption)
	if !(secondOption.Value.Raw == "    HostName laptop.lan" &&
		secondOption.Key.Value.Raw == "HostName" &&
		secondOption.OptionValue.Value.Raw == "laptop.lan" &&
		secondOption.LocationRange.Start.Line == 1 &&
		secondOption.LocationRange.Start.Character == 0 &&
		secondOption.Key.LocationRange.Start.Character == 4 &&
		secondOption.LocationRange.End.Line == 1 &&
		secondOption.LocationRange.End.Character == 23) {
		t.Errorf("Expected second entry to be HostName laptop.lan, but got: %v", secondOption)
	}

	rawThirdEntry, _ := p.Options.Get(uint32(3))
	secondBlock := rawThirdEntry.(*SSHMatchBlock)
	if !(secondBlock.Options.Size() == 1 &&
		secondBlock.LocationRange.Start.Line == 3 &&
		secondBlock.LocationRange.Start.Character == 0 &&
		secondBlock.LocationRange.End.Line == 4 &&
		secondBlock.LocationRange.End.Character == 23) {
		t.Errorf("Expected second entry to be Match originalhost laptop exec \"[[ $(/usr/bin/dig +short laptop.lan) == '' ]]\", but got: %v", secondBlock)
	}

	if !(secondBlock.MatchOption.LocationRange.End.Character == 78) {
		t.Errorf("Expected second entry to be Match originalhost laptop exec \"[[ $(/usr/bin/dig +short laptop.lan) == '' ]]\", but got: %v", secondBlock)
	}

	rawFourthEntry, _ := secondBlock.Options.Get(uint32(4))
	thirdOption := rawFourthEntry.(*SSHOption)
	if !(thirdOption.Value.Raw == "    HostName laptop.sdn" &&
		thirdOption.Key.Value.Raw == "HostName" &&
		thirdOption.OptionValue.Value.Raw == "laptop.sdn" &&
		thirdOption.LocationRange.Start.Line == 4 &&
		thirdOption.LocationRange.Start.Character == 0 &&
		thirdOption.Key.LocationRange.Start.Character == 4 &&
		thirdOption.LocationRange.End.Line == 4 &&
		thirdOption.LocationRange.End.Character == 23) {
		t.Errorf("Expected third entry to be HostName laptop.sdn, but got: %v", thirdOption)
	}
}

func TestIncompleteExample(
	t *testing.T,
) {
	input := utils.Dedent(`
User 
`)
	p := NewSSHConfig()

	errors := p.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 1) {
		t.Errorf("Expected 1 option, but got: %v", p.Options.Size())
	}

	if !(len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected no comment lines, but got: %v", len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)
	if !(firstEntry.Value.Raw == "User " && firstEntry.Key.Value.Raw == "User") {
		t.Errorf("Expected first entry to be User, but got: %v", firstEntry)
	}

	if !(firstEntry.OptionValue != nil && firstEntry.OptionValue.Value.Raw == "") {
		t.Errorf("Expected first entry to have an empty value, but got: %v", firstEntry)
	}
}

func TestIncompleteMatch(
	t *testing.T,
) {
	input := utils.Dedent(`
Match 
`)
	p := NewSSHConfig()

	errors := p.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 1) {
		t.Errorf("Expected 1 option, but got: %v", p.Options.Size())
	}

	if !(len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected no comment lines, but got: %v", len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHMatchBlock)
	if !(firstEntry.MatchOption.Key.Value.Raw == "Match") {
		t.Errorf("Expected first entry to be User, but got: %v", firstEntry)
	}

	if !(firstEntry.MatchOption.OptionValue != nil && firstEntry.MatchOption.OptionValue.Value.Raw == "") {
		t.Errorf("Expected first entry to have an empty value, but got: %v", firstEntry)
	}
}

func TestMatchWithIncompleteEntry(
	t *testing.T,
) {
	input := utils.Dedent(`
Match user 
`)
	p := NewSSHConfig()

	errors := p.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 1) {
		t.Errorf("Expected 1 option, but got: %v", p.Options.Size())
	}

	if !(len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected no comment lines, but got: %v", len(p.CommentLines))
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHMatchBlock)
	if !(firstEntry.MatchOption.Key.Value.Raw == "Match") {
		t.Errorf("Expected first entry to be User, but got: %v", firstEntry)
	}

	if !(firstEntry.MatchOption.OptionValue != nil && firstEntry.MatchOption.OptionValue.Value.Raw == "user ") {
		t.Errorf("Expected first entry to have an empty value, but got: %v", firstEntry)
	}

	if !(firstEntry.MatchValue.Entries[0].Criteria.Type == matchparser.MatchCriteriaTypeUser) {
		t.Errorf("Expected first entry to have a user criteria, but got: %v", firstEntry)
	}
}

func TestInvalidMatchExample(
	t *testing.T,
) {
	input := utils.Dedent(`
Match us
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) == 0 {
		t.Fatalf("Expected errors, got none")
	}
}

func TestComplexBigExample(
	t *testing.T,
) {
	// From https://gist.github.com/zeloc/b9455b793b07898025db
	input := utils.Dedent(`
### default for all ##
Host *
     ForwardAgent no
     ForwardX11 no
     ForwardX11Trusted yes
     User nixcraft
     Port 22
     Protocol 2
     ServerAliveInterval 60
     ServerAliveCountMax 30
 
## override as per host ##
Host server1
     HostName server1.cyberciti.biz
     User nixcraft
     Port 4242
     IdentityFile /nfs/shared/users/nixcraft/keys/server1/id_rsa
 
## Home nas server ##
Host nas01
     HostName 192.168.1.100
     User root
     IdentityFile ~/.ssh/nas01.key
 
## Login AWS Cloud ##
Host aws.apache
     HostName 1.2.3.4
     User wwwdata
     IdentityFile ~/.ssh/aws.apache.key
 
## Login to internal lan server at 192.168.0.251 via our public uk office ssh based gateway using ##
## $ ssh uk.gw.lan ##
Host uk.gw.lan uk.lan
     HostName 192.168.0.251
     User nixcraft
     ProxyCommand  ssh nixcraft@gateway.uk.cyberciti.biz nc %h %p 2> /dev/null
 
## Our Us Proxy Server ##
## Forward all local port 3128 traffic to port 3128 on the remote vps1.cyberciti.biz server ##
## $ ssh -f -N  proxyus ##
Host proxyus
    HostName vps1.cyberciti.biz
    User breakfree
    IdentityFile ~/.ssh/vps1.cyberciti.biz.key
    LocalForward 3128 127.0.0.1:3128
`)
	p := NewSSHConfig()

	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 6 && len(utils.KeysOfMap(p.CommentLines)) == 9) {
		t.Errorf("Expected 6 options and no comment lines, but got: %v options, %v comment lines", p.Options.Size(), len(p.CommentLines))
	}

	// Validate each Host block and its options
	rawFirstEntry, _ := p.Options.Get(uint32(1))
	firstBlock := rawFirstEntry.(*SSHHostBlock)
	if !(firstBlock.Options.Size() == 8) {
		t.Errorf("Expected 8 options for Host *, but got: %v", firstBlock.Options.Size())
	}

	rawSecondEntry, _ := p.Options.Get(uint32(12))
	secondBlock := rawSecondEntry.(*SSHHostBlock)
	if !(secondBlock.Options.Size() == 4) {
		t.Errorf("Expected 4 options for Host server1, but got: %v", secondBlock.Options.Size())
	}

	rawThirdEntry, _ := p.Options.Get(uint32(19))
	thirdBlock := rawThirdEntry.(*SSHHostBlock)
	if !(thirdBlock.Options.Size() == 3) {
		t.Errorf("Expected 2 options for Host nas01, but got: %v", thirdBlock.Options.Size())
	}

	rawFourthEntry, _ := p.Options.Get(uint32(25))
	fourthBlock := rawFourthEntry.(*SSHHostBlock)
	if !(fourthBlock.Options.Size() == 3) {
		t.Errorf("Expected 2 options for Host aws.apache, but got: %v", fourthBlock.Options.Size())
	}

	rawFifthEntry, _ := p.Options.Get(uint32(32))
	fifthBlock := rawFifthEntry.(*SSHHostBlock)
	if !(fifthBlock.Options.Size() == 3) {
		t.Errorf("Expected 3 options for Host uk.gw.lan uk.lan, but got: %v", fifthBlock.Options.Size())
	}

	rawSixthEntry, _ := p.Options.Get(uint32(40))
	sixthBlock := rawSixthEntry.(*SSHHostBlock)
	if !(sixthBlock.Options.Size() == 4) {
		t.Errorf("Expected 4 options for Host proxyus, but got: %v", sixthBlock.Options.Size())
	}
}
