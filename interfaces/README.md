# Interfaces

Os tipos de variaveis podem ser separadas em 2 tipos, tipo concreto e tipo de interface.
O Tipo concreto são variaveis normais como map, struct, int, string, etc, alem das variaveis criadas com os tipos bases como o tipo deck que foi criado em exercicios anteriores.
O tipo interface são variaveis criadas que 'herdam' as funções do tipo interface e que contem dentro delas, regras para determinar se um tipo pertence ou não pertence tambem a este tipo (interface).
São principalmente usadas para criar funções utilizadas por tipos que tambem contem o tipo interface, limitando assim a duplicação de código e deixando-o mais limpo.

Interfaces não podem ser utilizados como valores reais, porem são utilizados como possuidores de funções em comum entre tipos que pertencem a ele.
