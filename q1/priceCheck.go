package priceCheck

import ()

type ()

func PriceCheckMap(products []string, productPrices []float64, productSold []string, soldPrice []float64) (errNum int) {
	// ====
	if len(products) != len(productPrices) {
		// should be checked, but it does not ask in the conditions
	}

	given := make(map[string]float64)
	for i := range products {
		given[products[i]] = productPrices[i]
	}

	// ====
	if len(productSold) != len(soldPrice) {
		// should be checked, but it does not ask in the conditions
	}

	sold := make(map[string]float64)
	for i := range productSold {
		sold[productSold[i]] = soldPrice[i]
	}

	// ===
	for i := range sold {
		if sold[i] != given[i] {
			errNum++
		}
	}

	return
}

func PriceCheckArray(products []string, productPrices []float64, productSold []string, soldPrice []float64) (errNum int) {
	if len(products) != len(productSold) {
		// 
	}

	for i := range products {
		for y := range productSold {
			if products[i] == productSold[y] {
				if productPrices[i] != soldPrice[y] {
					errNum++
				}
			}
		}
	}
	return
}