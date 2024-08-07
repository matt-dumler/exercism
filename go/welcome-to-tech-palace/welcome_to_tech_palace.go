// The techpalace package implements welcome message functions.
package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    cutset := "\n*"
    oldMsg = strings.TrimLeft(oldMsg, cutset)
    oldMsg = strings.TrimRight(oldMsg, cutset)
    return strings.TrimSpace(oldMsg)
}
