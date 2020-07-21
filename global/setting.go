package global

import "loki/pkg/setting"

// 管理配置信息和应用程序
var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
)
