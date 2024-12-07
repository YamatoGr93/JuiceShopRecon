package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

func getDNSRecords(domain string) map[string]string {
	results := make(map[string]string)
	var aRecords, mxRecords, txtRecords, cnameRecord string

	// A Records
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		aRecords += fmt.Sprintf("- %s\n", ip)
	}

	// CNAME Record
	cname, _ := net.LookupCNAME(domain)
	cnameRecord = cname

	// MX Records
	mx, _ := net.LookupMX(domain)
	for _, m := range mx {
		mxRecords += fmt.Sprintf("- %s %d\n", m.Host, m.Pref)
	}

	// TXT Records
	txt, _ := net.LookupTXT(domain)
	for _, t := range txt {
		txtRecords += fmt.Sprintf("- %s\n", t)
	}

	results["A_RECORDS"] = strings.TrimSpace(aRecords)
	results["CNAME_RECORD"] = cnameRecord
	results["MX_RECORDS"] = strings.TrimSpace(mxRecords)
	results["TXT_RECORDS"] = strings.TrimSpace(txtRecords)
	return results
}

func applyTemplate(templatePath string, replacements map[string]string) string {
	template, err := os.ReadFile(templatePath)
	if err != nil {
		return "Error reading template: " + err.Error()
	}

	content := string(template)
	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, "{{"+placeholder+"}}", value)
	}
	return content
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain>")
		return
	}

	rawTarget := os.Args[1]
	domain := rawTarget
	if strings.Contains(rawTarget, "http") {
		u, err := url.Parse(rawTarget)
		if err == nil {
			domain = u.Hostname()
		}
	}

	fmt.Printf("Fetching DNS records for: %s\n\n", domain)
	templatePath := "dns_template.txt"
	reportPath := "dns_report.txt"

	data := getDNSRecords(domain)
	data["TARGET"] = domain
	data["DATE"] = time.Now().Format(time.RFC1123)

	report := applyTemplate(templatePath, data)

	err := os.WriteFile(reportPath, []byte(report), 0644)
	if err != nil {
		fmt.Printf("Error saving report: %v\n", err)
		return
	}

	fmt.Printf("DNS report saved to %s\n", reportPath)
}
