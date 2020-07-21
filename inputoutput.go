package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Carrega ficheiros do tipo CSV e retorna uma slice de slices de strings
func carregarFicheiro(inputFile string) [][]string {
	// Ler ficheiro train.csv
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// cria um reader
	r := csv.NewReader(file)

	// Ler Cada linha do ficheiro csv
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// Exporta uma slice de slices de strings para um ficheiro CSV
func exportarCSV(output [][]string) {
	// cria ficheiro submission.csv na pasta do executável
	file, _ := os.Create("submission.csv")

	// cria um writer
	w := csv.NewWriter(file)

	// Escreve no ficheiro linha a linha
	for _, record := range output {
		if err := w.Write(record); err != nil {
			log.Fatalln("Erro a escrever para o ficheiro csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func exportarTxt(texto string) {
	// cria ficheiro decisionTree.txt na pasta do executável
	file, _ := os.Create("decisionTree.txt")

	// cria um writer
	w := bufio.NewWriter(file)

	_, err := fmt.Fprintf(w, "%s", texto)
	if err != nil {
		panic(err)
	}

	w.Flush()
	file.Close()
}
