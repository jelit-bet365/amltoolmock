package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Read existing user_statuses.json
	data, err := os.ReadFile("data/user_statuses.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal into a generic slice so we don't depend on model struct tags
	var records []map[string]interface{}
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal error: %v\n", err)
		os.Exit(1)
	}

	// Languages to distribute across records
	langs := []string{"EN", "DE", "FR", "ES", "IT"}

	for i := range records {
		records[i]["language"] = langs[i%len(langs)]
	}

	// Marshal back to pretty JSON
	out, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal error: %v\n", err)
		os.Exit(1)
	}

	// Write file back with trailing newline
	if err := os.WriteFile("data/user_statuses.json", append(out, '\n'), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write error: %v\n", err)
		os.Exit(1)
	}
}
