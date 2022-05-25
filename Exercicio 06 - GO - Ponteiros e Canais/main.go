package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/** Exercise 1
Exercício 1 - Rede social

Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
senha
E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha
**/

type User1 struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

func (u *User1) SetFullname(name, lastName string) {
	u.Name = name
	u.Lastname = lastName
}

func (u *User1) SetAge(age int) {
	u.Age = age
}

func (u *User1) SetEmail(email string) {
	u.Email = email
}

func (u *User1) SetPassword(password string) {
	u.Password = password
}

func doExercise1() {
	u1 := User1{Name: "Rodrigu", Age: 31, Lastname: "Auves", Email: "rodrigo@alves.com", Password: "mypass1234"}
	fmt.Println(u1)
	u1.SetFullname("Rodrigo", "Alves")
	u1.SetAge(30)
	fmt.Println(u1)
	u1.SetEmail("rodrigo.alves@gmail.com")
	u1.SetPassword("my@pass@1234")
	fmt.Println(u1)
}

/** Exercise 2
Exercício 2 - E-commerce (Parte II)
Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.
Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.
**/

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

type User2 struct {
	Name     string
	Lastname string
	Email    string
	Products []Product
}

func AddProduct(user *User2, product Product, quantity int) {
	product.Quantity = quantity
	user.Products = append(user.Products, product)
}

func DeleteProducts(user *User2) {
	user.Products = nil
}

func doExercise2() {
	milk := NewProduct("Milk", 4.90)
	chocolate := NewProduct("Chocolate", 6.99)

	me := User2{Name: "Vinnicyus", Lastname: "Gracindo", Email: "foo@bar.com"}
	AddProduct(&me, *milk, 3)
	AddProduct(&me, *chocolate, 1)
	fmt.Println(me)
	DeleteProducts(&me)
	fmt.Println(me)
}

/** Exercise 3
Exercício 3 - Calcular Preço (Part II)
Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

Precisamos de 3 estruturas:
- Produtos: nome, preço, quantidade.
- Serviços: nome, preço, minutos trabalhados.
- Manutenção: nome, preço.
Precisamos de 3 funções:
- Somar Produtos: recebe um array de produto e devolve o preço total (preço *
quantidade).
- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
tivesse trabalhado meia hora).
- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
tela (somando o total dos 3).
**/

type Product2 struct {
	Name     string
	Price    float64
	Quantity int
}

func (p Product2) Total(c chan float64) {
	c <- p.Price * float64(p.Quantity)
}

func SumProducts(products []Product2, total chan float64) {
	c := make(chan float64)
	var sum float64
	for _, product := range products {
		go product.Total(c)
	}

	for range products {
		sum += <-c
	}
	total <- sum
}

type Service struct {
	Name       string
	Price      float64
	WorkedTime int
}

func (s Service) Total(c chan float64) {
	if s.WorkedTime < 30 {
		s.WorkedTime = 30
	}
	c <- s.Price * float64(s.WorkedTime)
}

func SumServices(services []Service, total chan float64) {
	var sum float64
	c := make(chan float64)
	for _, service := range services {
		go service.Total(c)
	}

	for range services {
		sum += <-c
	}

	total <- sum
}

type Maintenance struct {
	Name  string
	Price float64
}

func (m Maintenance) Total(c chan float64) {
	c <- m.Price
}

func SumMaintenances(maintenances []Maintenance, total chan float64) {
	var sum float64
	c := make(chan float64)
	for _, maintenance := range maintenances {
		go maintenance.Total(c)
	}

	for range maintenances {
		sum += <-c
	}

	total <- sum
}

func doExercise3() {
	sumProducts := make(chan float64)
	sumServices := make(chan float64)
	sumMaintenances := make(chan float64)

	products := []Product2{
		{"Banana", 6.99, 12},
		{"Arroz", 4.90, 5},
		{"Feijão", 6.29, 2},
		{"Óleo de Soja", 11.99, 1},
	}

	services := []Service{
		{"Jardinagem", 1, 15},
		{"Cabeamento Estruturado", 0.22, 120},
	}

	maintenances := []Maintenance{
		{"Manutenção Preventiva de Condicionadores de Ar", 220.10},
		{"Manutenção de Computador", 199},
		{"Manutenção de Hidrantes", 499.90},
	}

	go SumProducts(products, sumProducts)
	go SumServices(services, sumServices)
	go SumMaintenances(maintenances, sumMaintenances)

	fmt.Printf("Soma de todos os produtos: R$%.2f\n", <-sumProducts)
	fmt.Printf("Soma de todos os serviços: R$%.2f\n", <-sumServices)
	fmt.Printf("Soma de todos as manutenções: R$%.2f\n", <-sumMaintenances)
}

/** Exercise 4
Exercício 4 - Ordenamento
Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus
serviços.
Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
- uma matriz de inteiros com 100 valores
- uma matriz de inteiros com 1000 valores
- uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand:
package main
import (
"math/rand"
)

func main() {
variavel := rand.Perm(100)
variave2 := rand.Perm(1000)
variave3 := rand.Perm(10000)
}
Cada um deve ser ordenado por:
- Inserção
- grupo
- seleção

Uma rotina para cada execução de classificação
Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de
1000 e depois a ordenação de 10000.
Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual
ordenação ficou melhor para cada arranjo.
**/

func ExecTime(sort func([]int), items []int) {
	start := time.Now()
	sort(items)
	elapsed := time.Since(start)
	log.Printf("Tempo Gasto: %s", elapsed)
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func doExercise4() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)

	fmt.Println("\n\nExecutando insertionSort no array de 100 inteiros")
	ExecTime(insertionsort, variavel)
	fmt.Println("Executando selectionsort no array de 100 inteiros")
	ExecTime(selectionsort, variavel)

	fmt.Println("********************")

	fmt.Println("\n\nExecutando insertionSort no array de 1000 inteiros")
	ExecTime(insertionsort, variavel2)
	fmt.Println("Executando selectionsort no array de 1000 inteiros")
	ExecTime(selectionsort, variavel)
	fmt.Println("********************")

	fmt.Println("\n\nExecutando insertionSort no array de 10000 inteiros")
	ExecTime(insertionsort, variavel3)
	fmt.Println("Executando selectionsort no array de 10000 inteiros")
	ExecTime(selectionsort, variavel)
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
