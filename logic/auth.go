package logic

import (
	"errors"
	"fuxiaochen-api-with-go/dao/mysql"
	"fuxiaochen-api-with-go/model"
	"fuxiaochen-api-with-go/util"
	"gorm.io/gorm"
)

func Login(params model.ParamsLogin) (string, error) {
	user, err := mysql.GetUserByName(params.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrorUsernameOrPassword
		}

		return "", err
	}

	// 检查密码是否正确
	if err = util.ComparePassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return "", ErrorUsernameOrPassword
	}

	token, err := util.NewToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
