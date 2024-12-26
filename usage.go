package main

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintUsage() {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green("\nCommands:"))
	fmt.Println(green("  -Add <description>                      Add a new todo item with description"))
	fmt.Println(green("  -Edit <index> <new-description>         Edit an existing todo item by index"))
	fmt.Println(green("  -Delete <index>                         Delete a todo item by index"))
	fmt.Println(green("  -Toggle <index>                         Toggle the completion status of a todo item by index"))
	fmt.Println(green("  -List                                   List all todo items"))
	fmt.Println(green("\nExamples:"))
	fmt.Println(green("  todo-cli -Add \"Buy milk\""))
	fmt.Println(green("  todo-cli -Edit 1 \"Buy eggs\""))
	fmt.Println(green("  todo-cli -Delete 2"))
	fmt.Println(green("  todo-cli -Toggle 1"))
	fmt.Println(green("  todo-cli -List"))
}
