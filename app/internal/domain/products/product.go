package Products

type Product struct {
	ProductId   ProductId
	name        ProductName
	publishYear ProductPublishYear
	TotalSales  ProductTotalSales
}

func newProduct(productId ProductId, name ProductName, publishYear ProductPublishYear, ProductTotalSales ProductTotalSales) *Product {
	return &Product{
		ProductId:   productId,
		name:        name,
		publishYear: publishYear,
		TotalSales:  ProductTotalSales,
	}
}

func NewFiled(productId string, name string, publishYear string, productTotalSales string) *Product {
	return newProduct(
		ProductId{value: productId},
		ProductName{value: name},
		ProductPublishYear{value: publishYear},
		ProductTotalSales{value: productTotalSales},
	)
}
