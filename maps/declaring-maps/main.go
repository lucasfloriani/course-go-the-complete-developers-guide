package main

import "fmt"

func main() {
	// (Primeiro modo de criar um map)
	// Um map com keys do tipo string
	// e valores do tipo string
	fmt.Println("===== Primeiro Modo =====")
	var colors_1 map[string]string
	fmt.Println(colors_1)

	// (Segundo modo de criar um map)
	// Um map com keys do tipo string
	// e valores do tipo string
	fmt.Println("\n===== Segundo Modo =====")
	colors_2 := map[string]string{
		"red":  "#FF0000",
		"blue": "#00FF00",
	}
	fmt.Println(colors_2)

	// (Terceiro modo de criar um map)
	// Não se pode acessar os dados igual em struct
	// com colors_3.white, use [] para isso.
	// Isto não poderia ser usado em map porque
	// a chave tem um tipo de variavel, diferente
	// da struct que não tem
	fmt.Println("\n===== Terceiro Modo =====")
	colors_3 := make(map[string]string)
	fmt.Println(colors_3)
	colors_3["white"] = "#FFFFFF"
	fmt.Println(colors_3)
	// Remover um valor do map
	delete(colors_3, "white")
	fmt.Println(colors_3)
}
