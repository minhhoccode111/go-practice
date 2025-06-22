package ledger

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// Currency formatting configuration
type currencyConfig struct {
	symbol    string
	thousands string
	decimal   string
	negative  string
	positive  string
}

// Locale configuration
type localeConfig struct {
	dateFormat string
	header     string
	currency   currencyConfig
}

// Configuration maps
var localeConfigs = map[string]localeConfig{
	"en-US": {
		dateFormat: "01/02/2006",
		header:     "Date       | Description               | Change",
		currency: currencyConfig{
			thousands: ",",
			decimal:   ".",
			negative:  "()",
			positive:  " ",
		},
	},
	"nl-NL": {
		dateFormat: "02-01-2006",
		header:     "Datum      | Omschrijving              | Verandering",
		currency: currencyConfig{
			thousands: ".",
			decimal:   ",",
			negative:  "-",
			positive:  " ",
		},
	},
}

var currencySymbols = map[string]string{
	"USD": "$",
	"EUR": "€",
}

func Header(locale string) string {
	config, exists := localeConfigs[locale]
	if !exists {
		panic("Invalid locale")
	}
	return config.header + "\n"
}

// formatDate formats a date string according to the locale
func formatDate(dateStr, locale string) (string, error) {
	config, exists := localeConfigs[locale]
	if !exists {
		return "", fmt.Errorf("Invalid locale")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", errors.New("Date is not valid")
	}

	return date.Format(config.dateFormat), nil
}

// formatDescription formats and truncates description to fit column width
func formatDescription(desc string) string {
	const maxLength = 25
	if len(desc) > maxLength {
		return desc[:22] + "..."
	}
	return desc + strings.Repeat(" ", maxLength-len(desc))
}

// formatCurrency formats the change amount according to locale and currency
func formatCurrency(change int, currency, locale string) (string, error) {
	config, exists := localeConfigs[locale]
	if !exists {
		return "", fmt.Errorf("Invalid locale")
	}

	symbol, exists := currencySymbols[currency]
	if !exists {
		return "", fmt.Errorf("Invalid currency")
	}

	// Handle negative amounts
	negative := change < 0
	cents := change
	if negative {
		cents = -cents
	}

	// Convert to string and ensure proper decimal places
	centsStr := strconv.Itoa(cents)
	for len(centsStr) < 3 {
		centsStr = "0" + centsStr
	}

	// Split into dollars and cents
	dollars := centsStr[:len(centsStr)-2]
	centsPart := centsStr[len(centsStr)-2:]

	// Add thousands separators
	dollars = addThousandsSeparator(dollars, config.currency.thousands)

	// Build the formatted string, use strings.Builder for performance reason
	// in Go, every time we concatenate strings with `+` Go creates a new string
	// and copies all the data. strings.Builder is significantly faster for
	// multiple string operations
	var result strings.Builder

	if locale == "en-US" && negative {
		result.WriteString("(")
	}

	result.WriteString(symbol)
	if locale == "nl-NL" {
		result.WriteString(" ")
	}

	result.WriteString(dollars)
	result.WriteString(config.currency.decimal)
	result.WriteString(centsPart)

	if locale == "nl-NL" {
		if negative {
			result.WriteString("-")
		} else {
			result.WriteString(" ")
		}
	} else if locale == "en-US" && negative {
		result.WriteString(")")
	} else {
		result.WriteString(" ")
	}

	return result.String(), nil
}

// addThousandsSeparator adds thousands separators to a number string
func addThousandsSeparator(num, separator string) string {
	if len(num) <= 3 {
		return num
	}

	var parts []string
	for len(num) > 3 {
		parts = append(parts, num[len(num)-3:])
		num = num[:len(num)-3]
	}
	if len(num) > 0 {
		parts = append(parts, num)
	}

	// Reverse and join
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, separator)
}

// formatEntry formats a single ledger entry
func formatEntry(entry Entry, currency, locale string) (string, error) {
	// Format date
	dateStr, err := formatDate(entry.Date, locale)
	if err != nil {
		return "", err
	}

	// Format description
	desc := formatDescription(entry.Description)

	// Format currency
	changeStr, err := formatCurrency(entry.Change, currency, locale)
	if err != nil {
		return "", err
	}

	const changeWidth = 13

	// Use rune count for correct width, because of the € currency character
	changeLength := utf8.RuneCountInString(changeStr)
	changePadding := strings.Repeat(" ", changeWidth-changeLength)

	return fmt.Sprintf("%s | %s | %s%s\n", dateStr, desc, changePadding,
		changeStr), nil
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// Validate inputs
	if locale != "nl-NL" && locale != "en-US" {
		return "", fmt.Errorf("Invalid locale")
	}
	if currency != "EUR" && currency != "USD" {
		return "", fmt.Errorf("Invalid currency")
	}

	// Create a copy and sort entries base on Date > Description > Change
	entriesCopy := slices.Clone(entries)
	slices.SortFunc(entriesCopy, func(a, b Entry) int {
		if cmp := strings.Compare(a.Date, b.Date); cmp != 0 {
			return cmp
		}
		if cmp := strings.Compare(a.Description, b.Description); cmp != 0 {
			return cmp
		}
		return a.Change - b.Change
	})

	// Build the result, use strings.Builder for better performance
	var result strings.Builder
	result.WriteString(Header(locale))

	for _, entry := range entriesCopy {
		entryStr, err := formatEntry(entry, currency, locale)
		if err != nil {
			return "", err
		}
		result.WriteString(entryStr)
	}

	return result.String(), nil
}
