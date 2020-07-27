package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	Parsetime    bool
	MaxIdleConns int
	MaxOpenConns int
	LogModeBool  bool
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type LogrusSettingS struct {
	Log_FILE_PATH string
	LOG_FILE_NAME string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
