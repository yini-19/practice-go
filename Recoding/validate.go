// ...existing code...
package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
    if len(banner) == 0 {
        return fmt.Errorf("banner map is empty")
    }
    if len(banner) != 95 {
        return fmt.Errorf("banner must contain 95 characters (printable ASCII 32..126), got %d", len(banner))
    }
    for key, value := range banner {
        if key < rune(32) || key > rune(126) {
            return fmt.Errorf("invalid key: %q (must be printable ASCII 32..126)", key)
        }
        if len(value) != 8 {
            return fmt.Errorf("banner for %q must have 8 rows, got %d", key, len(value))
        }
    }
    return nil
}
// ...existing code...