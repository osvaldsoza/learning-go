package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"osf/util"
)

func main() {
	fmt.Println("Hello World")
	util.ShowMessage() // public function
	erro := checkmail.ValidateFormat("osvald.soza@gmail.com")
	// [go mod tidy] retira todas dependencias que nao estao sendo usadas
	fmt.Println(erro)
}
