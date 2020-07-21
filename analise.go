package main

import (
	"fmt"
	"strconv"
)

// Verifies if the received passenger survived or not
func analisePassenger(passenger []string, decisionTree *tree) string {
	passengerInfo := passenger[2:]
	age, _ := strconv.ParseFloat(passengerInfo[2], 64)
	fare, _ := strconv.ParseFloat(passengerInfo[5], 64)
	result := "-1"
	node := decisionTree.root

	for result == "-1" {
		if node.attributeID == 0 { // Pclass
			if passengerInfo[0] == "1" {
				node = node.leaves[0]
			} else if passengerInfo[0] == "2" {
				node = node.leaves[1]
			} else {
				node = node.leaves[2]
			}
		} else if node.attributeID == 1 { // Sex
			if passengerInfo[0] == "female" {
				node = node.leaves[0]
			} else {
				node = node.leaves[1]
			}
		} else if node.attributeID == 2 { // Age
			if age <= 10 {
				node = node.leaves[0]
			} else if age <= 20 {
				node = node.leaves[1]
			} else if age <= 30 {
				node = node.leaves[2]
			} else if age <= 40 {
				node = node.leaves[3]
			} else if age <= 50 {
				node = node.leaves[4]
			} else if age <= 60 {
				node = node.leaves[5]
			} else if age <= 70 {
				node = node.leaves[6]
			} else {
				node = node.leaves[7]
			}
		} else if node.attributeID == 3 { // SibSp
			if passengerInfo[3] == "0" {
				node = node.leaves[0]
			} else if passengerInfo[3] == "1" {
				node = node.leaves[1]
			} else if passengerInfo[3] == "2" {
				node = node.leaves[2]
			} else if passengerInfo[3] == "3" {
				node = node.leaves[3]
			} else if passengerInfo[3] == "4" {
				node = node.leaves[4]
			} else if passengerInfo[3] == "5" {
				node = node.leaves[5]
			} else {
				node = node.leaves[6]
			}
		} else if node.attributeID == 4 { // Parch
			if passengerInfo[4] == "0" {
				node = node.leaves[0]
			} else if passengerInfo[4] == "1" {
				node = node.leaves[1]
			} else if passengerInfo[4] == "2" {
				node = node.leaves[2]
			} else if passengerInfo[4] == "3" {
				node = node.leaves[3]
			} else if passengerInfo[4] == "4" {
				node = node.leaves[4]
			} else if passengerInfo[4] == "5" {
				node = node.leaves[5]
			} else {
				node = node.leaves[6]
			}
		} else if node.attributeID == 5 { // Fare
			if fare <= 9 {
				node = node.leaves[0]
			} else if fare <= 16 {
				node = node.leaves[1]
			} else if fare <= 26 {
				node = node.leaves[2]
			} else if fare <= 200 {
				node = node.leaves[3]
			} else {
				node = node.leaves[4]
			}
		} else if node.attributeID == 6 { // Embarked
			if passengerInfo[6] == "C" {
				node = node.leaves[0]
			} else if passengerInfo[6] == "Q" {
				node = node.leaves[1]
			} else {
				node = node.leaves[2]
			}
		} else if node.attributeID == 7 {
			result = "1"
		} else {
			result = "0"
		}
	}

	return result
}

// Analises and prints the passenger survival state from the train file
func analiseTrain(train [][]string, decisionTree *tree) {
	incorrects := .0
	corrects := .0
	success := .0
	for i := 0; i < len(train); i++ {
		passenger := train[i]

		survivalTest := analisePassenger(passenger, decisionTree)
		if survivalTest == passenger[1] {
			corrects++
		} else {
			incorrects++
		}
	}
	success = corrects / (corrects + incorrects)
	fmt.Println("Number of correct cases: ", corrects)
	fmt.Println("Number of incorrect cases: ", incorrects)
	fmt.Println("Survival tax: ", success)
}

// Analises and returns a list of passengers with respective survival state based on a list of received passengers
func analisePassengers(passengers [][]string, decisionTree *tree) [][]string {
	output := make([][]string, 0, 419)
	outputHeader := make([]string, 0, 2)
	outputHeader = append(outputHeader, "PassengerId")
	outputHeader = append(outputHeader, "Survived")
	output = append(output, outputHeader)

	for _, row := range passengers {
		outputRow := make([]string, 0, 2)
		outputRow = append(outputRow, row[0])
		survivalTest := analisePassenger(row, decisionTree)
		outputRow = append(outputRow, survivalTest)
		output = append(output, outputRow)
	}

	return output
}
