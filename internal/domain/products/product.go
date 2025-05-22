package Products

type Product struct {
	productId   ProductId
	name        ProductName
	publishYear ProductPublishYear
	totalSales  ProductTotalSales
}

func newProduct(productId ProductId, name ProductName, publishYear ProductPublishYear, ProductTotalSales ProductTotalSales) *Product {
	return &Product{
		productId:   productId,
		name:        name,
		publishYear: publishYear,
		totalSales:  ProductTotalSales,
	}
}

func NewProduct(productId string, name string, publishYear string, productTotalSales string) *Product {
	return newProduct(
		ProductId{value: productId},
		ProductName{value: name},
		ProductPublishYear{value: publishYear},
		ProductTotalSales{value: productTotalSales},
	)
}
