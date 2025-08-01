package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

	var unit rune

	if match[6] != -1 {
		unit = rune(value[match[6]])
	} else {
		unit = ' '
	}

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
) func(string, uint64) []*InvalidValue {
	minBytes := simpleParseValueToByte(min, float64(base))
	maxBytes := simpleParseValueToByte(max, float64(base))

	validator := func(value string, byteAmount uint64) []*InvalidValue {
		if byteAmount < minBytes || byteAmount > maxBytes {
			return []*InvalidValue{
				{
					Err:   fmt.Errorf("Value '%s' is out of range (%s - %s)", value, min, max),
					Start: 0,
					End:   uint32(len(value)),
				},
			}
		}

		return nil
	}

	return validator
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
	Validator (func(string, uint64) []*InvalidValue)
}

func (v DataAmountValue) generateUnitSuggestions() []protocol.CompletionItem {
	units := make([]protocol.CompletionItem, 0)
	completionKind := protocol.CompletionItemKindUnit

	for unit := range v.AllowedUnits {
		unitStr := string(unit)

		units = append(units, protocol.CompletionItem{
			Label:         unitStr,
			Kind:          &completionKind,
			Documentation: unitDocumentationMap[unicode.ToLower(unit)],
		})
	}

	return units
}

func (v DataAmountValue) calculateBytesAmount(line string) (uint64, error) {
	match := laxPattern.FindStringSubmatch(line)

	if len(match) == 0 {
		return 0, fmt.Errorf("invalid numeric value '%s'", line)
	}

	rawAmount := match[1]
	rawDecimal := match[2]

	// Parse float
	if rawDecimal != "" && v.AllowDecimal {
		rawAmount += "." + rawDecimal
	}

	amount, err := strconv.ParseFloat(rawAmount, 64)

	if err != nil {
		return 0, fmt.Errorf("invalid amount '%s': %w", rawAmount, err)
	}

	var unit rune

	if match[3] != "" {
		unit = rune(match[3][0])
	} else {
		unit = ' '
	}

	var suffix string

	if match[4] != "" {
		suffix = match[4]
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
		return fmt.Sprintf("'%c' (%s)", unit, unitDocumentationMap[unicode.ToLower(unit)])
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

func (v DataAmountValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
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
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}
	if unitStart != -1 && unitEnd != -1 {
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
				Err: errors.New("Only whole numbers are allowed."),
				// `- 1` to include the decimal point in the error range
				Start: uint32(decimalStart) - 1,
				End:   uint32(decimalEnd),
			},
		}
	}

	if v.Validator != nil {
		// Calculate the byte amount
		byteAmount, err := v.calculateBytesAmount(value)

		if err != nil {
			return []*InvalidValue{
				{
					Err:   fmt.Errorf("Invalid numeric value: %s", err.Error()),
					Start: 0,
					End:   uint32(len(value)),
				},
			}
		}

		// Run the Validator
		validationErrors := v.Validator(value, byteAmount)

		if len(validationErrors) > 0 {
			return validationErrors
		}
	}

	return nil
}

func (v DataAmountValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if value == "" {
		return GenerateBase10Completions(value)
	}

	lastChar := rune(cursor.GetCharacterBefore(value))
	isDigit := lastChar >= '0' && lastChar <= '9'
	isDecimal := lastChar == '.'
	isUnit := utils.KeyExists(v.AllowedUnits, lastChar)

	valueUntilNow := value[:cursor]

	if isDigit {
		// Possible scenarios:
		// `5` - suggest unit and decimal point
		// `5.5` - suggest unit

		// Suggest unit
		completions = append(completions, v.generateUnitSuggestions()...)

		if v.AllowDecimal && !strings.Contains(valueUntilNow, ".") {
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

		completions = append(completions, GenerateBase10Completions(value)...)
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

	return utils.AddSubstrToCompletionItems(completions, valueUntilNow)
}

func (v DataAmountValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	///// Calculate the byte amount from the value
	var byteAmountMessage string
	bytesAmount, err := v.calculateBytesAmount(line)

	if err == nil {
		byteAmountMessage = fmt.Sprintf("%s = %d bytes", line, bytesAmount)
	}

	messages := []string{
		"Numeric value representing a data amount",
	}

	if byteAmountMessage != "" {
		messages = append(messages, byteAmountMessage)
	}

	return messages
}
