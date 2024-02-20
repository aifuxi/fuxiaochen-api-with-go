package param

type ParamsGetCategories struct {
	ParamsPagination
	Name string `form:"name"`
}

type ParamsCreateCategory struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type ParamsUpdateCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
