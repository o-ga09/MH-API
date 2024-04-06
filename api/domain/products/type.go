package Products

type ProductId struct{ value string }
type ProductName struct{ value string }
type ProductPublishYear struct{ value string }
type ProductTotalSales struct{ value string }

func (f *ProductId) GetID() string            { return f.value }
func (f *ProductName) GetName() string        { return f.value }
func (f *ProductPublishYear) GetYear() string { return f.value }
func (f *ProductTotalSales) GetSales() string { return f.value }
