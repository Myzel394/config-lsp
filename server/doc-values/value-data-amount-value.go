package docvalues

import (
	"config-lsp/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func simpleParseValueToByte(
	value string,
	base float64,
) uint64 {
	match := laxPattern.FindStringSubmatchIndex(value)

	if len(match) == 0 {
		panic(fmt.Sprintf("Invalid numeric value '%s'", value))
	}

	amount := value[match[2]:match[3]]
	amountFloat, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		panic(fmt.Sprintf("Invalid numeric value '%s': %v", value, err))
	}

	unit := rune(value[match[6]])

	var suffix string
	if match[8] != -1 && match[9] != -1 {
		suffix = value[match[8]:match[9]]
	} else {
		suffix = ""
	}

	byteAmount, err := utils.CalculateNumericValueToByte(
		amountFloat,
		unit,
		suffix,
		base,
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to calculate byte amount from '%s': %v", value, err))
	}

	return byteAmount
}

// Validate that the value is within a specific range.
// The range is inclusive, meaning that the min and max values are valid.
// The min and max values are expected to be in a data amount format
func CreateDARangeValidator(
	min string,
	max string,
	base DataAmountValueBase,
) *func(string, uint64) *[]InvalidValue {
	minBytes := simpleParseValueToByte(min, float64(base))
	maxBytes := simpleParseValueToByte(max, float64(base))

	validator := func(value string, byteAmount uint64) *[]InvalidValue {
		if byteAmount < minBytes || byteAmount > maxBytes {
			return &[]InvalidValue{
				{
					Err:   fmt.Errorf("Value '%s' is out of range (%s - %s)", value, min, max),
					Start: 0,
					End:   uint32(len(value)),
				},
			}
		}

		return nil
	}

	return &validator
}

var unitDocumentationMap = map[rune]string{
	'k': "Kilobytes",
	'm': "Megabytes",
	'g': "Gigabytes",
	't': "Terabytes",
	'e': "Exabytes",
	'p': "Petabytes",
	'z': "Zettabytes",
}

// A lax regex pattern to further validate the data amount value.
var laxPattern = regexp.MustCompile(`(?i)^(?<amount>\d+)(?:\.(?<decimal_amount>\d+))?(?<unit>[a-z])?(?<suffix>b)?$`)

type DataAmountValueBase uint32

const (
	DataAmountValueBase1024 DataAmountValueBase = 1024
	DataAmountValueBase1000 DataAmountValueBase = 1000
)

// We store the raw value and indexes so that we can later parse them
// Why don't we parse the value directly?
// -> Because we want to avoid parsing the value multiple times
// -> And only a few values will need to be parsed to calculate the byte amount
type cachedValue struct {
	rawValue string

	amountStart int
	amountEnd   int

	decimalStart int
	decimalEnd   int

	unitStart int
	unitEnd   int

	suffixStart int
	suffixEnd   int
}

type DataAmountValue struct {
	// The rune set that is allowed for this value.
	// Valid options are:
	// - 'k' for kilobytes
	// - 'm' for megabytes
	// - 'g' for gigabytes
	// - 't' for terabytes
	// - 'e' for exabytes
	// - 'p' for petabytes
	// - 'z' for zettabytes
	// All units should be lowercase.
	AllowedUnits map[rune]struct{}
	// Whether to allow `b` or `B` as a suffix for bytes.
	// Default = false
	AllowByteSuffix bool
	// Whether to allow decimal values.
	// Default = false
	AllowDecimal bool

	// Set the base to either 1000 or 1024.
	// Note: Currently only 1024 is supported
	// Default = 1024
	Base DataAmountValueBase

	// An extra validator to run after the initial validation.
	// This should be a pointer to a function that takes the raw value and the byte amount,
	//
	// The first argument is the raw value as a string,
	// The second argument is the byte amount as uint64.
	// The first argument is guaranteed to be a valid data amount value,
	// The second argument may be 0 if the value is not a valid data amount value.
	//
	// The return value should be a slice of InvalidValue pointers,
	// where each InvalidValue represents an invalid part of the value.
	// If the validator returns nil, or an empty slice, the value is considered valid.
	Validator *(func(string, uint64) *[]InvalidValue)

	_cachedValue *cachedValue
}

func (v DataAmountValue) generateUnitSuggestions() []protocol.CompletionItem {
	units := make([]protocol.CompletionItem, 0)
	completionKind := protocol.CompletionItemKindUnit

	for unit := range v.AllowedUnits {
		unitStr := string(unit)

		units = append(units, protocol.CompletionItem{
			Label:         unitStr,
			Kind:          &completionKind,
			Documentation: unitDocumentationMap[unit],
		})
	}

	return units
}

func (v DataAmountValue) calculateBytesAmount() (uint64, error) {
	if v._cachedValue == nil {
		return 0, errors.New("value is not cached, please call DeprecatedCheckIsValid first")
	}

	// Parse float
	rawAmount := v._cachedValue.rawValue[v._cachedValue.amountStart:v._cachedValue.amountEnd]

	if v._cachedValue.decimalStart != -1 {
		rawAmount += "." + v._cachedValue.rawValue[v._cachedValue.decimalStart:v._cachedValue.decimalEnd]
	}

	amount, err := strconv.ParseFloat(rawAmount, 64)

	if err != nil {
		return 0, fmt.Errorf("invalid amount '%s': %w", rawAmount, err)
	}

	unit := rune(v._cachedValue.rawValue[v._cachedValue.unitStart])

	var suffix string

	if v._cachedValue.suffixStart != -1 {
		suffix = v._cachedValue.rawValue[v._cachedValue.suffixStart:v._cachedValue.suffixEnd]
	} else {
		suffix = ""
	}

	var base float64

	if v.Base == DataAmountValueBase1024 {
		base = 1024
	} else {
		base = 1000
	}

	byteAmount, err := utils.CalculateNumericValueToByte(
		amount,
		unit,
		suffix,
		base,
	)

	if err != nil {
		return 0, fmt.Errorf("failed to calculate byte amount from '%s': %w", rawAmount, err)
	}

	return byteAmount, nil
}

func (v DataAmountValue) GetTypeDescription() []string {
	description := []string{
		"Byte amount",
		"Example: 512, 2K, 1M",
	}

	allowedUnits := utils.MapMapToSlice(v.AllowedUnits, func(unit rune, _ struct{}) string {
		return fmt.Sprintf("'%c' (%s)", unit, unitDocumentationMap[unit])
	})

	description = append(description, fmt.Sprintf("* Allowed units: %s", strings.Join(allowedUnits, ", ")))

	if v.AllowDecimal {
		description = append(description, "* Decimal values are allowed (e.g. 1.5K)")
	} else {
		description = append(description, "* Decimal values are not allowed")
	}

	if v.AllowByteSuffix {
		description = append(description, "* Byte suffix is allowed (b or B)")
	} else {
		description = append(description, "* Byte suffix is not allowed")
	}

	if v.Base == DataAmountValueBase1024 {
		description = append(description, "* Base `1024` is used (e.g. 1K = 1024 bytes)")
	} else {
		description = append(description, "* Base `1000` is used (e.g. 1K = 1000 bytes)")
	}

	return description
}

func (v *DataAmountValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	v._cachedValue = nil

	match := laxPattern.FindStringSubmatchIndex(value)

	if len(match) == 0 {
		return []*InvalidValue{
			{
				Err:   errors.New("Invalid numeric value"),
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}

	amountStart := match[2]
	amountEnd := match[3]
	decimalStart := match[4]
	decimalEnd := match[5]
	unitStart := match[6]
	unitEnd := match[7]
	suffixStart := match[8]
	suffixEnd := match[9]

	if amountStart == -1 || amountEnd == -1 {
		return []*InvalidValue{
			{
				Err:   errors.New("Amount is missing"),
				Start: uint32(amountStart),
				End:   uint32(amountEnd),
			},
		}
	}
	if unitStart == -1 && unitEnd != -1 {
		return []*InvalidValue{
			{
				Err:   errors.New("Unit is missing"),
				Start: uint32(unitStart),
				End:   uint32(unitEnd),
			},
		}
	}

	unit := rune(value[unitStart])
	if !utils.KeyExists(v.AllowedUnits, unit) {
		allowedUnitsString := strings.Join(utils.MapMapToSlice(v.AllowedUnits, func(unit rune, _ struct{}) string {
			return string(unit)
		}), ", ")
		return []*InvalidValue{
			{
				Err:   fmt.Errorf("Unit '%c' is not allowed; It must be one of: %s", unit, allowedUnitsString),
				Start: uint32(unitStart),
				End:   uint32(unitEnd),
			},
		}
	}

	if !v.AllowByteSuffix && suffixStart != -1 && suffixEnd != -1 {
		return []*InvalidValue{
			{
				Err:   errors.New("Byte suffix is not allowed"),
				Start: uint32(suffixStart),
				End:   uint32(suffixEnd),
			},
		}
	}

	if !v.AllowDecimal && decimalStart != -1 && decimalEnd != -1 {
		return []*InvalidValue{
			{
				Err:   errors.New("Decimal part is not allowed"),
				Start: uint32(decimalStart),
				End:   uint32(decimalEnd),
			},
		}
	}

	// Validation done, store cached value
	v._cachedValue = &cachedValue{
		rawValue:     value,
		amountStart:  amountStart,
		amountEnd:    amountEnd,
		decimalStart: decimalStart,
		decimalEnd:   decimalEnd,
		unitStart:    unitStart,
		unitEnd:      unitEnd,
		suffixStart:  suffixStart,
		suffixEnd:    suffixEnd,
	}

	return nil
}

func (v DataAmountValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if line == "" {
		return GenerateBase10Completions(line)
	}

	lastChar := []rune(line)[cursor]
	isDigit := lastChar >= '0' && lastChar <= '9'
	isDecimal := lastChar == '.'
	isUnit := utils.KeyExists(v.AllowedUnits, lastChar)

	lineUntilNow := line[:cursor]

	if isDigit {
		// Possible scenarios:
		// `5` - suggest unit and decimal point
		// `5.5` - suggest unit

		// Suggest unit
		completions = append(completions, v.generateUnitSuggestions()...)

		if v.AllowDecimal && !strings.Contains(lineUntilNow, ".") {
			kind := protocol.CompletionItemKindValue
			completions = append(completions, protocol.CompletionItem{
				Label:         ".",
				Kind:          &kind,
				Documentation: "Decimal point",
			})
		}
	} else if isDecimal && v.AllowDecimal {
		// Possible scenarios:
		// `5.` - suggest numbers

		completions = append(completions, GenerateBase10Completions(line)...)
	} else if isUnit && utils.KeyExists(v.AllowedUnits, lastChar) && v.AllowByteSuffix {
		// Possible scenarios:
		// `5K` - suggest byte suffix
		// `5M` - suggest byte suffix
		// `5G` - suggest byte suffix

		kind := protocol.CompletionItemKindUnit
		completions = append(completions, protocol.CompletionItem{
			Label:         "b",
			Kind:          &kind,
			Documentation: "Bit suffix",
		})
		completions = append(completions, protocol.CompletionItem{
			Label:         "B",
			Kind:          &kind,
			Documentation: "Byte suffix",
		})
	}

	return completions
}

func (v DataAmountValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	///// Calculate the byte amount from the value
	var byteAmountMessage string
	bytesAmount, err := v.calculateBytesAmount()

	if err == nil {
		byteAmountMessage = fmt.Sprintf("%s = %d bytes", v._cachedValue.rawValue, bytesAmount)
	}

	messages := []string{
		"Numeric value representing a data amount",
	}

	if byteAmountMessage != "" {
		messages = append(messages, byteAmountMessage)
	}

	return messages
}
