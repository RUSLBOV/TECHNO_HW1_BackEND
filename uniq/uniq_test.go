package uniq

import (
	"reflect"
	"testing"
)

func TestUniqLines(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		opts   Options
		expect []string
	}{
		{
			name:   "simple unique",
			input:  []string{"a", "a", "b", "b", "c"},
			opts:   Options{},
			expect: []string{"a", "b", "c"},
		},
		{
			name:   "count occurrences",
			input:  []string{"a", "a", "b"},
			opts:   Options{Count: true},
			expect: []string{"2 a", "1 b"},
		},
		{
			name:   "duplicates only",
			input:  []string{"a", "a", "b", "c", "c"},
			opts:   Options{Duplicates: true},
			expect: []string{"a", "c"},
		},
		{
			name:   "unique only",
			input:  []string{"a", "b", "b", "c"},
			opts:   Options{Unique: true},
			expect: []string{"a", "c"},
		},
		{
			name:   "ignore case",
			input:  []string{"apple", "Apple", "banana", "BANANA", "cherry"},
			opts:   Options{IgnoreCase: true},
			expect: []string{"apple", "banana", "cherry"},
		},
		{
			name:   "skip fields",
			input:  []string{"one apple", "two apple", "one banana", "two banana"},
			opts:   Options{SkipFields: 1},
			expect: []string{"one apple", "one banana"}, // учитываются только соседние совпадения после пропуска первого поля
		},
		{
			name:   "skip chars",
			input:  []string{"apple", "apricot", "banana"},
			opts:   Options{SkipChars: 2},
			expect: []string{"apple", "apricot", "banana"}, // соседние строки после пропуска символов разные → все выводятся
		},
		{
			name:   "count with ignore case",
			input:  []string{"apple", "Apple", "banana", "BANANA", "banana"},
			opts:   Options{Count: true, IgnoreCase: true},
			expect: []string{"2 apple", "3 banana"},
		},
		{
			name:   "complex case with repeats",
			input:  []string{"apple", "Apple", "banana", "banana", "BANANA", "cherry", "apple", "cherry"},
			opts:   Options{Count: true, IgnoreCase: true},
			expect: []string{"2 apple", "3 banana", "1 cherry", "1 apple", "1 cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := UniqLines(tt.input, tt.opts)
			if !reflect.DeepEqual(out, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, out)
			}
		})
	}
}
