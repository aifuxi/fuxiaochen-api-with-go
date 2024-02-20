package model

type ParamsCreateUser struct {
	Name            string `json:"name" binding:"required"`
	IsAdmin         bool   `json:"isAdmin"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
}

type ParamsUpdateUser struct {
	IsAdmin bool `json:"isAdmin"`
}

type ParamsLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
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

type ParamsCreateCategory struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type ParamsUpdateCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ParamsCreatePost struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Slug   string `json:"slug" binding:"required"`
	Cover  string `json:"cover"`
	Author string `json:"author" binding:"required"`
	Type   int    `json:"type" binding:"required"`
	Status int    `json:"status" binding:"required"`
	Secret string `json:"secret"`
	//CategoryID string   `json:"categoryID"`
	TagIDs []string `json:"tagIDs"`
}

type ParamsUpdatePost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Desc   string `json:"desc"`
	Slug   string `json:"slug"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
	Type   int    `json:"type"`
	Status int    `json:"status"`
	Secret string `json:"secret"`
	//CategoryID string   `json:"categoryID"`
	TagIDs []string `json:"tagIDs"`
}
