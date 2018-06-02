package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://askhdkajsd.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Normal loop channel syntax
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative loop channel syntax
	// Faz a mesma coisa que o loop normal
	// porem com um código mais entendivel.
	//
	// Essa função itera sobre o channel
	// porem ele espera uma resposta no range c,
	// então quando ele retorna, o valor vai para
	// a variavel link, sendo então utilizada na criação
	// de uma outra go routine, fazendo assim um loop
	// infinito
	for link := range c {
		go checkLink(link, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
