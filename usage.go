package main

import "fmt"

func PrintUsage() {
	fmt.Println("\nCommands:")
	fmt.Println("  -Add <description>                      Add a new todo item with description")
	fmt.Println("  -Edit <index> <new-description>         Edit an existing todo item by index")
	fmt.Println("  -Delete <index>                         Delete a todo item by index")
	fmt.Println("  -Toggle <index>                         Toggle the completion status of a todo item by index")
	fmt.Println("  -List                                   List all todo items")
	fmt.Println("\nExamples:")
	fmt.Println("  todo-cli -Add \"Buy milk\"")
	fmt.Println("  todo-cli -Edit 1 \"Buy eggs\"")
	fmt.Println("  todo-cli -Delete 2")
	fmt.Println("  todo-cli -Toggle 1")
	fmt.Println("  todo-cli -List")
}
