// Package twofer provides a function that produces formatted strings based on
// its input.
package twofer

import "fmt"

// ShareWith return formatted strings based on its input.
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }

	return fmt.Sprintf("One for %s, one for me.", name)
}
