package docvalues

import (
	"net/netip"
	"testing"
)

func TestIPIPv4ValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
	}
	errs := value.DeprecatedCheckIsValid("192.168.1.1")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPIPv6ValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv6: true,
	}
	errs := value.DeprecatedCheckIsValid("fe80::890a")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPv6ExpandedValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv6: true,
	}
	errs := value.DeprecatedCheckIsValid("fe80:0000:0000:0000:890a:0000:0000:890a")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPv4DisabledExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: false,
		AllowIPv6: true,
	}
	errs := value.DeprecatedCheckIsValid("192.168.1.1")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPv6DisabledExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowIPv6: false,
	}
	errs := value.DeprecatedCheckIsValid("fe80::890a")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPv6BracketNotationPortValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv6: true,
		AllowPort: AllowedStatusAllowed,
	}
	errs := value.DeprecatedCheckIsValid("[fe80::890a]:80")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPWithIPRangeValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowedIPs: &[]netip.Prefix{
			netip.MustParsePrefix("192.168.0.0/16"),
		},
	}
	errs := value.DeprecatedCheckIsValid("192.168.1.1")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPWithIPRangeInvalidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowedIPs: &[]netip.Prefix{
			netip.MustParsePrefix("192.168.0.0/16"),
		},
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPWithDisallowedIPRangeValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		DisallowedIPs: &[]netip.Prefix{
			netip.MustParsePrefix("192.168.0.0/16"),
		},
	}

	errs := value.DeprecatedCheckIsValid("192.168.1.1")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPWithDisallowedIPRangeInvalidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		DisallowedIPs: &[]netip.Prefix{
			netip.MustParsePrefix("192.168.0.0/16"),
		},
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPWithPort(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowPort: AllowedStatusAllowed,
	}

	errs := value.DeprecatedCheckIsValid("1.1.1.1:80")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPWithRequiredPort(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowPort: AllowedStatusRequired,
	}

	errs := value.DeprecatedCheckIsValid("1.1.1.1:80")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPWithRequiredButNoPort(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowPort: AllowedStatusRequired,
	}

	errs := value.DeprecatedCheckIsValid("1.1.1.1")

	if len(errs) == 0 {
		t.Errorf("Expected errors, got none")
	}
}

func TestIPWithPortInvalid(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowPort: AllowedStatusRequired,
	}

	errs := value.DeprecatedCheckIsValid("1.1.1.1:999999")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPWithPortButNoPort(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
		AllowPort: AllowedStatusAllowed,
	}

	errs := value.DeprecatedCheckIsValid("1.1.1.1")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPRangeValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4:  true,
		AllowRange: AllowedStatusAllowed,
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1/24")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPRangeInvalidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4:  true,
		AllowRange: AllowedStatusRequired,
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1/33")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestIPNoRangeValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4:  true,
		AllowRange: AllowedStatusDisallowed,
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1/24")

	if len(errs) == 0 {
		t.Errorf("Expected errors, got none")
	}
}

func TestIPv6RangeValidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv6:  true,
		AllowRange: AllowedStatusAllowed,
	}

	errs := value.DeprecatedCheckIsValid("fe80::890a/64")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestIPNoPortAllowedButSpecifiedInvalidExample(t *testing.T) {
	value := IPAddressValue{
		AllowIPv4: true,
	}

	errs := value.DeprecatedCheckIsValid("10.0.0.1:80")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}
