package admin

type ProductsHandler struct{}

func (handler ProductsHandler) GetData() string {
	return "This is the products handler"
}
