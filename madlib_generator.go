package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// Define structure for JSON file
type MadlibData struct {
	Nouns      []string  `json:"Nouns"`
	Numbers    []float64 `json:"Numbers"`
	Adjectives []string  `json:"Adjectives"`
}

// Function to get a random element from a list
func getRandomElement(list interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	switch v := list.(type) {
	case []string:
		return v[rand.Intn(len(v))]
	case []float64:
		return v[rand.Intn(len(v))]
	default:
		return nil
	}
}

func main() {
	// Read and parse JSON file
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data MadlibData
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Read the Madlib text file
	madlibFile, err := ioutil.ReadFile("madlib.txt")
	if err != nil {
		fmt.Println("Error reading madlib file:", err)
		return
	}

	madlibText := string(madlibFile)

	// Replace placeholders with random values
	replacements := map[string]interface{}{
		"(noun)":      getRandomElement(data.Nouns),
		"(adjective)": getRandomElement(data.Adjectives),
		"(number)":    getRandomElement(data.Numbers),
	}

	// Regular expressions to replace placeholders in the madlib
	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(placeholder)
		madlibText = re.ReplaceAllString(madlibText, fmt.Sprintf("%v", replacement))
	}

	// Print the modified Madlib text
	fmt.Println(madlibText)
}
