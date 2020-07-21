package main

import (
	"fmt"
	"strconv"
)

// Node structure definition
type node struct {
	attributeID int     // Identifies node: 0=Pclass, 1=Sex, 2=Age, 3=SibSp, 4=Parch, 5=Fare, 6=Embarked, 7=Survided, 8=!Survived
	parent      *node   // Node parent
	path        []int   // Path from root to node
	leaves      []*node // Node's leaves
}

// Tree structure defenition
type tree struct {
	root *node
}

// Prints node
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

// Survival state nodes definition
func survival(train [][]string) int {
	attrID := 0
	surv := 0
	nSurv := 0
	for _, row := range train {
		if row[1] == "1" {
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

// Creates a node
func createNode(attrID int, nodeInput *node) *node {
	node := &node{
		attributeID: attrID,
		parent:      nodeInput,
		path:        nodeInput.path,
		leaves:      make([]*node, 0),
	}
	return node
}

// Selects the rows of the CSV file where the attribute's value is equal to the specified value
func selectValue(train [][]string, attributeID int, attributeValue string) [][]string {
	aux := make([][]string, 0)
	if attributeID == 4 { // If attribute is Age
		aux2 := make([]float64, 0)

		for _, row := range train {
			i, _ := strconv.ParseFloat(row[4], 64)
			aux2 = append(aux2, i)
		}
		if attributeValue == "x<=10" {
			for i, row := range train {
				if aux2[i] <= 10 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "10<x<=20" {
			for i, row := range train {
				if aux2[i] <= 20 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "20<x<=30" {
			for i, row := range train {
				if aux2[i] <= 30 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "30<x<=40" {
			for i, row := range train {
				if aux2[i] <= 40 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "40<x<=50" {
			for i, row := range train {
				if aux2[i] <= 50 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "50<x<=60" {
			for i, row := range train {
				if aux2[i] <= 60 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "60<x<=70" {
			for i, row := range train {
				if aux2[i] <= 70 {
					aux = append(aux, row)
				}
			}
		} else {
			for i, row := range train {
				if aux2[i] > 70 {
					aux = append(aux, row)
				}
			}
		}
	} else if attributeID == 7 { // If attribute is Fare
		aux2 := make([]float64, 0)

		for _, row := range train {
			i, _ := strconv.ParseFloat(row[7], 64)
			aux2 = append(aux2, i)
		}
		if attributeValue == "x<=9" {
			for i, row := range train {
				if aux2[i] <= 9 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "9<x<=16" {
			for i, row := range train {
				if aux2[i] <= 16 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "16<x<=26" {
			for i, row := range train {
				if aux2[i] <= 26 {
					aux = append(aux, row)
				}
			}
		} else if attributeValue == "26<x<=200" {
			for i, row := range train {
				if aux2[i] <= 200 {
					aux = append(aux, row)
				}
			}
		} else {
			for i, row := range train {
				if aux2[i] > 200 {
					aux = append(aux, row)
				}
			}
		}
	} else { // Remaining attributes
		for _, row := range train {
			if row[attributeID] == attributeValue {
				aux = append(aux, row)
			}
		}
	}
	return aux
}

// Adds leaves to a node - calls function from relevance.go file
func populateNode(nodeInput *node, train [][]string) {
	// If actual node is a final state node, returns
	if nodeInput.attributeID == 7 || nodeInput.attributeID == 8 {
		return
	}

	// Adds the actual node to the node's path
	nodeInput.path = append(nodeInput.path, nodeInput.attributeID)

	// Auxiliary function that:
	// 1 - Picks attribute for each leaf of the actual node;
	// 2 - If relevant attributes don't exist calculates survival and uses the result as the attribute;
	// 3 - Creates the node with that attribute;
	// 4 - Adds the new node to the list of leaves of the actual node;
	// 5 - Calls populateNode function using the created node as argument.
	aux := func(trains [][][]string) {
		for i, t := range trains {
			attrID := pickAttribute(t, nodeInput.path, train)
			if attrID == -1 {
				attrID = survival(train)
			}
			nodeAux := createNode(attrID, nodeInput)
			nodeInput.leaves = append(nodeInput.leaves, nodeAux)
			populateNode(nodeInput.leaves[i], t)
		}
	}

	// Separates the passengers based on the attribute and sends
	// each one of the new group of passengers to auxiliary function
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

// Creates a decision tree
func decisionTreeCreation(decisionTree *tree, train [][]string) {
	if decisionTree.root == nil {
		attrID := pickAttribute(train, nil, nil)
		decisionTree.root = &node{
			attributeID: attrID,
			parent:      nil,
			path:        nil,
			leaves:      make([]*node, 0),
		}
		populateNode(decisionTree.root, train)
	} else {
		fmt.Println("Finished tree")
	}
}
