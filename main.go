package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <target_url>")
		return
	}

	target := os.Args[1]
	fmt.Println("Starting reconnaissance for:", target)

	fmt.Println("Running DNS Mapper...")
	if err := runCommand("go", "run", "dns_mapper/main.go", target); err != nil {
		fmt.Printf("DNS Mapper failed: %v\n", err)
	}

	fmt.Println("Running Fuzzing Tool...")
	if err := runCommand("go", "run", "fuzzing_tool/main.go", target, "fuzzing_tool/payloads.txt"); err != nil {
		fmt.Printf("Fuzzing Tool failed: %v\n", err)
	}

	fmt.Println("Running Port Scanner...")
	if err := runCommand("go", "run", "nmap_scan/main.go", "localhost", "3000", "3100"); err != nil {
		fmt.Printf("Port Scanner failed: %v\n", err)
	}

	fmt.Println("Generating final summary report...")
	generateSummaryReport()
}

func generateSummaryReport() {
	outputFile := "final_recon_report.txt"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating summary report: %v\n", err)
		return
	}
	defer file.Close()

	sections := map[string]string{
		"DNS Report":       "dns_report.txt",
		"Fuzzing Report":   "fuzzing_report.txt",
		"Port Scan Report": "nmap_report.txt",
	}

	for title, reportFile := range sections {
		content, err := os.ReadFile(reportFile)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", reportFile, err)
			continue
		}
		file.WriteString(fmt.Sprintf("\n--- %s ---\n", title))
		file.Write(content)
		file.WriteString("\n")
	}

	fmt.Println("Summary report saved to final_recon_report.txt")
}
