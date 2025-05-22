package Products

type Products []Product

type ProductId struct{ value string }
type ProductName struct{ value string }
type ProductPublishYear struct{ value string }
type ProductTotalSales struct{ value string }

func (f *Product) GetID() string    { return f.productId.value }
func (f *Product) GetName() string  { return f.name.value }
func (f *Product) GetYear() string  { return f.publishYear.value }
func (f *Product) GetSales() string { return f.totalSales.value }
