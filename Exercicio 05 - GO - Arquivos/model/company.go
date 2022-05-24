package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type company struct {
	name string
}

func (c company) createDatabaseIfDoesNotExist() bool {
	_, err := os.ReadFile(fmt.Sprintf("./%s.txt", c.name))

	if err != nil {
		os.WriteFile(fmt.Sprintf("./%s.txt", c.name), []byte{}, 0644)
		c.insertIntoFile(fmt.Sprintf("%s\t%s\t%s\t\n", "ID", "Preco", "Quantidade"))
		return false
	}

	return true
}

func (c company) insertIntoFile(line string) error {
	file, err := os.OpenFile(fmt.Sprintf("./%s.txt", c.name), os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Printf("failed reading file: %s\n", err)
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(line); err != nil {
		log.Println(err)
		return err
	}

	log.Println("Saved.")
	return nil
}

func (c company) newSale(id string, price float64, quantity int) {
	err := c.insertIntoFile(fmt.Sprintf("%s\t%.2f\t%d\t\n", id, price, quantity))
	if err != nil {
		log.Fatalf("failed: %s", err.Error())
	}
}

func (c company) NewSales(items string) {
	for _, item := range strings.Split(items, ";") {
		result := strings.Split(item, ",")
		if len(result) != 3 {
			fmt.Printf("Error in that item %v. The line must have 3 variables separeted by comma (,).\n", result)
		} else if len(result) > 1 {
			price, _ := strconv.ParseFloat(result[1], 32)
			quantity, _ := strconv.Atoi(result[2])
			c.newSale(result[0], price, quantity)
		}
	}
}

func (c company) TotalSales() {
	file, err := os.Open(fmt.Sprintf("./%s.txt", c.name))
	if err != nil {
		log.Fatalln("The database doesn't exists.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var total float64 = 0

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		price, err := strconv.ParseFloat(strings.Split(text, "\t")[1], 64)
		if err == nil {
			total += price
		}
	}

	fmt.Printf("\t%.2f\n", total)

}

func NewCompany(name string) company {
	c := company{name: name}
	c.createDatabaseIfDoesNotExist()
	return c
}
