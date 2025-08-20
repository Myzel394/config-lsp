package ini

import (
	"config-lsp/utils"
	"testing"
)

func TestParserWorks(t *testing.T) {
	sample := utils.Dedent(`
# A comment at the very top


[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1
PostUp = iptables -I FORWARD -i wg0 -j ACCEPT; iptables -I INPUT -i wg0 -j ACCEPT



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	config := NewConfig()

	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(utils.KeyExists(config.CommentLines, 0) && utils.KeyExists(config.CommentLines, 13)) {
		t.Errorf("Parse: Expected comments to be present on lines 0 and 13")
	}

	if !(config.Sections[0].Start.Line == 3 && config.Sections[0].End.Line == 9) {
		t.Errorf("Parse: Expected section 0 to be present on lines 3 and 9, but it is: %v", config.Sections[0].End)
	}

	if !(config.Sections[0].Start.Character == 0 && config.Sections[0].End.Character == 0) {
		t.Errorf("Parse: Expected section 0 to be present on characters 0 and 0, but it is: %v", config.Sections[0].End)
	}

	if !(config.Sections[0].Header.Name == "Interface" && config.Sections[1].Header.Name == "Peer") {
		t.Errorf("Parse: Expected sections to be present on lines 0, 1, and 2")
	}

	rawFourthProperty, _ := config.Sections[0].Properties.Get(uint32(4))
	fourthProperty := rawFourthProperty.(*Property)
	if !(fourthProperty.Key.Name == "PrivateKey" && fourthProperty.Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 4 to be correct")
	}

	rawFifthProperty, _ := config.Sections[0].Properties.Get(uint32(5))
	fifthProperty := rawFifthProperty.(*Property)
	if !(fifthProperty.Key.Name == "Address" && fifthProperty.Value.Value == "10.0.0.1") {
		t.Errorf("Parse: Expected property line 5 to be correct")
	}

	rawTenthProperty, _ := config.Sections[1].Properties.Get(uint32(11))
	tenthProperty := rawTenthProperty.(*Property)
	if !(tenthProperty.Key.Name == "PublicKey" && tenthProperty.Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 11 to be correct")
	}

	rawPostUpProperty, _ := config.Sections[0].Properties.Get(uint32(6))
	postUpProperty := rawPostUpProperty.(*Property)
	if !(postUpProperty.Key.Name == "PostUp" && postUpProperty.Value.Value == "iptables -I FORWARD -i wg0 -j ACCEPT; iptables -I INPUT -i wg0 -j ACCEPT") {
		t.Errorf("Parse: Expected PostUp property to be correct; %v; %v", postUpProperty, postUpProperty.Value.Value)
	}
}

func TestRootPropertiesWorks(t *testing.T) {
	sample := utils.Dedent(`
ImAtRoot = 123

[Interface]
PrivateKey = aaa
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if len(config.Sections) != 2 {
		t.Fatalf("Parse: Expected 2 sections, but got %d", len(config.Sections))
	}

	if config.Sections[0].Header != nil {
		t.Errorf("Parse: Expected first section to be a root section, but it has a header: %v", config.Sections[0].Header)
	}

	if config.Sections[0].Properties.Size() != 1 {
		t.Errorf("Parse: Expected first section to have 1 property, but it has %d", config.Sections[0].Properties.Size())
	}

	firstProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	if firstProperty.(*Property).Key.Name != "ImAtRoot" || firstProperty.(*Property).Value.Value != "123" {
		t.Errorf("Parse: Expected root property to be 'ImAtRoot = 123', but got '%s = %s'", firstProperty.(*Property).Key.Name, firstProperty.(*Property).Value.Value)
	}
}

func TestRootPropertiesNotALlowed(t *testing.T) {
	sample := utils.Dedent(`
ImAtRoot = 123

[Interface]
PrivateKey = aaa
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: false,
	}
	errors := config.Parse(sample)

	if len(errors) == 0 {
		t.Fatalf("Parse: Expected errors, but got none")
	}

	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}
	if config.Sections[0].Header == nil {
		t.Fatalf("Parse: Expected first section to have a header, but it is nil")
	}
}

func TestOnlyRootPropertiesWorks(t *testing.T) {
	sample := utils.Dedent(`
ImAtRoot = 123
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(len(config.Sections) == 1) {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}

	if config.Sections[0].Header != nil {
		t.Errorf("Parse: Expected first section to be a root section, but it has a header: %v", config.Sections[0].Header)
	}

	rawFirstProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	firstProperty := rawFirstProperty.(*Property)
	if !(firstProperty.Key.Name == "ImAtRoot" && firstProperty.Value.Value == "123" && firstProperty.Separator.Start.Character == 9) {
		t.Errorf("Parse: Expected root property to be 'ImAtRoot = 123', but got '%s = %s'", firstProperty.Key.Name, firstProperty.Value.Value)
	}
}

func TestHalfTypedProperty(t *testing.T) {
	sample := utils.Dedent(`
PrivateKey = 
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}
	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}
	if config.Sections[0].Header != nil {
		t.Fatalf("Parse: Expected first section to have a header, but it is nil")
	}

	rawProperty, found := config.Sections[0].Properties.Get(uint32(0))
	if !found {
		t.Fatalf("Parse: Expected property to be present, but it is not")
	}

	property := rawProperty.(*Property)
	if !(property.Key.Name == "PrivateKey" && property.Value == nil && property.Separator != nil) {
		t.Errorf("Parse: Expected property to be 'PrivateKey ='")
	}
}

func TestEmptyConfig(t *testing.T) {
	sample := utils.Dedent(`
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(len(config.Sections) == 1) {
		t.Fatalf("Parse: Expected no sections, but got %d", len(config.Sections))
	}

	if len(config.CommentLines) != 0 {
		t.Fatalf("Parse: Expected no comment lines, but got %d", len(config.CommentLines))
	}
}

func TestIncompleteSection(t *testing.T) {
	sample := utils.Dedent(`
watch
`)

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if !(len(errors) == 0) {
		t.Fatalf("Parse: Expected errors, but got none")
	}

	if !(len(config.Sections) == 1) {
		t.Fatalf("Parse: Expected no sections, but got %d", len(config.Sections))
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	property := rawProperty.(*Property)

	if !(property.Key.Name == "watch" && property.Value == nil && property.Separator == nil && property.Start.Line == 0 && property.Start.Character == 0 && property.End.Line == 0 && property.End.Character == 5) {
		t.Errorf("Parse: Expected property to be 'watch', but got '%s'", property.Key.Name)
	}
}

func TestExactLines(t *testing.T) {
	sample := `hello = world`
	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}
	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}
	if config.Sections[0].Header != nil {
		t.Fatalf("Parse: Expected first section to have a header, but it is nil")
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	property := rawProperty.(*Property)
	if !(property.Start.Line == 0 && property.Start.Character == 0 && property.End.Line == 0 && property.End.Character == 13) {
		t.Errorf("Parse: Expected property to be 'hello = world', but got '%s = %s'", property.Key.Name, property.Value.Value)
	}
}

func TestExactLines2(t *testing.T) {
	sample := `hello = world
check = true`
	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}
	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}
	if config.Sections[0].Header != nil {
		t.Fatalf("Parse: Expected first section to have a header, but it is nil")
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(1))
	property := rawProperty.(*Property)
	if !(property.Start.Line == 1 && property.Start.Character == 0 && property.End.Line == 1 && property.End.Character == 12) {
		t.Errorf("Parse: Expected property to be at line 1, but got '%v = %v'", property.Start, property.End)
	}
}

func TestExactLines3(t *testing.T) {
	sample := `[main]
server=1


`
	config := NewConfig()
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(config.Sections[0].End.Line == 4 && config.Sections[0].End.Character == 0) {
		t.Errorf("Parse: Expected section to end at line 16, but got %v", config.Sections[1].End)
	}
}

func TestSectionHeaderJustOpened(t *testing.T) {
	sample := `[`

	config := NewConfig()
	errors := config.Parse(sample)

	if !(len(errors) == 0) {
		t.Fatalf("Parse: Expected errors, but got none")
	}

	if !(len(config.Sections) == 1 && config.Sections[0].Header.Name == "" && config.Sections[0].Header.RawValue == "[" && config.Sections[0].Properties.Size() == 0) {
		t.Fatalf("Parse: Expected one section with no header and no properties, but got %d sections", len(config.Sections))
	}
}

func TestSectionHeaderEmpty(t *testing.T) {
	sample := `[]`

	config := NewConfig()
	errors := config.Parse(sample)

	if !(len(errors) == 0) {
		t.Fatalf("Parse: Expected errors, but got none")
	}

	if !(len(config.Sections) == 1 && config.Sections[0].Header.Name == "" && config.Sections[0].Header.RawValue == "[]" && config.Sections[0].Properties.Size() == 0) {
		t.Fatalf("Parse: Expected one section with no header and no properties, but got %d sections", len(config.Sections))
	}
}

func TestPropertyValueInQuotes(t *testing.T) {
	sample := `hello = "world"`

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	property := rawProperty.(*Property)
	if !(property.Key.Name == "hello" && property.Value.Value == "world" && property.Value.Raw == `"world"`) {
		t.Errorf("Parse: Expected property to be 'hello = \"world\"', but got '%s = %s'; RAW: %s", property.Key.Name, property.Value.Value, property.Value.Raw)
	}
}

func TestPropertyValueInQuotesWithEscapedQuotes(t *testing.T) {
	sample := `hello = "world \"escaped\""`

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if len(config.Sections) != 1 {
		t.Fatalf("Parse: Expected 1 section, but got %d", len(config.Sections))
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	property := rawProperty.(*Property)
	if !(property.Key.Name == "hello" && property.Value.Value == `world "escaped"` && property.Value.Raw == `"world \"escaped\""` && property.Separator != nil) {
		t.Errorf("Parse: Expected property to be correct, got %s = %s", property.Key.Name, property.Value.Raw)
	}
}

func TestIncompleteProperty(t *testing.T) {
	sample := `=world`

	config := NewConfig()
	config.XParseConfig = INIParseConfig{
		AllowRootProperties: true,
	}
	errors := config.Parse(sample)

	if !(len(errors) == 0) {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	rawProperty, _ := config.Sections[0].Properties.Get(uint32(0))
	property := rawProperty.(*Property)
	if !(property.Key == nil && property.Value != nil && property.Value.Value == "world" && property.Separator != nil) {
		t.Errorf("Parse: Expected property to be 'hello =', but got '%s = %v'", property.Key.Name, property.Value)
	}
}
