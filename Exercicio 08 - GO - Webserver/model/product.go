package model

/*
Exercício 1 - Estruturar um JSON

Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de acordo
com o tema.
Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico), publicação
(sim-não), data de criação.
Os usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de
criação.
Transações: id, código da transação (alfanumérico), moeda, valor, emissor (string), receptor
(string), data da transação.
1. Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema
escolhido, ex: products.json.
2. Dentro dele escrevi um JSON que permite ter uma matriz de produtos, usuários ou
transações com todas as suas variantes.
*/

type Product struct {
	ID        int     `form:"id" json:"id"`
	Name      string  `form:"name" json:"name"`
	Color     string  `form:"color" json:"color"`
	Price     float64 `form:"price" json:"price"`
	Stock     int     `form:"stock" json:"stock"`
	Code      string  `form:"code" json:"code"`
	Published bool    `form:"published" json:"published"`
	CreatedAt string  `form:"created_at" json:"created_at"`
}
