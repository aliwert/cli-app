package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/fatih/color"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

// Add a new Todo to the list
func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

// Validate the index in the Todos slice
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		color.Red(fmt.Sprintf("Error: %v", err))
		return err
	}
	return nil
}

// Delete a Todo from the Todos slice by index
func (todos *Todos) Delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		color.Red(fmt.Sprintf("Error: %v", err))
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

// Toggle the completed status of a Todo by index
func (todos *Todos) Toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		color.Red(fmt.Sprintf("Error: %v", err))
		return err
	}
	todo := &(*todos)[index] // Get a pointer to the specific Todo item

	// Toggle the completion status
	todo.Completed = !todo.Completed
	if todo.Completed {
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	} else {
		todo.CompletedAt = nil
	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		color.Red(fmt.Sprintf("Error: %v", err))
		return err

	}
	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := "✖️"
		if todo.Completed {
			completed = "✔️"
		}
		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC3339)
		}
		table.AddRow(
			color.RedString(strconv.Itoa(index)),
			color.RedString(todo.Title),
			color.RedString(completed),
			color.RedString(todo.CreatedAt.Format(time.RFC3339)),
			color.RedString(completedAt),
		)
	}
	table.Render()
}
