package transformer

/*
加载系统配置
*/

type Conf struct {
	App      App
	Mysql    Mysql
	Mongodb  Mongodb
	Redis    Redis
	TestData TestData
}

type App struct {
	Name          string
	Url           string
	LoggerLevel   string
	DriverType    string
	CreateSysData bool
}

type Mysql struct {
	Connect string
	Name    string
	TName   string
}

type Mongodb struct {
	Connect string
}


type Redis struct {
	Addr     string
	Password string
	DB       string
}

type TestData struct {
	UserName string
	Name     string
	Pwd      string
}
