package mongo

import (
	"github.com/kyugao/gy.logger"
	"github.com/kyugao/gy.util.json"
)

var (
	dialContext *DialContext
)

var Config struct {
	DBUrl       string
	DBName       string
	DBConnPoolSize int
}

func init() {
	// load mongo config
	err := ujson.FromFile("conf/mongo.json", &Config)
	if err != nil {
		logger.Errorf("load mongo config error %s.", ujson.ToJsonString(err))
		Config.DBUrl = "mongodb://127.0.0.1:27017"
		Config.DBName = "default_schema"
		Config.DBConnPoolSize = 5
		logger.Errorf("Set mongo config to default value %v.", Config)
	}

	mongodbContext, err := Dial(Config.DBUrl, Config.DBConnPoolSize)
	if err != nil {
		logger.Debug("Could not establish connection with db:", err)
	} else {
		logger.Debugf("Connected to db %s.", Config.DBUrl)
		dialContext = mongodbContext
		logger.Debugf("%v, %v", dialContext, mongodbContext)
	}
}