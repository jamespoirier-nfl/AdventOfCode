package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Constants for cube counts
const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

func main() {
	// Read input from file
	input, err := os.ReadFile("day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Process the input
	lines := strings.Split(string(input), "\n")

	// Variable to track the total sum of IDs for games where all rounds are possible
	totalSum := 0

	// Variable to track the total cube power for all games
	totalPower := 0

	// Process each line
	for _, line := range lines {
		// Your code logic for each line goes here
		fmt.Println(line)

		// Check if the line represents a game and process it
		if strings.HasPrefix(line, "Game") {
			// Process the game and find the minimum cubes required
			minCubes, isPossible := processGame(line)

			// Calculate and accumulate the cube power for the game, even if it has non-valid rounds
			gamePower := calculateCubePower(minCubes)
			fmt.Printf("Cube Power for Game: %d\n", gamePower)
			totalPower += gamePower

			// If all rounds are possible, add the game number to the total sum
			if isPossible {
				// Extract the game number and add it to the total sum
				gameNumber := extractGameNumber(line)
				totalSum += gameNumber
			}
		}
	}

	// Print the total sum of IDs for games where all rounds are possible
	fmt.Printf("Total Sum of IDs for Possible Games: %d\n", totalSum)

	// Print the total cube power for all games
	fmt.Printf("Total Cube Power for All Games: %d\n", totalPower)
}

// Function to process a game and return the minimum cubes required and whether all rounds are possible
func processGame(line string) (minCubes int, isPossible bool) {
	// Extract the game number and subgames
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	parts := strings.Split(line, ":")
	subgames := strings.Split(parts[1], ";") // Split into individual subgames

	// Variables to track the maximum count of each color across all rounds in a game
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	// Process each subgame
	for _, subgame := range subgames {
		// Process each round in the subgame and update the maximum count for each color
		maxRed, maxGreen, maxBlue = processSubgame(subgame, maxRed, maxGreen, maxBlue)
	}

	// Calculate the minimum cubes required for the game
	minCubes = maxRed * maxGreen * maxBlue

	// Check if all rounds are possible in the game
	return minCubes, maxRed <= RedCubes && maxGreen <= GreenCubes && maxBlue <= BlueCubes
}

// Function to calculate the cube power for a game
func calculateCubePower(minCubes int) int {
	return minCubes
}

// Function to process each subgame
func processSubgame(subgame string, maxRed, maxGreen, maxBlue int) (int, int, int) {
	// Split subgame into individual cube counts
	// e.g., 3 blue, 4 red
	cubeCounts := strings.Split(strings.TrimSpace(subgame), ";")

	// Process each round in the subgame
	for _, round := range cubeCounts {
		// Variables to track total cube count for each color in the round
		totalRed := 0
		totalGreen := 0
		totalBlue := 0

		// Split the round into individual cube counts
		cubeCountsInRound := strings.Split(strings.TrimSpace(round), ",")

		// Process each cube count in the round
		for _, cubeCount := range cubeCountsInRound {
			parts := strings.Split(strings.TrimSpace(cubeCount), " ")

			// Check if there are exactly two parts (count and color)
			if len(parts) != 2 {
				fmt.Printf("Invalid cube count: %s\n", cubeCount)
				return maxRed, maxGreen, maxBlue
			}

			// Extract the color and count
			color := strings.TrimSpace(parts[1])
			countStr := strings.TrimSpace(parts[0])

			// Check if the count is a valid integer
			cubeCount, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Printf("Invalid cube count: %s\n", countStr)
				return maxRed, maxGreen, maxBlue
			}

			// Update the maximum count for each color
			switch color {
			case "red":
				if cubeCount > maxRed {
					maxRed = cubeCount
				}
				totalRed += cubeCount
			case "green":
				if cubeCount > maxGreen {
					maxGreen = cubeCount
				}
				totalGreen += cubeCount
			case "blue":
				if cubeCount > maxBlue {
					maxBlue = cubeCount
				}
				totalBlue += cubeCount
			default:
				// Handle invalid color
				fmt.Printf("Invalid color: %s\n", color)
				return maxRed, maxGreen, maxBlue
			}
		}

		// Print whether the round is valid or not
		fmt.Printf("Round: %s\n", isValidRound(totalRed, totalGreen, totalBlue))
	}

	// Return the updated maximum count for each color
	return maxRed, maxGreen, maxBlue
}

// Function to determine if a round is valid
func isValidRound(totalRed, totalGreen, totalBlue int) string {
	if totalRed <= RedCubes && totalGreen <= GreenCubes && totalBlue <= BlueCubes {
		return "Valid"
	}
	return "Not Valid"
}

// Function to extract the game number
func extractGameNumber(line string) int {
	// Extract the game number from the line
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	parts := strings.Split(line, ":")
	gameNumber, err := strconv.Atoi(strings.TrimSpace(parts[0][5:]))
	if err != nil {
		log.Fatal(err)
	}
	return gameNumber
}
