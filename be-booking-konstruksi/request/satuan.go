package request

type Satuan struct {
	Name string `json:"name" form:"name" binding:"required"`
}
