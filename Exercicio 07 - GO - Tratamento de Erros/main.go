package main

import (
	"errors"
	"fmt"
)

/** Exercise 1
Exercício 1 - Imposto sobre o salário #1
1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
tipo “int”.
2. Crie um erro personalizado com uma struct que implemente “Error()” com a
mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
console a mensagem “Necessário pagamento de imposto”.
**/

const MINIMUM_WAGE = 15000

type MinimumWageError struct{}

func (m MinimumWageError) Error() string {
	return fmt.Sprintf("Error: The minimum wage cannot be less than %d.", MINIMUM_WAGE)
}

func validateWage(wage int) error {
	if wage < MINIMUM_WAGE {
		return MinimumWageError{}
	}
	return nil
}

func doExercise1() {
	var wage int

	wage = 10000
	err := validateWage(wage)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}

	wage = 16000
	err = validateWage(wage)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}

/** Exercise 2
Exercício 2 - Imposto sobre o salário #2
Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
“Error()”, seja implementado “errors.New()”.
**/

func doExercise2() {
	var wage int

	wage = 10000
	err := validateWage(wage)
	if err != nil {
		fmt.Println(errors.New("Error: The minimum wage cannot be less than 15000"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}

	wage = 16000
	err = validateWage(wage)
	if err != nil {
		fmt.Println(errors.New("Error: The minimum wage cannot be less than 15000"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}

/** Exercise 3
Exercício 3 - Imposto sobre o salário #3

Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).
**/

func doExercise3() {
	var wage int

	wage = 10000
	err := validateWage(wage)
	if err != nil {

		fmt.Println(errors.New("Error: The minimum wage cannot be less than 15000"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}

	wage = 16000
	err = validateWage(wage)
	if err != nil {
		fmt.Println(errors.New("Error: The minimum wage cannot be less than 15000"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}

/** Exercise 4

**/

func doExercise4() {

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
