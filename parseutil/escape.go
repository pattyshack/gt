package parseutil

// This unescapes the content of a interpreted string/rune literal token.
//
// Note that this assumes the string was parsed by the tokenizer, and hence
// only contain valid utf8 runes and escape sequences. The content string
// excludes the string token leading/trailing delimiters.
//
// Note: we can't use strconv.Unquote due to different escape behaviors:
//   - this allows multiline string content
//   - go uses different escape sequence set depending on which delimiter
//     character is used.  This uses the same escape sequence set for all
//     the delimiters.
func Unescape(content string) string {
	remaining := []byte(content)
	result := make([]byte, 0, len(remaining))

	for len(remaining) > 0 {
		char := remaining[0]
		if char != '\\' {
			remaining = remaining[1:]
			result = append(result, char)
			continue
		}

		// handle escaped sequence

		basicEscape := true
		switch remaining[1] {
		case 'a':
			result = append(result, '\a')
		case 'b':
			result = append(result, '\b')
		case 'f':
			result = append(result, '\f')
		case 'n':
			result = append(result, '\n')
		case 'r':
			result = append(result, '\r')
		case 't':
			result = append(result, '\t')
		case 'v':
			result = append(result, '\v')
		case '\\':
			result = append(result, '\\')
		case '\'':
			result = append(result, '\'')
		case '"':
			result = append(result, '"')
		case '`':
			result = append(result, '`')
		case '\n': // line continuation
		default:
			basicEscape = false
		}

		if basicEscape {
			remaining = remaining[2:]
			continue
		}

		char = remaining[1]
		if IsOctalDigit(rune(char)) {
			value := rune(char-'0') << 6
			value |= rune(remaining[2]-'0') << 3
			value |= rune(remaining[3] - '0')
			result = append(result, []byte(string(value))...)
			remaining = remaining[4:]
			continue
		}

		var sequence []byte
		if char == 'x' {
			sequence = remaining[2:4]
			remaining = remaining[4:]
		} else if char == 'u' {
			sequence = remaining[2:6]
			remaining = remaining[6:]
		} else if char == 'U' {
			sequence = remaining[2:10]
			remaining = remaining[10:]
		} else {
			panic("should never happen")
		}

		var value rune
		for _, b := range sequence {
			value <<= 4
			if '0' <= b && b <= '9' {
				value |= rune(b - '0')
			} else if 'a' <= b && b <= 'f' {
				value |= rune(b - 'a' + 10)
			} else {
				value |= rune(b - 'A' + 10)
			}
		}
		result = append(result, []byte(string(value))...)
	}

	return string(result)
}
