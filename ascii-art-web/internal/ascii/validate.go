package ascii

import "fmt"

func ValidInput(input string) error {

	for _, r := range input {
		if r == '\n' || r == '\r' {
			continue
		}
		if r < 32 || r > 126 {
			return fmt.Errorf("%v is not a valid ascii character", r)
		}
	}
	return nil
}
