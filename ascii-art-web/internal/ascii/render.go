package ascii

import(
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string{

	result := []string{}
	for i := 0; i <= 7; i++{
		var output strings.Builder
		for _, ch := range input{
			output.WriteString(banner[ch][i])
		}
		result = append(result, output.String())
	}
	return result
}