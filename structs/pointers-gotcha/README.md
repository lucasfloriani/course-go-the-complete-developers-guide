# Informações sobre slice e array

Slices são um tipo de variavel que utiliza o mesmo principio dos arrays, porem não contendo um valor fixo de valores que pode receber.

Slice contem dentro de sua estrutura os valores:

* capacidade (quantos elementos pode conter)
* largura (dos elementos do slice)
* ponteiro (para um array por baixo dos panos)

## Como um slice guarda na memória RAM (bem resumido)

Primeiramente seria guardado os dados do ponteiro, da largura e da capacidade em um local da memória RAM, e o array criado pelo slice ficaria em outro local da memória.

Quando é enviado por parâmetro de uma função o slice, o Go copia os valores do slice e adiciona em um novo endereço de memória, porem ele somente vai copiar os 3 dados (pinteiro, largura e capacidade) e não criará um novo local de memoria para os valores do array.

## Outros tipos de variaveis que trabalha igual (por referência)

* slice
* maps
* channels
* pointers
* functions

## Tipos de variaveis que são passadas por valor

Precisam utilizar os comandos & ou * para trocar valores da mesma variavel

* int
* float
* string
* bool
* structs