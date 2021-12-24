package romannumerals

import (
	"errors"
	"strings"
)

var ErrArabicNotInRange = errors.New("arabic value is not in range [0, 4000)")
var ErrRomanNotValid = errors.New("roman value is not valid")

func ConvertToRoman(arabic int) (string, error) {
	if arabic < 0 || arabic > 3999 {
		return "", ErrArabicNotInRange
	}
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String(), nil
}

func ConvertToArabic(roman string) (int, error) {
	total := 0
	symbols := windowedRoman(roman)
	if !symbols.isValid() {
		return 0, ErrRomanNotValid
	}
	for _, s := range symbols.Symbols() {
		total += allRomanNumerals.ValueOf(s...)
	}
	return total, nil
}

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var baseRomanNumerals = func(numerals romanNumerals) romanNumerals {
	var baseRomanNumerals romanNumerals
	for _, numeral := range numerals {
		if len(numeral.Symbol) == 1 {
			baseRomanNumerals = append(baseRomanNumerals, numeral)
		}
	}
	return baseRomanNumerals
}(allRomanNumerals)

type windowedRoman string

func (w windowedRoman) isValid() bool {
	for i := range w {
		if !contains(w[i], baseRomanNumerals) {
			return false
		}
	}
	return true
}

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)
		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func contains(symbol byte, numerals romanNumerals) bool {
	ssymbol := string(symbol)
	for _, numeral := range numerals {
		if ssymbol == numeral.Symbol {
			return true
		}
	}
	return false
}

func isSubtractive(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
