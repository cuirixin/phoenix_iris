package main

import (
	"fmt"
	"github.com/kataras/iris"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/cuirixin/phoenix_iris/config"
)

func main() {
	//f := NewLogFile()
	//defer f.Close()

	api := NewApp()
	//api.Logger().SetOutput(f) //记录日志

	url := config.GetAppUrl()
	conf := config.GetIrisConf()
	if err := api.Run(iris.Addr(url), iris.WithConfiguration(conf)); err != nil {
		color.Yellow(fmt.Sprintf("项目运行结束: %v", err))
	}
}