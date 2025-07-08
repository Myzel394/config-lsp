package docvalues

import "testing"

func TestDAParseValidExample1(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1k")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1k", 0))

	bytesAmount, err := value.calculateBytesAmount()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if bytesAmount != 1024 {
		t.Errorf("Expected 1024 bytes, got: %d", bytesAmount)
	}
}

func TestDAParseValidExample2(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		AllowDecimal: true,
		Base:         DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1.5k")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1.5k", 0))

	bytesAmount, err := value.calculateBytesAmount()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if bytesAmount != 1536 {
		t.Errorf("Expected 1536 bytes, got: %d", bytesAmount)
	}
}

func TestDAParseValidExampleByteSuffix(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		AllowByteSuffix: true,
		AllowDecimal:    true,
		Base:            DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1.5kB")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1.5t", 0))

	byteAmount, err := value.calculateBytesAmount()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if byteAmount != 1536 {
		t.Errorf("Expected 1536 bytes, got: %d", byteAmount)
	}
}

func TestDAParseValidExampleBitSuffix(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		AllowByteSuffix: true,
		AllowDecimal:    true,
		Base:            DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1.5kb")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1.5t", 0))

	byteAmount, err := value.calculateBytesAmount()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if byteAmount != 192 {
		t.Errorf("Expected 192 bytes, got: %d", byteAmount)
	}
}

func TestDAParseValidExampleNoSuffix(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1024")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1024", 0))

	bytesAmount, err := value.calculateBytesAmount()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if bytesAmount != 1024 {
		t.Errorf("Expected 1024 bytes, got: %d", bytesAmount)
	}
}

func TestDAParseInvalidExample1(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1x")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1x", 0))
}

func TestDAParseInvalidExampleNoUnit(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1.5t")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	} else {
		print(errs[0].Err.Error())
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1.5t", 0))
}

func TestDAParseInvalidExampleNotAllowedUnit(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		AllowDecimal: false,
		Base:         DataAmountValueBase1024,
	}

	errs := value.DeprecatedCheckIsValid("1.5t")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	} else {
		print(errs[0].Err.Error())
	}

	print(value.GetTypeDescription())
	print(value.DeprecatedFetchHoverInfo("1.5t", 0))
}

func TestDAEmptyCompletions(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	completions := value.DeprecatedFetchCompletions("", 0)

	if len(completions) != 10 {
		t.Errorf("Expected no completions, got: %v", completions)
	}
}

func TestDACompletionsWithValidInput(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	completions := value.DeprecatedFetchCompletions("1", 1)

	if len(completions) == 0 {
		t.Error("Expected completions, got none")
	}
}

func TestDACompletionsAtEnd(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base: DataAmountValueBase1024,
	}

	completions := value.DeprecatedFetchCompletions("1k", 2)

	if len(completions) != 0 {
		t.Error("Expected completions, got none")
	}
}

func TestDACompletionsNoDecimal(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base:         DataAmountValueBase1024,
		AllowDecimal: false,
	}

	completions := value.DeprecatedFetchCompletions("1.5", 2)

	if len(completions) != 0 {
		t.Error("Expected no completions, got some")
	}
}

func TestDaCompletionsDecimal(t *testing.T) {
	value := DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
			'g': {},
		},
		Base:         DataAmountValueBase1024,
		AllowDecimal: true,
	}

	completions := value.DeprecatedFetchCompletions("1.5", 2)

	if len(completions) == 0 {
		t.Error("Expected completions, got none")
	}
}
