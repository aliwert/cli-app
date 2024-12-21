# Todo CLI Application

- This is a simple CLI application for managing a todo list from the command line.

## Installation

- Clone this repository and install the necessary dependencies:

```sh
git clone https://github.com/aliwert/cli-app
cd cli-app
go mod tidy
```

## Usage

- To run the application, use the following command:

```sh
go run ./
```

## Commands
- Add <description>: Add a new todo item with the given description.
- Edit <index> <new-description>: Edit an existing todo item by index.
- Delete <index>: Delete a todo item by index.
- Toggle <index>: Toggle the completion status of a todo item by index.
- List: List all todo items.

## Examples

- To add a new todo item:

```sh
go run main.go -Add "Buy milk"
```

- To edit an existing todo item:

```sh
go run main.go -Edit 1 "Buy eggs"
```

- To delete a todo item:

```sh
go run main.go -Delete 2
```

- To toggle the completion status of a todo item:

```sh
go run main.go -Toggle 1
```

- To list all todo items:

```sh
go run main.go -List
```
