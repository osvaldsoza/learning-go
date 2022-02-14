package util

import "fmt"

var Teste int

const NOME = "OSvaldo"

func ShowMessage() {
	fmt.Println("Message 1")
	showMessage() // private function
}
