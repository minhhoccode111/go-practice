package ledger

import (
	"errors"
	"fmt"
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

func Header(locale string) string {
	switch locale {
	case "nl-NL":
		return "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	case "en-US":
		return "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	default:
		panic("")
	}
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if locale != "nl-NL" && locale != "en-US" {
		return "", fmt.Errorf("Invalid locale")
	}

	if currency != "EUR" && currency != "USD" {
		return "", fmt.Errorf("Invalid currency")
	}

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

	s := Header(locale)

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
		}

		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}

		var changeString string
		switch locale {
		case "nl-NL":
			switch currency {
			case "EUR":
				changeString += "€"
			case "USD":
				changeString += "$"
			}
			changeString += " "
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
				changeString += parts[i] + "."
			}
			changeString = changeString[:len(changeString)-1]
			changeString += ","
			changeString += centsStr[len(centsStr)-2:]
			if negative {
				changeString += "-"
			} else {
				changeString += " "
			}
		case "en-US":
			if negative {
				changeString += "("
			}
			switch currency {
			case "EUR":
				changeString += "€"
			case "USD":
				changeString += "$"
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
				changeString += parts[i] + ","
			}
			changeString = changeString[:len(changeString)-1]
			changeString += "."
			changeString += centsStr[len(centsStr)-2:]
			if negative {
				changeString += ")"
			} else {
				changeString += " "
			}
		}
		var al int
		for range changeString {
			al++
		}
		s += dateString + strings.Repeat(" ", 10-len(dateString)) + " | " + desc + " | " + strings.Repeat(" ", 13-al) + changeString + "\n"
	}
	return s, nil
}
