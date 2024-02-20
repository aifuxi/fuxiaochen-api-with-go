package controller

type ResponseCode int

const (
	CodeSuccess             ResponseCode = 0
	CodeInternalServerError ResponseCode = 1

	CodeIncorrectUsernameOrPassword ResponseCode = 10001
	CodeIncorrectRequestParams      ResponseCode = 10002
	CodeIncorrectIDParams           ResponseCode = 10003

	CodeTokenNotFound        ResponseCode = 10004
	CodeIncorrectToken       ResponseCode = 10005
	CodeIncorrectTokenFormat ResponseCode = 10006
)

var codeMsg = map[ResponseCode]string{
	CodeSuccess:             "请求成功",
	CodeInternalServerError: "服务器内部错误",

	CodeIncorrectUsernameOrPassword: "用户名或密码错误",
	CodeIncorrectRequestParams:      "请求参数错误",
	CodeIncorrectIDParams:           "id错误",

	CodeTokenNotFound:        "token不存在",
	CodeIncorrectToken:       "无效的token",
	CodeIncorrectTokenFormat: "bearer token格式错误",
}

func (code ResponseCode) GetMsg() string {
	msg, ok := codeMsg[code]
	if !ok {
		return codeMsg[CodeInternalServerError]
	}

	return msg
}
