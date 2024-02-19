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
