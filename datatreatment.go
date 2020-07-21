package main

import (
	"math"
	"sort"
	"strconv"
)

// Calculate the average of a slice
func average(slice []float64) float64 {
	total := .0
	size := float64(len(slice))

	for _, row := range slice {
		total += row
	}
	avg := total / size
	return avg
}

// Calculate the median of a slice
func median(slice []float64) float64 {
	size := len(slice)
	half := size / 2
	intHalf := int(half)
	med := .0

	if size%2 == 0 {
		med = (slice[intHalf-1] + slice[intHalf]) / 2
	} else {
		med = slice[intHalf-1]
	}

	return med
}

// Calculate the standard deviation of a slice
func standardDeviation(slice []float64, average float64) float64 {
	stdDev := .0
	size := float64(len(slice))

	for _, row := range slice {
		stdDev += math.Pow(row-average, 2)
	}
	stdDev = math.Sqrt(stdDev / size)

	return stdDev
}

// Adds the survived column to test
func addSurvivedColumn(testInitial [][]string) [][]string {
	aux := make([][]string, 0, len(testInitial))

	for _, row := range testInitial {
		rowAux := make([]string, 0)
		rowAux = append(rowAux, row[0])
		rowAux = append(rowAux, "-1")
		rowAux = append(rowAux, row[1:]...)
		aux = append(aux, rowAux)
	}
	return aux
}

// Fills the age list crossing the Sex and Pclass attributes and orders it
func fillsAge(sex string, pclass string, ageList []float64, train [][]string) []float64 {
	aux := make([]float64, 0)

	for i, row := range train {
		if row[4] == sex && row[2] == pclass && row[5] != "" {
			aux = append(aux, ageList[i])
		}
	}
	// Orders list
	sort.Float64s(aux)

	return aux
}

// Changes the missing data to age attribute crossing the Sex and Pclass attributes
func changeMissingAge(train [][]string) {
	aux := make([]float64, 0)

	// Map of median of ages crossing the Sex and Pclass attributes
	ageMedian := map[string]float64{
		"female1": .0,
		"female2": .0,
		"female3": .0,
		"male1":   .0,
		"male2":   .0,
		"male3":   .0,
	}

	// Fills the aux list with all the ages
	for _, row := range train {
		i, _ := strconv.ParseFloat(row[5], 64)
		aux = append(aux, i)
	}

	// Fills the auxiliary lists crossing the Sex and Pclass attributes
	auxFemale1 := fillsAge("female", "1", aux, train)
	auxFemale2 := fillsAge("female", "2", aux, train)
	auxFemale3 := fillsAge("female", "3", aux, train)
	auxMale1 := fillsAge("male", "1", aux, train)
	auxMale2 := fillsAge("male", "2", aux, train)
	auxMale3 := fillsAge("male", "3", aux, train)

	// Median calculation
	ageMedian["female1"] = median(auxFemale1)
	ageMedian["female2"] = median(auxFemale2)
	ageMedian["female3"] = median(auxFemale3)
	ageMedian["male1"] = median(auxMale1)
	ageMedian["male2"] = median(auxMale2)
	ageMedian["male3"] = median(auxMale3)

	// Converts missing data with respective value
	for _, row := range train {
		if row[5] == "" {
			if row[4] == "female" {
				if row[2] == "1" {
					row[5] = strconv.FormatFloat(ageMedian["female1"], 'E', -1, 64)
				} else if row[2] == "2" {
					row[5] = strconv.FormatFloat(ageMedian["female2"], 'E', -1, 64)
				} else {
					row[5] = strconv.FormatFloat(ageMedian["female3"], 'E', -1, 64)
				}
			} else {
				if row[2] == "1" {
					row[5] = strconv.FormatFloat(ageMedian["male1"], 'E', -1, 64)
				} else if row[2] == "2" {
					row[5] = strconv.FormatFloat(ageMedian["male2"], 'E', -1, 64)
				} else {
					row[5] = strconv.FormatFloat(ageMedian["male3"], 'E', -1, 64)
				}
			}
		}
	}
}

// Changes the missing data of indicated values for the asymmetry coefficient (option 1 = age, 2 = fare)
func changeMissingAsymmetryCoeficient(train [][]string, option int) {
	aux := make([]float64, 0)

	if option == 1 {
		for _, row := range train {
			if row[5] != "" {
				i, _ := strconv.ParseFloat(row[5], 64)
				aux = append(aux, i)
			}
		}
	} else {
		for _, row := range train {
			if row[9] != "" {
				i, _ := strconv.ParseFloat(row[9], 64)
				aux = append(aux, i)
			}
		}
	}

	// Order list
	sort.Float64s(aux)
	// Auxiliary calculus
	average := average(aux)
	median := median(aux)
	standardDeviation := standardDeviation(aux, average)

	// Calculate asymmetry coefficient using Pearson's 2nd Skewness algorithm
	asymmetryCoefficient := (3 * (average - median)) / standardDeviation

	// Defines the value that's going to be used to fill the missing data
	var fillIn string
	if asymmetryCoefficient == 0 { // The data is symmetrical
		fillIn = strconv.FormatFloat(average, 'E', -1, 64) // Missing values are equal to average
	} else { // The data is asymmetrical
		fillIn = strconv.FormatFloat(median, 'E', -1, 64) // Missing values are equal to median
	}

	// Convert the missing values for the respective value
	if option == 1 {
		for _, row := range train {
			if row[5] == "" {
				row[5] = fillIn
			}
		}
	} else {
		for _, row := range train {
			if row[9] == "" {
				row[9] = fillIn
			}
		}
	}
}

// Convert Embarked missing data for S since it is the majority of the cases
func changeMissingEmbarked(train [][]string) {
	for _, row := range train {
		if row[11] == "" {
			row[11] = "S"
		}
	}
}

// Convert missing values to pretended values
func missingDataChange(train [][]string) {
	//changeMissingAsymetryCoeficient(train, 1)
	changeMissingAge(train)
	changeMissingEmbarked(train)
	changeMissingAsymmetryCoeficient(train, 2)
}

// Delete irrelevant attributes for the decision tree creation
func deleteAttribute(trainInitial [][]string) [][]string {
	aux := make([][]string, 0, len(trainInitial))
	for _, row := range trainInitial {
		newRow := make([]string, 0, 9)
		for i := 0; i < 12; i++ {
			if !(i == 3 || i == 8 || i == 10) {
				newRow = append(newRow, row[i])
			}
		}
		aux = append(aux, newRow)
	}
	return aux
}
