package main

import "fmt"

func main() {
	// A propriedade definida para a variável tem que permanecer veradedeira
	// antes, durante e após o loop
	n := 10
	soma := 0
	for i := 1; i < n; i++ {
		soma += 1
	}
	fmt.Println(soma)
	// Agora com o algoritmo definido, escolhemos uma propriedade
	// para fazer a prova, e ela deve ser invariante ao loop
	// Propriedade: a variável soma sempre resultará em
	// número inteiro positivo
	// Então se essa propriedade for verdadeira antes, durante e depois do
	// loop, o algoritmo está certo.
	// Inicialização: Soma é um inteiro positivo
	// Execução: Soma continua sendo um inteiro positivo
	// Término do Loop: Continua sendo positivo
	// Então está provado a corretude do algoritmo

}
