package priceCheck

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestPriceCheck(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 400.89}

	v := PriceCheck(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}

func TestPriceCheck2(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 345.99}

	v := PriceCheck(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}

func TestPriceCheck3(t *testing.T) {
	products := []string{"eggs", "milk", "cheese"}
	productPrices := []float64{2.89, 3.29, 5.79}
	productSold := []string{"eggs", "eggs", "cheese", "milk"}
	soldPrice := []float64{2.89, 2.99, 5.97, 3.29}

	v := PriceCheck(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}
