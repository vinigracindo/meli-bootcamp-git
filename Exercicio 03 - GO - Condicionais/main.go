package main

import "fmt"

/**
Exercício 1 - Letras de uma palavra

A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
letras que ela contém.
2. Em seguida, imprimi cada uma das letras.
**/

func doExercise1() {
	fmt.Println("Digite a palavra: ")
	word := "Brasil"

	fmt.Printf("O tamanho da palavra \"%v\" é %d\n", word, len(word))

	fmt.Printf("Soletrando \"%v\": ", word)
	for _, c := range word {
		fmt.Printf("%c ", c)
	}
}

/**
Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
**/

func doExercise2() {
	idade := 30
	tempoAtividadeEmMeses := 12
	salario := 100000

	if idade < 22 || tempoAtividadeEmMeses < 12 {
		fmt.Println("Cliente não elegível para receber empréstimo.")
	} else if salario <= 100000 {
		fmt.Println("Cliente elegível para empréstimo com taxa de juros.")
	} else {
		fmt.Println("Cliente elegível para empréstimo com isenção de taxa de juros.")
	}
}

/**
Exercício 3 - A que mês corresponde?

Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.
**/

func doExercise3() {
	switch mes := 8; mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês inválido.")
	}
}

/**
Exercício 4 - Quantos anos tem...

Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.
**/

func doExercise4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	employee := "Benjamin"

	// imprimir a idade de Benjamin.
	fmt.Printf("A idade de %v é %v.\n", employee, employees[employee])

	// Saiba quantos de seus funcionários têm mais de 21 anos.
	employeesOver21 := 0
	for _, value := range employees {
		if value > 21 {
			employeesOver21++
		}
	}

	fmt.Printf("O número de funcionários acima de 21 anos: %v\n", employeesOver21)

	// Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
	employees["Frederico"] = 25
	fmt.Println(employees)

	// Excluir Pedro do mapa.
	delete(employees, "Pedro")
	fmt.Println(employees)

}

func main() {
	fmt.Println("\nExercise 1...")
	doExercise1()
	fmt.Println("\n-------------------")
	fmt.Println("\nExercise 2...")
	doExercise2()
	fmt.Println("\n-------------------")
	fmt.Println("\nExercise 3...")
	doExercise3()
	fmt.Println("\n-------------------")
	fmt.Println("\nExercise 4...")
	doExercise4()
}
