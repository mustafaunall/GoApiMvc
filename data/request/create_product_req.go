package request

type CreateProductRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}
