package main

import (
	"strings"
)

const (
	novaLinha       = "\n"
	espacoVasio     = "    "
	itemMeio        = "├── "
	itemContinuacao = "│   "
	itemFinal       = "└── "
)

var ids = map[int]string{
	0: "Pclass",
	1: "Sex",
	2: "Age",
	3: "SibSp",
	4: "Parch",
	5: "Fare",
	6: "Embarked",
	7: "Survived",
	8: "Died",
}

// Desenha uma arvore
func desenharArvore(nodeInput *node) string {
	return ids[nodeInput.attributeID] + novaLinha + imprimeItems(nodeInput.leaves, []bool{})
}

// Desenha uma linha do desenho final
func imprimeTexto(texto string, espacos []bool, ultimo bool) string {
	var resultado string
	for _, espaco := range espacos {
		if espaco {
			resultado += espacoVasio
		} else {
			resultado += itemContinuacao
		}
	}

	indicador := itemMeio
	if ultimo {
		indicador = itemFinal
	}

	var output string
	lines := strings.Split(texto, "\n")
	for i := range lines {
		texto := lines[i]
		if i == 0 {
			output += resultado + indicador + texto + novaLinha
			continue
		}
		if ultimo {
			indicador = espacoVasio
		} else {
			indicador = itemContinuacao
		}
		output += resultado + indicador + texto + novaLinha
	}
	return output
}

// Faz os espaçamentos entre linhas do desenho final
func imprimeItems(leaves []*node, espacos []bool) string {
	var resultado string
	for i, leaf := range leaves {
		ultimo := i == len(leaves)-1
		resultado += imprimeTexto(ids[leaf.attributeID], espacos, ultimo)
		if len(leaf.leaves) > 0 {
			espacosChild := append(espacos, ultimo)
			resultado += imprimeItems(leaf.leaves, espacosChild)
		}
	}
	return resultado
}
