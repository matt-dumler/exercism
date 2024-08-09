package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/02/2006 15:04:05", date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, _ := time.Parse("January 2, 2006 15:04:05", date)

    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

    hour := t.Hour()

    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t, _ := time.Parse("1/2/2006 15:04:05", date)

    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
        t.Weekday().String(),
        t.Month(),
        t.Day(),
        t.Year(),
        t.Hour(),
        t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    anniversary := fmt.Sprintf("9/15/%d 00:00:00", time.Now().Year())

    t, _ := time.Parse("1/2/2006 15:04:05", anniversary)

	return t
}
