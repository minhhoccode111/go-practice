package markdown

import "strings"

func ReplaceHeader(input string) string {
	switch {
	case strings.HasPrefix(input, "###### "):
		return "<h6>" + strings.TrimPrefix(input, "###### ") + "</h6>"
	case strings.HasPrefix(input, "##### "):
		return "<h5>" + strings.TrimPrefix(input, "##### ") + "</h5>"
	case strings.HasPrefix(input, "#### "):
		return "<h4>" + strings.TrimPrefix(input, "#### ") + "</h4>"
	case strings.HasPrefix(input, "### "):
		return "<h3>" + strings.TrimPrefix(input, "### ") + "</h3>"
	case strings.HasPrefix(input, "## "):
		return "<h2>" + strings.TrimPrefix(input, "## ") + "</h2>"
	case strings.HasPrefix(input, "# "):
		return "<h1>" + strings.TrimPrefix(input, "# ") + "</h1>"
	default:
		return "<p>" + input + "</p>"
	}
}

func ReplaceTextStyle(input string) string {
	input = strings.Replace(input, "__", "<strong>", 1)
	input = strings.Replace(input, "__", "</strong>", 1)
	input = strings.Replace(input, "_", "<em>", 1)
	input = strings.Replace(input, "_", "</em>", 1)
	return input
}

func ReplaceList(input string) string {
	return "<li>" + strings.TrimPrefix(input, "* ") + "</li>"
}

// Render translates markdown to HTML
func Render(s string) string {
	splitedNewLines := strings.Split(s, "\n")
	result := make([]string, len(splitedNewLines))
	firstLi := -1
	lastLi := -1
	for i, line := range splitedNewLines {
		line = ReplaceTextStyle(line)
		switch {
		case strings.HasPrefix(line, "#"):
			result[i] = ReplaceHeader(line)
		case strings.HasPrefix(line, "*"):
			result[i] = ReplaceList(line)
			lastLi = i
			if firstLi == -1 {
				firstLi = i
			}
		default:
			result[i] = "<p>" + line + "</p>"
		}
	}
	if firstLi != -1 {
		result[firstLi] = "<ul>" + result[firstLi]
		result[lastLi] = result[lastLi] + "</ul>"
	}
	return strings.Join(result, "")
}
