package conf

import (
	"github.com/go-ini/ini"
	"time"
)

type SqlDataBase struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	CharSet  string
	Prefix   string
}

type Jwt struct {
	SecretKey string
}

type Project struct {
	StaticUrlMapPath string
	TemplateGlob     string
	MedisaFilePath   string
}

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var (
	DataBase     = &SqlDataBase{}
	JwtSecretKey = &Jwt{}
	ProjectCfg   = &Project{}
	HttpServer   = &Server{}
)

func Setup() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
	}
	if err := cfg.Section("mysql").MapTo(DataBase); err != nil {
		panic(err)
	}
	if err := cfg.Section("jwt").MapTo(JwtSecretKey); err != nil {
		panic(err)
	}
	if err := cfg.Section("project").MapTo(ProjectCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("server").MapTo(HttpServer); err != nil {
		panic(err)
	}
}
