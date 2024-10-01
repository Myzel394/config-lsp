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
}
