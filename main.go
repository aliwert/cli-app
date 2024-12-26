package main

import (
	"os"
)

func main() {

	PrintLogo()

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	// create a new Todos slice
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	CommandFlags := NewCommandFlags()
	CommandFlags.Execute(&todos)
	todos.Toggle(0)
	todos.print()
	storage.Save(todos)
}
