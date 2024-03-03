package main

import "fmt"

func commandHelp() error {
	fmt.Print("\n")
	fmt.Print("Welcome to Pokedex!\nUsage:\n")
	fmt.Printf("help: %v \nexit: %v \n", cliFunc()["help"].description, cliFunc()["exit"].description)

	return nil
}