package svc

import (
	"context"
	"cron/service/task/api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var ServiceContextObj *ServiceContext

type ServiceContext struct {
	Config        config.Config
	MongoClient   *mongo.Client
	MongoClientDb *mongo.Database
}

func NewServiceContext(c config.Config) *ServiceContext {
	srvCtx := &ServiceContext{
		Config: c,
	}

	//mongo初始化
	if c.Mongodb.Uri != "" {
		client, db, err := srvCtx.initMongoDB(c)
		if err != nil {
			panic("Mongodb init error" + err.Error())
		}
		srvCtx.MongoClient = client
		srvCtx.MongoClientDb = db
	}
	ServiceContextObj = srvCtx
	return srvCtx
}

// 初始化MongoDB
func (s *ServiceContext) initMongoDB(c config.Config) (*mongo.Client, *mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//设置mongo参数
	options := options.Client().
		ApplyURI(c.Mongodb.Uri).
		SetMaxPoolSize(c.Mongodb.MaxPoolSize)
	//常见数据库连接
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	pref := readpref.ReadPref{}
	err = client.Ping(ctx, &pref)

	db := client.Database(c.Mongodb.Db)
	if err != nil {
		return nil, nil, err
	}
	return client, db, nil
}
