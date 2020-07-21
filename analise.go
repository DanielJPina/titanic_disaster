package main

import (
	"fmt"
	"strconv"
)

// Verifica se o passageiro recebido sobreviveu ou não
func analisarPassageiro(passageiro []string, decisionTree *tree) string {
	infoPassageiro := passageiro[2:]
	age, _ := strconv.ParseFloat(infoPassageiro[2], 64)
	fare, _ := strconv.ParseFloat(infoPassageiro[5], 64)
	resultado := "-1"
	node := decisionTree.root

	for resultado == "-1" {
		//fmt.Println("Atributo a verificar: ", node.attributeID)
		if node.attributeID == 0 { // Se o atributo do node for Pclass
			if infoPassageiro[0] == "1" { // Se o Pclass do passageiro for 1
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if infoPassageiro[0] == "2" { // Se o Pclass do passageiro for 2
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else { // Caso contrário
				node = node.leaves[2] // O node passa a ser a terceira folha
			}
		} else if node.attributeID == 1 { // Se o atributo do node for Sex
			if infoPassageiro[0] == "female" { // Se o Sex do passageiro for female
				node = node.leaves[0] // O node passa a ser a primeira
			} else { // Caso contrário
				node = node.leaves[1] // O node passa a ser a segunda folha
			}
		} else if node.attributeID == 2 { // Se o atributo do node for Age
			if age <= 10 { // Se a idade do passageiro for menor ou igual a 10
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if age <= 20 { // Se a idade do passageiro for superior a 10 e inferior ou igual a 20
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else if age <= 30 { // Se a idade do passageiro for superior a 20 e inferior ou igual a 30
				node = node.leaves[2] // O node passa a ser a terceira folha
			} else if age <= 40 { // Se a idade do passageiro for superior a 30 e inferior ou igual a 40
				node = node.leaves[3] // O node passa a ser a quarta folha
			} else if age <= 50 { // Se a idade do passageiro for superior a 40 e inferior ou igual a 50
				node = node.leaves[4] // O node passa a ser a quinta folha
			} else if age <= 60 { // Se a idade do passageiro for superior a 50 e inferior ou igual a 60
				node = node.leaves[5] // O node passa a ser a sexta folha
			} else if age <= 70 { // Se a idade do passageiro for superior a 60 e inferior ou igual a 70
				node = node.leaves[6] // O node passa a ser a setima folha
			} else { // Caso contrário
				node = node.leaves[7] // O node passa a ser a outava folha
			}
		} else if node.attributeID == 3 { // Se o atributo do node for SibSp
			if infoPassageiro[3] == "0" { // Se o numero de irmaos/conjugues for 0
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if infoPassageiro[3] == "1" { // Se o numero de irmaos/conjugues for 1
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else if infoPassageiro[3] == "2" { // Se o numero de irmaos/conjugues for 2
				node = node.leaves[2] // O node passa a ser a terceira folha
			} else if infoPassageiro[3] == "3" { // Se o numero de irmaos/conjugues for 3
				node = node.leaves[3] // O node passa a ser a quarta folha
			} else if infoPassageiro[3] == "4" { // Se o numero de irmaos/conjugues for 4
				node = node.leaves[4] // O node passa a ser a quinta folha
			} else if infoPassageiro[3] == "5" { // Se o numero de irmaos/conjugues for 5
				node = node.leaves[5] // O node passa a ser a sexta folha
			} else { // Caso contrário
				node = node.leaves[6] // O node passa a ser a setima folha
			}
		} else if node.attributeID == 4 { // Se o atributo do node for Parch
			if infoPassageiro[4] == "0" { // Se o numero de pais/filhos for 0
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if infoPassageiro[4] == "1" { // Se o numero de pais/filhos for 1
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else if infoPassageiro[4] == "2" { // Se o numero de pais/filhos for 2
				node = node.leaves[2] // O node passa a ser a terceira folha
			} else if infoPassageiro[4] == "3" { // Se o numero de pais/filhos for 3
				node = node.leaves[3] // O node passa a ser a quarta folha
			} else if infoPassageiro[4] == "4" { // Se o numero de pais/filhos for 4
				node = node.leaves[4] // O node passa a ser a quinta folha
			} else if infoPassageiro[4] == "5" { // Se o numero de pais/filhos for 5
				node = node.leaves[5] // O node passa a ser a sexta folha
			} else { // Caso contrário
				node = node.leaves[6] // O node passa a ser a setima folha
			}
		} else if node.attributeID == 5 { // Se o atributo do node for Fare
			if fare <= 9 { // Se a tarifa que o passageiro pagou foi menor ou igual a 9
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if fare <= 16 { // Se a tarifa que o passageiro pagou foi superior a 9 e inferior ou igual a 16
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else if fare <= 26 { // Se a tarifa que o passageiro pagou foi superior a 16 e inferior ou igual a 26
				node = node.leaves[2] // O node passa a ser a terceira folha
			} else if fare <= 200 { // Se a tarifa que o passageiro pagou foi superior a 26 e inferior ou igual a 200
				node = node.leaves[3] // O node passa a ser a quarta folha
			} else { // Caso contrário
				node = node.leaves[4] // O node passa a ser a quinta folha
			}
		} else if node.attributeID == 6 { // Se o atributo do node for Embarked
			if infoPassageiro[6] == "C" { // Se o porto de embarque tiver sido Cherbourg
				node = node.leaves[0] // O node passa a ser a primeira folha
			} else if infoPassageiro[6] == "Q" { // Se o porto de embarque tiver sido Queenstown
				node = node.leaves[1] // O node passa a ser a segunda folha
			} else { // Caso contrário
				node = node.leaves[2] // O node passa a ser a terceira folha
			}
		} else if node.attributeID == 7 {
			resultado = "1"
		} else {
			resultado = "0"
		}
	}

	return resultado
}

// Analisa e imprime a sobrevivencia ou não da lista de passageiros do ficheiro Train
func analisarTreino(train [][]string, decisionTree *tree) {
	incorrectos := .0
	correctos := .0
	sucesso := .0
	for i := 0; i < len(train); i++ {
		passageiro := train[i]

		testeSobrevivencia := analisarPassageiro(passageiro, decisionTree)
		if testeSobrevivencia == passageiro[1] {
			correctos++
		} else {
			incorrectos++
		}
	}
	sucesso = correctos / (correctos + incorrectos)
	fmt.Println("Numero de casos correctos: ", correctos)
	fmt.Println("Numero de casos incorrectos: ", incorrectos)
	fmt.Println("Taxa de sucesso: ", sucesso)
}

// Analisa e devolve uma lista dos passageiros sobreviventes ou não baseados na lista de passageiros recebidos
func analisarPassageiros(passageiros [][]string, decisionTree *tree) [][]string {
	output := make([][]string, 0, 419)
	outputHeader := make([]string, 0, 2)
	outputHeader = append(outputHeader, "PassengerId")
	outputHeader = append(outputHeader, "Survived")
	output = append(output, outputHeader)

	for _, linha := range passageiros {
		raw := make([]string, 0, 2)
		raw = append(raw, linha[0])
		testeSobrevivencia := analisarPassageiro(linha, decisionTree)
		raw = append(raw, testeSobrevivencia)
		output = append(output, raw)
	}

	return output
}
