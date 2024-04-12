package main

import "fmt"

func commandHelp(*config, string) error {
	fmt.Print("\n")
	fmt.Print("Welcome to Pokedex!\n\nUsage:\n\n")
	fmt.Printf("help: %v \n\nexit: %v \n\nmap: %v \n\nmapb: %v \n\nexplore: %v\n", cliFunc()["help"].description, cliFunc()["exit"].description, cliFunc()["map"].description, cliFunc()["mapb"].description, cliFunc()["explore"].description);

	return nil
}