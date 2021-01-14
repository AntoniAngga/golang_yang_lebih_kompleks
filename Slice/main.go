package main

import "fmt"

func main() {
	var gammingConsole []string

	gammingConsole = append(gammingConsole, "Playstation4")
	gammingConsole = append(gammingConsole, "Nintendo Switch")
	gammingConsole = append(gammingConsole, "Xbox One")

	for _, console := range gammingConsole {
		fmt.Println("Console: ", console)
	}
}