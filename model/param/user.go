package param

type ParamsGetUsers struct {
	ParamsPagination
	Name string `form:"name"`
}

type ParamsCreateUser struct {
	Name            string `json:"name" binding:"required"`
	IsAdmin         bool   `json:"isAdmin"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
}

type ParamsUpdateUser struct {
	IsAdmin bool `json:"isAdmin"`
}
