package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	hl "github.com/JanSzymanski/goacademy-todo-cli/todostorelib"
)

var clearingBytes = [...]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A}

func clearScreen() {
	os.Stdout.Write(clearingBytes[:])
}

func displayVaultInfo(tds *hl.TodoStore) {
	vaultInfo := tds.GetVaultInfo()
	fmt.Println("Vault name: ", vaultInfo["name"])
	fmt.Printf("ToDo count: %s\n\n", vaultInfo["counter"])
}

func printMenu(tds *hl.TodoStore) int {
	clearScreen()
	displayVaultInfo(tds)
	fmt.Println("==== Main Menu ====")
	fmt.Println("1. Display ToDos")
	fmt.Println("2. Display details of a ToDo")
	fmt.Println("3. Add ToDo")
	fmt.Println("4. Exit")

	scanner := bufio.NewScanner(os.Stdin)
	var input int
	var err error
	for input == 0 {
		fmt.Printf("\nPlease chose one option: ")
		scanner.Scan()
		input, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid user input, please provide numbers")
		}
	}
	return input

}
func displayTodos(tds *hl.TodoStore) {
	clearScreen()
	var exit bool
	for !exit {
		fmt.Println("Listing all ToDos:")
		fmt.Println("===================================")
		todos := tds.GetTodos(0, 20)

		for _, todo := range todos {
			fmt.Printf("Id: %s | Message: %-20s | Status: %s\n", todo["id"], todo["message"], todo["status"])
		}
		exit = displayTodo(tds)

	}
}
func displayTodo(tds *hl.TodoStore) bool {
	// clearScreen()
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	var err error
	for id == 0 {
		fmt.Println("\nPlease chose ToDo id or type 'mm' to return to main menu ")
		scanner.Scan()
		if scanner.Text() == "mm" {
			return true
		}
		id, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid user input, please provide numbers")
			scanner.Scan()
			clearScreen()
			return false
		}
	}
todoloop:
	for {
		todo, err := tds.GetTodo(id)
		if err != nil {
			fmt.Println("No todo with given id: ", id)
			scanner.Scan()
			clearScreen()
			return false
		}
		clearScreen()
		fmt.Println("Id: ", todo["id"])
		fmt.Println("Message: ", todo["message"])
		fmt.Println("Status: ", todo["status"])
		var close_sub_menu bool
		for !close_sub_menu {
			switch sub_option := displayTodoSubmenu(); sub_option {
			case "mm":
				return true
				// break todoloop
			case "gb":
				clearScreen()
				close_sub_menu = true
				return false
			case "cm":
				displayChangeMessageMenu(tds, id)
				close_sub_menu = true
			case "csa":
				tds.ChangeTodoStatus(id, hl.Active)
				fmt.Println("Status changed to 'active'")
				close_sub_menu = true
			case "csi":
				tds.ChangeTodoStatus(id, hl.Inactive)
				fmt.Println("Status changed to 'inactive'")
				close_sub_menu = true
			case "csd":
				tds.ChangeTodoStatus(id, hl.Done)
				fmt.Println("Status changed to 'done'")
				close_sub_menu = true
			case "del":
				deleteTodoConfirmation(tds, id)
				break todoloop
			default:
				fmt.Println("Option not yet implemented")
			}
		}
	}
	return true
}

func deleteTodoConfirmation(tds *hl.TodoStore, id int) {
	fmt.Println("Confirm deleting by typing 'delete'")
	fmt.Println("Any other input will cancel")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "delete" {
		tds.DeleteTodo(id)
	}
}

func displayChangeMessageMenu(tds *hl.TodoStore, id int) {
	fmt.Println("Empty message cancels")
	fmt.Println("\nPlease type in a new message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		tds.ChangeTodoMessagge(id, input)
	}

}
func displayTodoSubmenu() string {
	fmt.Println("\nSelect option: ")
	fmt.Println("gb - go back")
	fmt.Println("mm - return to main menu")
	fmt.Println("cm - change message")
	fmt.Println("csa - change status to active")
	fmt.Println("csi - change status to inactive")
	fmt.Println("csd - change status to done")
	fmt.Println("del - delete todo")

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for input == "" {
		fmt.Printf("\nPlease chose one option: ")
		scanner.Scan()
		input = scanner.Text()
	}
	return input
}
func addTodo(tds *hl.TodoStore) {
	clearScreen()
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
	scanner := bufio.NewScanner(os.Stdin)
	var exit bool

	for !exit {
		switch userInput := printMenu(todostore); userInput {
		case 1:
			displayTodos(todostore)
		case 2:
			displayTodo(todostore)
		case 3:
			addTodo(todostore)
		case 4:
			exit = true
			fmt.Println("Closing ToDo app")
		default:
			fmt.Println("Option not yet implemented")
			scanner.Scan()

		}
	}

}
