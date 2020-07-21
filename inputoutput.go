package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Load CSV files and return slice of slices of strings
func loadFiles(inputFile string) [][]string {
	// read train.csv file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// create a reader
	r := csv.NewReader(file)

	// Read each row from the CSV file
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// Exports a slice of slices of strings from a CSV file
func exportCSV(output [][]string) {
	// Creat submission.csv file on the executable's folder
	file, _ := os.Create("submission.csv")

	// Creates a writer
	w := csv.NewWriter(file)

	// Writes on file row by row
	for _, record := range output {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error to write to csv file:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// Exports a string to a .txt file
func exportTxt(texto string) {
	// Creates decisionTree.txt file on the executable's folder
	file, _ := os.Create("decisionTree.txt")

	// Creates a writer
	w := bufio.NewWriter(file)

	_, err := fmt.Fprintf(w, "%s", texto)
	if err != nil {
		panic(err)
	}

	w.Flush()
	file.Close()
}
