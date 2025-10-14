package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 1. Create a custom type based on time.Time
type UnixTime time.Time

// 2. Implement the json.Unmarshaler interface for the custom type
func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	// If the value is an empty object {} or null, treat it as zero time
	if s == "{}" || s == "null" {
		*ut = UnixTime(time.Time{}) // Sets to the zero value for time
		return nil
	}

	// Otherwise, proceed with parsing it as a number
	timestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// If it fails to parse as a number, also treat as zero time
		*ut = UnixTime(time.Time{})
		return nil // Return nil to avoid stopping the whole unmarshal process
	}

	*ut = UnixTime(time.Unix(timestamp, 0))

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (ut UnixTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ut)

	// Handle the zero-value case
	if t.IsZero() {
		return []byte("null"), nil
	}

	// 1. Format the time into a string using a standard layout
	//    time.RFC3339 is "2006-01-02T15:04:05Z07:00"
	formattedString := t.Format(time.RFC3339)

	// 2. Wrap the string in quotes to make it a valid JSON string
	quotedString := fmt.Sprintf(`"%s"`, formattedString)

	return []byte(quotedString), nil
}
