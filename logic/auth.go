package logic

import (
	"errors"
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/model/param"
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

func Login(params param.ParamsLogin) (token string, user model.User, err error) {

	if user, err = mysql.GetUserByName(params.Name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", model.User{}, ErrorUsernameOrPassword
		}

		return "", model.User{}, err
	}

	// 检查密码是否正确
	if err = util.ComparePassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return "", model.User{}, ErrorUsernameOrPassword
	}

	if token, err = util.NewToken(user.ID); err != nil {
		return "", model.User{}, err
	}

	return token, user, nil
}
