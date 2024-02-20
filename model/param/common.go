package param

type ParamsPagination struct {
	Page  int `form:"page" binding:"required,min=1"`
	Limit int `form:"limit" binding:"required,min=10,max=50"`
}
