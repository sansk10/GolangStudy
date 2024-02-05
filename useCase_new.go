package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var a = map[string][]string{
	"7750": {"SR-1e", "SR-14S", "SR-12", "SR-7"},
	"7250": {"IXR-6", "IXR-R6", "IXR-R6dl"},
	"7705": {"SAR-8", "SAR-M", "SAR-A", "SAR-W"},
}

func addNewFamily() {
	var numbers []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a family name: ")
	family, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}

	fmt.Printf("Enter a list of nodes to be added in the entered family %s (separated by spaces): ", strings.TrimSpace(family))
	numbersStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading numbers:", err)
		return
	}

	numbersStr = strings.TrimSpace(numbersStr)
	numStrings := strings.Fields(numbersStr)

	for _, numStr := range numStrings {
		numbers = append(numbers, numStr)
	}
	family = strings.TrimSpace(family)

	if existingNumbers, ok := a[family]; ok {
		a[family] = append(existingNumbers, numbers...)
	} else {
		a[family] = numbers
	}
	fmt.Println("Family and nodes after adding:")
	displayFamily()
}

func checkTheNode() {
	var nodeToCheck string
	var familyName string

	fmt.Print("Enter a familyName to check: ")
	_, err := fmt.Scanf("%s\n", &familyName)
	fmt.Print("Enter a node that you want to check in this family %s: ", familyName)
	_, err = fmt.Scanf("%s\n", &nodeToCheck)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if values, ok := a[familyName]; ok {
		for _, v := range values {
			if v == nodeToCheck {
				fmt.Printf("Node %s exists in the array for family %s\n", nodeToCheck, familyName)
				return
			}
		}
		fmt.Printf("Node %s does not exist in the array for family %s\n", nodeToCheck, familyName)
	} else {
		fmt.Printf("Family %s does not exist in the map\n", familyName)
	}

	fmt.Println("Family -> Nodes")
	displayFamily()
}

func nodeToDelete() {
	var familyName string
	var nodeToDelete string

	fmt.Print("Enter a familyName to delete a node from: ")
	_, err := fmt.Scanf("%s\n", &familyName)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter a node to delete from family %s: ", familyName)
	_, err = fmt.Scanf("%s\n", &nodeToDelete)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if values, ok := a[familyName]; ok {
		// Find the index of the value to delete
		indexToDelete := -1
		for i, v := range values {
			if v == nodeToDelete {
				indexToDelete = i
				break
			}
		}

		// If the value is found, create a new slice without that value
		if indexToDelete != -1 {
			a[familyName] = append(values[:indexToDelete], values[indexToDelete+1:]...)
			fmt.Printf("Deleted value %s from key %s\n", nodeToDelete, familyName)
		} else {
			fmt.Printf("Value %s not found in the array for key %s\n", nodeToDelete, familyName)
		}
	} else {
		fmt.Printf("Key %s does not exist in the map\n", familyName)
	}

	fmt.Println("Family and nodes after deleting:")
	displayFamily()
}

func displayFamily() {
	fmt.Println("Family -> Nodes")
	for key, values := range a {
		fmt.Printf(" %s -> %v\n", key, values)
	}
}

func main() {
	x := true
	for x {
		var option int
		fmt.Println("Start")
		fmt.Println("1. Add node family with Node type")
		fmt.Println("2. Check the node is available in the required family or not")
		fmt.Println("3. Delete the node from a particular family")
		fmt.Println("4. Display the Map")
		fmt.Println("5. Exit")
		fmt.Print("Enter your option: ")
		fmt.Scanln(&option)
		
		switch option {
		case 1:
			addNewFamily()
		case 2:
			checkTheNode()
		case 3:
			nodeToDelete()
		case 4:
			displayFamily()
		case 5:
			fmt.Println("End")
			x = false
			
		default:
			continue
		}

	}
	
}
