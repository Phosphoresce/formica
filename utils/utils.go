package utils

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"os"
	"os/exec"
)

// UI utilities that should be made into a library
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintOpts(options ...string) {
	for x := 0; x < len(options); x++ {
		fmt.Printf("%v: %v\n", x+1, options[x])
	}
}

func EnterTo() {
	fmt.Printf("Press enter to continue... ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func HandleErr(err error) {
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
