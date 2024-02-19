package mysql

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"gorm.io/gorm/clause"
)

func GetUsers() (users []model.User, err error) {
	result := global.DB.Find(&users)

	return users, result.Error
}

func GetUserByID(id int64) (user model.User, err error) {
	result := global.DB.First(&user, id)

	return user, result.Error
}

func GetUserByName(name string) (user model.User, err error) {
	result := global.DB.First(&user, model.User{Name: name})

	return user, result.Error
}

func DeleteUserByID(id int64) (user model.User, err error) {
	user, err = GetUserByID(id)
	if err != nil {
		return model.User{}, err
	}

	result := global.DB.Clauses(clause.Returning{}).Delete(&user)

	return user, result.Error
}

func UpdateUserByID(id int64, params model.ParamsUpdateUser) (user model.User, err error) {
	user, err = GetUserByID(id)
	if err != nil {
		return model.User{}, err
	}

	result := global.DB.Model(&user).Updates(map[string]any{"IsAdmin": params.IsAdmin})

	return user, result.Error
}

func CreateUser(params model.ParamsCreateUser) (user model.User, err error) {
	user = model.User{
		Name:     params.Name,
		Password: params.Password,
		IsAdmin:  params.IsAdmin,
	}
	result := global.DB.Create(&user)

	return user, result.Error
}
