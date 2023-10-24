package gu2en

import (
	"testing"
)

func TestTransliterate(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"ગુજરાતી ભાષા", "Gujarātī bhāṣhā"},
		{"નમસ્તે", "Namaste"},
		{"કેમ છો?", "Kem chho?"},
		{"સારું છું", "Sārun chhun"},
		{"તમે કેમ છો", "Tame kem chho"},
		{"આપણે", "Āpaṇe"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output := Transliterate(test.input)
			if output != test.output {
				t.Errorf("Expected %s, got %s", test.output, output)
			}
		})
	}
}
