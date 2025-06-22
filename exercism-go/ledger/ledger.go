package ledger

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Info struct {
	index int
	str   string
	err   error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	slices.SortFunc(entriesCopy, func(a, b Entry) int {
		if cmp := strings.Compare(a.Date, b.Date); cmp != 0 {
			return cmp
		}
		if cmp := strings.Compare(a.Description, b.Description); cmp != 0 {
			return cmp
		}
		return a.Change - b.Change
	})

	var s string
	switch locale {
	case "nl-NL":
		s = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	case "en-US":
		s = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	default:
		return "", errors.New("")
	}

	usFormat := "01/02/2006"
	nlFormat := "02-01-2006"
	inputFormat := "2006-01-02"
	for _, entry := range entriesCopy {
		date, err := time.Parse(inputFormat, entry.Date)
		if err != nil {
			return "", errors.New("Date is not valid")
		}
		desc := entry.Description
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		} else {
			desc = desc + strings.Repeat(" ", 25-len(desc))
		}
		var dateString string
		switch locale {
		case "nl-NL":
			dateString = date.Format(nlFormat)
		case "en-US":
			dateString = date.Format(usFormat)
		default:
			return "", errors.New("Locale is not valid")
		}
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}
		var a string
		switch locale {
		case "nl-NL":
			switch currency {
			case "EUR":
				a += "€"
			case "USD":
				a += "$"
			default:
				return "", errors.New("Currency is not valid")
			}
			a += " "
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + "."
			}
			a = a[:len(a)-1]
			a += ","
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += "-"
			} else {
				a += " "
			}
		case "en-US":
			if negative {
				a += "("
			}
			switch currency {
			case "EUR":
				a += "€"
			case "USD":
				a += "$"
			default:
				return "", errors.New("Cuerrency is not valid")
			}
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + ","
			}
			a = a[:len(a)-1]
			a += "."
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += ")"
			} else {
				a += " "
			}
		default:
			return "", errors.New("Locale is not valid")
		}
		var al int
		for range a {
			al++
		}
		s += dateString + strings.Repeat(" ", 10-len(dateString)) + " | " + desc + " | " + strings.Repeat(" ", 13-al) + a + "\n"
	}
	return s, nil
}
