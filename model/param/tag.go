package param

type ParamsGetTags struct {
	ParamsPagination
	Name string `form:"name"`
}

type ParamsCreateTag struct {
	Name    string   `json:"name" binding:"required"`
	Icon    string   `json:"icon"`
	Slug    string   `json:"slug" binding:"required"`
	PostIDs []string `json:"postIDs"`
}

type ParamsUpdateTag struct {
	Name    string   `json:"name"`
	Icon    string   `json:"icon"`
	Slug    string   `json:"slug"`
	PostIDs []string `json:"postIDs"`
}
