package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	LOGIN_INVALID_PARAMS:           "登录参数错误",
	LOGIN_FAILED:                   "登录失败",
	DUPLICATE_USERNAME:             "用户名重复",
	USER_NOT_FOUND:                 "用户不存在",
	USERINFO_CHECK_JWT_FAILED:      "用户信息校验失败",
	HEADER_NEED_TOKEN:              "缺少token",
	MODIFY_VERIFY_PASSWORD_FAILED:  "修改密码验证失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
