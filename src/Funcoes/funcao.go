package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { //funcao com retorno
	return n1 + n2
}

func calculoMatematicos(n1, n2 int8) (int8, int8) { // funcao com mais de um retorno
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func dadosPessoa(nome string, idade int8) (string, int8) {
	return nome, idade
}

func main() {
	resultado := somar(12, 20)

	fmt.Println(resultado)

	var fun = func(txt string) string {
		return txt
	}

	txt := fun("Retorno da funcao")

	fmt.Println(txt)

	resultSoma, resultSub := calculoMatematicos(10, 15)
	fmt.Println(resultSoma, resultSub)

	resultSoma1, _ := calculoMatematicos(13, 15) //descartando o segundo retorno
	fmt.Println(resultSoma1)

	_, resultSub1 := calculoMatematicos(8, 15) //descartando o primeiro retorno
	fmt.Println(resultSub1)

	nome, idade := dadosPessoa("Osvaldo", 43)

	fmt.Printf("%s tem %d anos de idade", nome, idade)
}
