package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tomghaines/GoTLE/pkg/tle"
)

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt for satellite name
	fmt.Print("Enter the satellite name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	name = name[:len(name)-1] // Remove trailing newline character

	// Prompt for line 1
	fmt.Print("Enter the first line of TLE data: ")
	line1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	line1 = line1[:len(line1)-1] // Remove trailing newline character

	// Prompt for line 2
	fmt.Print("Enter the second line of TLE data: ")
	line2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	line2 = line2[:len(line2)-1] // Remove trailing newline character

	// Parse the TLE data
	tleData, err := tle.ParseTLE(name, line1, line2)
	if err != nil {
		log.Fatalf("Error parsing TLE: %v", err)
	}

	// Format the TLE data
	formattedTLE := tle.FormatTLE(tleData)
	fmt.Println("\nFormatted TLE Data:")
	fmt.Println(formattedTLE)
}
