package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	hl "github.com/JanSzymanski/goacademy-todo-cli/helperlib"
)

var clearingBytes = [...]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A}

func printMenu() int {
	os.Stdout.Write(clearingBytes[:])
	fmt.Println("==== Main Menu ====")
	fmt.Println("1. Display all ToDos")
	fmt.Println("2. Display details of ToDos")
	fmt.Println("3. Add ToDo")
	fmt.Println("4. Save to file")
	fmt.Println("5. Load from file")
	fmt.Println("6. Create new vault")
	fmt.Println("7. Exit")
	fmt.Println("8. Clear screen")

	scanner := bufio.NewScanner(os.Stdin)
	var input int
	var err error
	for input == 0 {
		fmt.Printf("\nPlease chode one option: ")
		scanner.Scan()
		input, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid user input, please provide numbers")
		}
	}
	return input

}
func displayTodos(tds *hl.TodoStore) {
	os.Stdout.Write(clearingBytes[:])
	fmt.Println("Listing all ToDos:")
	fmt.Println("===================================")
	todos := tds.GetTodos()

	for _, todo := range todos {
		fmt.Printf("Id: %s | Message: %-20s | Status: %s\n", todo[0], todo[1], todo[2])
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press any key to return to Main Menu")
	scanner.Scan()
}
func addTodo(tds *hl.TodoStore) {
	os.Stdout.Write(clearingBytes[:])
	fmt.Println("Adding ToDo to: ", tds.GetVaultName())
	fmt.Println("===================================")
	fmt.Println("Empty message will cancel operation.")
	fmt.Print("Please enter ToDo message: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		tds.AddTodo(input)
	}
}

func main() {

	fmt.Println("No database detected ")
	fmt.Println("Creating new         ")
	todostore := hl.NewTodoStore("Jan's vault")
	var exit bool

	for !exit {
		switch userInput := printMenu(); userInput {
		case 1:
			fmt.Println("Option: ", userInput)
			displayTodos(todostore)
		case 2:
			fmt.Println("Option: ", userInput)
		case 3:
			fmt.Println("Option: ", userInput)
			addTodo(todostore)
		case 4:
			fmt.Println("Option: ", userInput)
		case 5:
			fmt.Println("Option: ", userInput)
		case 6:
			fmt.Println("Option: ", userInput)
		case 7:
			fmt.Println("Option: ", userInput)
			exit = true
			fmt.Println("Closing ToDo app")
		case 8:
			os.Stdout.Write(clearingBytes[:])
		default:
			fmt.Println("Option not yet implemented")

		}
	}

}
