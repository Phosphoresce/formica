package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// UI utilities that should be made into a library
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printOpts(options ...string) {
	for x := 0; x < len(options); x++ {
		fmt.Printf("%v: %v\n", x+1, options[x])
	}
}

func handleErr(err error) {
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
// -- end -- //

// EC2 functions that should be libraried
func entityCount() {
	regions := []string{"us-east-1", "us-west-1", "us-west-2"}
	for x := 0; x < len(regions); x++ {
		client := ec2.New(&aws.Config{Region: regions[x]})
		resp, err := client.DescribeInstances(nil)
		handleErr(err)
		fmt.Printf("Region: %v\n-----------------\n", regions[x])
		fmt.Printf("%15.15v %15.15v %10.10v %16.16v\n", "Instance:", "Key Pair:", "State:", "Security Groups:")
		for y := 0; y < len(resp.Reservations); y++ {
			for z := 0; z < len(resp.Reservations[y].Instances); z++ {
				var tag, state, key string
				if resp.Reservations[y].Instances[z].Tags[0].Value != nil {
					tag = *resp.Reservations[y].Instances[z].Tags[0].Value
				}
				if resp.Reservations[y].Instances[z].State.Name != nil {
					state = *resp.Reservations[y].Instances[z].State.Name
				}
				if resp.Reservations[y].Instances[z].KeyName != nil {
					key = *resp.Reservations[y].Instances[z].KeyName
				}
				fmt.Printf("%15.15v %15.15v %10.10v [", tag, key, state)
				for a := 0; a < len(resp.Reservations[y].Instances[z].SecurityGroups); a++ {
					var group string
					if resp.Reservations[y].Instances[z].SecurityGroups[a] != nil {
						group = *resp.Reservations[y].Instances[z].SecurityGroups[a].GroupName
					}
					if a < len(resp.Reservations[y].Instances[z].SecurityGroups)-1 {
						fmt.Printf("%v, ", group)
					} else {
						fmt.Printf("%v", group)
					}
				}
				fmt.Printf("]\n")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Press enter to continue... ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearScreen()
}

func serviceHealth() {
	// temp array maybe pass this in??
	regions := []string{"us-east-1", "us-west-1", "us-west-2"}
	for x := 0; x < len(regions); x++ {
		client := ec2.New(&aws.Config{Region: regions[x]})
		resp, err := client.DescribeAvailabilityZones(nil)
		handleErr(err)
		fmt.Printf("Region: %v\n-----------------\n", regions[x])
		for y := 0; y < len(resp.AvailabilityZones); y++ {
			fmt.Printf("Zone: %v... %v\n", *resp.AvailabilityZones[y].ZoneName, *resp.AvailabilityZones[y].State)
		}
		fmt.Println()
	}
	fmt.Printf("Press enter to continue... ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearScreen()
}
// -- end -- //

func ec2menu() {
	// ec2menu variables
	var input string
	options := []string{"Entities", "Service Status", "Back", "Quit"}

	for {
		fmt.Printf("EC2 Menu\n\n")
		// print options
		printOpts(options...)
		fmt.Printf("option... ")
		fmt.Scan(&input)
		switch input {
		case "1", "entities":
			// Check entities in region
			fmt.Printf("Checking entities...\n")
			clearScreen()
			entityCount()
		case "2", "service", "status":
			// Check ec2 service status
			fmt.Printf("Checking EC2 Service Status...\n")
			clearScreen()
			serviceHealth()
		case "3", "back":
			// back to main menu
			fmt.Printf("Returning to main menu...\n")
			clearScreen()
			return
		case "4", "q", "quit":
			// exit gracefully
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		default:
			clearScreen()
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
	clearScreen()
	for {
		fmt.Printf("Main Menu\n\n")
		// print options
		printOpts(options...)
		fmt.Printf("option... ")
		fmt.Scan(&input)
		switch input {
		case "1", "ec2", "EC2":
			// run ec2
			fmt.Printf("Loading EC2...\n")
			clearScreen()
			ec2menu()
		case "2", "q", "quit":
			// exit gracefully
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
