package main

import (
	"math"
	"sort"
	"strconv"
)

// Calcula media de um slice
func media(slice []float64) float64 {
	total := .0
	tamanho := float64(len(slice))

	for _, linha := range slice {
		total += linha
	}
	med := total / tamanho
	return med
}

// Calcula mediana de um slice
func mediana(slice []float64) float64 {
	tamanho := len(slice)
	metade := tamanho / 2
	intMetade := int(metade)
	med := .0

	if tamanho%2 == 0 {
		med = (slice[intMetade-1] + slice[intMetade]) / 2
	} else {
		med = slice[intMetade-1]
	}

	return med
}

// Calcula desvio padrão de um slice
func desvioPadrao(slice []float64, media float64) float64 {
	desvP := .0
	tamanho := float64(len(slice))

	for _, linha := range slice {
		desvP += math.Pow(linha-media, 2)
	}
	desvP = math.Sqrt(desvP / tamanho)

	return desvP
}

// Adiciona a coluna de survived a test
func adicionarColunaSurvived(testInical [][]string) [][]string {
	aux := make([][]string, 0, len(testInical))

	for _, linha := range testInical {
		linhaAux := make([]string, 0)
		linhaAux = append(linhaAux, linha[0])
		linhaAux = append(linhaAux, "-1")
		linhaAux = append(linhaAux, linha[1:]...)
		aux = append(aux, linhaAux)
	}
	return aux
}

// Preenche lista de idades consoante os atributos Sex e Pclass e ordena por ordem crescente
func preencherIdade(sex string, pclass string, listaIdades []float64, train [][]string) []float64 {
	aux := make([]float64, 0)

	for i, linha := range train {
		if linha[4] == sex && linha[2] == pclass && linha[5] != "" {
			aux = append(aux, listaIdades[i])
		}
	}
	// Ordenar lista
	sort.Float64s(aux)

	return aux
}

// Altera os espaços a branco para o atributo Age tendo em conta os atributos Sex e Pclass
func alterarBrancoIdade(train [][]string) {
	aux := make([]float64, 0)

	// Mapa das medianas das idades tendo em conta os atributos Sex e Pclass
	medianaIdades := map[string]float64{
		"feminino1":  .0,
		"feminino2":  .0,
		"feminino3":  .0,
		"masculino1": .0,
		"masculino2": .0,
		"masculino3": .0,
	}

	// Preenche a lista aux com todas as idades
	for _, linha := range train {
		i, _ := strconv.ParseFloat(linha[5], 64)
		aux = append(aux, i)
	}

	// Preenche as listas auxiliares tendo em conta os atributos Sex e Pclass
	auxFeminino1 := preencherIdade("female", "1", aux, train)
	auxFeminino2 := preencherIdade("female", "2", aux, train)
	auxFeminino3 := preencherIdade("female", "3", aux, train)
	auxMasculino1 := preencherIdade("male", "1", aux, train)
	auxMasculino2 := preencherIdade("male", "2", aux, train)
	auxMasculino3 := preencherIdade("male", "3", aux, train)

	// Calculo das medianas
	medianaIdades["feminino1"] = mediana(auxFeminino1)
	medianaIdades["feminino2"] = mediana(auxFeminino2)
	medianaIdades["feminino3"] = mediana(auxFeminino3)
	medianaIdades["masculino1"] = mediana(auxMasculino1)
	medianaIdades["masculino2"] = mediana(auxMasculino2)
	medianaIdades["masculino3"] = mediana(auxMasculino3)

	// Converter calores a branco pelo respetivo valor
	for _, linha := range train {
		if linha[5] == "" {
			if linha[4] == "female" {
				if linha[2] == "1" {
					linha[5] = strconv.FormatFloat(medianaIdades["feminino1"], 'E', -1, 64)
				} else if linha[2] == "2" {
					linha[5] = strconv.FormatFloat(medianaIdades["feminino2"], 'E', -1, 64)
				} else {
					linha[5] = strconv.FormatFloat(medianaIdades["feminino3"], 'E', -1, 64)
				}
			} else {
				if linha[2] == "1" {
					linha[5] = strconv.FormatFloat(medianaIdades["masculino1"], 'E', -1, 64)
				} else if linha[2] == "2" {
					linha[5] = strconv.FormatFloat(medianaIdades["masculino2"], 'E', -1, 64)
				} else {
					linha[5] = strconv.FormatFloat(medianaIdades["masculino3"], 'E', -1, 64)
				}
			}
		}
	}
}

// Altera os espaços a branco para os valores indicados pelo coeficiente de simetria (opcao 1 = age, 2 = fare)
func alterarBrancoCoeficienteSimetria(train [][]string, opcao int) {
	aux := make([]float64, 0)

	if opcao == 1 {
		for _, linha := range train {
			if linha[5] != "" {
				i, _ := strconv.ParseFloat(linha[5], 64)
				aux = append(aux, i)
			}
		}
	} else {
		for _, linha := range train {
			if linha[9] != "" {
				i, _ := strconv.ParseFloat(linha[9], 64)
				aux = append(aux, i)
			}
		}
	}

	// Ordenar lista
	sort.Float64s(aux)
	// Calculos auxiliares
	media := media(aux)
	mediana := mediana(aux)
	desvioPadrao := desvioPadrao(aux, media)

	// Calcula coeficiente de simetria utilizando o algoritmo Pearson's 2nd Skewness
	coeficienteSimetria := (3 * (media - mediana)) / desvioPadrao

	// Define o valor que vai ser utilizado para substituir os valores a branco
	var preenchimento string
	if coeficienteSimetria == 0 { // Os dados são simétricos
		preenchimento = strconv.FormatFloat(media, 'E', -1, 64) // Valores a brancos são iguais à média
	} else { // Os dados não são simétricos
		preenchimento = strconv.FormatFloat(mediana, 'E', -1, 64) // Valores a brancos são iguais à mediana
	}

	// Converter calores a branco pelo respetivo valor
	if opcao == 1 {
		for _, linha := range train {
			if linha[5] == "" {
				linha[5] = preenchimento
			}
		}
	} else {
		for _, linha := range train {
			if linha[9] == "" {
				linha[9] = preenchimento
			}
		}
	}
}

// Converter Locais de embarque a branco para S visto ser a grande maioria dos casos
func alterarBrancoEmbarque(train [][]string) {
	for _, linha := range train {
		if linha[11] == "" {
			linha[11] = "S"
		}
	}
}

// Converte valores em branco para valores pretendidos
func alterarBranco(train [][]string) {
	//alterarBrancoCoeficienteSimetria(train, 1)
	alterarBrancoIdade(train)
	alterarBrancoEmbarque(train)
	alterarBrancoCoeficienteSimetria(train, 2)
}

// Elimina atributos irrelevantes para construção de arvore de decisão
func eliminarAtributos(trainInicial [][]string) [][]string {
	aux := make([][]string, 0, len(trainInicial))
	for _, linha := range trainInicial {
		novaLinha := make([]string, 0, 9)
		for i := 0; i < 12; i++ {
			if !(i == 3 || i == 8 || i == 10) {
				novaLinha = append(novaLinha, linha[i])
			}
		}
		aux = append(aux, novaLinha)
	}
	return aux
}
