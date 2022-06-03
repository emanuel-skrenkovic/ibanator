package iban

import (
	"strconv"
	"strings"
)

func mod97(digits string) string {
	number, _ := strconv.Atoi(digits)
	mod := number % 97

	modDigits := strconv.Itoa(mod)
	return modDigits
}

// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Modulo_operation_on_IBAN
func ValidateIBAN(iban string) bool {
	lowercase := strings.ToLower(iban)

	parsePosition := 9
	firstDigits := lowercase[:parsePosition]

	modDigits := mod97(firstDigits)

	ibanLen := len(iban)
	for parsePosition < ibanLen {
		var parseEnd int

		// Why doesn't Go have a min int function?
		if parsePosition+7 < ibanLen {
			parseEnd = parsePosition + 7
		} else {
			parseEnd = ibanLen
		}

		modDigits = modDigits + lowercase[parsePosition:parseEnd]
		modDigits = mod97(modDigits)

		parsePosition += 7
	}

	finalRemainder, err := strconv.Atoi(modDigits)
	if err != nil {
		// If we fail to parse the number after all
		// the previous calculations, then something
		// went horribly wrong.
		panic(err) // TODO: still need a better way
	}

	return finalRemainder == 1
}
