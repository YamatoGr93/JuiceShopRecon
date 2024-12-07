package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func scanPort(protocol, hostname string, port int, results chan string) {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		results <- fmt.Sprintf("Port %d: Closed", port)
		return
	}
	defer conn.Close()

	// Attempt basic banner grabbing
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	banner := strings.TrimSpace(string(buffer))

	if err == nil && banner != "" {
		results <- fmt.Sprintf("Port %d: Open - Banner: %s", port, banner)
	} else {
		results <- fmt.Sprintf("Port %d: Open - Service: Unknown", port)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <hostname> <start_port> <end_port>")
		return
	}

	hostname := os.Args[1]
	startPort, _ := strconv.Atoi(os.Args[2])
	endPort, _ := strconv.Atoi(os.Args[3])

	reportFile := "nmap_report.txt"
	file, err := os.Create(reportFile)
	if err != nil {
		fmt.Printf("Error creating report file: %v\n", err)
		return
	}
	defer file.Close()

	write := func(data string) {
		fmt.Println(data)
		file.WriteString(data + "\n")
	}

	write(fmt.Sprintf("Scanning %s from port %d to %d...\n", hostname, startPort, endPort))
	results := make(chan string)

	for port := startPort; port <= endPort; port++ {
		go scanPort("tcp", hostname, port, results)
	}

	for port := startPort; port <= endPort; port++ {
		write(<-results)
	}

	write("Scan complete.")
	fmt.Printf("Port scan report saved to %s\n", reportFile)
}
