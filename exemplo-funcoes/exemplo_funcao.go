package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	texto := "Barbosa tem 18 anos."
	expr := regexp.MustCompile("\\d")

	fmt.Println(expr.ReplaceAllString(texto, "2"))

}

func ImprimirVersao() {
	fmt.Println("1.12")
}

func ImprimirSaudacao(nome string) {
	fmt.Printf("Olá, %s\n", nome)
}

func ImprimirDados(nome string, idade int) {
	fmt.Printf("%s, %d anos. \n", nome, idade)
}

func ImprimirSoma(a, b int, texto string) {
	fmt.Printf("%s: %d\n", texto, a+b)
}

func Somar(a, b int) int {
	return a + b
}

func PrecoFinal(precoCusto float64) (float64, float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoFinalDolar := precoCusto * fatorLucro
	return precoFinalDolar, precoFinalDolar * taxaConversao
}

func PrecoFinalNomeado(precoCusto float64) (
	precoDolar float64,
	precoReal float64) {

	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return
}

func CriarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")

		arq, err := os.Create(caminhoArquivo)

		// defer obriga função executar o Close antes de retornar
		defer arq.Close()

		if err != nil {
			fmt.Printf("Erro ao criar arquivo %s: %v\n", nome, err)
			os.Exit(1)
		}

		fmt.Printf("Arquivo %s criado.\n", arq.Name())
	}
}

func Fibonacci() {
	a, b := 0, 1

	fib := func() int {
		a, b = b, a+b

		return a
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", fib())
	}
}

func Agregar(
	valores []int,
	inicial int,
	fn func(n, m int) int,
) int {
	agregado := inicial

	for _, v := range valores {
		agregado = fn(v, agregado)
	}

	return agregado
}
