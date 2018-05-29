package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// resp vai ser uma variavel do tipo *Response (custom struct).
	//
	// Response tem uma variavel chamada Body que é responsavel
	// por retornar o Body da requisição, porem este valor é retornado
	// como tipo io.ReadCloser
	//
	// ReadCloser é uma interface que contem Reader e Closer,
	// estes dois valores são 2 interfaces que acabam formando o
	// ReadCloser.
	// Esta interface é criada através do uso de composição e não de
	// herança.
	//
	// Reader é uma interface que contem o método Read que recebe
	// um parametro []byte e retorna um valor de int e um error.
	//
	// Closer é uma interface que contem o método Close que retorna
	// um valor error.
	//
	// Resumindo, a variavel Body da struct Response poderá receber
	// qualquer valor atrelado a ela somente se atender os requerimentos
	// da interface atrelada a ela, que no caso seria as funções Read e Close

	// Reader é uma interface usada para ler qualquer tipo de entrada
	// com a função Read como imagens, requests http, arquivos de texto, etc.
	// Principalmente usada como uma função em comum que centraliza uma função
	// que poderia ser duplicada se não a utiliza-se.

	// make é uma função da biblioteca padrão que recebe um slice
	// e no segundo parâmetro recebe um integer de quantos espacos em branco o slice
	// vai ter, algo como deixar um valor especifico de inicio como um array.
	// É criado um tamanho grande para o tamanho pq a função Read não aumenta o tamanho
	// do slice durante a sua execução.
	//
	// Não é utilizado deste modo a criação de slice de bytes, nas bibliotecas padrões
	// do Go existem funções próprias para isto.
	bs := make([]byte, 99999)
	// Na função Read ele ira ler os dados do Body do request e passar para o slice
	// de bytes criado.
	resp.Body.Read(bs)
	// []byte é essencialmente uma string, porem com os valores ASCI dos caracteres.
	// Imprime todo o html do google.com
	fmt.Println(string(bs))

	// Resumindo, a interface Read nos auxilia para não ter que criar várias funções
	// para ler os dados de cada input no sistema, como por exemplo:
	// * func lerImagem()
	// * func lerArquivoDeTexto()
	// * func lerAudio()
	// e assim por diante, demonstrando assim o pode da utilização de interfaces.
}
