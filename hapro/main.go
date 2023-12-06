package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const haproxyConfigPath = "/etc/haproxy/haproxy.cfg"

func main() {
	if reusedPorts, err := findReusedPorts(haproxyConfigPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else if len(reusedPorts) > 1 {
		fmt.Printf("Reused ports found in HAProxy configuration:\n%s\n", reusedPorts)
	} else {
		fmt.Println("No reused ports found in HAProxy configuration.")
	}
}

func findReusedPorts(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reusedPorts []string
	portSet := make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isBindLine(line) {
			port := extractPortFromBind(line)
			if _, exists := portSet[port]; exists {
				reusedPorts = append(reusedPorts, port)
			} else {
				portSet[port] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reusedPorts, nil
}

func isBindLine(line string) bool {
	return regexp.MustCompile(`^\s*bind\s`).MatchString(line)
}

func extractPortFromBind(line string) string {
	re := regexp.MustCompile(`\*\:(\d+)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
