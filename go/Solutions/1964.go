package Solutions

import "strings"

func reformatNumber(number string) string {
	number = strings.ReplaceAll(strings.ReplaceAll(number, " ", ""), "-", "")
	sb := strings.Builder{}
	length := len(number)
	for i := 0; i < length; {
		if length-i > 4 {

			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			sb.WriteByte('-')
			continue
		}
		if length-i == 4 {
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			sb.WriteByte('-')
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			break
		}
		if length-i == 3 {
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			break
		}
		if length-i == 2 {
			sb.WriteByte(number[i])
			i++
			sb.WriteByte(number[i])
			i++
			break
		}

	}
	return sb.String()
}
