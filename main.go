package main

import (
	"flag"
	"fmt"
)

var (
	userName string
)

func init() {
	flag.StringVar(&userName, "userName", "userName", "Input your user name")
}

func main() {
	flag.Parse()
	fmt.Println(flag.Args(), validateArgs())
}

func validateArgs() error {
	flag.Parse()
	if !validateUserName() {
		return fmt.Errorf("error: Invalid user name %s", userName)
	}
	return nil
}

// TODO: Implement user check logic
func validateUserName() bool {
	return userName == "niwaniwa"
}
