package priceCheck

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestPriceCheckArray(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 400.89}

	v := PriceCheckArray(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}


func TestPriceCheckArray_2(t *testing.T) {
	products := []string{"eggs", "milk", "cheese"}
	productPrices := []float64{2.89, 3.29, 5.79}
	productSold := []string{"eggs", "eggs", "cheese", "milk"}
	soldPrice := []float64{2.89, 2.99, 5.97, 3.29}

	v := PriceCheckArray(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}


func TestPriceCheckMap(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 400.89}

	v := PriceCheckMap(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}

func TestPriceCheckMap_2(t *testing.T) {
	products := []string{"eggs", "milk", "cheese"}
	productPrices := []float64{2.89, 3.29, 5.79}
	productSold := []string{"eggs", "eggs", "cheese", "milk"}
	soldPrice := []float64{2.89, 2.99, 5.97, 3.29}

	v := PriceCheckMap(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}

func TestPriceCheckArrayWithError(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 345.99}

	v := PriceCheckArray(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}

func TestPriceCheckMapWithError(t *testing.T) {
	products := []string{"rice", "sugar", "wheat", "cheese"}
	productPrices := []float64{16.89, 56.92, 20.89, 345.99}
	productSold := []string{"rice", "cheese"}
	soldPrice := []float64{18.99, 345.99}

	v := PriceCheckMap(products, productPrices, productSold, soldPrice)
	if v != 2 {
		t.Fatalf("\t%s\tExpected 2, got %d", failed, v)
	}

	t.Logf("\t%s\tTest passed", success)
}