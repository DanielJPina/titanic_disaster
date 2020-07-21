package main

import (
	"strings"
)

const (
	newLine          = "\n"
	emptySpace       = "    "
	middleItem       = "├── "
	continuationItem = "│   "
	finalItem        = "└── "
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

// Draws a Tree
func drawTree(nodeInput *node) string {
	return ids[nodeInput.attributeID] + newLine + printItems(nodeInput.leaves, []bool{})
}

// Draws a row from final drawing
func printText(text string, spaces []bool, last bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continuationItem
		}
	}

	indicator := middleItem
	if last {
		indicator = finalItem
	}

	var output string
	lines := strings.Split(text, "\n")
	for i := range lines {
		text := lines[i]
		if i == 0 {
			output += result + indicator + text + newLine
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continuationItem
		}
		output += result + indicator + text + newLine
	}
	return output
}

// Makes final drawing's spaces between rows
func printItems(leaves []*node, spaces []bool) string {
	var result string
	for i, leaf := range leaves {
		last := i == len(leaves)-1
		result += printText(ids[leaf.attributeID], spaces, last)
		if len(leaf.leaves) > 0 {
			spacesChild := append(spaces, last)
			result += printItems(leaf.leaves, spacesChild)
		}
	}
	return result
}
