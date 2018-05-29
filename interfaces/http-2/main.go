package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Executa a mesma coisa que o código:
	//// bs := make([]byte, 99999)
	//// resp.Body.Read(bs)
	//// fmt.Println(string(bs))
	//
	// Esta função diferente do que utilizamos antes,
	// le os dados passados no parâmetro e realiza o output
	// ao inves de fazer o input destes valores.
	//
	// io.Copy recebe 2 parâmetros, o primeiro é uma variavel
	// que implementa a interface Writer, que no caso necessita
	// de uma função Write. O segundo parâmetro é uma variavel
	// implementa a interface Reader, a qual contem necessita da
	// função Read.
	// Basicamente os próprios parâmetros já dizem o que a função
	// realizado, que no caso é ler os dados de um input e salvar
	// em algum output utilizando os parâmetros que contem as interfaces
	// que realizam escrita e leitura.
	//
	// Um tipo de variavel que implementa está função é o os.Stdout,
	// sua função é imprimir dados num dos 3 arquivos principais de
	// output do Go, sendo este o arquivo de saida. Existem outros 2
	// arquivos com a mesma função, sendo que o de entrada se chama
	// os.Stdin e o de saida de erros se chama Stderr.
	io.Copy(os.Stdout, resp.Body)
}
