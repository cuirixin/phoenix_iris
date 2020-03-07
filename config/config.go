//此包用于获取配置，
//iris 框架本身的配置处理已经比较完善，
//增加这些方法主要是增加配置使用的灵活性
package config

import (
	"errors"
	"github.com/kataras/iris"
	"path"
	"sync"
	"time"

	gf "github.com/cuirixin/phoenix_corelib/libs/gotransformer"
	"github.com/cuirixin/phoenix_corelib/utils"
	"github.com/kataras/iris/v12"
	"phoenix_iris/transformer"
)

type config struct {
	Tc  *transformer.Conf
	Isc iris.Configuration
}

var cfg *config
var once sync.Once
var curpath = utils.CallerSourcePath()

func getConfig() *config {
	once.Do(func() {
		isc := iris.TOML(path.Join(curpath, ".", "/conf.tml")) // 加载配置文件
		tc := getTfConf(isc)
		cfg = &config{Tc: tc, Isc: isc}
	})
	return cfg
}

func getTfConf(isc iris.Configuration) *transformer.Conf {
	app := transformer.App{}
	g := gf.NewTransform(&app, isc.Other["App"], time.RFC3339)
	_ = g.Transformer()

	db := transformer.Mysql{}
	g.OutputObj = &db
	g.InsertObj = isc.Other["Mysql"]
	_ = g.Transformer()

	mongodb := transformer.Mongodb{}
	g.OutputObj = &mongodb
	g.InsertObj = isc.Other["Mongodb"]
	_ = g.Transformer()

	redis := transformer.Redis{}
	g.OutputObj = &redis
	g.InsertObj = isc.Other["Redis"]
	_ = g.Transformer()

	testData := transformer.TestData{}
	g.OutputObj = &testData
	g.InsertObj = isc.Other["TestData"]
	_ = g.Transformer()

	return &transformer.Conf{
		App:      app,
		Mysql:    db,
		Mongodb:  mongodb,
		Redis:    redis,
		TestData: testData,
	}
}

func GetIrisConf() iris.Configuration {
	return getConfig().Isc
}

func getTc() *transformer.Conf {
	return getConfig().Tc
}

func GetAppName() string {
	return getTc().App.Name
}

func GetAppUrl() string {
	return getTc().App.Url
}

func GetAppLoggerLevel() string {
	return getTc().App.LoggerLevel
}

func GetAppDriverType() string {
	return getTc().App.DriverType
}

func GetAppCreateSysData() bool {
	return getTc().App.CreateSysData
}

func GetMysqlConnect() string {
	return getTc().Mysql.Connect
}

func GetMysqlName() string {
	return getTc().Mysql.Name
}

func GetMysqlTName() string {
	return getTc().Mysql.TName
}

func GetMongodbConnect() string {
	return getTc().Mongodb.Connect
}


func GetTestDataUserName() string {
	return getTc().TestData.UserName
}

func GetTestDataName() string {
	return getTc().TestData.Name
}

func GetTestDataPwd() string {
	return getTc().TestData.Pwd
}

func SetAppName(arg string) error {
	if len(arg) == 0 {
		return errors.New("AppName is not be empty")
	}
	getTc().App.Name = arg
	return nil
}

func SetAppUrl(arg string) error {
	if len(arg) == 0 {
		return errors.New("AppUrl is not be empty")
	}
	getTc().App.Url = arg
	return nil
}

func SetAppLoggerLevel(arg string) error {
	if len(arg) == 0 {
		return errors.New("AppLoggerLevel is not be empty")
	}
	getTc().App.LoggerLevel = arg
	return nil
}

func SetAppDriverType(arg string) error {
	if len(arg) == 0 {
		return errors.New("DriverType is not be empty")
	}
	if arg != "Sqlite" && arg != "Mysql" {
		return errors.New("DriverType only support Sqlite or Mysql")
	}
	getTc().App.DriverType = arg
	return nil
}

func SetAppCreateSysData(arg bool) error {

	getTc().App.CreateSysData = arg

	return nil
}

func SetMysqlConnect(arg string) error {
	if len(arg) == 0 {
		return errors.New("MysqlConnect is not be empty")
	}

	getTc().Mysql.Connect = arg
	return nil
}

func SetMysqlName(arg string) error {
	if len(arg) == 0 {
		return errors.New("MysqlName is not be empty")
	}
	getTc().Mysql.Name = arg
	return nil
}

func SetMysqlTName(arg string) error {
	if len(arg) == 0 {
		return errors.New("MysqlTName is not be empty")
	}
	getTc().Mysql.TName = arg
	return nil
}

func SetMongodbConnect(arg string) error {
	if len(arg) == 0 {
		return errors.New("MongodbConnect is not be empty")
	}
	getTc().Mongodb.Connect = arg
	return nil
}

func SetTestDataUserName(arg string) error {
	if len(arg) < 6 {
		return errors.New("DataUserName is not be empty")
	}
	getTc().TestData.UserName = arg
	return nil
}

func SetTestDataName(arg string) error {
	if len(arg) < 6 {
		return errors.New("DataName 必须大于6个字符")
	}
	getTc().TestData.Name = arg
	return nil
}

func SetTestDataPwd(arg string) error {
	if len(arg) < 6 {
		return errors.New("DataPwd 必须大于6个字符")
	}
	getTc().TestData.Pwd = arg
	return nil
}