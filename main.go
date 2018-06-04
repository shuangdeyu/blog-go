package main

import (
	"github.com/cihub/seelog"
	"flag"
	"jyj-go/conf"
	"jyj-go/model"
	"github.com/gin-gonic/gin"
	"jyj-go/controller"
	"html/template"
)

func main(){
	configPath := flag.String("c","conf/conf.yaml","config file path") // 系统配置文件
	logPath := flag.String("l","conf/log.xml","log file path") // 日志配置文件

	// 加载日志配置模板
	logger, err := seelog.LoggerFromConfigAsFile(*logPath) // 用配置文件创建日志记录器
	if err != nil {
		seelog.Error("加载日志配置错误", err)
		return
	}
	seelog.ReplaceLogger(logger) // 替换日志记录器
	defer seelog.Flush() // 刷新立即处理所有当前排队的消息和所有当前缓冲的消息

	// 加载配置文件
	if err := conf.LoadConfiguration(*configPath); err != nil {
		seelog.Error("加载系统配置错误", err)
		return
	}

	// 初始化数据库连接
	db, err := model.InitDB()
	if err != nil {
		seelog.Error("数据库连接出错", err)
		return
	}
	defer db.Close()

	// 初始化连接redis
	/*c, err := model.InitRedis()
	if err != nil {
		seelog.Error("redis连接出错", err)
		return
	}
	defer c.Close()*/

	gin.SetMode(gin.ReleaseMode) // 设置运行环境
	r := gin.Default() // 启动gin

	// 定义模板相关设置
	setTemplate(r)

	// 设置静态文件目录
	r.Static("/public/static","./public/static")

	// 路径错误视图
	r.NoRoute()

	// 主页
	r.GET("/",controller.HomePage)
	r.GET("/index",controller.HomePage)

	// 我
	r.GET("/card", controller.CardPage)

	// 微博
	r.GET("/weibo", controller.WeiboPage)
	r.GET("/weibo/:page", controller.WeiboPage)

	// 作品集
	photo := r.Group("/photo")
	photo.GET("/shan", controller.ShanPage)
	photo.GET("/ren", controller.RenPage)
	photo.GET("/xian", controller.XianPage)
	photo.GET("/view", controller.ViewPage)
	photo.GET("/net", controller.NetPage)

	// 文章
	r.GET("article", controller.ArticlePage)
	r.GET("article/:page", controller.ArticlePage)
	r.GET("/detail", controller.DetailPage)

	r.Run(conf.GetConfiguration().Port)
}

func setTemplate(route *gin.Engine){
	// 设置自定义模板函数
	funcmap := template.FuncMap{
		"dateFormat": controller.DateFormat,
		"unescaped": controller.Unescaped,
	}
	route.SetFuncMap(funcmap)
	// 设置视图目录格式
	route.LoadHTMLGlob("./views/**/*")
}
