# Informações básicas sobre Go

## Comandos Go Cli

* go build: compila os arquivos e cria um .exe
* go run: compila e executa um ou mais arquivos
* go fmt: formata todo o código em cada arquivo no diretório atual
* go install: compila e instala os pacotes referênciados no código no import
* go get: utilizado para baixar arquivos de pacotes terceiros
* go test: roda qualquer teste associado com o projeto atual

## Pacotes

Pacotes são criados para juntar arquivos que tem propósito parecido em sua codificação, onde podem ser separados em 2 tipos:

* Executavel: Gera um arquivo que pode ser rodado (Pacote com nome 'main' e precisa ter uma função chamada 'main')
* Reusavel: Código usado como 'helper'. Ótimo lugar para reutilização de lógica.
