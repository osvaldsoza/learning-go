package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"osf/util"
)

func main() {
	fmt.Println("Hello World")
	util.ShowMessage()
	erro := checkmail.ValidateFormat("osvald.soza@gmail.com")
	fmt.Println(erro)
}
