package main

import (
	"errors"
	"fmt"
)

/** Exercise 1
Exercício 1 - Impostos de salário

Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.
**/

const (
	TAX_LESS_THAN_150K  = 0.17
	TAX_ABOVE_THAN_150K = TAX_LESS_THAN_150K + 0.10
)

func calcTax(salary float64) float64 {
	if salary >= 150000 {
		return salary * TAX_ABOVE_THAN_150K
	} else {
		return salary * TAX_LESS_THAN_150K
	}
}

func doExercise1() {
	var salary float64 = 10000
	fmt.Printf("O imposto aplicado no salário de R$%.2f é de R$%.2f e representa %.0f%% do salário.\n", salary, calcTax(salary), (100 * calcTax(salary) / salary))
	salary = 100000
	fmt.Printf("O imposto aplicado no salário de R$%.2f é de R$%.2f e representa %.0f%% do salário.\n", salary, calcTax(salary), (100 * calcTax(salary) / salary))
	salary = 150000
	fmt.Printf("O imposto aplicado no salário de R$%.2f é de R$%.2f e representa %.0f%% do salário.\n", salary, calcTax(salary), (100 * calcTax(salary) / salary))
	salary = 200000
	fmt.Printf("O imposto aplicado no salário de R$%.2f é de R$%.2f e representa %.0f%% do salário.\n", salary, calcTax(salary), (100 * calcTax(salary) / salary))
}

/** Exercise 2
Exercício 2 - Calcular média
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
**/

func calcClassAverage(grades ...float64) float64 {
	var sum float64 = 0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func doExercise2() {
	mean := calcClassAverage(10, 5, 5, 7, 9.5)
	fmt.Printf("A média da turma é: %0.2f\n", mean)
	mean = calcClassAverage(4.5, 10, 9.8, 9.5, 0, 5.9, 7.1)
	fmt.Printf("A média da turma é: %0.2f\n", mean)
	mean = calcClassAverage(5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10)
	fmt.Printf("A média da turma é: %0.2f\n", mean)
}

/** Exercise 3
Exercício 3 - Calcular salário

Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria do funcionário.
Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais
**/

const (
	CATEGORY_C rune = 'C'
	CATEGORY_B rune = 'B'
	CATEGORY_A rune = 'A'
)

const (
	HOURLY_SALARY_C float64 = 1000
	HOURLY_SALARY_B float64 = 1500
	HOURLY_SALARY_A float64 = 3000
)

const (
	OVERTIME_SALARY_C float64 = 1.0
	OVERTIME_SALARY_B float64 = 1.2
	OVERTIME_SALARY_A float64 = 1.5
)

func calcSalary(hoursWorked uint, category rune) (float64, error) {
	hours := float64(hoursWorked)
	switch category {
	case CATEGORY_C:
		return HOURLY_SALARY_C * hours, nil
	case CATEGORY_B:
		if hoursWorked <= 160 {
			return HOURLY_SALARY_B * hours, nil
		} else {
			return (HOURLY_SALARY_B * OVERTIME_SALARY_B) * hours, nil
		}
	case CATEGORY_A:
		if hoursWorked <= 160 {
			return HOURLY_SALARY_A * hours, nil
		} else {
			return (HOURLY_SALARY_A * OVERTIME_SALARY_A) * hours, nil
		}
	}
	return 0, errors.New("Categoria inválida.")
}

func doExercise3() {
	salary, _ := calcSalary(162, CATEGORY_C)
	fmt.Printf("O funcionário de categoria C que trabalhou %v horas vai receber R$%.2f\n", 162, salary)
	salary, _ = calcSalary(176, CATEGORY_B)
	fmt.Printf("O funcionário de categoria B que trabalhou %v horas vai receber R$%.2f\n", 176, salary)
	salary, _ = calcSalary(172, CATEGORY_A)
	fmt.Printf("O funcionário de categoria A que trabalhou %v horas vai receber R$%.2f\n", 172, salary)
}

/** Exercise 4
Exercício 4 - Cálculo de estatísticas

Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior
Exemplo:

const (
minimum = "minimum"
average = "average"
maximum = "maximum"
)

...
minhaFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...
minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
**/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(grades ...float64) (min float64) {
	min = grades[0]
	for _, grade := range grades[1:] {
		if grade < min {
			min = grade
		}
	}
	return min
}

func averageFunc(grades ...float64) (average float64) {
	var sum float64 = 0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func maxFunc(grades ...float64) (max float64) {
	max = grades[0]
	for _, grade := range grades[1:] {
		if grade > max {
			max = grade
		}
	}
	return max
}

func operation(op string) (func(grades ...float64) float64, error) {
	switch op {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("Operação inválida.")
	}
}

func doExercise4() {
	minFn, err := operation(minimum)
	averageFn, err := operation(average)
	maxFn, err := operation(maximum)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		min := minFn(2, 3, 3, 4, 10, 2, 4, 5)
		average := averageFn(2, 3, 3, 4, 1, 2, 4, 5)
		max := maxFn(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println(min)
		fmt.Println(average)
		fmt.Println(max)
	}
}

/**
Exercício 5 - Cálculo da quantidade de alimento
Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
Exemplo:

const (
dog = "dog"
cat = "cat"
)

...
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)

**/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animalDog(quantity int) float64 {
	return float64(quantity) * 10
}

func animalCat(quantity int) float64 {
	return float64(quantity) * 5
}

func animalHamster(quantity int) float64 {
	return float64(quantity) * 0.25
}

func animalTarantula(quantity int) float64 {
	return float64(quantity) * 0.15
}

func Animal(animal string) (func(quantity int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("Animal inválido.")
	}
}

func doExercise5() {
	dogFn, err := Animal(dog)
	catFn, err := Animal(cat)
	hamsterFn, err := Animal(hamster)
	tarantulaFn, err := Animal(tarantula)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		var amount float64
		amount += dogFn(5)
		fmt.Println(amount)
		amount += catFn(8)
		fmt.Println(amount)
		amount += hamsterFn(4)
		fmt.Println(amount)
		amount += tarantulaFn(10)
		fmt.Println(amount)
	}

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
	fmt.Println("\nExercise 5...")
	doExercise5()
}
