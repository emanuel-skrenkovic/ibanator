package iban

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ValidateIBAN
// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Modulo_operation_on_IBAN
func ValidateIBAN(iban string) (bool, error) {
	ibanLen := len(iban)

	if len(iban) == 0 {
		return false, errors.New("empty IBAN")
	}

	if ibanLen < 5 {
		return false, errors.New("IBAN length cannot be less than 5 characters")
	}

	if len(iban) > 34 {
		return false, errors.New("IBAN length cannot exceed 34 characters")
	}

	normalized := strings.ToUpper(iban)
	if err := validateCountryFormat(normalized); err != nil {
		return false, err
	}

	var buffer bytes.Buffer
	appendIBANDigits(normalized[4:], &buffer)
	appendIBANDigits(normalized[:4], &buffer)

	validFormat := mod97(buffer.String()) == 1
	if !validFormat {
		return false, errors.New("invalid IBAN check digits MOD-97-10 as per ISO/IEC 7064:2003")
	}

	return true, nil
}

// TODO: BigInt instead?
// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN
func mod97(input string) int {
	parsePosition := 9
	firstDigits := input[:parsePosition]

	modDigits := smallIntMod97(firstDigits)

	ibanLen := len(input)
	for parsePosition < ibanLen {
		var parseEnd int

		// Why, oh why, doesn't Go have a min int function?
		if parsePosition+7 < ibanLen {
			parseEnd = parsePosition + 7
		} else {
			parseEnd = ibanLen
		}

		modDigits = modDigits + input[parsePosition:parseEnd]
		modDigits = smallIntMod97(modDigits)

		parsePosition += 7
	}

	finalRemainder, err := strconv.Atoi(modDigits)
	if err != nil {
		// Should not fail after all the previous calculations.
		// Still need a better solution.
		panic(err)
	}

	return finalRemainder
}

func smallIntMod97(input string) string {
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	mod := number % 97

	modDigits := strconv.Itoa(mod)
	return modDigits
}

func validateCountryFormat(iban string) error {
	countryCode := iban[:2]

	err := validateLength(countryCode, iban)
	if err != nil {
		return err
	}

	// There is much more to validating IBANs with country-specific
	// rules, but keeping it small.

	return nil
}

func validateLength(cc string, iban string) error {
	requiredLength, found := countryIBANLengths[cc]
	if !found {
		errorMessage := fmt.Sprintf("country '%s' is not supported", cc)
		return errors.New(errorMessage)
	}

	if len(iban) != requiredLength {
		errorMessage := fmt.Sprintf("IBAN is of incorrect length for country with code: '%s'", cc)
		return errors.New(errorMessage)
	}

	return nil
}

func appendIBANDigits(input string, buffer *bytes.Buffer) {
	for _, r := range input {
		if r >= 48 && r <= 57 { // 48 -> 57 = ASCII 0 - 9. We only like ASCII here.
			val := string(r)
			buffer.WriteString(val)
		} else {
			val := strconv.Itoa(int(r - 55))
			buffer.WriteString(val)
		}
	}
}

var countryIBANLengths = map[string]int{
	"AL": 28,
	"AD": 24,
	"AT": 20,
	"AZ": 28,
	"BH": 22,
	"BY": 28,
	"BE": 16,
	"BA": 20,
	"BR": 29,
	"BG": 22,
	"CR": 22,
	"HR": 21,
	"CY": 28,
	"CZ": 24,
	"DK": 18,
	"DO": 28,
	"TL": 23,
	"EG": 29,
	"SV": 28,
	"EE": 20,
	"FO": 18,
	"FI": 18,
	"FR": 27,
	"GE": 22,
	"DE": 22,
	"GI": 23,
	"GR": 27,
	"GL": 18,
	"GT": 28,
	"HU": 28,
	"IS": 26,
	"IQ": 23,
	"IE": 22,
	"IL": 23,
	"IT": 27,
	"JO": 30,
	"KZ": 20,
	"XK": 20,
	"KW": 30,
	"LV": 21,
	"LB": 28,
	"LY": 25,
	"LI": 21,
	"LT": 20,
	"LU": 20,
	"MK": 19,
	"MT": 31,
	"MR": 27,
	"MU": 30,
	"MC": 27,
	"MD": 24,
	"ME": 22,
	"NL": 18,
	"NO": 15,
	"PK": 24,
	"PS": 29,
	"PL": 28,
	"PT": 25,
	"QA": 29,
	"RO": 24,
	"LC": 32,
	"SM": 27,
	"ST": 25,
	"SA": 24,
	"RS": 22,
	"SC": 31,
	"SK": 24,
	"SI": 19,
	"ES": 24,
	"SD": 18,
	"SE": 24,
	"CH": 21,
	"TN": 24,
	"TR": 26,
	"UA": 29,
	"AE": 23,
	"GB": 22,
	"VA": 22,
	"VG": 24,
}
