# Channels

O código principal de um programa em Go é executado dentro de uma go routine, sendo esta a principal dentre outras que podem ser criadas durante seu processo de execução.

Com a keyword 'go' nós podemos inicializar uma go routine (thread) em nosso programa, na maioria das vezes acompanhado com este comando nós adicionamos uma função que sera executada como uma thread.

```go
  go checkLink(link)
```

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
