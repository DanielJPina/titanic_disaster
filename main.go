package main

func main() {
	// Load files - calls functions from inputoutput.go file
	trainCSV := loadFiles("train.csv")
	testCSV := loadFiles("test.csv")

	// CSV files data without header
	trainInitial := trainCSV[1:]
	testInitial := testCSV[1:]

	// Add survived column testInitial - calls function from datatreatment.go file
	completedTestInitial := addSurvivedColumn(testInitial)

	// Treatment of files missing data - calls function from datatreatment.go file
	missingDataChange(trainInitial)
	missingDataChange(completedTestInitial)

	// Non relevant data deletion from files - calls function from datatreatment.go file
	train := deleteAttribute(trainInitial)
	test := deleteAttribute(completedTestInitial)

	// Decision tree inicialization
	decisionTree := &tree{
		root: nil,
	}

	// Decision tree creation algorithm - calls function from decisiontree.go file
	decisionTreeCreation(decisionTree, train)

	// Draw a tree - calls function from drawtree.go file
	drawing := drawTree(decisionTree.root)

	// Print a drawing to a txt file - calls function from inputoutput.go file
	exportTxt(drawing)

	// Analise passengers from train - calls function from analise.go file
	analiseTrain(train, decisionTree)

	// Analise passengers from test - calls function from analise.go file
	output := analisePassengers(test, decisionTree)

	// Convert output to submission.csv file - calls function from inputoutput.go file
	exportCSV(output)
}
