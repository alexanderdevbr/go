package main

import (
	"database/sql"
	"fmt"

	"github.com/alexanderdevbr/go/internal/entity"
	"github.com/alexanderdevbr/go/internal/infra/database"
	"github.com/alexanderdevbr/go/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// Metodo
func (c *Car) Start() {
	println(c.Model + " start!")
}
func (c *Car) ChangeColor(color string) {
	c.Color = color
}

// Funcao
func Soma(x, y int) int {
	return x + y
}

func TestNewOrder() {
	order, err := entity.NewOrder("1", -10, 1)
	if err != nil {
		println(err.Error())
	}
	if order != nil {
		println(order.Id)
	}
}

func TestCar() {
	car := Car{
		Model: "Ferrari",
		Color: "Vermelho",
	}

	car.Start()
	println(car.Color)
	car.ChangeColor("Blue")
	println(car.Color)
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	ord_rep := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(ord_rep)

	input := usecase.OrderInput{
		Id:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output.FinalPrice)

}
