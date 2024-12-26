package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type CommandFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new todo")

	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo")

	flag.IntVar(&cf.Delete, "Delete", -1, "Delete a todo by index")

	flag.IntVar(&cf.Toggle, "Toggle", -1, "Toggle the completed status of a todo by index")

	flag.BoolVar(&cf.List, "List", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, " ", 2)
		if len(parts) != 2 {
			color.Red("Invalid edit format")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			color.Red("Invalid index")
		}
		todos.edit(index, parts[1])
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	default:
	}
}
