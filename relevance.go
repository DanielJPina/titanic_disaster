package main

import (
	"math"
	"strconv"
)

// Fills attribute map
func fillMap(statistics map[string][]float64, train [][]string, attributeID int) map[string][]float64 {
	attribute := attributeID + 2
	// Calculate survived passengers and number of total passengers of the respective type
	if attribute == 4 { // Age
		aux := make([]float64, 0)

		for _, row := range train {
			i, _ := strconv.ParseFloat(row[4], 64)
			aux = append(aux, i)
		}

		for i := 0; i < len(aux); i++ {
			if aux[i] <= 10 {
				if train[i][1] == "1" {
					statistics["x<=10"][0]++
				}
				statistics["x<=10"][1]++
			} else if aux[i] <= 20 {
				if train[i][1] == "1" {
					statistics["10<x<=20"][0]++
				}
				statistics["10<x<=20"][1]++
			} else if aux[i] <= 30 {
				if train[i][1] == "1" {
					statistics["20<x<=30"][0]++
				}
				statistics["20<x<=30"][1]++
			} else if aux[i] <= 40 {
				if train[i][1] == "1" {
					statistics["30<x<=40"][0]++
				}
				statistics["30<x<=40"][1]++
			} else if aux[i] <= 50 {
				if train[i][1] == "1" {
					statistics["40<x<=50"][0]++
				}
				statistics["40<x<=50"][1]++
			} else if aux[i] <= 60 {
				if train[i][1] == "1" {
					statistics["50<x<=60"][0]++
				}
				statistics["50<x<=60"][1]++
			} else if aux[i] <= 70 {
				if train[i][1] == "1" {
					statistics["60<x<=70"][0]++
				}
				statistics["60<x<=70"][1]++
			} else {
				if train[i][1] == "1" {
					statistics["x>70"][0]++
				}
				statistics["x>70"][1]++
			}
		}
	} else if attribute == 7 { // Fare
		aux := make([]float64, 0)

		for _, row := range train {
			i, _ := strconv.ParseFloat(row[7], 64)
			aux = append(aux, i)
		}

		for i := 0; i < len(aux); i++ {
			if aux[i] <= 9 {
				if train[i][1] == "1" {
					statistics["x<=9"][0]++
				}
				statistics["x<=9"][1]++
			} else if aux[i] <= 16 {
				if train[i][1] == "1" {
					statistics["9<x<=16"][0]++
				}
				statistics["9<x<=16"][1]++
			} else if aux[i] <= 26 {
				if train[i][1] == "1" {
					statistics["16<x<=26"][0]++
				}
				statistics["16<x<=26"][1]++
			} else if aux[i] <= 200 {
				if train[i][1] == "1" {
					statistics["26<x<=200"][0]++
				}
				statistics["26<x<=200"][1]++
			} else {
				if train[i][1] == "1" {
					statistics["x>200"][0]++
				}
				statistics["x>200"][1]++
			}
		}
	} else { // Remaining attributes
		for _, row := range train {
			if row[1] == "1" {
				statistics[row[attribute]][0]++
			}
			statistics[row[attribute]][1]++
		}
	}

	// Calculus of survival tax of each type
	for _, row := range statistics {
		row[2] = row[0] / row[1]
	}

	return statistics
}

// Calculates the major and minor survival tax
func calculateTaxes(statistics map[string][]float64) (float64, float64) {
	maxTax := .0
	minTax := 100.0
	for _, row := range statistics {
		if maxTax < row[2] {
			maxTax = row[2]
		}
		if minTax > row[2] {
			minTax = row[2]
		}
	}
	return maxTax, minTax
}

// Calculates the number of passengers relative to the type with tax equal to maximum tax and minimum tax
func calculateQuantities(statistics map[string][]float64, maxTax float64, minTax float64) (float64, float64) {
	maxQt := .0
	minQt := .0
	for _, row := range statistics {
		if row[2] == maxTax {
			maxQt += row[1]
		}
		if row[2] == minTax {
			minQt += row[1]
		}
	}
	return maxQt, minQt
}

// Creates statistics' map
func createMap(attributeID int) (statistics map[string][]float64) {
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

// Major number of cases resolved algorithm
func majorNumberOfCases(statistics map[string][]float64) float64 {
	// Calculates the major tax of survival and the minimum tax of survival
	maxTax, minTax := calculateTaxes(statistics)

	// Calculates the number of passengers of the type equal to the maximum and minimum taxes
	maxQt, minQt := calculateQuantities(statistics, maxTax, minTax)

	return maxQt*maxTax + minQt*(1-minTax)
}

// Gini impurity algorithm
func giniImpurity(survivers float64, nSurvivers float64) float64 {
	if survivers != 0 || nSurvivers != 0 {
		var survivalTax float64 = survivers / (survivers + nSurvivers)
		var nSurvivalTax float64 = nSurvivers / (survivers + nSurvivers)
		return 1 - math.Pow(survivalTax, 2) - math.Pow(nSurvivalTax, 2)
	}
	return 0
}

// Gini impurity gain algorithm
func giniGain(statistics map[string][]float64) float64 {
	gain := .0
	totalPassengers := .0
	impurities := make([]float64, 0)
	weightedAverages := make([]float64, 0)

	for _, row := range statistics {
		totalPassengers += row[1]
	}

	for _, row := range statistics {
		survivers := row[0]
		nSurvivers := row[1] - row[0]
		impurities = append(impurities, giniImpurity(survivers, nSurvivers))
		if totalPassengers != 0 {
			weightedAverages = append(weightedAverages, row[1]/totalPassengers)
		}
	}

	if totalPassengers != 0 {
		for i, impurity := range impurities {
			gain += impurity * weightedAverages[i]
		}
	}

	gain = 1 - gain
	return gain
}

// Entropy(q) function where q = survival probability
func entropy(q float64) float64 {
	if q > 0 && (1-q) > 0 {
		return -(q*math.Log2(q) + (1-q)*math.Log2(1-q))
	}
	return .0
}

// remainder(A) function where A = attribute being analised
func remainder(statistics map[string][]float64, p float64, n float64) float64 {
	result := .0
	for _, row := range statistics {
		pk := row[0]
		nk := row[1] - row[0]
		if (pk + nk) != 0 {
			aux1 := ((pk + nk) / (p + n))
			aux2 := entropy(pk / (pk + nk))
			result += aux1 * aux2
		}
	}
	return result
}

// Entropy information gain algorithm
func entropyGain(statistics map[string][]float64) float64 {
	p := .0
	n := .0
	gain := .0
	for _, row := range statistics {
		p += row[0]
		n += row[1] - row[0]
	}
	if (p + n) != 0 {
		aux1 := entropy(p / (p + n))
		aux2 := remainder(statistics, p, n)
		gain = aux1 - aux2
	}
	return gain
}

// Calculates the received attribute relevance
func attributeRelevance(train [][]string, attributeID int, method int) float64 {
	relevance := .0
	// Survivers, total and survival tax map
	statistics := createMap(attributeID)

	// Fills statistics' map
	fillMap(statistics, train, attributeID)

	// Calculates the relevance of the attribute using the following algorithm:
	if method == 1 { // Major number of cases resolved
		relevance = majorNumberOfCases(statistics)
	} else if method == 2 || method == 3 { // Gini impurity gain
		relevance = giniGain(statistics)
	} else { // Entropy gain
		relevance = entropyGain(statistics)
	}

	return relevance
}

// Verifies if a number is contained on an array
func contain(s []int, element int) bool {
	for _, sElement := range s {
		if sElement == element {
			return true
		}
	}
	return false
}

// Picks the attribute that gives more gain on the node. Doesn't analise the attributes that are received as 0
func pickAttribute(train [][]string, dontAnalise []int, trainParent [][]string) int {
	var relevance [7]float64 // Relevance of each attribute
	best := 0
	aux := .0

	// Select the method of information gain
	// 1: Major number of cases resolved
	// 2: Gini impurity gain without stop condition
	// 3: Gini impurity gain with stop condition
	// 4: Entropy gain
	method := 4

	// Calculate each attribute's relevance
	for i := 0; i < 7; i++ {
		if !contain(dontAnalise, i) {
			relevance[i] = attributeRelevance(train, i, method)
		}
	}

	// If no attribute has relevance returns -1
	var noRelevance [7]float64
	if relevance == noRelevance {
		return -1
	}

	// Picks the attribute with more relevance
	for i, rel := range relevance {
		if aux < rel {
			aux = rel
			best = i
		}
	}

	// If purity of the parent node is equal or higher than the gini gain from the current node returns -1
	if method == 3 && trainParent != nil {
		survivers := .0
		nSurvivers := .0
		for _, row := range trainParent {
			if row[1] == "1" {
				survivers++
			} else {
				nSurvivers++
			}
		}
		parentPurity := 1 - giniImpurity(survivers, nSurvivers)
		if parentPurity >= aux {
			return -1
		}
	}

	return best
}
