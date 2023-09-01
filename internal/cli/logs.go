package cli

import "fmt"

func LogError(s string) {
	fmt.Printf("\nError: %s\n", s)
}
