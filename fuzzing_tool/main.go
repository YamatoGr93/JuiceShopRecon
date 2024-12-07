package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func fuzzEndpoint(url string, payloads []string, reportFile string) {
	file, err := os.Create(reportFile)
	if err != nil {
		fmt.Printf("Error creating report file: %v\n", err)
		return
	}
	defer file.Close()

	write := func(data string) {
		fmt.Print(data)
		file.WriteString(data + "\n")
	}

	write(fmt.Sprintf("Fuzzing URL: %s\n", url))
	write(fmt.Sprintf("Date: %s\n", time.Now().Format(time.RFC1123)))
	for _, payload := range payloads {
		fullURL := fmt.Sprintf("%s%s", url, payload)
		resp, err := http.Get(fullURL)
		if err != nil {
			write(fmt.Sprintf("Payload: %s\nError: %v\n", payload, err))
			continue
		}
		write(fmt.Sprintf("Payload: %s\nStatus Code: %d\n", payload, resp.StatusCode))
		resp.Body.Close()
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <url> <payload_file>")
		return
	}

	url := os.Args[1]
	payloadFile := os.Args[2]
	reportFile := "fuzzing_report.txt"

	file, err := os.Open(payloadFile)
	if err != nil {
		fmt.Printf("Error reading payload file: %v\n", err)
		return
	}
	defer file.Close()

	payloads := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payloads = append(payloads, scanner.Text())
	}

	fmt.Printf("Starting fuzzing on %s\n", url)
	fuzzEndpoint(url, payloads, reportFile)
	fmt.Printf("Fuzzing report saved to %s\n", reportFile)
}
