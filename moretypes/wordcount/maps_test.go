package wordcount_test

import (
	"reflect"
	"testing"

	"github.com/dmitris/gotour-solutions/moretypes/wordcount"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "I am learning Go!",
			expected: map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1},
		},
		{
			input:    "The quick brown fox jumped over the lazy dog.",
			expected: map[string]int{"The": 1, "brown": 1, "dog.": 1, "fox": 1, "jumped": 1, "lazy": 1, "over": 1, "quick": 1, "the": 1},
		},
		{
			input:    "I ate a donut. Then I ate another donut.",
			expected: map[string]int{"I": 2, "Then": 1, "a": 1, "another": 1, "ate": 2, "donut.": 2},
		},
		{
			input:    "A man a plan a canal panama.",
			expected: map[string]int{"A": 1, "a": 2, "canal": 1, "man": 1, "panama.": 1, "plan": 1},
		},
	}
	for _, tt := range tests {
		result := wordcount.WordCount(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("bad result on input [%s]: got %v, want %v", tt.input, result, tt.expected)
		}
	}
}
