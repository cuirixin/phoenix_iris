package main
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"time"
)

func myAuthMiddlewareHandler(ctx iris.Context){
	ctx.WriteString("Authentication failed")
	ctx.Next()//继续执行后续的handler
}
func userProfileHandler(ctx iris.Context) {
	id:=ctx.Params().Get("id")
	ctx.Writef("id=%s now=%s", id, time.Now())
}
func userMessageHandler(ctx iris.Context){
	id:=ctx.Params().Get("id")
	ctx.WriteString(id)
}

func notFound(ctx iris.Context) {
	ctx.WriteString("Resource not found")
}

//当出现错误的时候，再试一次
func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

const maxSize = 5 << 20 // 5MB

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 中间件
	app.Use(recover.New())
	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	//app.Get("/", func(ctx iris.Context){})
	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/health") })

	//输出字符串的心跳检查
	// 类似于 app.Handle("Any", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/health
	app.Any("/health", func(ctx iris.Context) {
		ctx.WriteString("success")
		// ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
	})
	//输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// 分组
	app.PartyFunc("/users", func(users iris.Party) {
		users.Use(myAuthMiddlewareHandler)
		// http://localhost:8080/users/42/profile
		users.Get("/{id:int}/profile", cache.Handler(10*time.Second), userProfileHandler)
		// http://localhost:8080/users/messages/1
		users.Get("/inbox/{id:int}", userMessageHandler)
	})

	// 错误页面
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	// 启动
	app.Run(
		// 绑定端口和捕获退出
		iris.Addr(":8080", func(h *iris.Supervisor) {
			h.RegisterOnShutdown(func() {
				println("server terminated")
			})
		}),
		// 初始化配置
		iris.WithConfiguration(iris.YAML("./conf/iris.yml")))
}