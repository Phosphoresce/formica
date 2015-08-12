package main

import (
	"fmt"
	"os"
	"github.com/phosphoresce/formica/ecc"
	"github.com/phosphoresce/formica/utils"
)

func ec2menu() {
	// ec2menu variables
	var input string
	options := []string{"Entities", "Service Status", "Events", "Back", "Quit"}

	for {
		fmt.Printf("EC2 Menu\n\n")
		// print options
		utils.PrintOpts(options...)
		fmt.Printf("option... ")
		fmt.Scan(&input)
		switch input {
		case "1", "entities":
			// Check entities in region
			fmt.Printf("Checking entities...\n")
			utils.ClearScreen()
			ecc.EntityCount()
		case "2", "service", "status":
			// Check ec2 service status
			fmt.Printf("Checking EC2 Service Status...\n")
			utils.ClearScreen()
			ecc.ServiceHealth()
		case "3", "events":
			// check for amazon schedule events
			fmt.Printf("Checking for AWS scheduled events...\n")
			utils.ClearScreen()
			ecc.ScheduledEvents()
		case "4", "back":
			// back to main menu
			fmt.Printf("Returning to main menu...\n")
			utils.ClearScreen()
			return
		case "5", "q", "quit":
			// exit gracefully
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		default:
			utils.ClearScreen()
		}
	}
}

func main() {
	// main variables
	var input string
	options := []string{"EC2", "Quit"}

	// Runtime arguments

	// Interactive
	// Fresh terminal to start
	utils.ClearScreen()
	for {
		fmt.Printf("Main Menu\n\n")
		// print options
		utils.PrintOpts(options...)
		fmt.Printf("option... ")
		fmt.Scan(&input)
		switch input {
		case "1", "ec2", "EC2":
			// run ec2
			fmt.Printf("Loading EC2...\n")
			utils.ClearScreen()
			ec2menu()
		case "2", "q", "quit":
			// exit gracefully
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		default:
			utils.ClearScreen()
		}
	}
}
