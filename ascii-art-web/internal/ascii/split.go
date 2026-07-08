package ascii

import "strings"

func SplitText(str string) []string {
	output := strings.ReplaceAll(str, "\r\n", "\n")
	return strings.Split(output, "\n")

}
