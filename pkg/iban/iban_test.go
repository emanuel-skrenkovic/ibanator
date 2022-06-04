package iban_test

import (
	"fmt"
	"github.com/sula0/ibanator/v2/pkg/iban"
	"testing"
)

func TestValidateIBAN(t *testing.T) {
	// https://www.iban.com/testibans
	// Keeping only values for which the validation
	// checks are written.
	var tests = []struct {
		iban string
		want bool
	}{
		{"", false},
		{"GB82WEST123456987654325555555555555", false},
		{"GB82WEST12345698765432", true},
		{"GB33BUKB20201555555555", true},
		{"GB94BARC10201530093459", true},
		{"GB96BARC202015300934591", false},
		{"US64SVBKUS6S3300958879", false},
	}

	// https://github.com/golang/go/wiki/TableDrivenTests
	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.iban)

		t.Run(testName, func(t *testing.T) {
			ans, _ := iban.ValidateIBAN(tt.iban)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
