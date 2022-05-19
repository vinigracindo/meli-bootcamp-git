package main

import "fmt"

/**
Exercício 1 - Imprimindo o nome na tela
1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
2. Imprima no terminal o valor de cada variável.
**/

func doExercicio1() {
	nome := "Maria Joana"
	idade := 28

	fmt.Println(nome, idade)
}

/**
Exercício 2 - Clima
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.

1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?
**/

func doExercicio2() {
	var (
		temp    float32 = 32.5
		umidade float32 = 19.8
		atm     float32 = 760
	)

	fmt.Printf("Temperatura: %v, Umidade: %v e Pressão atmosférica: %v\n", temp, umidade, atm)
}

/**
Exercício 3 - Declaração de variáveis

Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.

var 1nome string
var sobrenome string
var int idade
1sobrenome := 6
var licenca_para_dirigir = true
var estatura da pessoa int
quantidadeDeFilhos := 2
**/

func doExercicio3() {
	var nome string                 // O nome da varíavel não pode começar com número.
	var sobrenome string            // Correto
	var idade int                   // O Tipo vem após o nome da variável
	sobrenome = "6"                 // Já declarada como string
	var licenca_para_dirigir = true // Correto
	var estaturaDaPessoa int        // O nome da variável não pode conter espaços em branco.
	quantidadeDeFilhos := 2

	fmt.Println(nome, sobrenome, idade, licenca_para_dirigir, estaturaDaPessoa, quantidadeDeFilhos)
}

/**
Exercício 4 - Tipos de dados

Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
pediu a ajuda de um desenvolvedor experiente que pode:
1. Revisar o código e realizar as devidas correções.

var sobrenome string = "Silva"
var idade int = "25"
boolean := "false";
var salario string = 4585.90
var nome string = "Fellipe"
**/

func doExercicio4() {
	var sobrenome string = "Silva" // Correto
	var idade int = 25             // Com aspas é string.
	boolean := false               // Com aspas é string.
	var salario float32 = 4585.90  // Deve ser float.
	var nome string = "Fellipe"    // Correto

	fmt.Println(sobrenome, idade, boolean, salario, nome)
}

// Main

func main() {
	fmt.Println("Executando Exercício 1 - Imprimindo o nome na tela")
	doExercicio1()
	fmt.Println("Executando Exercício 2 - Clima")
	doExercicio2()
	fmt.Println("Executando Exercício 3 - Declaração de variáveis")
	doExercicio3()
	fmt.Println("Executando Exercício 4 - Tipos de dados")
	doExercicio4()
}
