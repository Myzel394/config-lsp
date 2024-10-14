package common

import (
	"config-lsp/utils"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSplitIntoVirtualLinesSimpleExample(
	t *testing.T,
) {
	input := utils.Dedent(`
Hello
World \
how are you
`)
	expected := []VirtualLine{
		{
			LocationRange: LocationRange{
				Start: Location{
					Line:      0,
					Character: 0,
				},
				End: Location{
					Line:      0,
					Character: 5,
				},
			},
			Parts: []VirtualLinePart{
				{
					LocationRange: LocationRange{
						Start: Location{
							Line:      0,
							Character: 0,
						},
						End: Location{
							Line:      0,
							Character: 5,
						},
					},
					Text: "Hello",
				},
			},
		},
		{
			LocationRange: LocationRange{
				Start: Location{
					Line:      1,
					Character: 0,
				},
				End: Location{
					Line:      2,
					Character: 11,
				},
			},
			Parts: []VirtualLinePart{
				{
					LocationRange: LocationRange{
						Start: Location{
							Line:      1,
							Character: 0,
						},
						End: Location{
							Line:      1,
							Character: 6,
						},
					},
					Text: "World ",
				},
				{
					LocationRange: LocationRange{
						Start: Location{
							Line:      2,
							Character: 0,
						},
						End: Location{
							Line:      2,
							Character: 11,
						},
					},
					Text: "how are you",
				},
			},
		},
	}

	lines := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, lines) {
		t.Fatalf("Expected %v, got %v", expected, lines)
	}

	expectedText := "World how are you"
	actualText := lines[1].GetText()

	if expectedText != actualText {
		t.Fatalf("Expected %v, got %v", expectedText, actualText)
	}

	expectedText = "rld how are"
	actualText = lines[1].GetText()[2:13]

	if expectedText != actualText {
		t.Fatalf("Expected %v, got %v", expectedText, actualText)
	}

	expectedRanges := []LocationRange{
		{
			Start: Location{
				Line:      1,
				Character: 2,
			},
			End: Location{
				Line:      1,
				Character: 6,
			},
		},
		{
			Start: Location{
				Line:      2,
				Character: 0,
			},
			End: Location{
				Line:      2,
				Character: 7,
			},
		},
	}
	actualRanges := lines[1].ConvertRangeToTextRange(2, 13)

	if !cmp.Equal(expectedRanges, actualRanges) {
		t.Fatalf("Expected %v, got %v", expectedRanges, actualRanges)
	}
}

// func TestSplitIntoVirtualLinesIndentedExample(
// 	t *testing.T,
// ) {
// 	// 4 spaces
// 	input := utils.Dedent(`
//     Hello
// `)
// 	expected := []VirtualLine{
// 		{
// 			LocationRange: LocationRange{
// 				Start: Location{
// 					Line:      0,
// 					Character: 4,
// 				},
// 				End: Location{
// 					Line:      0,
// 					Character: 9,
// 				},
// 			},
// 			Parts: []VirtualLinePart{
// 				{
// 					LocationRange: LocationRange{
// 						Start: Location{
// 							Line:      0,
// 							Character: 4,
// 						},
// 						End: Location{
// 							Line:      0,
// 							Character: 9,
// 						},
// 					},
// 					Text: "Hello",
// 				},
// 			},
// 		},
// 	}
//
// 	actual := SplitIntoVirtualLines(input)
//
// 	for index, line := range actual {
// 		actual[index] = line.AsTrimmed()
// 	}
//
// 	if !cmp.Equal(expected, actual) {
// 		t.Fatalf("Expected %v, got %v", expected, actual)
// 	}
// }
