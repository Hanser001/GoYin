package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Salt     string `mapstructure:"salt" json:"salt"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type NsqConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	Port         string         `mapstructure:"port" json:"port"`
	MysqlInfo    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	NsqInfo      NsqConfig      `mapstructure:"nsq" json:"nsq"`
	OtelInfo     OtelConfig     `mapstructure:"otel" json:"otel"`
	VideoSrvInfo VideoSrvConfig `mapstructure:"video_srv" json:"video_srv"`
}

type VideoSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
