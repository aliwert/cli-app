package main

import (
	"fmt"

	"github.com/fatih/color"
)

const logo = `
 ____  _____  ____   _____     ___  __    ____      __    ____  ____ 
(_  _)(  _  )(  _ \ (  _  )   / __)(  )  (_  _)    /__\  (  _ \(  _ \
  )(   )(_)(  )(_) ) )(_)(   ( (__  )(__  _)(_    /(__)\  )___/ )___/
 (__) (_____)(____/ (_____)   \___)(____)(____)  (__)(__)(__)  (__)  
`

func PrintLogo() {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Print(green(logo + "\n"))
	fmt.Printf("%40s%s\n", "", green("A simple CLI todo list application"))
}
