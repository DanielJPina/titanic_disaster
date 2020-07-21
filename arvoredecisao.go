package main

import (
	"fmt"
	"strconv"
)

// Definição da estrutura do nó
type node struct {
	attributeID int     // Identifica o node: 0=Pclass, 1=Sex, 2=Age, 3=SibSp, 4=Parch, 5=Fare, 6=Embarked, 7=Survided, 8=!Survived
	parent      *node   // Pai do nó
	path        []int   // Caminho da root até ao nó
	leaves      []*node // Folhas do nó
}

// Definição da estrutura da arvore
type tree struct {
	root *node
}

// Imprime nó
func (n *node) Print() {
	fmt.Println("attributeID: ", n.attributeID)
	fmt.Println("parent: ", n.parent.attributeID)
	fmt.Println("path: ", n.path)
	for i, leaf := range n.leaves {
		if leaf != nil {
			fmt.Println("leaf ", i, " = ", leaf.attributeID)
		} else if leaf == nil && i == 0 {
			fmt.Println("This node doesnt have leaves")
		}
	}
}

// Define nós de sobrevivencia ou não sobrevivencia
func sobrevivencia(train [][]string) int {
	attrID := 0
	surv := 0
	nSurv := 0
	for _, linha := range train {
		if linha[1] == "1" {
			surv++
		} else {
			nSurv++
		}
	}
	if surv > nSurv {
		attrID = 7
	} else {
		attrID = 8
	}
	return attrID
}

// cria um nó
func criarNo(attrID int, nodeInput *node) *node {
	node := &node{
		attributeID: attrID,
		parent:      nodeInput,
		path:        nodeInput.path,
		leaves:      make([]*node, 0),
	}
	return node
}

// Selecciona as linhas do ficheiro csv cujo valor do atributo especifico é igual ao especificado
func selectValue(train [][]string, attributeID int, attributeValue string) [][]string {
	aux := make([][]string, 0)
	if attributeID == 4 { // Se o atributo for Age
		aux2 := make([]float64, 0)

		for _, linha := range train {
			i, _ := strconv.ParseFloat(linha[4], 64)
			aux2 = append(aux2, i)
		}
		if attributeValue == "x<=10" {
			for i, linha := range train {
				if aux2[i] <= 10 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "10<x<=20" {
			for i, linha := range train {
				if aux2[i] <= 20 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "20<x<=30" {
			for i, linha := range train {
				if aux2[i] <= 30 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "30<x<=40" {
			for i, linha := range train {
				if aux2[i] <= 40 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "40<x<=50" {
			for i, linha := range train {
				if aux2[i] <= 50 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "50<x<=60" {
			for i, linha := range train {
				if aux2[i] <= 60 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "60<x<=70" {
			for i, linha := range train {
				if aux2[i] <= 70 {
					aux = append(aux, linha)
				}
			}
		} else {
			for i, linha := range train {
				if aux2[i] > 70 {
					aux = append(aux, linha)
				}
			}
		}
	} else if attributeID == 7 { // Se o atributo for Fare
		aux2 := make([]float64, 0)

		for _, linha := range train {
			i, _ := strconv.ParseFloat(linha[7], 64)
			aux2 = append(aux2, i)
		}
		if attributeValue == "x<=9" {
			for i, linha := range train {
				if aux2[i] <= 9 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "9<x<=16" {
			for i, linha := range train {
				if aux2[i] <= 16 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "16<x<=26" {
			for i, linha := range train {
				if aux2[i] <= 26 {
					aux = append(aux, linha)
				}
			}
		} else if attributeValue == "26<x<=200" {
			for i, linha := range train {
				if aux2[i] <= 200 {
					aux = append(aux, linha)
				}
			}
		} else {
			for i, linha := range train {
				if aux2[i] > 200 {
					aux = append(aux, linha)
				}
			}
		}
	} else { // Restantes atributos
		for _, linha := range train {
			if linha[attributeID] == attributeValue {
				aux = append(aux, linha)
			}
		}
	}
	return aux
}

// Adiciona folhas a um nó - chama função do ficheiro relevancia.go
func popularNo(nodeInput *node, train [][]string) {
	// Caso o nó atual seja um nó de estado final, retorna
	if nodeInput.attributeID == 7 || nodeInput.attributeID == 8 {
		return
	}

	// Adiciona o nó atual ao caminho do nó
	nodeInput.path = append(nodeInput.path, nodeInput.attributeID)

	// Função auxiliar que:
	// 1 - escolhe atributo para cada folha do nó atual;
	// 2 - caso não existam atributos relevantes calcula sobrevivencia e utiliza o resultado como atributo;
	// 3 - cria o nó com esse atributo;
	// 4 - adiciona o novo nó à lista de folhas do nó atual;
	// 5 - volta a chamar a função utilizando o nó criado como argumento.
	aux := func(trains [][][]string) {
		for i, t := range trains {
			attrID := escolherAtributo(t, nodeInput.path, train)
			if attrID == -1 {
				attrID = sobrevivencia(train)
			}
			nodeAux := criarNo(attrID, nodeInput)
			nodeInput.leaves = append(nodeInput.leaves, nodeAux)
			popularNo(nodeInput.leaves[i], t)
		}
	}

	// Faz a separação dos passageiros consoante o atributo e envia cada um dos novos grupos de
	// passageiros para a função auxiliar
	if nodeInput.attributeID == 0 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train1 = selectValue(train, 2, "1")
		train2 = selectValue(train, 2, "2")
		train3 = selectValue(train, 2, "3")
		trains := [][][]string{train1, train2, train3}
		aux(trains)
	} else if nodeInput.attributeID == 1 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train1 = selectValue(train, 3, "female")
		train2 = selectValue(train, 3, "male")
		trains := [][][]string{train1, train2}
		aux(trains)
	} else if nodeInput.attributeID == 2 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train4 := make([][]string, 0)
		train5 := make([][]string, 0)
		train6 := make([][]string, 0)
		train7 := make([][]string, 0)
		train8 := make([][]string, 0)
		train1 = selectValue(train, 4, "x<=10")
		train2 = selectValue(train, 4, "10<x<=20")
		train3 = selectValue(train, 4, "20<x<=30")
		train4 = selectValue(train, 4, "30<x<=40")
		train5 = selectValue(train, 4, "40<x<=50")
		train6 = selectValue(train, 4, "50<x<=60")
		train7 = selectValue(train, 4, "60<x<=70")
		train8 = selectValue(train, 4, "x>70")
		trains := [][][]string{train1, train2, train3, train4, train5, train6, train7, train8}
		aux(trains)
	} else if nodeInput.attributeID == 3 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train4 := make([][]string, 0)
		train5 := make([][]string, 0)
		train6 := make([][]string, 0)
		train7 := make([][]string, 0)
		train1 = selectValue(train, 5, "0")
		train2 = selectValue(train, 5, "1")
		train3 = selectValue(train, 5, "2")
		train4 = selectValue(train, 5, "3")
		train5 = selectValue(train, 5, "4")
		train6 = selectValue(train, 5, "5")
		train7 = selectValue(train, 5, "8")
		trains := [][][]string{train1, train2, train3, train4, train5, train6, train7}
		aux(trains)
	} else if nodeInput.attributeID == 4 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train4 := make([][]string, 0)
		train5 := make([][]string, 0)
		train6 := make([][]string, 0)
		train7 := make([][]string, 0)
		train1 = selectValue(train, 6, "0")
		train2 = selectValue(train, 6, "1")
		train3 = selectValue(train, 6, "2")
		train4 = selectValue(train, 6, "3")
		train5 = selectValue(train, 6, "4")
		train6 = selectValue(train, 6, "5")
		train7 = selectValue(train, 6, "6")
		trains := [][][]string{train1, train2, train3, train4, train5, train6, train7}
		aux(trains)
	} else if nodeInput.attributeID == 5 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train4 := make([][]string, 0)
		train5 := make([][]string, 0)
		train1 = selectValue(train, 7, "x<=9")
		train2 = selectValue(train, 7, "9<x<=16")
		train3 = selectValue(train, 7, "16<x<=26")
		train4 = selectValue(train, 7, "26<x<=200")
		train5 = selectValue(train, 7, "x>200")
		trains := [][][]string{train1, train2, train3, train4, train5}
		aux(trains)
	} else if nodeInput.attributeID == 6 {
		train1 := make([][]string, 0)
		train2 := make([][]string, 0)
		train3 := make([][]string, 0)
		train1 = selectValue(train, 8, "C")
		train2 = selectValue(train, 8, "Q")
		train3 = selectValue(train, 8, "S")
		trains := [][][]string{train1, train2, train3}
		aux(trains)
	}
}

// Cria uma arvore de decisão
func criacaoArvoreDecisao(arvoreDecisao *tree, train [][]string) {
	if arvoreDecisao.root == nil {
		attrID := escolherAtributo(train, nil, nil)
		arvoreDecisao.root = &node{
			attributeID: attrID,
			parent:      nil,
			path:        nil,
			leaves:      make([]*node, 0),
		}
		popularNo(arvoreDecisao.root, train)
	} else {
		fmt.Println("Arvore acabada")
	}
}
