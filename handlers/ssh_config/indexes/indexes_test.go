package indexes

import (
	// "config-lsp/handlers/sshd_config/ast"
	// "config-lsp/utils"
	// "testing"
)

// func TestComplexExample(
// 	t *testing.T,
// ) {
// 	input := utils.Dedent(`
// IdentityFile ~/.ssh/id_rsa
//
// Host server1
//      HostName server1.cyberciti.biz
//      User nixcraft
//      Port 4242
//      IdentityFile /nfs/shared/users/nixcraft/keys/server1/id_rsa
//  
// ## Home nas server ##
// Host nas01
//      HostName 192.168.1.100
//      User root
//      IdentityFile ~/.ssh/nas01.key
// `)
// 	config := ast.NewSSHDConfig()
// 	errors := config.Parse(input)
//
// 	if len(errors) > 0 {
// 		t.Fatalf("Expected no errors, but got %v", len(errors))
// 	}
//
// 	indexes, errors := CreateIndexes(*config)
//
// 	if !(len(errors) == 1) {
// 		t.Fatalf("Expected one errors, but got %v", len(errors))
// 	}
//
// 	firstMatchBlock := config.FindMatchBlock(uint32(6))
// 	opts := indexes.AllOptionsPerName["PermitRootLogin"]
// 	if !(len(opts) == 2 &&
// 		len(opts[nil]) == 1 &&
// 		opts[nil][0].Value.Value == "PermitRootLogin yes" &&
// 		opts[nil][0].Start.Line == 0 &&
// 		len(opts[firstMatchBlock]) == 1 &&
// 		opts[firstMatchBlock][0].Value.Value == "\tPermitRootLogin no" &&
// 		opts[firstMatchBlock][0].Start.Line == 6 &&
// 		opts[firstMatchBlock][0].Key.Key == "PermitRootLogin") {
// 		t.Errorf("Expected 3 PermitRootLogin options, but got %v", opts)
// 	}
// }

