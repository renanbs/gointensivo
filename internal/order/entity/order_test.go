package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyID(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestEmptyPrice(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestEmptyTax(t *testing.T) {
	order := Order{ID: "123", Price: 123}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestShouldHaveOrder(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.1}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.1, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestShouldHaveOrder_when_instantiate(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.1)

	assert.Nil(t, order.IsValid())
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.1, order.Tax)
}

func TestShouldNotHaveAnOrder(t *testing.T) {

	order, err := NewOrder("123", 10.0, 0)
	assert.Error(t, err, "invalid tax")
	assert.Nil(t, order)
}

func TestCalculateFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2)
	assert.Nil(t, err)
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}
