package param

type ParamsGetPosts struct {
	ParamsPagination
	Title      string   `form:"title"`
	Author     string   `form:"author"`
	Types      []string `form:"types[]"`
	Statuses   []string `form:"statuses[]"`
	CategoryID string   `form:"categoryID"`
	TagIDs     []string `form:"tagIDs[]"`
}

type ParamsCreatePost struct {
	Title      string   `json:"title" binding:"required"`
	Body       string   `json:"body" binding:"required"`
	Desc       string   `json:"desc" binding:"required"`
	Slug       string   `json:"slug" binding:"required"`
	Cover      string   `json:"cover"`
	Author     string   `json:"author" binding:"required"`
	Type       int      `json:"type" binding:"required"`
	Status     int      `json:"status" binding:"required"`
	Secret     string   `json:"secret"`
	CategoryID string   `json:"categoryID"`
	TagIDs     []string `json:"tagIDs"`
}

type ParamsUpdatePost struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Desc       string   `json:"desc"`
	Slug       string   `json:"slug"`
	Cover      string   `json:"cover"`
	Author     string   `json:"author"`
	Type       int      `json:"type"`
	Status     int      `json:"status"`
	Secret     string   `json:"secret"`
	CategoryID string   `json:"categoryID"`
	TagIDs     []string `json:"tagIDs"`
}
