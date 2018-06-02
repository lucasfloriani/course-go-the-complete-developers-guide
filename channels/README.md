# Channels

Channels são utilizados para que diferentes go routines comuniquem entre si, sendo o unico meio para executar esta função.

Channel é básicamente um valor em um tipo de variavel que poderá ser passada entra as go routines existentes.

```go
c := make(chan string)
```

Para enviar dados atraves deste channel utilizamos a seguinte sintaxe

```go
channel <- 5
```

Para receber os dados deste channel podemos utilizar a seguinte sintaxe

```go
myNumber <- channel
```

Se necessário podemos passar para funções

```go
fmt.Println(<- channel)
```

OBS: Quando chamamos o retorno do channel, estamos bloqueando a execução de código esperando a resposta desta channel criada, algo como async/await do javascript.

OBS 2: Quando retornamos menos valores da channel que estamos criando de go routines, somente a quantidade de vezes do código de retorno serão utilizadas, porem caso seja criado mais retornos que go routines existentes o programa estará travado esperando uma resposta do channel.

O código principal de um programa em Go é executado dentro de uma go routine, sendo esta a principal dentre outras que podem ser criadas durante seu processo de execução.

Com a keyword 'go' nós podemos inicializar uma go routine (thread) em nosso programa, onde somente deve ser utilizado quando esta acompanhado com uma função que sera executada como uma thread.

```go
  go checkLink(link)
```

Quem decide a execução das go routines é o go scheduler.

Normalmente o Go utilizará somente um core para trabalhar com go routines, porem isto pode ser alterado.

## Concurrency

É a habilidade de rodar
multiplas funções quase ao mesmo tempo,
porem intercalando entre elas para
rodar os vários processos em modo de
concorrencia, disputando qual go routine
vai ter a vez de rodar o código.
Utilizado principalmente com um core.

## Parellelism

Utiliza dois ou mais cores
para executar códigos ao mesmo tempo.
