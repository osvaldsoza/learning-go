package util

import "fmt"

func ShowMessage() {
	fmt.Println("Message 1")
	showMessage() // private function
}
