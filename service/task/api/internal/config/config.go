package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mongodb MongoConf
	Email   EmailConf
}

type MongoConf struct {
	Uri         string
	Db          string
	MaxPoolSize uint64
}

type EmailConf struct {
	Host     string
	UserName string
	Password string
}
