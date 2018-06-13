package main


import (
	"fmt"

	"rsc.io/pdf"
)

func main() {
	file, err := pdf.Open("demo.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.NumPage())
}