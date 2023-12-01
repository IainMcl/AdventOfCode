package utils

import (
	"net/http"
	"os"
)

// ReadInput
// Reads the input for the current day of adventofcode.com
// Input can be found at https://adventofcode.com/2023/day/{day}/input where {day} is the current day as an integer
// Returns a string containing the input
func ReadInput(day rune) string {
	// Get session cookie from .env file
	sessionCookie := os.Getenv("SessionCookie")
	if sessionCookie == "" {
		panic("Session cookie not found in .env file")
	}

	url := "https://adventofcode.com/2023/day/" + string(day) + "/input"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Cookie", "session="+sessionCookie)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var input string
	for {
		buf := make([]byte, 1024)
		n, err := resp.Body.Read(buf)
		if n == 0 || err != nil {
			break
		}
		input += string(buf[:n])
	}
	return input
}
