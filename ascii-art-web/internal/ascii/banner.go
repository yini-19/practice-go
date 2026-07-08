package ascii

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func BannerLoader(input string) (map[rune][]string, error) {
	data, err := os.ReadFile(input)
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	if err != nil {
		fmt.Println("error while reading file")
		os.Exit(1)
	}

	line := strings.Split(string(data), "\n")
	if len(line) > 856 {
		return nil, errors.New("incomplete banner file")
	}
	font := make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		begin := (int(char) - 32) * 9
		font[char] = line[begin+1 : begin+9]
	}
	return font, nil
}
