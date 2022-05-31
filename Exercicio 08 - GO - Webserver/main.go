package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"webserver/model"

	"github.com/gin-gonic/gin"
)

type DatabaseInMemory struct {
	Products []model.Product
}

var databaseInMemory DatabaseInMemory

/*
Exercício 2.1 - Olá {nome}
1. Crie dentro da pasta go-web um arquivo chamado main.go
2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
“mensagem” e diga Olá seguido do seu nome.
3. Acesse o end-point para verificar se a resposta está correta.
*/

/*
Exercício 1.2 - Vamos filtrar nosso endpoint

Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser
capaz de filtrar todos os campos.
1. Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
2. Em seguida, ele gera a lógica do filtro para nossa matriz.
3. Retorne a matriz filtrada por meio do endpoint.

*/

/*
Exercício 2.2 - Get one endpoint

Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema
sempre tem que ser plural). Uma vez que o id é recebido, ele retorna a posição
correspondente.
1. Gere uma nova rota.
2. Gera um manipulador para a rota criada.
3. Dentro do manipulador, procure o item que você precisa.
4. Retorna o item de acordo com o id.
Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.
*/

func loadProductsInMemoryFromJSONFile() error {
	jsonFile, err := ioutil.ReadFile("./go-web/products.json")
	if err != nil {
		return errors.New(err.Error())
	}

	err = json.Unmarshal(jsonFile, &databaseInMemory)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func searchProducts(key string, values []string) ([]model.Product, error) {
	result := make([]model.Product, 0)
	for _, product := range databaseInMemory.Products {
		for _, value := range values {
			valueAttrOfProduct := reflect.Indirect(reflect.ValueOf(product))
			if valueAttrOfProduct.IsValid() {
				v := valueAttrOfProduct.FieldByName(strings.Title(key)).String()
				if strings.ToLower(v) == strings.ToLower(value) {
					result = append(result, product)
				}
			} else {
				return nil, fmt.Errorf("Invalid field: %s", string(key))
			}

		}
	}
	return result, nil
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func getProducts(c *gin.Context) {
	var err error
	copyProducts := databaseInMemory.Products

	searches := c.Request.URL.Query()

	for key, values := range searches {
		copyProducts, err = searchProducts(key, values)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, copyProducts)
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product *model.Product

	for _, p := range databaseInMemory.Products {
		if fmt.Sprint(p.ID) == id {
			product = &p
			break
		}
	}

	c.JSON(http.StatusOK, product)
}

func main() {
	err := loadProductsInMemoryFromJSONFile()
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	router.GET("/ping", ping)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)

	router.Run()
}
