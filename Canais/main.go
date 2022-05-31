package main

import (
	"fmt"
	"time"
)

func Hello(message string) {
	fmt.Println(message)
}

/*
Hipótese 1 -
Somente o Main é capaz de finalizar todos os processos em aberto através de go routines.
*/

func Hipotese1() {
	fmt.Println("[H1] -Hipótese 1. Chamando [func].")
	go Hello("[Hello] - Olá mundo. Executando uma função.")
	fmt.Println("[H1] -Terminando a Funcao")
}

/*
Hipótese 2 -
Chan é um espaço compartilhado de memória. O processo para na entrada do canal e só continua
sua execução após um valor ser escrito no canal.
*/

func Hello2(c chan string) {
	fmt.Println("[func] - Hello2")
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Hello %d\n", i)
	}
}

func Hipotese2() {
	c := make(chan string)
	fmt.Println("[H2] - Chamando [func].")
	fmt.Println("Canal")
	go Hello2(c)
	for i := 0; i < 10; i++ {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	opcao := "H2"
	switch opcao {
	case "H1":
		fmt.Println("[main] - Iniciando main.")
		Hipotese1()
		fmt.Println("[main] - esperar")
	case "H2":
		Hipotese2()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("[main] - Finalizando main...")
}
