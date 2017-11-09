package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
)

const dataFile = "data/data.json"

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// JsonData contains information we read from the Json file
type Bakery struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Batter  Batter  `json:"batters"`
	Toppings []ToppingsArray `json:"topping"`
}

type Batter struct {
	BatterId int `json:"battersId"`
	Batters []BattersArray `json:"batter"`
}

type BattersArray struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
}

type ToppingsArray struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
}

// main is the entry point for the program.
func main() {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Error occurred when trying to open the JSON data file %v/n", err.Error())
		os.Exit(1)
	}

	// Schedule the file to be closed once
	defer file.Close()

	// Decode the file into the Bakery type
	var bakery Bakery
	err = json.NewDecoder(file).Decode(&bakery)
	fmt.Printf("%v", bakery)
}