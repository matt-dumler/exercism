package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
    appIcons := map[string]rune{
        "recommendation": '\u2757',
        "search": '\U0001F50D',
        "weather": '\u2600',
    }

    for _, r := range log {
        for app, icon := range appIcons {
            if r == icon {
                return app
            }
        }
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= len([]rune(log))
}
