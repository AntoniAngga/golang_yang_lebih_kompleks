package main

import "fmt"

func main() {
	// Di isi Tidak Langsung
	//var languages [5]string
	//languages[0] = "Go"
	//languages[1] = "Ruby"
	//languages[2] = "Javascript"
	//languages[3] = "C#"
	//languages[4] = "Python"

	// Di isi Langsung
	languages := [...]string{
		"Ruby",
		"Python",
		"Javascript",
		"Go",
		"C#",
		"Kotlin",
	}

	fmt.Println(languages)
	length := len(languages)

	fmt.Println(length)

	for index, lang := range languages{
		fmt.Println("Index: ", index, " language : ", lang)
	}

}