package main

import "os"

func writeFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

