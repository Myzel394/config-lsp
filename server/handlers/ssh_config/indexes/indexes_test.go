package indexes

import (
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/utils"
	"testing"
)

func TestComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
IdentityFile ~/.ssh/id_rsa

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
`)
	config := ast.NewSSHConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	firstMatchBlock := config.FindBlock(uint32(6))
	opts := indexes.AllOptionsPerName["identityfile"]
	if !(len(opts) == 3 &&
		opts[nil][0].OptionValue.Value.Value == "~/.ssh/id_rsa" &&
		opts[firstMatchBlock][0].OptionValue.Value.Value == "/nfs/shared/users/nixcraft/keys/server1/id_rsa") {
		t.Errorf("Expected 3 IdentityFile options, but got %v", opts)
	}

}

// TODO: Add check for options that require other options to be present
func TestDoubleOptionExample(
	t *testing.T,
) {
	input := utils.Dedent(`
IdentityFile ~/.ssh/id_rsa

Host server1
	User nixcraft
	User root
`)
	config := ast.NewSSHConfig()

	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if !(len(errors) == 1) {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}

	if !(errors[0].Range.Start.Line == 4) {
		t.Errorf("Expected error on line 4, but got %v", errors[0].Range.Start.Line)
	}

	if !(len(indexes.AllOptionsPerName["user"]) == 1) {
		t.Errorf("Expected 1 User option, but got %v", indexes.AllOptionsPerName["user"])
	}
}

func TestIgnoredUnknownExample(
	t *testing.T,
) {
	input := utils.Dedent(`
IgnoreUnknown UseKeychain
User root
UseKeychain yes
`)
	config := ast.NewSSHConfig()

	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}

	firstOption, _ := config.Options.Get(uint32(0))
	option := firstOption.(*ast.SSHOption)
	if !(indexes.IgnoredOptions[nil].OptionValue == option) {
		t.Errorf("Expected IgnoredOptions to be first option, but got %v", option)
	}

	if !(len(indexes.IgnoredOptions[nil].IgnoredOptions) == 1 && utils.KeyExists(indexes.IgnoredOptions[nil].IgnoredOptions, "usekeychain")) {
		t.Errorf("Expected IgnoreOptions to contain 'UseKeychain', but got: %v", indexes.IgnoredOptions[nil].IgnoredOptions)
	}

	if !(indexes.IgnoredOptions[nil].IgnoredOptions["usekeychain"].Start.Line == 0 && indexes.IgnoredOptions[nil].IgnoredOptions["usekeychain"].Start.Character == 14 && indexes.IgnoredOptions[nil].IgnoredOptions["usekeychain"].End.Character == 25) {
		t.Errorf("Expected IgnoreOptions to contain 'UseKeychain' on line 0 and from position 14-24, but got: %v", indexes.IgnoredOptions[nil].IgnoredOptions)
	}
}

func TestTagsExample(
	t *testing.T,
) {
	input := utils.Dedent(`
Match tagged good_ip
	AddressFamily inet

Match tagged myuser
	User root
	Tag good_ip
`)

	config := ast.NewSSHConfig()

	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}

	if !(len(indexes.Tags) == 2) {
		t.Errorf("Expected 2 tags, but got %v", indexes.Tags)
	}

	rawFirstMatch, _ := config.Options.Get(uint32(0))
	firstMatch := rawFirstMatch.(*ast.SSHMatchBlock)
	if !(indexes.Tags["good_ip"].Block.Start.Line == firstMatch.Start.Line) {
		t.Errorf("Expected first tag to be 'good_ip', but got %v", indexes.Tags)
	}

	_, secondBlock := config.FindOption(uint32(3))
	secondMatch := secondBlock.(*ast.SSHMatchBlock)
	if !(indexes.Tags["myuser"].Block.Start.Line == secondMatch.Start.Line) {
		t.Errorf("Expected second tag to be 'myuser', but got %v", indexes.Tags)
	}

	if !(len(indexes.TagImports) == 1) {
		t.Errorf("Expected 1 tag import, but got %v", indexes.TagImports)
	}

	if !(len(indexes.TagImports["good_ip"]) == 1) {
		t.Errorf("Expected 1 tag import for 'good_ip', but got %v", indexes.TagImports["good_ip"])
	}

	tagOption, _ := config.FindOption(uint32(5))
	if !(indexes.TagImports["good_ip"][secondBlock].Start.Line == tagOption.Start.Line) {
		t.Errorf("Expected first tag import to be 'good_ip', but got %v", indexes.TagImports)
	}
}

func TestIncludeExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Include /etc/ssh/sshd_config.d/*.conf hello_world
`)
	config := ast.NewSSHConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	if !(len(indexes.Includes) == 1) {
		t.Fatalf("Expected 1 include, but got %v", len(indexes.Includes))
	}

	if !(len(indexes.Includes[1].Values) == 2) {
		t.Fatalf("Expected 2 include path, but got %v", len(indexes.Includes[1].Values))
	}

	if !(indexes.Includes[1].Values[0].Value == "/etc/ssh/sshd_config.d/*.conf" &&
		indexes.Includes[1].Values[0].Start.Line == 1 &&
		indexes.Includes[1].Values[0].End.Line == 1 &&
		indexes.Includes[1].Values[0].Start.Character == 8 &&
		indexes.Includes[1].Values[0].End.Character == 37) {
		t.Errorf("Expected '/etc/ssh/sshd_config.d/*.conf' on line 1, but got %v on line %v", indexes.Includes[1].Values[0].Value, indexes.Includes[1].Values[0].Start.Line)
	}

	if !(indexes.Includes[1].Values[1].Value == "hello_world" &&
		indexes.Includes[1].Values[1].Start.Line == 1 &&
		indexes.Includes[1].Values[1].End.Line == 1 &&
		indexes.Includes[1].Values[1].Start.Character == 38 &&
		indexes.Includes[1].Values[1].End.Character == 49) {
		t.Errorf("Expected 'hello_world' on line 1, but got %v on line %v", indexes.Includes[1].Values[1].Value, indexes.Includes[1].Values[1].Start.Line)
	}
}
