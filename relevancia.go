package main

import (
	"math"
	"strconv"
)

// Preenche mapa de atributos
func preencherMapa(statistics map[string][]float64, train [][]string, attributeID int) map[string][]float64 {
	atributo := attributeID + 2
	// Calcula sobreviventes e numero total de passageiros do respetivo tipo
	if atributo == 4 {
		aux := make([]float64, 0)

		for _, linha := range train {
			i, _ := strconv.ParseFloat(linha[4], 64)
			aux = append(aux, i)
		}

		for i := 0; i < len(aux); i++ {
			if aux[i] <= 10 { // Se o passageiro tem idade compreendida entre os 0 e os 10 anos
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["x<=10"][0]++
				}
				statistics["x<=10"][1]++
			} else if aux[i] <= 20 { // Se o passageiro tem idade superior a 10 e menor ou igual a 20
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["10<x<=20"][0]++
				}
				statistics["10<x<=20"][1]++
			} else if aux[i] <= 30 { // Se o passageiro tem idade superior a 20 e menor ou igual a 30
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["20<x<=30"][0]++
				}
				statistics["20<x<=30"][1]++
			} else if aux[i] <= 40 { // Se o passageiro tem idade superior a 30 e menor ou igual a 40
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["30<x<=40"][0]++
				}
				statistics["30<x<=40"][1]++
			} else if aux[i] <= 50 { // Se o passageiro tem idade superior a 40 e menor ou igual a 50
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["40<x<=50"][0]++
				}
				statistics["40<x<=50"][1]++
			} else if aux[i] <= 60 { // Se o passageiro tem idade superior a 50 e menor ou igual a 60
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["50<x<=60"][0]++
				}
				statistics["50<x<=60"][1]++
			} else if aux[i] <= 70 { // Se o passageiro tem idade superior a 60 e menor ou igual a 70
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["60<x<=70"][0]++
				}
				statistics["60<x<=70"][1]++
			} else { // Se o passageiro tem idade superior a 70
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["x>70"][0]++
				}
				statistics["x>70"][1]++
			}
		}
	} else if atributo == 7 {
		aux := make([]float64, 0)

		for _, linha := range train {
			i, _ := strconv.ParseFloat(linha[7], 64)
			aux = append(aux, i)
		}

		for i := 0; i < len(aux); i++ {
			if aux[i] <= 9 { // Se o passageiro pagou uma tarifa compreendida entre os 0 e os 9
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["x<=9"][0]++
				}
				statistics["x<=9"][1]++
			} else if aux[i] <= 16 { // Se o passageiro pagou uma tarifa superior a 9 menor ou igual a 16
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["9<x<=16"][0]++
				}
				statistics["9<x<=16"][1]++
			} else if aux[i] <= 26 { // Se o passageiro pagou uma tarifa superior a 16 menor ou igual a 26
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["16<x<=26"][0]++
				}
				statistics["16<x<=26"][1]++
			} else if aux[i] <= 200 { // Se o passageiro pagou uma tarifa superior a 26 menor ou igual a 200
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["26<x<=200"][0]++
				}
				statistics["26<x<=200"][1]++
			} else { // Se o passageiro pagou uma tarifa superior a 200
				if train[i][1] == "1" { // Se esse passageiro tiver sobrevivido
					statistics["x>200"][0]++
				}
				statistics["x>200"][1]++
			}
		}
	} else {
		for _, linha := range train {
			if linha[1] == "1" { // Se esse passageiro tiver sobrevivido
				statistics[linha[atributo]][0]++
			}
			statistics[linha[atributo]][1]++
		}
	}

	// Calculo da taxa de sobrevivencia respetivo a cada tipo
	for _, linha := range statistics {
		linha[2] = linha[0] / linha[1]
	}

	return statistics
}

// Calcula a maior e a menor taxa de sobrevivencia
func calcularTaxas(statistics map[string][]float64) (float64, float64) {
	maxTax := .0
	minTax := 100.0
	for _, linha := range statistics {
		if maxTax < linha[2] {
			maxTax = linha[2]
		}
		if minTax > linha[2] {
			minTax = linha[2]
		}
	}
	return maxTax, minTax
}

// Calcula o numero de passageiros relativos ao tipo com taxa igual à taxa maxima e taxa minima
func calcularQuantidades(statistics map[string][]float64, maxTax float64, minTax float64) (float64, float64) {
	maxQt := .0
	minQt := .0
	for _, linha := range statistics {
		if linha[2] == maxTax {
			maxQt += linha[1]
		}
		if linha[2] == minTax {
			minQt += linha[1]
		}
	}
	return maxQt, minQt
}

// Cria mapa de estatisticas
func criarMapa(attributeID int) (statistics map[string][]float64) {
	switch attributeID {
	case 0:
		statistics = map[string][]float64{
			"1": {.0, .0, .0},
			"2": {.0, .0, .0},
			"3": {.0, .0, .0},
		}
	case 1:
		statistics = map[string][]float64{
			"female": {.0, .0, .0},
			"male":   {.0, .0, .0},
		}
	case 2:
		statistics = map[string][]float64{
			"x<=10":    {.0, .0, .0},
			"10<x<=20": {.0, .0, .0},
			"20<x<=30": {.0, .0, .0},
			"30<x<=40": {.0, .0, .0},
			"40<x<=50": {.0, .0, .0},
			"50<x<=60": {.0, .0, .0},
			"60<x<=70": {.0, .0, .0},
			"x>70":     {.0, .0, .0},
		}
	case 3:
		statistics = map[string][]float64{
			"0": {.0, .0, .0},
			"1": {.0, .0, .0},
			"2": {.0, .0, .0},
			"3": {.0, .0, .0},
			"4": {.0, .0, .0},
			"5": {.0, .0, .0},
			"8": {.0, .0, .0},
		}
	case 4:
		statistics = map[string][]float64{
			"0": {.0, .0, .0},
			"1": {.0, .0, .0},
			"2": {.0, .0, .0},
			"3": {.0, .0, .0},
			"4": {.0, .0, .0},
			"5": {.0, .0, .0},
			"6": {.0, .0, .0},
		}
	case 5:
		statistics = map[string][]float64{
			"x<=9":      {.0, .0, .0},
			"9<x<=16":   {.0, .0, .0},
			"16<x<=26":  {.0, .0, .0},
			"26<x<=200": {.0, .0, .0},
			"x>200":     {.0, .0, .0},
		}
	default:
		statistics = map[string][]float64{
			"C": {.0, .0, .0},
			"Q": {.0, .0, .0},
			"S": {.0, .0, .0},
		}
	}
	return
}

// Algoritmo resolução de maior numero de casos
func maiorNumeroCasos(statistics map[string][]float64) float64 {
	// Calcula a maior taxa de sobrevivencia e a menor taxa de sobrevivencia
	maxTax, minTax := calcularTaxas(statistics)

	// Calcula o numero de passageiros do tipo igual ao máximo e ao minimo
	maxQt, minQt := calcularQuantidades(statistics, maxTax, minTax)

	return maxQt*maxTax + minQt*(1-minTax)
}

// Algoritmo gini impurity
func giniImpurity(sobreviventes float64, nSobreviventes float64) float64 {
	if sobreviventes != 0 || nSobreviventes != 0 {
		var txSobrevivencia float64 = sobreviventes / (sobreviventes + nSobreviventes)
		var txNSobrevivencia float64 = nSobreviventes / (sobreviventes + nSobreviventes)
		return 1 - math.Pow(txSobrevivencia, 2) - math.Pow(txNSobrevivencia, 2)
	}
	return 0
}

// Algoritmo de ganho de informação utilizando gini impurity
func giniGanho(statistics map[string][]float64) float64 {
	ganho := .0
	totalPassageiros := .0
	impuridades := make([]float64, 0)
	mediasPonderadas := make([]float64, 0)

	for _, linha := range statistics {
		totalPassageiros += linha[1]
	}

	for _, linha := range statistics {
		sobreviventes := linha[0]
		nSobreviventes := linha[1] - linha[0]
		impuridades = append(impuridades, giniImpurity(sobreviventes, nSobreviventes))
		if totalPassageiros != 0 {
			mediasPonderadas = append(mediasPonderadas, linha[1]/totalPassageiros)
		}
	}

	if totalPassageiros != 0 {
		for i, impuridade := range impuridades {
			ganho += impuridade * mediasPonderadas[i]
		}
	}

	ganho = 1 - ganho
	return ganho
}

// Função entropia(q) onde q = probabilidade de sobrevivencia
func entropia(q float64) float64 {
	if q > 0 && (1-q) > 0 {
		return -(q*math.Log2(q) + (1-q)*math.Log2(1-q))
	}
	return .0
}

// Função remainder(A) onde A é o atributo a ser analisado
func remainder(statistics map[string][]float64, p float64, n float64) float64 {
	resultado := .0
	for _, linha := range statistics {
		pk := linha[0]
		nk := linha[1] - linha[0]
		if (pk + nk) != 0 {
			aux1 := ((pk + nk) / (p + n))
			aux2 := entropia(pk / (pk + nk))
			resultado += aux1 * aux2
		}
	}
	return resultado
}

// Algoritmo de ganho de informação utilizando entropia
func entropiaGanho(statistics map[string][]float64) float64 {
	p := .0
	n := .0
	ganho := .0
	for _, linha := range statistics {
		p += linha[0]
		n += linha[1] - linha[0]
	}
	if (p + n) != 0 {
		aux1 := entropia(p / (p + n))
		aux2 := remainder(statistics, p, n)
		ganho = aux1 - aux2
	}
	return ganho
}

// Calcula relevancia para o atributo recebido
func relevanciaAtributo(train [][]string, attributeID int, metodo int) float64 {
	relevancia := .0
	// Mapa dos sobreviventes, total, taxa sobrevivencia de cada local de embarque
	statistics := criarMapa(attributeID)

	// Preenche mapa de estatisticas
	preencherMapa(statistics, train, attributeID)

	// Calcula a relevancia do atributo utilizando:
	if metodo == 1 { // O algoritmo de resolução de maior numero de casos
		relevancia = maiorNumeroCasos(statistics)
	} else if metodo == 2 || metodo == 3 { // O algoritmo gini impurity
		relevancia = giniGanho(statistics)
	} else { // O algoritmo da entropia
		relevancia = entropiaGanho(statistics)
	}

	return relevancia
}

// Verifica se um numero esta contido num array
func contem(s []int, element int) bool {
	for _, sElement := range s {
		if sElement == element {
			return true
		}
	}
	return false
}

// Escolhe o atributo que resolve mais casos para o nó, não analisa os atributos que recebe a 0
func escolherAtributo(train [][]string, dontAnalise []int, trainPai [][]string) int {
	var relevancia [7]float64 // Relevancia de cada atributo
	best := 0
	aux := .0

	// Escolher o método de ganho de informação
	// 1: maior numero de casos
	// 2: gini impurity sem condição de paragem
	// 3: gini impurity com condição de paragem
	// 4: entropia
	metodo := 4

	// Calcular relevancia de cada atributo
	for i := 0; i < 7; i++ {
		if !contem(dontAnalise, i) {
			relevancia[i] = relevanciaAtributo(train, i, metodo)
		}
	}

	// Caso nenhum atributo tenha relevancia retorna -1
	var nenhumaRelevancia [7]float64
	if relevancia == nenhumaRelevancia {
		return -1
	}

	// Seleccionar o atributo com maior relevancia
	for i, rel := range relevancia {
		if aux < rel {
			aux = rel
			best = i
		}
	}

	// Caso o ganho do nó pai utilizando gini impurity seja igual ou superior ao nó atual retorna -1
	if metodo == 3 && trainPai != nil {
		sobreviventes := .0
		nSobreviventes := .0
		for _, linha := range trainPai {
			if linha[1] == "1" {
				sobreviventes++
			} else {
				nSobreviventes++
			}
		}
		ganhoPai := 1 - giniImpurity(sobreviventes, nSobreviventes)
		if ganhoPai >= aux {
			return -1
		}
	}

	return best
}
