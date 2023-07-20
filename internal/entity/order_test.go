package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestIfGetsAnErro(t *testing.T) {
	order := Order{}
	if order.Validate() == nil {
		t.Error("ID is required")
	}
}*/

// Validacao de erros esperados
func TestIfGetsAnErroIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id requerido")
}

func TestIfGetsAnErroIfPriceIsBlank(t *testing.T) {
	order := Order{Id: "Abc"}
	assert.Error(t, order.Validate(), "preco menor ou igual a zero")
}

func TestIfGetsAnErroIfTaxIsBlank(t *testing.T) {
	order := Order{Id: "Abc", Price: 10.0}
	assert.Error(t, order.Validate(), "taxa invalida")
}

// Testando se o calculo do preco final esta retornando o valor esperado
func TestFinalPrice(t *testing.T) {
	order := Order{Id: "Abcd", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "Abcd", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
